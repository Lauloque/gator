// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"fmt"

	"github.com/Lauloque/gator/internal/config"
)

type state struct {
	configPtr *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	registeredCmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.registeredCmds[cmd.name]
	if !exists {
		return fmt.Errorf("Unknown command: '%s'", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCmds[name] = f
}
