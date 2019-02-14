# FortniteCLI
command line interface to analyze fortnite stats

`go get -u github.com/yaoalex/FortniteCLI`

required packages:
`go get -u "github.com/boltdb/bolt"`
`go get -u ""github.com/spf13/cobra"`

`fortnitecli stats <username> --save <note>`
the stats command gets your current fortnite stats and gives you the option of saving the stats into boltDB with a side note

`fortnitecli showstats --player <player>`
the showstats command shows all the saved stats from boltDB with the option of specifying which player to show
