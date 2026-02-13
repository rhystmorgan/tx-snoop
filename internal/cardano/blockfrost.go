package cardano

import (
	"context"

	"github.com/blockfrost/blockfrost-go"
)

func InitBlockfrost(key string) blockfrost.APIClient {
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
			ProjectID: key,
			Server:    blockfrost.CardanoPreview,
		},
	)

	return api
}

func GetTransactions(key string) []blockfrost.AddressUTXO {
	address := "addr_test1vqlhvhcwaddssxnkfugwlvmk69925xjdx7nc20j2nzuc0gq43pzgq"

	txCh := InitBlockfrost(key).AddressUTXOsAll(context.TODO(), address)

	txList := []blockfrost.AddressUTXO{}

	for result := range txCh {
		if result.Err != nil {
			return txList
		}
		for _, tx := range result.Res {
			txList = append(txList, tx)
		}
	}

	return txList
}
