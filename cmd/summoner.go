/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	api "github.com/breadinator/lolapi/lolapi"
	"github.com/breadinator/lolapi/lolapi/endpoints"
	"github.com/spf13/cobra"
)

// summonerCmd represents the summoner command
var summonerCmd = &cobra.Command{
	Use:   "summoner",
	Short: "Get a summoner from their name, PUUID or ID.",
	Long:  `Get a summoner from their name, PUUID or ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		regionString, err := cmd.Flags().GetString("region")
		if err != nil {
			panic(err)
		}
		region := api.RegionFromString(regionString)
		if region == api.RegionInvalid {
			fmt.Println("Invalid region " + regionString)
			os.Exit(1)
		}

		name, _ := cmd.Flags().GetString("name")
		puuid, _ := cmd.Flags().GetString("puuid")
		id, _ := cmd.Flags().GetString("id")

		conf, err := api.GetConfig()
		if err != nil {
			panic(err)
		}
		client := api.Client{
			Key:    conf.APIKey,
			Limit:  conf.Limit,
			Region: region,
		}

		var resp []byte
		var status int

		if name != "" {
			resp, status, err = client.Call(endpoints.GetSummonerByName, name)
		} else if puuid != "" {
			resp, status, err = client.Call(endpoints.GetSummonerByPUUID, puuid)
		} else if id != "" {
			resp, status, err = client.Call(endpoints.GetSummonerByAccountID, id)
		} else {
			fmt.Println("Requires either a summoner name, PUUID or ID.")
			os.Exit(1)
		}

		if err != nil || status != 200 {
			fmt.Printf("ERR: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(string(resp))
	},
}

func init() {
	getCmd.AddCommand(summonerCmd)
	summonerCmd.Flags().StringP("region", "r", "", "Summoner region (e.g. EUW)")
	summonerCmd.MarkFlagRequired("region")

	summonerCmd.Flags().String("name", "", "Summoner name")
	summonerCmd.Flags().String("puuid", "", "Summoner PUUID")
	summonerCmd.Flags().String("id", "", "Summoner ID")
}
