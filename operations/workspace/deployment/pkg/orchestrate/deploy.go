package orchestrate

import (
	"sync"

	log "github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/common"
	"github.com/gitpod-io/gitpod/ws-deployment/pkg/step"
)

func Deploy(context *common.ProjectContext, clusters []common.WorkspaceCluster) error {

	var wg sync.WaitGroup
	wg.Add(len(clusters))
	defer wg.Wait()

	for _, cluster := range clusters {
		go func(context *common.ProjectContext, cluster common.WorkspaceCluster) {
			defer wg.Done()
			err := step.CreateCluster(context, cluster)
			if err != nil {
				log.Infof("Error creating cluster %s: %s", cluster.Name, err)
			}
		}(context, cluster)
	}

	return nil
}
