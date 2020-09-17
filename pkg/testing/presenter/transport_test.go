package presenter

import (
	"net/http"
	"testing"
)

// command for benchmark go test -bench=. -benchtime 100x
// httprouter is much faster than gorilla mux
// but the difference not too far as i do benchmark, otherwise many people said it give much different

// this api using httprouter
func BenchmarkGetV2(b *testing.B) {
	client := http.Client{}

	for n := 0; n < b.N; n++ {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/template/api/testing", nil)
		client.Do(req)
	}

}

// this api using gorilla mux router
func BenchmarkGet(b *testing.B) {
	client := http.Client{}

	for n := 0; n < b.N; n++ {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/template/api/testing", nil)
		client.Do(req)
	}

}

func BenchmarkGetByParamv2(b *testing.B) {
	client := http.Client{}

	for n := 0; n < b.N; n++ {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/template/api/testing/1", nil)
		client.Do(req)
	}

}

func BenchmarkGetByParam(b *testing.B) {
	client := http.Client{}

	for n := 0; n < b.N; n++ {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/template/api/testing/1", nil)
		client.Do(req)
	}

}
