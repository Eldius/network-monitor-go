/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/Eldius/network-monitor-go/display"
	"github.com/Eldius/network-monitor-go/pingtools"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping a host",
	Long:  `Just ping a host`,
	Run: func(cmd *cobra.Command, args []string) {
		if *quiet {
			display.DisplayPingResponse(pingtools.Ping(*pingTargets, *packets))
		} else {
			display.DisplayPing(pingtools.Ping(*pingTargets, *packets))
		}
	},
}

var pingTargets *[]string
var packets *int
var quiet *bool

func init() {
	rootCmd.AddCommand(pingCmd)
	pingTargets = pingCmd.Flags().StringArrayP("target", "t", []string{"8.8.8.8"}, "Ping target")
	packets = pingCmd.Flags().IntP("packets", "p", 4, "Ping packets to send")
	quiet = pingCmd.Flags().BoolP("quiet", "q", false, "Executes without needs to interaction")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
