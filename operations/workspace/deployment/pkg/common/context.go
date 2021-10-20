// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package common

import "k8s.io/client-go/kubernetes"

// ProjectContext is a wrapper which contains information required to communicate to the
// right GCP project with correct inputs
type ProjectContext struct {
	GCPSACredFile string `yaml:"gcpSACredFile"`
	ProjectId     string `yaml:"projectId"`
	Environment   string `yaml:"environment"`
	Network       string `yaml:"network"`
	DNSZone       string `yaml:"dnsZone"`
}

// ClusterContext contains the context to access the cluster
type ClusterContext struct {
	KubeconfigPath string
	Client         *kubernetes.Clientset
}
