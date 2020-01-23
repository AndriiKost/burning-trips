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

func (server *Server) CreateStory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	story := models.Story{}
	err = json.Unmarshal(body, &story)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	story.Prepare()
	err = story.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != story.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	storyCreated, err := story.Create(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, storyCreated.ID))
	responses.JSON(w, http.StatusCreated, storyCreated)
}

func (server *Server) GetStories(w http.ResponseWriter, r *http.Request) {

	story := models.Story{}

	stories, err := story.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, stories)
}

func (server *Server) GetStory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storyId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	story := models.Story{}

	storyReceived, err := story.Find(server.DB, storyId)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, storyReceived)
}

func (server *Server) UpdateStory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	storyId, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the story exist
	story := models.Story{}
	err = server.DB.Debug().Model(models.Story{}).Where("id = ?", storyId).Take(&story).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("story not found"))
		return
	}

	// If a user attempt to update a story not belonging to him
	if uid != story.AuthorID {
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
	storyUpdate := models.Story{}
	err = json.Unmarshal(body, &storyUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != storyUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	storyUpdate.Prepare()
	err = storyUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	storyUpdate.ID = story.ID //this is important to tell the model the story id to update, the other update field are set above

	updatedStory, err := storyUpdate.Update(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedStory)
}

func (server *Server) DeleteStory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid story id given to us?
	storyId, err := strconv.ParseUint(vars["id"], 10, 64)
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

	// Check if the story exist
	story := models.Story{}
	err = server.DB.Debug().Model(models.Story{}).Where("id = ?", storyId).Take(&story).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this story?
	if uid != story.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = story.Delete(server.DB, storyId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", storyId))
	responses.JSON(w, http.StatusNoContent, "")
}
