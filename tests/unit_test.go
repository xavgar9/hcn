package tests

import (
	"hcn/myhandlers/teachers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router(url string, f func(http.ResponseWriter, *http.Request), method string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(url, f).Methods(method)
	return router
}

// Test bla bla...
type Test struct {
	Method     string                                   `json:"Method"`
	URL        string                                   `json:"URL"`
	Function   func(http.ResponseWriter, *http.Request) `json:"Function"`
	Body       teacher                                  `json:"Body"`
	StatusCode int                                      `json:"StatusCode"`
	Answer     string                                   `json:"Answer"`
}

type teacher struct {
	ID    int
	Name  string
	Email string
}

// AllTest bla bla...
type AllTest []Test

/*
func TestTeachers(t *testing.T) {

	allTest := AllTest{
		{
			Method:     "GET",
			URL:        "/Teachers/GetAllTeachers",
			Function:   teachers.GetAllTeachers,
			Body:       teacher{},
			StatusCode: 200,
			Answer:     "Expected status: 200",
		},
		{
			Method:     "GET",
			URL:        "/Teachers/GetTeacher/1",
			Function:   teachers.GetTeacher,
			Body:       teacher{},
			StatusCode: 200,
			Answer:     "Expected status: 200",
		},
		{
			Method:     "POST",
			URL:        "/Teachers/UpdateTeacher",
			Function:   teachers.UpdateTeacher,
			Body:       teacher{ID: 1, Name: "Juan Carlos", Email: "juan@email.com"},
			StatusCode: 200,
			Answer:     "Expected status: 200",
		},
		{
			Method:     "POST",
			URL:        "/Teachers/CreateTeacher",
			Function:   teachers.UpdateTeacher,
			Body:       teacher{ID: 1144, Name: "Ernesto Flores", Email: "ernesto@email.com"},
			StatusCode: 200,
			Answer:     "Expected status: 200",
		},
		{
			Method:     "DELETE",
			URL:        "/Teachers/DeleteTeacher",
			Function:   teachers.DeleteTeacher,
			Body:       teacher{ID: 1},
			StatusCode: 200,
			Answer:     "Expected status: 200",
		},
	}

	for _, test := range allTest {
		if test.Method == "GET" {
			w, _ := http.NewRequest(test.Method, test.URL, nil)
			r := httptest.NewRecorder()
			Router(test.URL, test.Function, test.Method).ServeHTTP(r, w)
			assert.Equal(t, test.StatusCode, r.Code, test.Answer)
		} else {
			jsonBody, _ := json.Marshal(test.Body)
			w, _ := http.NewRequest(test.Method, test.URL, bytes.NewBuffer(jsonBody))
			r := httptest.NewRecorder()
			Router(test.URL, test.Function, test.Method).ServeHTTP(r, w)
			assert.Equal(t, test.StatusCode, r.Code, test.Answer)
		}
	}
}
*/
func TestGetAllTeachers(t *testing.T) {
	test := Test{
		Method:     "GET",
		URL:        "/Teachers/GetAllTeachers",
		Function:   teachers.GetAllTeachers,
		Body:       teacher{},
		StatusCode: 200,
		Answer:     "Expected status: 200",
	}
	w, _ := http.NewRequest(test.Method, test.URL, nil)
	r := httptest.NewRecorder()
	Router(test.URL, test.Function, test.Method).ServeHTTP(r, w)
	assert.Equal(t, test.StatusCode, r.Code, test.Answer)
}

func TestGetTeacher(t *testing.T) {
	test := Test{
		Method:     "GET",
		URL:        "/Teachers/GetTeacher",
		Function:   teachers.GetTeacher,
		Body:       teacher{},
		StatusCode: 200,
		Answer:     "Expected status: 200",
	}
	w, _ := http.NewRequest(test.Method, test.URL, nil)
	w.Header.Add("id", "1")
	r := httptest.NewRecorder()
	Router(test.URL, test.Function, test.Method).ServeHTTP(r, w)
	assert.Equal(t, test.StatusCode, r.Code, test.Answer)
}
