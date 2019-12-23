package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/andriikost/burning-gateway/api/auth"
	"github.com/andriikost/burning-gateway/api/models"
	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/andriikost/burning-gateway/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateStop(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	stop := models.Stop{}
	err = json.Unmarshal(body, &stop)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	stop.Prepare()
	err = stop.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != stop.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	stopCreated, err := stop.Save(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, stopCreated.ID))
	responses.JSON(w, http.StatusCreated, stopCreated)
}

func (server *Server) GetStops(w http.ResponseWriter, r *http.Request) {

	stop := models.Stop{}

	stops, err := stop.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, stops)
}

func (server *Server) GetStop(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	stop := models.Stop{}

	stopReceived, err := stop.Find(server.DB, sid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, stopReceived)
}

func (server *Server) UpdateStop(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the stop id is valid
	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the stop exist
	stop := models.Stop{}
	err = server.DB.Debug().Model(models.Stop{}).Where("id = ?", sid).Take(&stop).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Stop not found"))
		return
	}

	// If a user attempt to update a stop not belonging to him
	if uid != stop.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data posted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	stopUpdate := models.Stop{}
	err = json.Unmarshal(body, &stopUpdate)
	// err = json.NewDecoder(r.Body).Decode(&stopUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != stopUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	stopUpdate.Prepare()
	err = stopUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	stopUpdate.ID = stop.ID //this is important to tell the model the stop id to update, the other update field are set above

	stopUpdated, err := stopUpdate.Update(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, stopUpdated)
}

func (server *Server) DeleteStop(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid stop id given to us?
	sid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the stop exist
	stop := models.Stop{}
	err = server.DB.Debug().Model(models.Stop{}).Where("id = ?", sid).Take(&stop).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this stop?
	if uid != stop.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = stop.Delete(server.DB, sid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", sid))
	responses.JSON(w, http.StatusNoContent, "")
}
