package poller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	submission "github.com/ryanralphs/contract-tester/backend/submission"
	validator "github.com/ryanralphs/contract-tester/backend/validator"
)

type Poller struct {
	Submission *submission.Submission
	Validator  *validator.Validation
}

type ApiResponse struct {
	Outcome string
}

func NewPoller(submission *submission.Submission, validator *validator.Validation) *Poller {
	return &Poller{
		Submission: submission,
		Validator:  validator,
	}
}

func (p *Poller) Run() bool {
	if p.Submission.Method == "GET" {
		outcome, err := p.Get()
		if err != nil {
			return false
		}
		return outcome
	}
	if p.Submission.Method == "POST" {
		p.Post()
	}
	if p.Submission.Method == "PUT" {
		p.Put()
	}
	if p.Submission.Method == "DELETE" {
		p.Delete()
	}
	return false
}

func (p *Poller) Get() (bool, error) {
	match := false
	resp, err := http.Get(p.Submission.Url)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	for key := range data {
		match := strings.Contains(p.Submission.Expected, key)
		fmt.Printf("Key: %s, Match: %t\n", key, match)
		return match, nil
	}
	return match, nil
}

func (p *Poller) Post() {
	resp, err := http.Post(p.Submission.Url, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func (p *Poller) Put() {
	resp, err := http.Post(p.Submission.Url, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func (p *Poller) Delete() {
	resp, err := http.Post(p.Submission.Url, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
