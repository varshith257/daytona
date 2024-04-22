// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"errors"

	"github.com/daytonaio/daytona/pkg/serverapiclient"
	"github.com/daytonaio/daytona/pkg/views/workspace/create"
)

func GetCreationDataFromPrompt(workspaceNames []string, userGitProviders []serverapiclient.GitProvider, manual bool, multiProject bool, advancedFlag bool) (string, []serverapiclient.CreateWorkspaceRequestProject, error) {
	var projectList []serverapiclient.CreateWorkspaceRequestProject
	var providerRepo serverapiclient.GitRepository
	var err error

	if !manual && userGitProviders != nil && len(userGitProviders) > 0 {
		providerRepo, err = getRepositoryFromWizard(userGitProviders, 0)
		if err != nil {
			return "", nil, err
		}
		if providerRepo == (serverapiclient.GitRepository{}) {
			return "", nil, nil
		}
	}

	if providerRepo.Url == nil {
		providerRepo.Url = new(string)
	}

	workspaceCreationPromptResponse, err := create.RunInitialForm(*providerRepo.Url, multiProject, advancedFlag)
	if err != nil {
		return "", nil, err
	}

	if workspaceCreationPromptResponse.PrimaryProject == (serverapiclient.CreateWorkspaceRequestProject{}) {
		return "", nil, errors.New("primary project is required")
	}

	projectList = []serverapiclient.CreateWorkspaceRequestProject{workspaceCreationPromptResponse.PrimaryProject}

	if multiProject {
		for i := 0; workspaceCreationPromptResponse.AddingMoreProjects; i++ {

			if !manual && userGitProviders != nil && len(userGitProviders) > 0 {
				providerRepo, err = getRepositoryFromWizard(userGitProviders, i+1)
				if err != nil {
					return "", nil, err
				}
				if providerRepo == (serverapiclient.GitRepository{}) {
					return "", nil, nil
				}

				primaryProject := serverapiclient.CreateWorkspaceRequestProject{
					Source: &serverapiclient.CreateWorkspaceRequestProjectSource{
						Repository: &providerRepo,
					},
				}

				workspaceCreationPromptResponse.SecondaryProjects = append(workspaceCreationPromptResponse.SecondaryProjects, primaryProject)
			}

			workspaceCreationPromptResponse, err = create.RunProjectForm(workspaceCreationPromptResponse)
			if err != nil {
				return "", nil, err
			}
		}
		projectList = append(projectList, workspaceCreationPromptResponse.SecondaryProjects...)
	}

	suggestedName := create.GetSuggestedWorkspaceName(*workspaceCreationPromptResponse.PrimaryProject.Source.Repository.Url)

	workspaceCreationPromptResponse, err = create.RunWorkspaceNameForm(workspaceCreationPromptResponse, suggestedName, workspaceNames)
	if err != nil {
		return "", nil, err
	}

	if workspaceCreationPromptResponse.WorkspaceName == "" {
		return "", nil, errors.New("workspace name is required")
	}

	return workspaceCreationPromptResponse.WorkspaceName, projectList, nil
}
