package internal

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)
import "github.com/charmbracelet/lipgloss/table"

var (
	HeaderStyle = lipgloss.Style{}.
			Background(lipgloss.ANSIColor(5)).
			Foreground(lipgloss.ANSIColor(99))
	NormalStyle = lipgloss.Style{}.
			Background(lipgloss.NoColor{}).
			Foreground(lipgloss.ANSIColor(10))
)

func Info(config RootConfig) string {
	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Width(100).
		StyleFunc(func(row, col int) lipgloss.Style {
			width := 0
			switch {
			case col == 0, col == 1:
				width = 10
			default:
				width = 20
			}

			switch {
			case row == 0:
				return HeaderStyle.Width(width)
			default:
				return NormalStyle.Width(width)
			}
		}).
		Headers("name", "value", "source")

	for _, c := range config.values() {
		t.Row(c.name, c.value, c.source)
	}

	return fmt.Sprintln(t)
}
