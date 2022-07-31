package decorator

import (
	"log"
	"net/http"
	"testing"
)

func TestDecorator(t *testing.T) {
	http.HandleFunc("/", method(address(cost(HelloHandler))))
	log.Printf("starting http server at: %s", "http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
