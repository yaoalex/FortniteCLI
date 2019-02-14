package cmd_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoalex/FortniteCLI/cmd"
)

type mockClient struct{}

func (mc *mockClient) MakeRequest(url string) (map[string]interface{}, error) {
	uidRequest := `https://fortnite-public-api.theapinetwork.com/prod09/users/id?username=%s`
	statsRequest := `https://fortnite-public-api.theapinetwork.com/prod09/users/public/br_stats?user_id=%s&platform=pc`

	if url == fmt.Sprintf(uidRequest, "invalid_player") {
		return nil, nil
	} else if url == fmt.Sprintf(uidRequest, "valid_player") {
		return map[string]interface{}{"uid": "valid_uid"}, nil
	} else if url == fmt.Sprintf(statsRequest, "valid_uid") {
		return map[string]interface{}{"kd_solo": 1.0, "kd_squads": 2}, nil
	} else {
		return nil, errors.New("Failed request")
	}
}

func TestGetPlayerStats(t *testing.T) {
	assert := assert.New(t)
	type TestCase struct {
		Name     string
		Player   string
		Error    bool
		Expected map[string]interface{}
	}
	testCases := []TestCase{
		{
			Name:     "getPlayerStats - Invalid player name",
			Player:   "invalid_player",
			Error:    false,
			Expected: nil,
		},
		{
			Name:     "getPlayerStats - Valid player name",
			Player:   "valid_player",
			Error:    false,
			Expected: map[string]interface{}{"kd_solo": 1.0, "kd_squads": 2},
		},
		{
			Name:     "getPlayerStats - Player name causes error",
			Player:   "error_player",
			Error:    true,
			Expected: nil,
		},
	}
	for _, tc := range testCases {
		ps := &cmd.PlayerStats{Client: &mockClient{}}
		stats, err := ps.GetPlayerStats(tc.Player)
		if tc.Error {
			assert.NotNil(err)
		}
		assert.Equal(stats["kd_solo"], tc.Expected["kd_solo"])
	}
}
