package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yaoalex/FortniteCLI/db"
)

var ps *PlayerStats

var showStatsCmd = &cobra.Command{
	Use:   "showstats",
	Short: "Show all stat entries",
	Long: (`Returns all BoltDB entries for player stats
	Stats will come of the form (kd/total kills)
	-- Example Usage --
	fortnitecli showstats -p maikeyaoyao`),
	Run: func(cmd *cobra.Command, args []string) {
		stats, err := db.AllStats()
		if err != nil {
			fmt.Println("Something went wrong accessing your stats: ", err)
			os.Exit(1)
		}
		player, _ := cmd.Flags().GetString("player")
		for _, stat := range stats {
			if player != "" && player != stat.Player {
				continue
			}
			fmt.Println("Player: ", stat.Player, "  Note: ", stat.Note, "  Time Added: ", time.Unix(int64(stat.Added), 0))
			printStats(stat.Value)
		}
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get the stats based on Fortnite player name",
	Long: `Calls the fortnite api
	-- Example Usage --
	fortnitecli stats maikeyaoyao -s "Learned 90s"`,
	Run: func(cmd *cobra.Command, args []string) {
		player := strings.Join(args, " ")
		// fmt.Println(save)
		stats, err := ps.GetPlayerStats(player)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		} else if stats == nil {
			fmt.Println("No stats for", player)
		} else {
			fmt.Println("Stats for", player)
			fmt.Println(stats)
			note, _ := cmd.Flags().GetString("save")
			if note != "" {
				err := db.CreateStat(player, note, stats)
				if err != nil {
					fmt.Println("Error trying to add to db: ", err)
					os.Exit(1)
				} else {
					fmt.Println("Successfully added to db")
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(showStatsCmd)
	showStatsCmd.Flags().StringP("player", "p", "", "Search for a specific player")
	statsCmd.Flags().StringP("save", "s", "", "Saved entry to database")
	ps = &PlayerStats{Client: &BasePlayerClient{}}
}
