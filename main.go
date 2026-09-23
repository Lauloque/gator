/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Lauloque/gator/internal/config"
	"github.com/Lauloque/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	var appState = state{}
	appState.configPtr = &cfg
	appState.dbPtr = dbQueries

	cmds := commands{
		registeredCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	// REPL
	if len(os.Args) < 2 {
		log.Fatal("Excpeced usage: 'gator <command> [args...]'")
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
