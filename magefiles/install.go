//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Install() error {
	return sh.Run("go", "install", "./cli/plan.go")
}
