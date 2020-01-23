package controllers

import middlewares "github.com/andriikost/burning-gateway/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/auth/loggedInUser", middlewares.SetMiddlewareJSON(s.GetLoggedInUser)).Methods("GET")

	// Users routes
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Stops routes
	s.Router.HandleFunc("/stop", middlewares.SetMiddlewareJSON(s.CreateStop)).Methods("POST")
	s.Router.HandleFunc("/stop", middlewares.SetMiddlewareJSON(s.GetStops)).Methods("GET")
	s.Router.HandleFunc("/stop/{id}", middlewares.SetMiddlewareJSON(s.GetStop)).Methods("GET")
	s.Router.HandleFunc("/stop/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateStop))).Methods("PUT")
	s.Router.HandleFunc("/stop/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteStop)).Methods("DELETE")

	// Route routes
	s.Router.HandleFunc("/route", middlewares.SetMiddlewareJSON(s.CreateRoute)).Methods("POST")
	s.Router.HandleFunc("/route", middlewares.SetMiddlewareJSON(s.GetRoutes)).Methods("GET")
	s.Router.HandleFunc("/route/{id}", middlewares.SetMiddlewareJSON(s.GetRoute)).Methods("GET")
	s.Router.HandleFunc("/route/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateRoute))).Methods("PUT")
	s.Router.HandleFunc("/route/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteRoute)).Methods("DELETE")

	// Story story
	s.Router.HandleFunc("/story", middlewares.SetMiddlewareJSON(s.CreateStory)).Methods("POST")
	s.Router.HandleFunc("/story", middlewares.SetMiddlewareJSON(s.GetStories)).Methods("GET")
	s.Router.HandleFunc("/story/{id}", middlewares.SetMiddlewareJSON(s.GetStory)).Methods("GET")
	s.Router.HandleFunc("/story/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateStory))).Methods("PUT")
	s.Router.HandleFunc("/story/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteStory)).Methods("DELETE")

	// Votes
	s.Router.HandleFunc("/votes/stop-votes", middlewares.SetMiddlewareAuthentication(s.UpdateStopVote)).Methods("POST")
	s.Router.HandleFunc("/votes/route-votes", middlewares.SetMiddlewareAuthentication(s.UpdateRouteVote)).Methods("POST")
	s.Router.HandleFunc("/votes/story-votes", middlewares.SetMiddlewareAuthentication(s.UpdateStoryVote)).Methods("POST")

	// File upload
	s.Router.HandleFunc("/file-upload/presign/{object-name}", middlewares.SetMiddlewareAuthentication((s.GetPresignedUploadUrl))).Methods("GET")
	s.Router.HandleFunc("/file-upload/upload-image/{object-name}", middlewares.SetMiddlewareJSON((s.UploadFile))).Methods("POST")
}
