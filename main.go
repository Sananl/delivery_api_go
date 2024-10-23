package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"your_module_name/router" // แทนที่ด้วยชื่อโมดูลของคุณ

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var client *firestore.Client

func init() {
	// กำหนดค่า Firestore client ที่นี่
	ctx := context.Background()
	var err error
	client, err = firestore.NewClient(ctx, "delivery-db-a893e", option.WithCredentialsFile(""))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port 5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
