package createclient

import (
	"github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreteClientUseCase_Execute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
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
