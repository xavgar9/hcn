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
			ExpectedBody: `[{"ID":10001,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"},{"ID":10002,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"},{"ID":10003,"Name":"Juan F. Gil","Email":"transfer10@email.com"},{"ID":10004,"Name":"Edgar Silva","Email":"ednosil@email.com"},{"ID":10005,"Name":"Juanita María Parra Villamíl","Email":"juanitamariap@email.com"},{"ID":10006,"Name":"Sebastián Rodríguez Osorio Silva","Email":"sebasrosorio98@email.com"},{"ID":10007,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}]`,
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
			URL:          "/Students/GetStudent?id=10001",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":10001,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=10002",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":10002,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Students/GetStudent?id=10003",
			Function:     students.GetStudent,
			Body:         "",
			ExpectedBody: `{"ID":10003,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
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
			Body:         `{"ID":10002,"Name":"xavier","Email":"xavier@email.com"}`,
			ExpectedBody: `{"ID":10002,"Name":"xavier","Email":"xavier@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":10003,"Name":"Juan Fernando Gil","Email":"transfer10@email.com"}`,
			ExpectedBody: `{"ID":10003,"Name":"Juan Fernando Gil","Email":"transfer10@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":10007,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			ExpectedBody: `{"ID":10007,"Name":"Andrés Lucumi","Email":"lucumi@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":10002,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			ExpectedBody: `{"ID":10002,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":10003,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
			ExpectedBody: `{"ID":10003,"Name":"Juan F. Gil","Email":"transfer10@email.com"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateStudent",
			Function:     students.UpdateStudent,
			Body:         `{"ID":10007,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}`,
			ExpectedBody: `{"ID":10007,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}`,
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
			Body:         `{"ID":10008,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			ExpectedBody: `{"ID":10008,"Name":"Andrés Ocoró","Email":"ocoro@email.com"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Students/CreateStudent",
			Function:     students.CreateStudent,
			Body:         `{"ID":10008,"Name":"Osquitar Zapata","Email":"osquitar@email.com"}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '10008' for key 'PRIMARY'`,
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
			Body:         `{"ID":100010}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/DeleteStudent",
			Function:     students.DeleteStudent,
			Body:         `{"ID":10008}`,
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
