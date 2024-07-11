package gateway

import "github.com/matheus-santos-souza/go-ms-walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
