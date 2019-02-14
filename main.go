package main

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/yaoalex/FortniteCLI/cmd"
	"github.com/yaoalex/FortniteCLI/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "fortnitestats.db")
	err := db.Init(dbPath)
	if err != nil {
		fmt.Println("Something went wrong initializing db: ", err.Error())
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
