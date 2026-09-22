/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"log"
	"os"

	"github.com/Lauloque/gator/internal/config"
)

func main() {
	// Read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var appState = state{}
	appState.configPtr = &cfg

	cmds := commands{
		registeredCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

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
