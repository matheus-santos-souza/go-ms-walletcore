package gateway

import "github.com/matheus-santos-souza/go-ms-walletcore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
