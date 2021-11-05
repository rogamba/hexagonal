package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/rogamba/hexagonal/domain/services"
)

type Handler interface {
	GetUser(http.ResponseWriter, *http.Request)
}

type handler struct {
	userService services.UserService
}

func NewHandler(userService services.UserService) Handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	userId := chi.URLParam(r, "id")
	user, _ := h.userService.FetchUserDetails(userId)
	responseBody, err := jsonify(user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusFound)
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		fmt.Println(err)
	}
}

func jsonify(object interface{}) ([]byte, error) {
	rawMsg, err := json.Marshal(object)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
