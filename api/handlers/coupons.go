package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	sugar "sugar/data"
	"sugar/helpers/response"
	"sugar/helpers/utils"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func (h *Handler) HandleGetDomainCoupons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		domain := r.PathValue("domain")

		codes, err := h.queries.GetCouponsByDomain(r.Context(), &domain)
		if err != nil {

			return
		}

		response.Success(w, "Successfully retrieved codes", codes)
	}
}

func (h *Handler) HandleCreateDomainCoupon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type CreateDomainCouponRequest struct {
			Domain string `json:"domain"`
			Code   string `json:"code"`
		}

		var request CreateDomainCouponRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.BadRequest(w, "Invalid request")
			return
		}

		if request.Domain == "" {
			response.BadRequest(w, "Missing required field: domain")
			return
		}

		if request.Code == "" {
			response.BadRequest(w, "Missing required field: code")
			return
		}

		if !utils.ValidDomain(request.Domain) {
			response.BadRequest(w, "Invalid field: domain")
			return
		}

		createCouponParams := sugar.CreateCouponParams{
			Code:   &request.Code,
			Domain: &request.Domain,
		}

		_, err := h.queries.CreateCoupon(r.Context(), createCouponParams)
		if err != nil {
			var sqliteErr sqlite3.Error
			if errors.As(err, &sqliteErr) {
				if sqliteErr.Code == sqlite3.ErrConstraint || sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
					response.Conflict(w, "Code already exists for this domain.")
					return
				}
			}

			response.InternalServerError(w, err, "Something went wrong")
			return
		}

		response.Success(w, "Successfully created code", nil)
	}
}
