package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/foobar", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("form is in an invalid state (false), expected valid (true)")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/foobar", nil)
	form := New(r.PostForm)

	form.Required("name", "age")
	if form.Valid() {
		t.Error("required fields are not set but form is valid state (true), expected invalid (false)")
	}

	postedData := url.Values{}
	postedData.Add("name", "Joe")
	postedData.Add("age", "55")

	r = httptest.NewRequest("POST", "/foobar", nil)
	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("name", "age")
	if !form.Valid() {
		t.Error("required fields are set correctly but form is in an invalid state (false), expected valid (true)")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("name", "Joe")
	postedData.Add("age", "")

	form := New(postedData)
	h := form.Has("name")
	if !h {
		t.Error("name field not set or empty but value is present in r.PostForm")
	}

	h = form.Has("age")
	if h {
		t.Error("age field is not empty but although is set as empty string in r.PostForm")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("name", "Joe")

	form := New(postedData)
	form.MinLength("name", 4)

	if form.Valid() {
		t.Error("form is in a valid state but violates MinLength == 1 on name field")
	}

	isError := form.Errors.Get("name")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}

	form = New(postedData)

	form.MinLength("name", 3)
	if !form.Valid() {
		t.Error("form is in an invalid state when it should be valid as MinLength rule on name field is not violated")
	}

	isError = form.Errors.Get("name")
	if isError != "" {
		t.Error("Should not have an error, but one was returned")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "Joe")

	form := New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("form is in a valid state but has invalid email field")
	}

	postedData.Del("email")
	postedData.Add("email", "joe@example.com")

	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form is in an invalid state but has valid email field")
	}
}
