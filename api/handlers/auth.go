package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	sugar "sugar/data"
	"sugar/globals/auth"
	"sugar/helpers/response"
	"sugar/helpers/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

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

		if !utils.ValidEmail(request.Email) {
			response.BadRequest(w, "Invalid field: email")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), auth.PasswordCost)
		if err != nil {
			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		userParams := sugar.CreateUserParams{
			Email:    request.Email,
			Password: string(hashedPassword),
		}

		user, err := h.queries.CreateUser(r.Context(), userParams)
		if err != nil {
			log.Print(err)
			return
		}

		sessionId := uuid.New()

		sessionParams := sugar.CreateSessionParams{
			UserID:    sql.NullInt64{Int64: user.ID, Valid: true},
			SessionID: sql.NullString{String: sessionId.String(), Valid: true},
			CreatedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
			ExpiresAt: sql.NullInt64{Int64: time.Now().Add(24 * 30 * time.Hour).Unix(), Valid: true},
		}

		session, err := h.queries.CreateSession(r.Context(), sessionParams)
		if err != nil {
			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		response.Success(w, "Successfully registered user.", session.SessionID)
	}
}

func (h *Handler) HandleEmailLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type EmailLoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var request EmailLoginRequest

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

		if !utils.ValidEmail(request.Email) {
			response.BadRequest(w, "Invalid field: email")
			return
		}

		user, err := h.queries.GetUserByEmail(r.Context(), request.Email)
		if err != nil {
			response.Unauthorized(w, "Invalid email or password")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			response.Unauthorized(w, "Invalid email or password")
			return
		}

		sessionId := uuid.New()

		sessionParams := sugar.CreateSessionParams{
			UserID:    sql.NullInt64{Int64: user.ID, Valid: true},
			SessionID: sql.NullString{String: sessionId.String(), Valid: true},
			CreatedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
			ExpiresAt: sql.NullInt64{Int64: time.Now().Add(24 * 30 * time.Hour).Unix(), Valid: true},
		}

		session, err := h.queries.CreateSession(r.Context(), sessionParams)
		if err != nil {
			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		response.Success(w, "Successfully logged in user.", session.SessionID)
	}
}
