package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// กำหนดเส้นทางต่าง ๆ ที่นี่
	router.HandleFunc("/register", registerUserHandler).Methods("POST")

	return router
}

// handler สำหรับการลงทะเบียนผู้ใช้
func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// สร้าง handler ของคุณที่นี่
}
