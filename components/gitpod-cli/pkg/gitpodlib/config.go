// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package gitpodlib

type GitpodImage struct {
	File    string
	Context string `yaml:"context,omitempty"`
}

type gitpodPort struct {
	Number int32 `yaml:"port"`
}

type gitpodTask struct {
	Command string `yaml:"command,omitempty"`
	Init    string `yaml:"init,omitempty"`
}

type GitpodFile struct {
	Image             interface{}  `yaml:"image,omitempty"`
	Ports             []gitpodPort `yaml:"ports,omitempty"`
	Tasks             []gitpodTask `yaml:"tasks,omitempty"`
	CheckoutLocation  string       `yaml:"checkoutLocation,omitempty"`
	WorkspaceLocation string       `yaml:"workspaceLocation,omitempty"`
}

//SetImageName 按名称配置预先构建的 docker 镜像
func (cfg *GitpodFile) SetImageName(name string) {
	cfg.Image = name
}

// SetImage 将 Dockerfile 配置为工作区映像
func (cfg *GitpodFile) SetImage(img GitpodImage) {
	cfg.Image = img
}

// AddPort 将端口添加到暴露端口列表中
func (cfg *GitpodFile) AddPort(port int32) {
	cfg.Ports = append(cfg.Ports, gitpodPort{
		Number: port,
	})
}

// AddTask 添加工作区启动任务
func (cfg *GitpodFile) AddTask(task ...string) {
	if len(task) > 1 {
		cfg.Tasks = append(cfg.Tasks, gitpodTask{
			Command: task[0],
			Init:    task[1],
		})
	} else {
		cfg.Tasks = append(cfg.Tasks, gitpodTask{
			Command: task[0],
		})
	}
}
