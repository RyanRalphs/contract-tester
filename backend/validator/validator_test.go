package validator_test

import (
	"testing"

	submission "github.com/ryanralphs/contract-tester/backend/submission"
	validator "github.com/ryanralphs/contract-tester/backend/validator"
)

var v *validator.Validation

func setupTest() func() {
	v = validator.NewValidation()

	return func() {
		v = nil
	}
}

func TestIsValidURL(t *testing.T) {
	defer setupTest()()
	validURLs := []string{"http://example.com", "https://www.example.com", "https://example.com/api", "ftp://ftp.example.com/file.txt"}
	invalidURLs := []string{"example.com", "www.example.com", "example.com/api", "file.txt"}

	for _, url := range validURLs {
		if !v.IsValidURL(url) {
			t.Errorf("%v should be valid but got invalid", url)
		}
	}
	for _, url := range invalidURLs {
		if v.IsValidURL(url) {
			t.Errorf("%v should be invalid but got valid", url)
		}
	}
}

func TestIsValidJSON(t *testing.T) {
	defer setupTest()()
	validJSONs := []string{`{"name":"John", "age":30, "city":"New York"}`, `{"data": {"name": "John"}}`}
	invalidJSONs := []string{`{"name":John", "age":30, "city":"New York"}`, `{"name" "John"}`, ``, `invalid json`}

	for _, json := range validJSONs {
		if !v.IsValidJSON(json) {
			t.Errorf("%v should be valid but got invalid", json)
		}
	}
	for _, json := range invalidJSONs {
		if v.IsValidJSON(json) {
			t.Errorf("%v should be invalid but got valid", json)
		}
	}
}

func TestIsValidMethod(t *testing.T) {
	defer setupTest()()
	validMethods := []string{"GET", "POST", "PUT", "DELETE"}
	invalidMethods := []string{"get", "post", "put", "delete", "PATCH", "HEAD", "OPTIONS"}

	for _, method := range validMethods {
		if !v.IsValidMethod(method) {
			t.Errorf("%v should be valid but got invalid", method)
		}
	}
	for _, method := range invalidMethods {
		if v.IsValidMethod(method) {
			t.Errorf("%v should be invalid but got valid", method)
		}
	}
}

func TestValidateSubmission(t *testing.T) {
	defer setupTest()()
	sub := submission.Submission{
		Url:      "https://example.com/api",
		Expected: `{"name":"John", "age":30, "city":"New York"}`,
		Method:   "POST",
	}
	if !v.ValidateSubmission(sub) {
		t.Errorf("Submission should be valid but got invalid")
	}

	sub.Url = "example.com/api"
	if v.ValidateSubmission(sub) {
		t.Errorf("Submission should be invalid but got valid")
	}
}
