package table

import (
	"fmt"
	"rhystmorgan/tx-snoop/internal/colours"

	"github.com/blockfrost/blockfrost-go"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func TableStyle(t *table.Model) {
	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(colours.Colours.Lavender))

	s.Selected = s.Selected.
		Foreground(lipgloss.Color(colours.Colours.Base)).
		Background(lipgloss.Color(colours.Colours.Lavender))

	t.SetStyles(s)
}

func MakeTable(transactions []blockfrost.AddressUTXO, height int) table.Model {
	columns := []table.Column{
		{Title: "No"},
		{Title: "TxHash"},
		{Title: "Ix"},
		{Title: "Assets"},
		{Title: "Block"},
	}

	qty := len(transactions)
	txRows := []table.Row{}

	for i := 0; i < qty; i++ {
		row := table.Row{
			fmt.Sprintf("%d", qty-i),
			transactions[i].TxHash,
			fmt.Sprintf("%d", transactions[i].OutputIndex),
			fmt.Sprintf("%d", len(transactions[i].Amount)),
			transactions[i].Block,
		}

		txRows = append(txRows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(txRows),
		table.WithFocused(true),
		table.WithHeight(height),
	)

	TableStyle(&t)

	return t
}
