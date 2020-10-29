/*
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tejabeta/kubectl-ttl/internal/options"
	"github.com/tejabeta/kubectl-ttl/internal/ttl"
)

var (
	nameSpace    string
	allResources bool
	time         uint64
	cfgFile      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-ttl",
	Short: "A tiny tool to add time to live option to k8s resources",
	Long: `A tiny kubectl plugin to add time to live option
to k8s resources within a namespace.

Tool helps to create a job within the specified namespace
to kill/clean the resources after certain time. 
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		options, err := options.GetOptions()
		if err != nil {
			log.Errorln(err)
			return
		}
		options.AllResources = allResources
		options.Namespace = nameSpace
		options.TimeToLive = time
		ttl.KubectlTTL(options)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&nameSpace, "ns", "n", "", "namespace to apply ttl")
	rootCmd.Flags().Uint64VarP(&time, "time", "t", 15, "time in minutes to keep the resource alive")
	rootCmd.Flags().BoolVarP(&allResources, "all", "a", false, "boolean flag to take all resources into consideration")
	rootCmd.MarkFlagRequired("ns")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kubectl-ttl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kubectl-ttl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
