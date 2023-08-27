package main

import (
	_ "embed"
)

//go:embed src/user.tmpl.html
var userTmpl string
