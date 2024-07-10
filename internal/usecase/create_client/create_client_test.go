package createclient

import (
	"testing"

	"github.com/matheus-santos-souza/go-ms-walletcore/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreteClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)
	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "Matheus",
		Email: "m@m.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "Matheus", output.Name)
	assert.Equal(t, "m@m.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
