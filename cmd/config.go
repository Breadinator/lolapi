package cmd

import (
	"fmt"

	api "github.com/breadinator/lolapi/lolapi"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get config location.",
	Long:  "Returns the path to the config file, and generates one if not available.",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := api.GetConfigPath()
		if err != nil {
			panic(err)
		}
		exists := api.CheckPathExists(path)

		if !exists {
			c := new(api.Config)
			api.GenerateNewConfig(c)
			api.WriteNewConfig(c, path, jsoniter.ConfigCompatibleWithStandardLibrary)
		}

		fmt.Printf("The config for this program is stored at %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
