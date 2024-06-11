package api

import (
	"News/internal"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	r *mux.Router
	i internal.Inter
}

func New(i internal.Inter) *API {
	api := API{
		r: mux.NewRouter(),
		i: i,
	}
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet)
	// веб-приложение
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./internal/webapp"))))
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
	orders, err := api.i.GetPosts(limit)
	// Отправка данных клиенту в формате JSON.
	json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
