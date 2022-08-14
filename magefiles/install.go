//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Install() error {
	if err := sh.Run("go", "install", "./cli/plan.go"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "-ldflags", "-s -w", "./pland/pland.go"); err != nil {
		return err
	}
	return nil
}
