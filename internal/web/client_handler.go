package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	createclient "github.com/matheus-santos-souza/go-ms-walletcore/internal/usecase/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create client")
	var dto createclient.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println("Bad request: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		fmt.Println("Internal Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		fmt.Println("Internal Error 2: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
