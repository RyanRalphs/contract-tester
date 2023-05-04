package submission_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	submission "github.com/ryanralphs/contract-tester/backend/submission"
)

var sub *submission.Submission

func setupTest() func() {
	sub = &submission.Submission{}

	return func() {
		sub = nil
	}
}

func TestNewSubmission(t *testing.T) {
	submit := submission.Submission{
		Url:      "https://example.com",
		Method:   "GET",
		Expected: `{"data": "example"}`,
	}
	jsonSubmit, err := json.Marshal(submit)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/submit", bytes.NewBuffer(jsonSubmit))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	sub, err := submission.NewSubmission(rec, req)
	if err != nil {
		t.Fatal(err)
	}

	if sub.Url != "https://example.com" {
		t.Errorf("got %s, expected %s", sub.Url, "https://example.com")
	}

	if sub.Method != "GET" {
		t.Errorf("got %s, expected %s", sub.Method, "GET")
	}

	if sub.Expected != `{"data": "example"}` {
		t.Errorf("got %s, expected %s", sub.Expected, `{"data": "example"}`)
	}
}

func TestSubmission_SetUrl(t *testing.T) {
	defer setupTest()()
	sub.SetUrl("https://example.com")
	if sub.Url != "https://example.com" {
		t.Errorf("got %s, expected %s", sub.Url, "https://example.com")
	}
}

func TestSubmission_SetExpected(t *testing.T) {
	defer setupTest()()
	sub.SetExpected(`{"data": "example"}`)
	if sub.Expected != `{"data": "example"}` {
		t.Errorf("got %s, expected %s", sub.Expected, `{"data": "example"}`)
	}
}

func TestSubmission_SetMethod(t *testing.T) {
	defer setupTest()()
	sub.SetMethod("GET")
	if sub.Method != "GET" {
		t.Errorf("got %s, expected %s", sub.Method, "GET")
	}
}

func TestSubmission_PopulateSubmissionFields(t *testing.T) {
	defer setupTest()()
	sub.PopulateSubmissionFields("https://example.com", `{"data": "example"}`, "GET")

	if sub.Url != "https://example.com" {
		t.Errorf("got %s, expected %s", sub.Url, "https://example.com")
	}

	if sub.Method != "GET" {
		t.Errorf("got %s, expected %s", sub.Method, "GET")
	}

	if sub.Expected != `{"data": "example"}` {
		t.Errorf("got %s, expected %s", sub.Expected, `{"data": "example"}`)
	}
}
