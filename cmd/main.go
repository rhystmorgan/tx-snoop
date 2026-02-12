package main

import (
	"fmt"
	"os"

	banner "rhystmorgan/tx-snoop/internal/banner"
	colour "rhystmorgan/tx-snoop/internal/colours"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	height int
	width  int
	header string
}

func bannerRenderer() string {
	makeBanner := ""

	for i := 0; i < len(banner.Banner); i++ {
		makeBanner += banner.Banner[i]
	}

	return makeBanner
}

func initModel() model {
	m := model{
		height: 24,
		width:  80,
		header: bannerRenderer(),
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func headerStyle(m model) string {
	decorationStyle := lipgloss.NewStyle().
		Width(m.width - 2).
		Height(m.height - 2).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color(colour.Colours.Peach)).
		Border(lipgloss.HiddenBorder()).
		SetString(m.header)

	return decorationStyle.Render()
}

func (m model) View() string {
	s := headerStyle(m)

	return s
}

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
