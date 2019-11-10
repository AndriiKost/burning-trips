package controllers

import (
	"net/http"

	"github.com/andriikost/burning-gateway/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Burning-Trips API")

}