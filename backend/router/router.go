package router




import (
 "awesomeProject/middleware"

"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/card", middleware.Card).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/guestledger", middleware.GuestLedger).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/quotes", middleware.Quotes).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/counter", middleware.Counter).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/fizzbuzz/{count}", middleware.FizzBuzz).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/morse", middleware.MorseCode).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
