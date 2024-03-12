package main

import (
	apiserver "k8s.io/apiserver/pkg/server"

	gpunums "github.com/Garrybest/k8s-example/cmd/gpunums"
	scheduletest "github.com/Garrybest/k8s-example/cmd/schedtest"
	"github.com/spf13/cobra"
)

func main() {
	ctx := apiserver.SetupSignalContext()

	// 创建一个根命令
	rootCmd := &cobra.Command{
		Use:   "k8s-testing",
		Short: "My k8s testing tool based on corba",
	}

	// 添加 "task" 子命令
	taskCmd := &cobra.Command{
		Use:   "task",
		Short: "Perform a task",
	}

	// modify your example app command here

	cmdScheduletest := scheduletest.NewCommand(ctx)
	cmdGPUNums := gpunums.NewCommand(ctx)
	taskCmd.AddCommand(cmdScheduletest)
	taskCmd.AddCommand(cmdGPUNums)

	rootCmd.AddCommand(taskCmd)

	// 执行根命令
	rootCmd.Execute()
}
