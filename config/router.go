package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type router struct {
	Mux    *mux.Router
	Server *http.Server
}

func NewRouter() router {
	mux := mux.NewRouter().PathPrefix(common.BaseUrl).Subrouter()

	return router{
		Mux:    mux,
		Server: createServer(mux),
	}
}

func wireWithCors(r *mux.Router) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"}),
	)(r)
}

func createServer(r *mux.Router) *http.Server {
	// using pointer bcs receiver is pointer
	// actually it's okay to use not pointer even receiver is pointer
	// bcs this struct not return an interface
	// but if struct return an interface you should return as pointer if it's not it will error
	return &http.Server{
		Handler:      wireWithCors(r),
		Addr:         fmt.Sprintf("%s:%s", common.Host, common.Port),
		ReadTimeout:  time.Second * 5, // time for read request header and request body (if exist)
		WriteTimeout: time.Second * 10,
	}
}

func (r router) Start() error {
	return r.Server.ListenAndServe()
}

// function to get to know that available router
// it is directly from mux router doc
func (r router) GetListRouterAvailable() {
	r.Mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}
