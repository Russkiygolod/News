package api

import (
	"News/pkg/storage/postgres"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r  *mux.Router    // маршрутизатор запросов
	db postgres.Store // база данных
}

// Конструктор API.
func New(db postgres.Store) *API {
	api := API{
		r:  mux.NewRouter(),
		db: db,
	}
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
	// веб-приложение
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

// возвращает n последних новостей
func (api *API) posts(w http.ResponseWriter, r *http.Request) {
	// Получение данных из БД.
	s := mux.Vars(r)["n"]
	limit, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orders, err := api.db.Posts(limit)
	// Отправка данных клиенту в формате JSON.
	json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
