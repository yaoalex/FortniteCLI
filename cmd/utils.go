package cmd

import "fmt"

func printStats(stats map[string]interface{}) {
	fmt.Printf("Solo: %.2f|%v    Duo: %.2f|%v    Squad: %.2f|%v \n", stats["kd_solo"], stats["kills_solo"],
		stats["kd_duo"], stats["kills_duo"], stats["kd_squad"], stats["kills_squad"])
}
