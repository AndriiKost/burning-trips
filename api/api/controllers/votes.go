package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andriikost/burning-gateway/api/auth"
	"github.com/andriikost/burning-gateway/api/models"
	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/andriikost/burning-gateway/api/utils/formaterror"
)

func (server *Server) UpdateStopVote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during reading request body")
		return
	}
	stopVote := models.StopVote{}
	err = json.Unmarshal(body, &stopVote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during stop vote extraction from request body")
		return
	}
	err = stopVote.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during stop validation")
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != stopVote.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	stopVoteCreated, err := stopVote.SaveStopVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, stopVote.ID))
	responses.JSON(w, http.StatusCreated, stopVoteCreated)
}

func (server *Server) UpdateRouteVote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during reading request body")
		return
	}
	routeVote := models.RouteVote{}
	err = json.Unmarshal(body, &routeVote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during route vote extraction from request body")
		return
	}
	err = routeVote.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during stop validation")
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != routeVote.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	routeVoteCreated, err := routeVote.SaveRouteVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, routeVote.ID))
	responses.JSON(w, http.StatusCreated, routeVoteCreated)
}

func (server *Server) UpdateStoryVote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during reading request body")
		return
	}
	stopVote := models.StopVote{}
	err = json.Unmarshal(body, &stopVote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during route vote extraction from request body")
		return
	}
	err = stopVote.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Println("Error during stop validation")
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != stopVote.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	stopVoteCreated, err := stopVote.SaveStopVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, stopVote.ID))
	responses.JSON(w, http.StatusCreated, stopVoteCreated)
}
