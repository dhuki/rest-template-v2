package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type router struct {
	Mux    *httprouter.Router
	Server *http.Server
}

func NewRouter() router {
	mux := httprouter.New()

	return router{
		Mux:    mux,
		Server: createServer(mux),
	}
}

func wireWithCors(r *httprouter.Router) http.Handler {
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"www.google.com"},
	// })

	return cors.Default().Handler(r)
}

func createServer(r *httprouter.Router) *http.Server {
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
// func (r router) GetListRouterAvailable() {
// 	r.Mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 		t, err := route.GetPathTemplate()
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(t)
// 		return nil
// 	})
// }
