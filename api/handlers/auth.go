package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	sugar "sugar/data"
	"sugar/helpers/response"
)

func (h *Handler) HandleEmailLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) HandleEmailRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type EmailRegisterRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var request EmailRegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.BadRequest(w, "Invalid request")
			return
		}

		if request.Email == "" {
			response.BadRequest(w, "Missing required field: email")
			return
		}

		if request.Password == "" {
			response.BadRequest(w, "Missing required field: password")
			return
		}

		userParams := sugar.CreateUserParams{
			Email:    request.Email,
			Password: request.Password,
		}

		user, err := h.queries.CreateUser(r.Context(), userParams)
		if err != nil {
			log.Print(err)
			return
		}

		sessionParams := sugar.CreateSessionParams{
			Userid: &user.ID,
			Createdat: ,
		}

		session, err := h.queries.CreateSession(r.Context())

		response.Success(w, "Successfully registered user.", nil)
	}
}
