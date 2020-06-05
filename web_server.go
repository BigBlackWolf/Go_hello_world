package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "testing"
	sslmode  = "disable"
)

func main() {
	connect_db()
	// create_router()
	// create_server()
}

func connect_db() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	connector, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	result, _ := connector.Exec("Select * from test")
	fmt.Println(result)
	defer connector.Close()
}

func create_router() {
	router := mux.NewRouter()
	sub := router.PathPrefix("/product/").Subrouter()

	awm := authenticationMiddleware{userToken: make(map[string]string)}
	awm.Create()
	sub.Use(awm.Middleware)
	sub.Use(middlewareCheck)

	sub.HandleFunc("/{category}/{id:[0-9]+}", productHandler)
	sub.HandleFunc("/", productInfoHandler)

	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)
}

func create_server() {
	http.ListenAndServe(":8080", nil)
}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Choose category")
	fmt.Fprint(w, response)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := vars["category"]
	response := fmt.Sprintf("Product %s Category: %s", id, category)
	fmt.Fprint(w, response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(map[string]string{"greetings": "Traveller"})
	// response := fmt.Sprintf("Hello world")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func middlewareCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type authenticationMiddleware struct {
	userToken map[string]string
}

func (token *authenticationMiddleware) Create() {
	token.userToken["000000"] = "user0"
	token.userToken["000001"] = "user1"
}

func (awm *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		for _, cookie := range r.Cookies() {
			fmt.Println(cookie)
		}

		if user, found := awm.userToken[token]; found {
			log.Printf("Authenticated user %s", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
