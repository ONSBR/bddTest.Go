package pageObject

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/stretchr/testify/mock"
	"net/http"
	"time"
)

type MockedHTTPHandler struct {
	mock.Mock
}

func (h *MockedHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<input type='button' id='elementId'/>HTML Content")
}

// initTestServer inits a test HTTP Server
func initTestServer() {
	server := manners.NewWithServer(&http.Server{
		Addr:           fmt.Sprintf("172.17.0.1:8080"),
		Handler:        &MockedHTTPHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	go server.ListenAndServe()
}
