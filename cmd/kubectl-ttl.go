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
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tejabeta/kubectl-ttl/internal/handler"
	"github.com/tejabeta/kubectl-ttl/internal/util"

	"github.com/ghodss/yaml"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	time    uint64
	cfgFile string
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
		initTTL()
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
	rootCmd.Flags().Uint64VarP(&time, "time", "t", 15, "time in minutes to keep the resource alive")
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

func initTTL() {
	stdin := os.Stdin

	info, err := stdin.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if info.Size() <= 0 {
		log.Fatalln("No input provided")
	}

	in, err := ioutil.ReadAll(stdin)
	if err != nil {
		log.Fatalln(err)
	}

	if !util.IsJSON(in) && !util.IsYAML(in) {
		log.Fatalln("Invalid input format")
	}

	if util.IsYAML(in) {
		in, err = yaml.YAMLToJSON(in)
		if err != nil {
			log.Errorf("Error while converting yaml to json : %v", err)
		}
	}

	resInfo := util.GetResDetails(string(in))
	if !isResValid(resInfo) {
		log.Fatalln("Invalid resource provided")
	}

	for _, v := range resInfo {
		if !handler.CheckSA(v.Namespace, strings.ToLower(v.Kind+"-"+v.Name+"-SA-ttl")) {
			handler.CreateSA(
				v.Namespace,
				strings.ToLower(v.Kind+"-"+v.Name+"-SA-ttl"),
			)
		}

		if !handler.CheckRole(v.Namespace, strings.ToLower(v.Kind+"-"+v.Name+"-Role-ttl")) {
			handler.CreateRole(
				v.Namespace,
				strings.ToLower(v.Kind+"-"+v.Name+"-Role-ttl"),
				roleResource(v.Kind),
				v.Name,
			)
		}

		if !handler.CheckRB(v.Namespace, strings.ToLower(v.Kind+"-"+v.Name+"-RB-ttl")) {
			handler.CreateRB(
				v.Namespace,
				strings.ToLower(v.Kind+"-"+v.Name+"-RB-ttl"),
				strings.ToLower(v.Kind+"-"+v.Name+"-SA-ttl"),
				strings.ToLower(v.Kind+"-"+v.Name+"-Role-ttl"),
			)
		}

		if !handler.CheckJob(v.Namespace, strings.ToLower(v.Kind+"-"+v.Name+"-Job-ttl")) {
			handler.CreateJob(
				v.Namespace,
				strings.ToLower(v.Kind+"-"+v.Name+"-Job-ttl"),
				v.Kind,
				v.Name,
				strings.ToLower(v.Kind+"-"+v.Name+"-SA-ttl"),
				time,
			)
		}
	}

}

func isResValid(r []util.ResInfo) bool {
	for _, v := range r {
		if !util.ValidResList[v.Kind] {
			return false
		}
	}
	return true
}

func roleResource(r string) string {
	switch r {
	case "Deployment":
		return "deployments"
	case "Pod":
		return "pods"
	case "Service":
		return "services"
	case "Ingress":
		return "ingresses"
	case "ConfigMap":
		return "configmaps"
	case "Secret":
		return "secrets"
	case "ReplicaSet":
		return "replicasets"
	case "PersistentVolumeClaim":
		return "persistentvolumeclaims"
	case "PersistentVolume":
		return "persistentvolumes"
	case "ServiceAccount":
		return "serviceaccounts"
	case "Job":
		return "jobs"
	}
	return ""
}
