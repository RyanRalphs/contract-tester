package submission

import (
	"encoding/json"
	"net/http"
)

type Submission struct {
	Url      string `json:"url"`
	Method   string `json:"method"`
	Expected string `json:"payload"`
}

func NewSubmission(w http.ResponseWriter, r *http.Request) (*Submission, error) {
	NewSubmission := &Submission{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewSubmission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return NewSubmission, err
	}
	return NewSubmission, nil
}

func (s *Submission) SetUrl(url string) {
	s.Url = url
}

func (s *Submission) SetExpected(expected string) {
	s.Expected = expected
}

func (s *Submission) SetMethod(method string) {
	s.Method = method
}

func (s *Submission) PopulateSubmissionFields(url, expected, method string) {
	s.SetUrl(url)
	s.SetExpected(expected)
	s.SetMethod(method)
}
