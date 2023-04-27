package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/leojasmim/kafka-producer-go/domain/models"
	"github.com/leojasmim/kafka-producer-go/service"
)

type MessageRequest struct {
	Message string `json:"message"`
}

type MessageResponse struct {
	Message *models.Message `json:"message,omitempty"`
	Error   string          `json:"error,omitempty"`
}

func (rp MessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg_request MessageRequest

	err := json.NewDecoder(r.Body).Decode(&msg_request)

	var response MessageResponse

	if err != nil {
		response.Error = err.Error()
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, response)
		return
	}

	response.Message = models.NewMessage(msg_request.Message)

	err = service.NewMessageService().SendMessage(*response.Message)

	if err != nil {
		response.Error = err.Error()
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, response)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, response)
}
