package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andriikost/burning-gateway/api/models"
	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetLandmarks(w http.ResponseWriter, r *http.Request) {

	landmark := models.Landmark{}

	landmarks, err := landmark.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, landmarks)
}

func (server *Server) GetLandmark(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	landmark := models.Landmark{}

	landmarkReceived, err := landmark.Find(server.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, landmarkReceived)
}

func (server *Server) SearchLandmarks(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var query models.SearchQuery

	err := decoder.Decode(&query)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	landmark := models.Landmark{}

	landmarks, err := landmark.SearchLandmarks(server.DB, query)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, landmarks)
}
