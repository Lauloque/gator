/* SPDX-License-Identifier: GPL-3.0-or-later */
package config

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	cfg, err := Read()
	if err != nil {
		t.Fatalf("Read() return an error: %v", err)
	}

	if cfg.DBURL == "" {
		fmt.Printf("%#v\n", cfg)
		t.Error("db_url was not populated")
	}

	username, err := GetUsername()
	if err != nil {
		t.Fatalf("GetUsername() returned an error: %v", err)
	}

	if cfg.CurrentUserName != username {
		t.Errorf("expected current_user_name to be '%s', got '%s'", username, cfg.CurrentUserName)
	}
}
