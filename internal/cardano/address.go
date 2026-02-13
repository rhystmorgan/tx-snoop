package cardano

import (
	"context"

	"github.com/blockfrost/blockfrost-go"
)

func AddressTransactions(address, key string, ctx context.Context) ([]blockfrost.AddressTransactions, error) {
	b := InitBlockfrost(key)
	params := blockfrost.APIQueryParams{
		Count: 1,
		Order: "desc",
	}
	transactions, err := b.AddressTransactions(ctx, address, params)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
