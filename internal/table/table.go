package table

import (
	"fmt"
	"github.com/blockfrost/blockfrost-go"
	"github.com/charmbracelet/bubbles/table"
)

type TableModel struct {
	table table.Model
}

func MakeTable(transactions []blockfrost.AddressUTXO, height int) table.Model {
	columns := []table.Column{
		{Title: "No", Width: 4},
		{Title: "TxHash", Width: 20},
		{Title: "Ix", Width: 4},
		{Title: "Assets", Width: 6},
		{Title: "Block", Width: 10},
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

	return t
}
