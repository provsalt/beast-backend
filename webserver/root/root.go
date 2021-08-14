package root

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("Hello, this server is the api endpoint for xmrvsbeast "))
}
