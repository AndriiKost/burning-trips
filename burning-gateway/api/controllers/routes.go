package controllers

import middlewares "github.com/andriikost/burning-gateway/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/auth/loggedInUser", middlewares.SetMiddlewareJSON(s.GetLoggedInUser)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Stops routes
	s.Router.HandleFunc("/stops", middlewares.SetMiddlewareJSON(s.CreateStop)).Methods("POST")
	s.Router.HandleFunc("/stops", middlewares.SetMiddlewareJSON(s.GetStops)).Methods("GET")
	s.Router.HandleFunc("/stops/{id}", middlewares.SetMiddlewareJSON(s.GetStop)).Methods("GET")
	s.Router.HandleFunc("/stops/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateStop))).Methods("PUT")
	s.Router.HandleFunc("/stops/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteStop)).Methods("DELETE")
}
