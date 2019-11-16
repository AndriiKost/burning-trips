package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/andriikost/burning-gateway/api/auth"
	"github.com/andriikost/burning-gateway/api/models"
	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/andriikost/burning-gateway/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err, user := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    os.Getenv("JWT_TOKEN"),
		Value:   token,
		Expires: expirationTime,
		Domain:  "localhost",
	})

	responses.JSON(w, http.StatusOK, user)
}

func (server *Server) SignIn(email, password string) (string, error, models.User) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err, user
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err, user
	}
	token, err := auth.CreateToken(user.ID)
	return token, err, user
}
