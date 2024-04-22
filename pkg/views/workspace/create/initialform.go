// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package create

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/serverapiclient"
	"github.com/daytonaio/daytona/pkg/views"
)

func NewModel(providerRepo serverapiclient.GitRepository, multiProject bool) Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	var primaryRepoUrl string

	if providerRepo.Url != nil {
		primaryRepoUrl = *providerRepo.Url
	}

	secondaryProjectsCountString := ""

	primaryRepoPrompt := huh.NewInput().
		Title("Primary project repository").
		Value(&primaryRepoUrl).
		Key("primaryProjectRepo").
		Validate(func(str string) error {
			result, err := util.GetValidatedUrl(str)
			if err != nil {
				return err
			}
			primaryRepoUrl = result
			return nil
		})

	secondaryProjectCountPrompt := huh.NewInput().
		Title("How many secondary projects?").
		Value(&secondaryProjectsCountString).
		Validate(func(str string) error {
			count, err := strconv.Atoi(str) // Try to convert the input string to an integer
			if err != nil {
				return errors.New("enter a valid number")
			}
			if count > maximumSecondaryProjects {
				return errors.New("maximum 8 secondary projects allowed")
			}
			return nil
		})

	dTheme := views.GetCustomTheme()

	m.form = huh.NewForm(
		huh.NewGroup(
			primaryRepoPrompt,
		).WithHide(primaryRepoUrl != ""),
		huh.NewGroup(
			secondaryProjectCountPrompt,
		).WithHide(!multiProject),
	).WithTheme(dTheme).
		WithWidth(maxWidth).
		WithShowHelp(false).
		WithShowErrors(true)

	m.altForm = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Choose a container registry"),
		),
	).WithTheme(dTheme).
		WithWidth(maxWidth).
		WithShowHelp(false).
		WithShowErrors(true)

	m.form.Init()
	m.altForm.Init()

	// err := m.form.Run()
	// if err != nil {
	// 	return m
	// }

	secondaryProjectsCount, err := strconv.Atoi(secondaryProjectsCountString)
	if err != nil {
		secondaryProjectsCount = 0
	}

	providerRepo.Url = &primaryRepoUrl
	m.workspaceCreationPromptResponse = WorkspaceCreationPromptResponse{
		PrimaryRepository:     providerRepo,
		SecondaryProjectCount: secondaryProjectsCount,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.altForm.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "esc":
			var cmd tea.Cmd
			if m.altscreen {
				cmd = tea.ExitAltScreen
				m.form.NextField()
			} else {
				cmd = tea.EnterAltScreen
				m.altForm.NextField()
			}
			m.altscreen = !m.altscreen
			return m, cmd
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.quitting {
		return "Bye!\n"
	}

	inlineMode := m.form.View()
	altscreenMode := m.altForm.View()

	var mode string
	if m.altscreen {
		mode = altscreenMode
	} else {
		mode = inlineMode
	}

	return mode
}

func ProjectConfigTry(providerRepo serverapiclient.GitRepository, multiProject bool) {
	m := NewModel(providerRepo, multiProject)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
