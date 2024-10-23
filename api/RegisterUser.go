package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type User struct {
	GPS     firestore.GeoPoint `json:"gps"`
	Address string             `json:"address"`
	Image   string             `json:"image"`
	Name    string             `json:"name"`
	Password string            `json:"password"`
	Phone   string             `json:"phone"`
	Uid     string             `json:"uid"`
}

var client *firestore.Client

func init() {
	ctx := context.Background()
	var err error
	// สร้าง Firestore client
	client, err = firestore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("path/to/your/serviceAccountKey.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// สร้าง Document ใน Firestore
	_, _, err := client.Collection("users").Add(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to add user to Firestore", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User registered successfully")
}

func main() {
	http.HandleFunc("/register", registerUserHandler)

	fmt.Println("Server is running on port 5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
