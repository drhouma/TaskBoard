package route

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"html_comments_system/internal/postgres"
	"net/http"
)

const (
	userQueryParam     = "user"
	usernameQueryParam = "user_name"
	messageQueryParam  = "message"
)

var db *postgres.DataBase

func Init(current *postgres.DataBase) {
	db = current
}

func New() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}))

	mux.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", addUser)

			r.Get("/exist", isUserExist)
		})
		r.Route("/comments", func(r chi.Router) {
			r.Get("/", getComments)
			r.Post("/", addComment)
			r.Delete("/", deleteComment)

			r.Get("/exist", isCommentExist)
		})
	})

	return mux
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := postgres.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := db.AddUser(user); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func isUserExist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	if !query.Has(userQueryParam) {
		WriteError(w, errors.New(`missing query parameter "user"`), http.StatusBadRequest)
		return
	}

	exist, err := db.IsUserExist(postgres.User{
		Name: query.Get(userQueryParam),
	})
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	resp := struct {
		Exist bool `json:"exist"`
	}{
		Exist: exist,
	}

	WriteResponseJson(w, resp)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments, err := db.GetComments()
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, comments)
}

func addComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comment := postgres.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := db.AddComment(comment); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comment := postgres.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := db.DeleteComment(comment); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func isCommentExist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	if !query.Has(usernameQueryParam) {
		WriteError(w, errors.New(`missing query parameter "user_name"`), http.StatusBadRequest)
		return
	}
	if !query.Has(messageQueryParam) {
		WriteError(w, errors.New(`missing query parameter "message"`), http.StatusBadRequest)
		return
	}

	username := query.Get(usernameQueryParam)
	message := query.Get(messageQueryParam)
	exist, err := db.IsCommentExist(postgres.Comment{
		UserName: &username,
		Message:  &message,
	})
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	resp := struct {
		Exist bool `json:"exist"`
	}{
		Exist: exist,
	}

	WriteResponseJson(w, resp)
}
