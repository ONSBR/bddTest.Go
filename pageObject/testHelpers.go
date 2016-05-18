package pageObject

import (
	"fmt"
	"net/http"
	"time"

	"github.com/braintree/manners"
	"github.com/stretchr/testify/mock"
)

type mockedHTTPHandler struct {
	mock.Mock
}

func (h *mockedHTTPHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<input type='button' id='buttonId'/><input type='text' id='txtId'/>HTML Content")
}

// InitTestServer inits a test HTTP Server
func InitTestServer() string {
	server := manners.NewWithServer(&http.Server{
		Addr:           fmt.Sprintf("172.17.0.1:8080"),
		Handler:        &mockedHTTPHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	go server.ListenAndServe()

	return "http://172.17.0.1:8080"
}
