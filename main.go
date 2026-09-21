/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"fmt"
	"log"

	"github.com/Lauloque/gator/internal/config"
)

func main() {
	fmt.Println("hi there")

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	username, err := config.GetUsername()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser(username)

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", cfg)
}
