package validator

import (
	"encoding/json"
	"net/url"

	submission "github.com/ryanralphs/contract-tester/backend/submission"
)

type Validation struct {
	url      bool
	expected bool
	method   bool
}

func NewValidation() *Validation {
	return &Validation{
		url:      false,
		expected: false,
		method:   false,
	}
}

func (v *Validation) IsValidURL(str string) bool {
	url, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	if url.Scheme == "" || url.Host == "" {
		return false
	}
	v.url = true
	return v.url
}

func (v *Validation) GetUrl() bool {
	return v.url
}

func (v *Validation) GetExpected() bool {
	return v.expected
}

func (v *Validation) GetMethod() bool {
	return v.method
}

func (v *Validation) SetUrl(b bool) {
	v.url = b
}

func (v *Validation) SetExpected(b bool) {
	v.expected = b
}

func (v *Validation) SetMethod(b bool) {
	v.method = b
}

func (v *Validation) IsValidJSON(s string) bool {
	var expectedJson map[string]interface{}
	err := json.Unmarshal([]byte(s), &expectedJson)
	v.expected = err == nil

	return v.expected
}

func (v *Validation) IsValidMethod(s string) bool {
	if s == "GET" || s == "POST" || s == "PUT" || s == "DELETE" {
		v.method = true
		return v.method
	}
	return false
}

func (v *Validation) IsValid() bool {
	return v.url && v.expected && v.method
}

func (v *Validation) ValidateSubmission(submission submission.Submission) bool {
	valid := v.IsValidURL(submission.Url)
	if !valid {
		return false
	}
	valid = v.IsValidJSON(submission.Expected)
	if !valid {
		return false
	}

	valid = v.IsValidMethod(submission.Method)
	if !valid {
		return false
	}
	return v.IsValid()
}
