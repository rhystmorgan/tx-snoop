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
	height  int
	width   int
	header  string
	panels  []string
	focused int
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
		height:  24,
		width:   80,
		header:  bannerRenderer(),
		panels:  []string{"ADDRESSES", "SUMMARY", "TRANSACTIONS"},
		focused: 0,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height - 2
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "tab":
			if len(m.panels) == m.focused+1 {
				m.focused = 0

				return m, nil
			} else {
				m.focused += 1

				return m, nil
			}
		}
	}

	return m, nil
}

func bannerStyle(m model) string {
	decorationStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(6).
		AlignHorizontal(lipgloss.Center).
		Foreground(lipgloss.Color(colour.Colours.Peach)).
		SetString(m.header)

	return decorationStyle.Render()
}

func calculateHorizontalSplit(height int) (addressHeight, summaryHeight, transactionHeight int) {
	bannerHeight := 8
	footerHeight := 3
	panelQty := 2
	panelBorderQty := 2 * panelQty

	workingSpace := height - (bannerHeight + panelBorderQty + footerHeight)

	addressHeight = height - (bannerHeight + 2 + footerHeight)
	summaryHeight = workingSpace / 3
	mainHeight := summaryHeight * 2

	return addressHeight, summaryHeight, mainHeight + (workingSpace % 3)

	// switch workingSpace {
	// case workingSpace % 3 == 0:
	// 	return (addressHeight, summaryHeight, transactionHeight)

	// case workingSpace % 3 == 1:
	// 	return (addressHeight, summaryHeight, transactionHeight + 1)

	// case workingHeight % 3 == 2:
	// 	return (addressHeight, summaryHeight, transactransactionHeight + 2)
	// }
}

func calculateVerticalSplit(width int) (addressWidth, rightWidth int) {
	qtyPanels := 2
	qtyPanelBorders := qtyPanels * 2

	workingSpace := width - qtyPanelBorders
	addressWidth = workingSpace / 4
	mainWidth := addressWidth * 3

	return addressWidth, mainWidth + (workingSpace % 4)
}

func makePanel(header string, index, height, width int) lipgloss.Style {
	panel := lipgloss.NewStyle().
		Height(height).
		Width(width).
		Border(lipgloss.NormalBorder()).
		BorderDecoration(
			lipgloss.NewBorderDecoration(
				lipgloss.BorderTop,
				lipgloss.Left,
				lipgloss.NewStyle().Foreground(lipgloss.Color(colour.Colours.Base)).Background(lipgloss.Color(colour.Colours.Peach)).SetString(fmt.Sprintf(" %d %s ", index, header)).String(),
			),
		).
		BorderForeground(lipgloss.Color(colour.Colours.Peach))

	return panel
}

func panelStyle(panels []string, height, width int) []lipgloss.Style {
	aHeight, sHeight, tHeight := calculateHorizontalSplit(height)
	aWidth, rWidth := calculateVerticalSplit(width)

	addressPanel := makePanel(panels[0], 0, aHeight, aWidth)
	summaryPanel := makePanel(panels[1], 1, sHeight, rWidth)
	transactionPanel := makePanel(panels[2], 2, tHeight, rWidth)

	return []lipgloss.Style{addressPanel, summaryPanel, transactionPanel}
}

func (m model) View() string {
	s := bannerStyle(m)

	p := panelStyle(m.panels, m.height, m.width)

	right := lipgloss.JoinVertical(lipgloss.Left, p[2].Render(), p[1].Render())

	body := lipgloss.JoinHorizontal(lipgloss.Top, p[0].Render(), right)

	return lipgloss.JoinVertical(lipgloss.Left, s, body)
}

func main() {
	p := tea.NewProgram(initModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
