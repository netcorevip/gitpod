// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package step

import (
	"fmt"
	"strings"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/common"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/runner"
	"golang.org/x/xerrors"
)

const (
	// DefaultTFModuleGeneratorScriptPath is the path to script that must be invoked
	// from its parent dir in order to generate terraform modules
	DefaultTFModuleGeneratorScriptPath = "dev/build-ws-cluster/build-ws-cluster.sh"

	// DefaultGeneratedTFModulePathTemplate represents the path template where the default module
	// would be generated
	//
	// deploy/{environment}/ws-{name}/terraform
	DefaultGeneratedTFModulePathTemplate = "deploy/%s/ws-%s/terraform"
)

func CreateCluster(context *common.ProjectContext, cluster *common.WorkspaceCluster) error {
	exists, err := doesClusterExist(context, cluster)
	// If we see an error finding out if cluster exists
	if err != nil {
		return xerrors.Errorf("Issue finding out if cluster exists: %s", err)
	}
	// If the cluster already exists
	if exists {
		return xerrors.Errorf("Cluster '%s' already exists", cluster.Name)
	}
	// If there is neither an error nor the cluster exist then continue
	err = generateTerraformModules(context, cluster)
	if err != nil {
		return err
	}
	err = applyTerraformModules(context, cluster)
	if err != nil {
		return err
	}
	return nil
}

func doesClusterExist(context *common.ProjectContext, cluster *common.WorkspaceCluster) (bool, error) {
	stdOut, stdErr, err := runner.ShellRun("gcloud", []string{})
	log.Log.Errorf("out, err, err: %s, %s, %s", stdOut, stdErr, err)
	commandToRun := "gcloud"
	// container clusters describe gp-stag-ws-us11-us-weswt1 --project gitpod-staging --region us-west1
	argsString := fmt.Sprintf("container clusters describe %s --project %s --region %s", cluster.Name, context.Id, cluster.Region)
	stdOut, stdErr, err = runner.ShellRun(commandToRun, strings.Split(argsString, " "))
	if err == nil {
		return true, nil
	}
	if strings.Contains(stdErr, "No cluster named") {
		return false, nil
	}
	log.Log.Errorf("Encountered an error while trying to describe cluster: %s, %s", stdErr, stdOut)
	return false, err
}

func generateTerraformModules(context *common.ProjectContext, cluster *common.WorkspaceCluster) error {
	stdOut, stdErr, err := runner.ShellRun(DefaultTFModuleGeneratorScriptPath, generateDefaultScriptArgs(context, cluster))
	log.Log.Errorf("Error generating Terraform modules: stdErr %s, stdOut: %s", stdErr, stdOut)
	return err
}

func generateDefaultScriptArgs(context *common.ProjectContext, cluster *common.WorkspaceCluster) []string {
	// example `-e staging -l europe-west1 -n us89 -t k3s -g gitpod-staging -w gitpod-staging -d gitpod-staging-com`
	argsString := fmt.Sprintf("-e %s -l %s -n %s -t %s -g %s -w %s -d %s",
		context.Environment, cluster.Region, cluster.Name, cluster.ClusterType, context.Id, context.Network, context.DNSZone,
	)
	return strings.Split(argsString, " ")
}

func applyTerraformModules(context *common.ProjectContext, cluster *common.WorkspaceCluster) error {
	tfModulesDir := fmt.Sprintf(DefaultGeneratedTFModulePathTemplate, context.Environment, cluster.Name)
	commandToRun := fmt.Sprintf("cd %s && terraform init && terraform apply", tfModulesDir)
	err := runner.ShellRunWithDefaultConfig("/bin/sh", []string{"-c", commandToRun})
	return err
}
