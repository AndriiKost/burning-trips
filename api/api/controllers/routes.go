package controllers

import middleware "github.com/andriikost/burning-gateway/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/auth/login", middleware.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/auth/loggedInUser", middleware.SetMiddlewareJSON(s.GetLoggedInUser)).Methods("GET")

	// Users routes
	s.Router.HandleFunc("/user", middleware.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user", middleware.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middleware.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", middleware.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Stops routes
	s.Router.HandleFunc("/stop", middleware.SetMiddlewareJSON(s.CreateStop)).Methods("POST")
	s.Router.HandleFunc("/search-stops", middleware.SetMiddlewareJSON(s.SearchStops)).Methods("POST")
	s.Router.HandleFunc("/stop", middleware.SetMiddlewareJSON(s.GetStops)).Methods("GET")
	s.Router.HandleFunc("/stop/{id}", middleware.SetMiddlewareJSON(s.GetStop)).Methods("GET")
	s.Router.HandleFunc("/stop/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateStop))).Methods("PUT")
	s.Router.HandleFunc("/stop/{id}", middleware.SetMiddlewareAuthentication(s.DeleteStop)).Methods("DELETE")

	// Landmarks routes
	s.Router.HandleFunc("/search-landmarks", middleware.SetMiddlewareJSON(s.SearchLandmarks)).Methods("POST")
	s.Router.HandleFunc("/landmark", middleware.SetMiddlewareJSON(s.GetLandmarks)).Methods("GET")
	s.Router.HandleFunc("/landmark/{id}", middleware.SetMiddlewareJSON(s.GetLandmark)).Methods("GET")


	// Route routes
	s.Router.HandleFunc("/route", middleware.SetMiddlewareJSON(s.CreateRoute)).Methods("POST")
	s.Router.HandleFunc("/route", middleware.SetMiddlewareJSON(s.GetRoutes)).Methods("GET")
	s.Router.HandleFunc("/route/{id}", middleware.SetMiddlewareJSON(s.GetRoute)).Methods("GET")
	s.Router.HandleFunc("/route/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateRoute))).Methods("PUT")
	s.Router.HandleFunc("/route/{id}", middleware.SetMiddlewareAuthentication(s.DeleteRoute)).Methods("DELETE")

	// Story story
	s.Router.HandleFunc("/story", middleware.SetMiddlewareJSON(s.CreateStory)).Methods("POST")
	s.Router.HandleFunc("/story", middleware.SetMiddlewareJSON(s.GetStories)).Methods("GET")
	s.Router.HandleFunc("/story/{id}", middleware.SetMiddlewareJSON(s.GetStory)).Methods("GET")
	s.Router.HandleFunc("/story/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateStory))).Methods("PUT")
	s.Router.HandleFunc("/story/{id}", middleware.SetMiddlewareAuthentication(s.DeleteStory)).Methods("DELETE")

	// Votes
	s.Router.HandleFunc("/votes/stop-votes", middleware.SetMiddlewareAuthentication(s.UpdateStopVote)).Methods("POST")
	s.Router.HandleFunc("/votes/route-votes", middleware.SetMiddlewareAuthentication(s.UpdateRouteVote)).Methods("POST")
	s.Router.HandleFunc("/votes/story-votes", middleware.SetMiddlewareAuthentication(s.UpdateStoryVote)).Methods("POST")

	// File upload
	s.Router.HandleFunc("/file-upload/presign/{object-name}", middleware.SetMiddlewareAuthentication((s.GetPresignedUploadUrl))).Methods("GET")
	s.Router.HandleFunc("/file-upload/upload-image/{object-name}", middleware.SetMiddlewareJSON((s.UploadFile))).Methods("POST")
}
