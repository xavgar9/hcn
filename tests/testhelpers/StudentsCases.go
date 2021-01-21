package testhelpers

import (
	"hcn/myhandlers/students"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllStudents bla bla...
func CasesGetAllStudents() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Students/GetAllStudents",
			Function:     students.GetAllStudents,
			Body:         "",
			ExpectedBody: `[{"ID":1,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"},{"ID":2,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"},{"ID":3,"Name":"Juan F. Gil","Email":"transfer10@email.com"},{"ID":4,"Name":"Edgar Silva","Email":"ednosil@email.com"},{"ID":5,"Name":"Juanita María Parra Villamíl","Email":"juanitamariap@email.com"},{"ID":6,"Name":"Sebastián Rodríguez Osorio Silva","Email":"sebasrosorio98@email.com"},{"ID":7,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetStudent bla bla...
func CasesGetStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?idddd=1",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=1",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":1,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=2",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":2,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=3",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":3,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=15",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateStudent bla bla...
func CasesUpdateStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":2,"Name":"xavier","Email":"xavier@email.com"}`,
			ExpectedBody: `{"ID":2,"Name":"xavier","Email":"xavier@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":3,"Name":"Juan Fernando Gil","Email":"transfer10@email.com"}`,
			ExpectedBody: `{"ID":3,"Name":"Juan Fernando Gil","Email":"transfer10@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":7,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			ExpectedBody: `{"ID":7,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":2,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			ExpectedBody: `{"ID":2,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":3,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
			ExpectedBody: `{"ID":3,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":7,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}`,
			ExpectedBody: `{"ID":7,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":55,"Name":"Ghost User","Email":"Ghost@email.com"}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateStudent bla bla...
func CasesCreateStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":8,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			ExpectedBody: `{"ID":8,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":8,"Name":"Osquitar Zapata","Email":"osquitar@email.com"}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '8' for key 'PRIMARY'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":Antonia,"Name":"Antonia Vélez","Email":"antonia@email.com"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{Name":"Antonia Vélez","Email":"antonia@email.com"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":15,"Email":"antonia@email.com"}`,
			ExpectedBody: `Name is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":21,"Name":"Antonia Vélez"}`,
			ExpectedBody: `Email is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteStudent bla bla...
func CasesDeleteStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Students/DeleteStudent",
			Function:     students.DeleteStudent,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/DeleteStudent",
			Function:     students.DeleteStudent,
			Body:         `{"ID":8}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/DeleteStudent",
			Function:     students.DeleteStudent,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
