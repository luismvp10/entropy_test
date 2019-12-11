package controllers

import "github.com/luismvp10/entropy_test/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	//s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	//s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	///Contacts routes
	s.Router.HandleFunc("/contacts", middlewares.SetMiddlewareAuthentication(s.CreateContact)).Methods("POST")
	s.Router.HandleFunc("/contacts/{id}", middlewares.SetMiddlewareAuthentication(s.GetContacts)).Methods("GET")
	s.Router.HandleFunc("/contact/{id}", middlewares.SetMiddlewareAuthentication(s.GetContact)).Methods("GET")
}
