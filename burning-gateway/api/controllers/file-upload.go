package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/andriikost/burning-gateway/api/auth"
	"github.com/andriikost/burning-gateway/api/responses"
)

type FileUploadRequest struct {
	fileName   string
	filePrefix string
}

func (fileUpload *FileUploadRequest) Prepare() {
	fileUpload.fileName = html.EscapeString(strings.TrimSpace(fileUpload.fileName))
	fileUpload.filePrefix = html.EscapeString(strings.TrimSpace(fileUpload.filePrefix))
}

func (fileUpload *FileUploadRequest) Validate() error {

	if fileUpload.fileName == "" {
		return errors.New("Required fileName")
	}
	if fileUpload.filePrefix == "" {
		return errors.New("Required filePrefix")
	}
	return nil
}

func (server *Server) FileUpload(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fileUpload := FileUploadRequest{}
	err = json.Unmarshal(body, &fileUpload)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fileUpload.Prepare()
	err = fileUpload.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	uploadUrl := "https://aws-sampleurl.com"

	responses.JSON(w, http.StatusCreated, uploadUrl)
}
