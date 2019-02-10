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
	Run: func(cmd *cobra.Command, args []string) {
		player := strings.Join(args, " ")
		uidURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/id?username=%s`, player)
		resp, err := http.Get(uidURL)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		var uidData map[string]interface{}
		err = json.Unmarshal(body, &uidData)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		statsURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/public/br_stats?user_id=%s&platform=pc`, uidData["uid"])
		resp, err = http.Get(statsURL)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		var statsData map[string]interface{}
		err = json.Unmarshal(body, &statsData)
		fmt.Printf("Stats for player \"%s\":", player)
		fmt.Println(statsData)
	},
}

func init() {
	RootCmd.AddCommand(statsCmd)
}
