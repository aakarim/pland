package common

import (
	"github.com/aakarim/pland/cli/internal/config"
	"github.com/charmbracelet/charm/ui/common"
)

type Model struct {
	Config   config.Config
	Quitting bool
	Err      error
}

var (
	Styles    = common.DefaultStyles()
	Paragraph = Styles.Paragraph.Render
	Keyword   = Styles.Keyword.Render
	Code      = Styles.Code.Render
	Subtle    = Styles.Subtle.Render
)
