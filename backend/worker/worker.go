package worker

import (
	"fmt"
	"log"
	"net/http"

	poller "github.com/ryanralphs/contract-tester/backend/poller"
	submission "github.com/ryanralphs/contract-tester/backend/submission"
	validator "github.com/ryanralphs/contract-tester/backend/validator"
)

func Run() {
	SetUpServer()
}

func ContractTestSubmission(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ContractTestSubmission ")

	NewSubmission, err := submission.NewSubmission(w, r)

	if err != nil {
		log.Fatal(err)
	}

	NewValidator := validator.NewValidation()
	NewPoller := poller.NewPoller(NewSubmission, NewValidator)

	if outcome := NewPoller.Validator.ValidateSubmission(*NewSubmission); !outcome {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		fmt.Println("Invalid submission")
		return
	}

	isMatch := NewPoller.Run()

	if !isMatch {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.Error(w, "Contract test failed", http.StatusBadRequest)
		fmt.Println("Contract test failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

}

func SetUpServer() {
	http.Handle("/", http.FileServer(http.Dir("./../client/build")))
	http.HandleFunc("/api", ContractTestSubmission)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
