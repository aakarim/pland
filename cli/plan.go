package main

import "github.com/aakarim/pland/cli/cmd"

//go:generate go run -mod=mod github.com/Khan/genqlient
func main() {
	cmd.Execute()
}
