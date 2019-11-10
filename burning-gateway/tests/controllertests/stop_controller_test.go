package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/andriikost/burning-gateway/api/models"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateStop(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}
	token, err := server.SignIn(user.Email, "password") //Note the password in the database is already hashed, we want unhashed
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	samples := []struct {
		inputJSON    string
		statusCode   int
		name         string
		content      string
		authorID     uint32
		tokenGiven   string
		errorMessage string
	}{
		{
			inputJSON:    `{"name":"The name", "content": "the content", "authorID": 1}`,
			statusCode:   201,
			tokenGiven:   tokenString,
			name:         "The name",
			content:      "the content",
			authorID:     user.ID,
			errorMessage: "",
		},
		{
			inputJSON:    `{"name":"The name", "content": "the content", "authorID": 1}`,
			statusCode:   500,
			tokenGiven:   tokenString,
			errorMessage: "The Name Already Taken",
		},
		{
			// When no token is passed
			inputJSON:    `{"name":"When no token is passed", "content": "the content", "authorID": 1}`,
			statusCode:   401,
			tokenGiven:   "",
			errorMessage: "Unauthorized",
		},
		{
			// When incorrect token is passed
			inputJSON:    `{"name":"When incorrect token is passed", "content": "the content", "authorID": 1}`,
			statusCode:   401,
			tokenGiven:   "This is an incorrect token",
			errorMessage: "Unauthorized",
		},
		{
			inputJSON:    `{"name": "", "content": "The content", "authorID": 1}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Name",
		},
		{
			inputJSON:    `{"name": "This is a name", "content": "", "authorID": 1}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Content",
		},
		{
			inputJSON:    `{"name": "This is an awesome name", "content": "the content"}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Author",
		},
		{
			// When user 2 uses user 1 token
			inputJSON:    `{"name": "This is an awesome name", "content": "the content", "authorID": 2}`,
			statusCode:   401,
			tokenGiven:   tokenString,
			errorMessage: "Unauthorized",
		},
	}
	for _, v := range samples {

		req, err := http.NewRequest("POST", "/stops", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateStop)

		req.Header.Set("Authorization", v.tokenGiven)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["name"], v.name)
			assert.Equal(t, responseMap["content"], v.content)
			assert.Equal(t, responseMap["authorID"], float64(v.authorID)) //just for both ids to have the same type
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetStops(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = seedUsersAndStops()
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/stops", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetStops)
	handler.ServeHTTP(rr, req)

	var Stops []models.Stop
	err = json.Unmarshal([]byte(rr.Body.String()), &Stops)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(Stops), 2)
}
func TestGetStopByID(t *testing.T) {

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatal(err)
	}
	stop, err := seedOneUserAndOneStop()
	if err != nil {
		log.Fatal(err)
	}
	postSample := []struct {
		id           string
		statusCode   int
		name         string
		content      string
		authorID     uint32
		errorMessage string
	}{
		{
			id:         strconv.Itoa(int(stop.ID)),
			statusCode: 200,
			name:       stop.Name,
			content:    stop.Content,
			authorID:   stop.AuthorID,
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
	}
	for _, v := range postSample {

		req, err := http.NewRequest("GET", "/stops", nil)
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetStop)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, stop.Name, responseMap["name"])
			assert.Equal(t, stop.Content, responseMap["content"])
			assert.Equal(t, float64(stop.AuthorID), responseMap["authorID"]) //the response author id is float64
		}
	}
}

func TestUpdateStop(t *testing.T) {

	var StopUserEmail, StopUserPassword string
	var AuthStopAuthorID uint32
	var AuthStopID uint64

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatal(err)
	}
	users, stops, err := seedUsersAndStops()
	if err != nil {
		log.Fatal(err)
	}
	// Get only the first user
	for _, user := range users {
		if user.ID == 2 {
			continue
		}
		StopUserEmail = user.Email
		StopUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	//Login the user and get the authentication token
	token, err := server.SignIn(StopUserEmail, StopUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	// Get only the first Stop
	for _, stop := range stops {
		if stop.ID == 2 {
			continue
		}
		AuthStopID = stop.ID
		AuthStopAuthorID = stop.AuthorID
	}
	// fmt.Printf("this is the auth Stop: %v\n", AuthStopID)

	samples := []struct {
		id           string
		updateJSON   string
		statusCode   int
		name         string
		content      string
		authorID     uint32
		tokenGiven   string
		errorMessage string
	}{
		{
			// Convert int64 to int first before converting to string
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "The updated Stop", "content": "This is the updated content", "authorID": 1}`,
			statusCode:   200,
			name:         "The updated Stop",
			content:      "This is the updated content",
			authorID:     AuthStopAuthorID,
			tokenGiven:   tokenString,
			errorMessage: "",
		},
		{
			// When no token is provided
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "This is still another name", "content": "This is the updated content", "authorID": 1}`,
			tokenGiven:   "",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
		{
			// When incorrect token is provided
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "This is still another name", "content": "This is the updated content", "authorID": 1}`,
			tokenGiven:   "this is an incorrect token",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},

		// {
		// 	//Note: "name 2" belongs to Stop 2, and name must be unique
		// 	id:           strconv.Itoa(int(AuthStopID)),
		// 	updateJSON:   `{"name": "Name 2", "content": "This is the updated content", "authorID": 1}`,
		// 	statusCode:   200,
		// 	tokenGiven:   tokenString,
		// 	errorMessage: "The Name Already Taken",
		// },

		{
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "", "content": "This is the updated content", "authorID": 1}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Name",
		},
		{
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "Awesome name", "content": "", "authorID": 1}`,
			statusCode:   422,
			tokenGiven:   tokenString,
			errorMessage: "Required Content",
		},
		{
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "This is another name", "content": "This is the updated content"}`,
			statusCode:   401,
			tokenGiven:   tokenString,
			errorMessage: "Unauthorized",
		},
		{
			id:         "unknwon",
			statusCode: 400,
		},
		{
			id:           strconv.Itoa(int(AuthStopID)),
			updateJSON:   `{"name": "This is still another name", "content": "This is the updated content", "authorID": 2}`,
			tokenGiven:   tokenString,
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("Stop", "/Stops", bytes.NewBufferString(v.updateJSON))
		if err != nil {
			t.Errorf("this is the error: %v\n", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.UpdateStop)

		req.Header.Set("Authorization", v.tokenGiven)

		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.Equal(t, responseMap["name"], v.name)
			assert.Equal(t, responseMap["content"], v.content)
			assert.Equal(t, responseMap["authorID"], float64(v.authorID)) //just to match the type of the json we receive thats why we used float64
		}
		if v.statusCode == 401 || v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestDeleteStop(t *testing.T) {

	var StopUserEmail, StopUserPassword string
	var StopUserID uint32
	var AuthStopID uint64

	err := refreshUserAndStopTable()
	if err != nil {
		log.Fatal(err)
	}
	users, Stops, err := seedUsersAndStops()
	if err != nil {
		log.Fatal(err)
	}
	//Let's get only the Second user
	for _, user := range users {
		if user.ID == 1 {
			continue
		}
		StopUserEmail = user.Email
		StopUserPassword = "password" //Note the password in the database is already hashed, we want unhashed
	}
	//Login the user and get the authentication token
	token, err := server.SignIn(StopUserEmail, StopUserPassword)
	if err != nil {
		log.Fatalf("cannot login: %v\n", err)
	}
	tokenString := fmt.Sprintf("Bearer %v", token)

	// Get only the second Stop
	for _, stop := range Stops {
		if stop.ID == 1 {
			continue
		}
		AuthStopID = stop.ID
		StopUserID = stop.AuthorID
	}
	postSample := []struct {
		id           string
		authorID     uint32
		tokenGiven   string
		statusCode   int
		errorMessage string
	}{
		{
			// Convert int64 to int first before converting to string
			id:           strconv.Itoa(int(AuthStopID)),
			authorID:     StopUserID,
			tokenGiven:   tokenString,
			statusCode:   204,
			errorMessage: "",
		},
		{
			// When empty token is passed
			id:           strconv.Itoa(int(AuthStopID)),
			authorID:     StopUserID,
			tokenGiven:   "",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
		{
			// When incorrect token is passed
			id:           strconv.Itoa(int(AuthStopID)),
			authorID:     StopUserID,
			tokenGiven:   "This is an incorrect token",
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
		{
			id:         "unknwon",
			tokenGiven: tokenString,
			statusCode: 400,
		},
		{
			id:           strconv.Itoa(int(1)),
			authorID:     1,
			statusCode:   401,
			errorMessage: "Unauthorized",
		},
	}
	for _, v := range postSample {

		req, _ := http.NewRequest("GET", "/Stops", nil)
		req = mux.SetURLVars(req, map[string]string{"id": v.id})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.DeleteStop)

		req.Header.Set("Authorization", v.tokenGiven)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 401 && v.errorMessage != "" {

			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}
