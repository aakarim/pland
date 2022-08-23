//go:build mage

package main

import (
	"fmt"

	"github.com/aakarim/pland/cli/ui/common"
	"github.com/magefile/mage/sh"
)

func Install() error {
	if err := sh.Run("go", "install", "./cli/plan.go"); err != nil {
		return err
	}
	if err := sh.Run("go", "install", "-ldflags", "-s -w", "./pland/pland.go"); err != nil {
		return err
	}
	fmt.Println(common.Styles.Keyword.Render("installed."))
	return nil
}

func ReinstallService() error {
	if err := sh.Run("plan", "service", "uninstall"); err != nil {
		return err
	}
	if err := sh.Run("plan", "service", "install"); err != nil {
		return err
	}
	if err := sh.Run("plan", "service", "start"); err != nil {
		return err
	}
	fmt.Println(common.Styles.Keyword.Render("reinstalled."))
	return nil
}
