package testhelpers

import (
	"hcn/myhandlers/teachers"
	"net/http"
)

// Test bla bla...
type Test struct {
	Method       string                                   `json:"Method"`
	URL          string                                   `json:"URL"`
	Function     func(http.ResponseWriter, *http.Request) `json:"Function"`
	Body         string                                   `json:"Body"`
	ExpectedBody string                                   `json:"BodyResponse"`
	StatusCode   int                                      `json:"StatusCode"`
}

// AllTest bla bla...
type AllTest []Test

// CasesGetAllTeachers bla bla...
func CasesGetAllTeachers() AllTest {
	return AllTest{
		{
			Method:       "GET",
			URL:          "/Teachers/GetAllTeachers",
			Function:     teachers.GetAllTeachers,
			Body:         "",
			ExpectedBody: `[{"ID":1,"Name":"Benjamín Calderón Silva","Email":"matlab@email.com"},{"ID":2,"Name":"Oscar David Hurtado Zapata","Email":"oscrdh@email.com"},{"ID":3,"Name":"Christian Camilo Ortiz","Email":"camilorto@email.com"}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetTeacher bla bla...
func CasesGetTeacher() AllTest {
	return AllTest{
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher?id=1",
			Function:     teachers.GetTeacher,
			Body:         "",
			ExpectedBody: `{"ID":1,"Name":"Benjamín Calderón Silva","Email":"matlab@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher?id=2",
			Function:     teachers.GetTeacher,
			Body:         "",
			ExpectedBody: `{"ID":2,"Name":"Oscar David Hurtado Zapata","Email":"oscrdh@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher?id=3",
			Function:     teachers.GetTeacher,
			Body:         "",
			ExpectedBody: `{"ID":3,"Name":"Christian Camilo Ortiz","Email":"camilorto@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher?id=4",
			Function:     teachers.GetTeacher,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateTeacher bla bla...
func CasesUpdateTeacher() AllTest {
	return AllTest{
		{
			Method:       "POST",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":1,"Name":"Benjamín Calderón Ramírez","Email":"matlab@email.com"}`,
			ExpectedBody: `{"ID":1,"Name":"Benjamín Calderón Ramírez","Email":"matlab@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":1,"Name":"Benjamín Calderón Silva","Email":"matlab@email.com"}`,
			ExpectedBody: `{"ID":1,"Name":"Benjamín Calderón Silva","Email":"matlab@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":4,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			ExpectedBody: `{"ID":4,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":55,"Name":"Ghost User","Email":"Ghost@email.com"}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateTeacher bla bla...
func CasesCreateTeacher() AllTest {
	return AllTest{
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":4,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			ExpectedBody: `{"ID":4,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":1,"Name":"Moscar David Hurtado Zapata","Email":"Moscar@email.com"}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1' for key 'PRIMARY'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":AA,"Name":"Mariana Ramos","Email":"mariana@email.com"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteTeacher bla bla...
func CasesDeleteTeacher() AllTest {
	return AllTest{
		{
			Method:       "DELETE",
			URL:          "/Teachers/DeleteTeacher",
			Function:     teachers.DeleteTeacher,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Teachers/DeleteTeacher",
			Function:     teachers.DeleteTeacher,
			Body:         `{"ID":4}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Teachers/DeleteTeacher",
			Function:     teachers.DeleteTeacher,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
