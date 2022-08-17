package sync

import (
	"errors"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type PlanDay string

type DayPlan struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	TS   int64     `json:"timestamp"`
	Hash uint64    `json:"hash"`
	Txt  string    `json:"txt"`
	Date time.Time `json:"date"`
}

func printFormatted(s ...string) {
	for _, ss := range s {
		fmt.Printf(paragraph(ss))
	}
	fmt.Println()
}

func (m model) sync() tea.Msg {
	homePlanPath, err := m.planService.GetLocalPlanPath()
	if err != nil {
		return err
	}
	_, err = m.planService.Sync()
	// create file and exit
	if errors.Is(err, os.ErrNotExist) {
		return freshPlan{path: homePlanPath}
	}
	if err != nil {
		return err
	}
	return syncCompleted{}
}
