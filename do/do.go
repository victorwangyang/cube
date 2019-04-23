package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	initCmd()
}

// InitCmd is initing commands
func initCmd() {

	var fileName string

	//Define Start-command to start a Master
	var cmdStart = &cobra.Command{
		Use:   "start ",
		Short: "start master",
		Long:  `create is for starting master server to listen to cli`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if fileName == "" {
				fmt.Println("Error: Yaml file is needed")
				return
			}
			//	cluster.StartCluster(fileName)
		},
	}
	cmdStart.Flags().StringVarP(&fileName, "file", "f", "", "file to start the cluster")

	//Define Kill-command to stop a Master
	var cmdKill = &cobra.Command{
		Use:   "stop ",
		Short: "stop master",
		Long:  `create is for stopping master server to listen to cli`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			//cluster.KillCluster()
		},
	}

	//Define getcluster-command to stop a Master
	var cmdGetcluster = &cobra.Command{
		Use:   "getcluster ",
		Short: "get cluster info",
		Long:  `getcluster is to get cluster nodes info`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			getCluster()
		},
	}

	var rootCmd = &cobra.Command{Use: "do"}
	rootCmd.AddCommand(cmdKill, cmdStart, cmdGetcluster)
	rootCmd.Execute()

}
