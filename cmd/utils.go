package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PlayerClient interface {
	MakeRequest(url string) (map[string]interface{}, error)
}

type PlayerStats struct {
	Client PlayerClient
}

type BasePlayerClient struct{}

func (ps *PlayerStats) GetPlayerStats(player string) (map[string]interface{}, error) {
	uidURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/id?username=%s`, player)
	uidData, err := ps.Client.MakeRequest(uidURL)
	if err != nil {
		return nil, err
	}
	uid := uidData["uid"]
	if uid == nil {
		return nil, nil
	}

	statsURL := fmt.Sprintf(`https://fortnite-public-api.theapinetwork.com/prod09/users/public/br_stats?user_id=%s&platform=pc`, uidData["uid"])
	statsData, err := ps.Client.MakeRequest(statsURL)
	if err != nil {
		return nil, err
	}
	return statsData, nil
}

func (bpc *BasePlayerClient) MakeRequest(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func printStats(stats map[string]interface{}) {
	fmt.Printf("Solo: %.2f|%v    Duo: %.2f|%v    Squad: %.2f|%v \n", stats["kd_solo"], stats["kills_solo"],
		stats["kd_duo"], stats["kills_duo"], stats["kd_squad"], stats["kills_squad"])
}
