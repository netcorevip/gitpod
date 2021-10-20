package orchestrate

import (
	"sync"

	"github.com/gitpod-io/gitpod/ws-deployment/pkg/check/prerun"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/common"
	v1 "github.com/gitpod-io/gitpod/ws-deployment/pkg/config/v1"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/step"
)

func Deploy(c *v1.Config, environment *common.Environment) error {
	createClustersSteps, err := buildCreateClustersStep(c, environment)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(createClustersSteps))
	defer wg.Wait()

	for i := range createClustersSteps {
		stepToExecute := createClustersSteps[i]
		go func() {
			defer wg.Done()
			stepToExecute.Run()
		}()
	}
	// TODO(prs): This will change once we have code handling error cases
	// for building cluster
	return nil
}

func buildCreateClustersStep(config *v1.Config, environment *common.Environment) ([]step.Step, error) {
	projectContext := &common.ProjectContext{
		Environment: *environment,
	}
	steps := make([]step.Step, 0, 0)
	for _, wc := range config.WorkspaceClusters {
		step, err := buildCreateClusterStep(projectContext, &wc)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}

func buildCreateClusterStep(projectContext *common.ProjectContext, workspaceCluster *common.WorkspaceCluster) (step.Step, error) {
	createClusterPreruns := &prerun.CreateClusterPreruns{
		Cluster:        workspaceCluster,
		ProjectContext: projectContext,
	}
	err := createClusterPreruns.CreatePreRuns()
	if err != nil {
		return nil, err
	}
	createClusterStep := &step.CreateClusterStep{
		ProjectContext: projectContext,
		Cluster:        workspaceCluster,
		PreRuns:        createClusterPreruns,
	}
	return createClusterStep, nil
}
