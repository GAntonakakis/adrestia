package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Response struct for generic responses
type Response struct {
	Message string `json:"message"`
}

// RequestData struct to represent incoming POST data
type RequestData struct {
	FirstName string `jason:"firstname"`
	LastName  string `jason:"lastname"`
	Age       string `jason:"age"`
	Gender    string `jason:"gender"`
	Ethnicity string `jason:"ethnicity"`
}

var db *sql.DB

// Initialize and connect to the PostgreSQL database
func initDB() {
	var err error
	connStr := "user=postgres dbname=postgres sslmode=disable password=12345678"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database successfully!")
}

// helloHandler to handle GET requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from Go"}
	json.NewEncoder(w).Encode(response)
}

// dataHandler handles incoming POST requests with JSON data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Read the raw body for making sure decoding works correctly
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Raw request body: %s", string(body)) // Log the raw request body

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Decoded data: %+v", requestData) // Log the decoded data

	// Insert data into PostgreSQL
	_, err = db.Exec("INSERT INTO users (firstname, lastname, age, gender, ethnicity) VALUES ($1, $2, $3, $4, $5)", requestData.FirstName, requestData.LastName, requestData.Age, requestData.Gender, requestData.Ethnicity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{Message: "Data received and stored successfully"}
	json.NewEncoder(w).Encode(response)
}

// corsHandler to handle CORS preflight requests
func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()
	fmt.Println("Successfully connected!")

	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/data", dataHandler)

	// Wrap handlers with CORS middleware
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux)))
	http.ListenAndServe(":8080", nil)
}
