package createaccount

import (
	"testing"

	"github.com/matheus-santos-souza/go-ms-walletcore/internal/entity"
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Matheus", "m@m.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
