package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"task_board/internal/postgres"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const (
	descriptionQueryParam = "description"
	userQueryParam        = "user"
	categoryQueryParam    = "category"
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
			r.Get("/", getUser)
		})
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", getTasks)
			r.Post("/", addTask)
			r.Delete("/", deleteTask)
			r.Patch("/", updateTask)
		})
	})

	return mux
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := db.GetTasks()
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, tasks)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	if !query.Has(userQueryParam) {
		WriteError(w, errors.New(`missing query parameter "user"`), http.StatusBadRequest)
		return
	}

	user, err := db.GetUser(query.Get(userQueryParam))
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, user)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	task := postgres.Task{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := db.AddTask(task); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	if !query.Has(userQueryParam) {
		WriteError(w, errors.New(`missing query parameter "user"`), http.StatusBadRequest)
		return
	}
	if !query.Has(descriptionQueryParam) {
		WriteError(w, errors.New(`missing query parameter "description"`), http.StatusBadRequest)
		return
	}

	user := query.Get(userQueryParam)
	description := query.Get(descriptionQueryParam)

	if err := db.DeleteTask(user, description); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	if !query.Has(userQueryParam) {
		WriteError(w, errors.New(`missing query parameter "user"`), http.StatusBadRequest)
		return
	}
	if !query.Has(descriptionQueryParam) {
		WriteError(w, errors.New(`missing query parameter "description"`), http.StatusBadRequest)
		return
	}
	if !query.Has(categoryQueryParam) {
		WriteError(w, errors.New(`missing query parameter "category"`), http.StatusBadRequest)
		return
	}

	user := query.Get(userQueryParam)
	description := query.Get(descriptionQueryParam)
	category := query.Get(categoryQueryParam)

	if err := db.UpdateTask(user, description, category); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
