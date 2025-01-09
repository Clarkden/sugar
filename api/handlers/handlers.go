package handlers

import (
	sugar "sugar/data"
)

type Handler struct {
	queries *sugar.Queries
}

func NewHandler(queries *sugar.Queries) *Handler {
	return &Handler{queries: queries}
}
