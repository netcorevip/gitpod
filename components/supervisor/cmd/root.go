// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/gitpod-io/gitpod/common-go/log"
)

// rootCmd 表示在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "supervisor",
	Short: "Workspace container init process",
}

var (
	// ServiceName is the name we use for tracing/logging
	ServiceName = "supervisor"
	// Version of this service - set during build
	Version = ""
)

// Execute 将所有子命令添加到根命令并适当地设置标志。
//这由 main.main() 调用。 它只需要对 rootCmd 发生一次。
func Execute() {
	log.Init(ServiceName, Version, true, false)
	log.Log.Logger.AddHook(fatalTerminationLogHook{})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type fatalTerminationLogHook struct{}

func (fatalTerminationLogHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.FatalLevel}
}

func (fatalTerminationLogHook) Fire(e *logrus.Entry) error {
	msg := e.Message
	if err := e.Data[logrus.ErrorKey]; err != nil {
		msg += ": " + err.(error).Error()
	}

	return os.WriteFile("/dev/termination-log", []byte(msg), 0644)
}
