package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get the stats based on Fortnite player name",
	Long:  "Calls the fortnite",
	Run: func(cmd *cobra.Command, args []string) {
		player := strings.Join(args, " ")
		// save, _ := cmd.Flags().GetString("save")
		// fmt.Println(save)
		stats, err := getPlayerStats(player)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		} else if stats == nil {
			fmt.Println("No stats for", player)
		} else {
			fmt.Println("Stats for", player)
			fmt.Println(stats)
		}
	},
}

func getPlayerStats(player string) (map[string]interface{}, error) {
	uidURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/id?username=%s`, player)
	resp, err := http.Get(uidURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var uidData map[string]interface{}
	err = json.Unmarshal(body, &uidData)
	if err != nil {
		return nil, err
	}
	uid := uidData["uid"]
	if uid == nil {
		return nil, nil
	}

	statsURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/public/br_stats?user_id=%s&platform=pc`, uidData["uid"])
	resp, err = http.Get(statsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var statsData map[string]interface{}
	err = json.Unmarshal(body, &statsData)
	if err != nil {
		return nil, err
	}
	return statsData, nil
}

func init() {
	RootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringP("save", "s", "", "Save entry to database")
}
