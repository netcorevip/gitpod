// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package step

import (
	"fmt"

	"github.com/gitpod-io/gitpod/ws-deployment/pkg/common"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/runner"
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

func CreateCluster(context *common.ProjectContext, cluster common.WorkspaceCluster) error {
	err := doesClusterExist(context, cluster)
	if err != nil {
		return err
	}
	err := generateTerraformModules(context, cluster)
	if err != nil {
		return err
	}
	err = applyTerraformModules(context, cluster)
	if err != nil {
		return err
	}
	return nil
}

func doesClusterExist(context *common.ProjectContext, cluster common.WorkspaceCluster) (bool, error) {
	runner.ShellRun()
}

func generateTerraformModules(context *common.ProjectContext, cluster common.WorkspaceCluster) error {
	err := runner.ShellRun(DefaultTFModuleGeneratorScriptPath, []string{generateDefaultScriptArgsString(context, cluster)})
	return err
}

func generateDefaultScriptArgsString(context *common.ProjectContext, cluster common.WorkspaceCluster) string {
	// example `-e staging -l europe-west1 -n us89 -t k3s -g gitpod-staging -w gitpod-staging -d gitpod-staging-com`
	argsString := fmt.Sprintf("-e %s -l %s -n %s -t %s -g %s -w %s -d %s",
		context.Environment, cluster.Region, cluster.Name, cluster.ClusterType, context.ProjectId, context.Network, context.DNSZone,
	)
	return argsString
}

func applyTerraformModules(context *common.ProjectContext, cluster common.WorkspaceCluster) error {
	tfModulesDir := fmt.Sprintf(DefaultGeneratedTFModulePathTemplate, context.Environment, cluster.Name)
	commandToRun := fmt.Sprintf("cd %s && terraform init && terraform apply -auto-approve", tfModulesDir)
	err := runner.ShellRun("/bin/sh", []string{"-c", commandToRun})
	return err
}
