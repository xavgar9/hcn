package testhelpers

import (
	"hcn/myhandlers/teachers"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllTeachers bla bla...
func CasesGetAllTeachers() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Teachers/GetAllTeachers",
			Function:     teachers.GetAllTeachers,
			Body:         "",
			ExpectedBody: `[{"ID":50001,"Name":"Gerardo Mauricio Sarria","Email":"gerardo@email.com","Password":"4024fb06e1423da90b80f0274e8e4476"},{"ID":50002,"Name":"Juan Carlos Martinez","Email":"juan@email.com","Password":"a94652aa97c7211ba8954dd15a3cf838"},{"ID":50003,"Name":"Jhoan Lozano Rojas","Email":"jhoan@email.com","Password":"88ca9791c0f2e27a503c23b74896b377"},{"ID":50004,"Name":"Xavier Garzón López","Email":"xavier@email.com","Password":"0f5366b3b19afc3184d23bc73d8cd311"}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetTeacher bla bla...
func CasesGetTeacher() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"IDDD":1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"ID":}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"ID":50001}`,
			ExpectedBody: `{"ID":50001,"Name":"Gerardo Mauricio Sarria","Email":"gerardo@email.com","Password":"4024fb06e1423da90b80f0274e8e4476"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"ID":50002}`,
			ExpectedBody: `{"ID":50002,"Name":"Juan Carlos Martinez","Email":"juan@email.com","Password":"a94652aa97c7211ba8954dd15a3cf838"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"ID":50003}`,
			ExpectedBody: `{"ID":50003,"Name":"Jhoan Lozano Rojas","Email":"jhoan@email.com","Password":"88ca9791c0f2e27a503c23b74896b377"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Teachers/GetTeacher",
			Function:     teachers.GetTeacher,
			Body:         `{"ID":155}`,
			ExpectedBody: "(db 2) element does not exist in db",
			StatusCode:   http.StatusNotFound,
		},
	}
}

// CasesUpdateTeacher bla bla...
func CasesUpdateTeacher() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":50001,"Name":"Gerardo Mauricio Sarria","Email":"gerardo@email.com","Password":"cambio"}`,
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":50001,"Name":"Gerardo Mauricio Sarria","Email":"gerardo@email.com","Password":"4024fb06e1423da90b80f0274e8e4476"}`,
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":55,"Name":"Ghost User","Email":"Ghost@email.com","Password":"4024fb06e1423da90b80f0274e8e4476"}`,
			ExpectedBody: "(db 2) element does not exist in db",
			StatusCode:   http.StatusNotFound,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"Name":"Ghost User","Email":"Ghost@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":55,"Email":"Ghost@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `Name is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":55,"Name":"Ghost User","Password":"conrtraseña"}`,
			ExpectedBody: `Email is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Teachers/UpdateTeacher",
			Function:     teachers.UpdateTeacher,
			Body:         `{"ID":55,"Name":"Ghost User","Email":"Ghost@email.com"}`,
			ExpectedBody: `Password is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesCreateTeacher bla bla...
func CasesCreateTeacher() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":50005,"Name":"Andrés Ocoró","Email":"ocoro@email.com","Password":"conrtraseña"}`,
			ExpectedBody: "",
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":50001,"Name":"Moscar David Hurtado Zapata","Email":"Moscar@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `(db 2) Error 1062: Duplicate entry '50001' for key 'PRIMARY'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":AA,"Name":"Mariana Ramos","Email":"mariana@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"Name":"Mariana Ramos","Email":"mariana@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":55,"Email":"mariana@email.com","Password":"conrtraseña"}`,
			ExpectedBody: `Name is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":56,"Name":"Mariana Ramos","Password":"conrtraseña"}`,
			ExpectedBody: `Email is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Teachers/CreateTeacher",
			Function:     teachers.CreateTeacher,
			Body:         `{"ID":55,"Name":"Mariana Ramos","Email":"mariana@email.com"}`,
			ExpectedBody: `Password is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteTeacher bla bla...
func CasesDeleteTeacher() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Teachers/DeleteTeacher",
			Function:     teachers.DeleteTeacher,
			Body:         `{"ID":500010}`,
			ExpectedBody: "(db 2) element does not exist in db",
			StatusCode:   http.StatusNotFound,
		},
		{
			Method:       "DELETE",
			URL:          "/Teachers/DeleteTeacher",
			Function:     teachers.DeleteTeacher,
			Body:         `{"ID":50005}`,
			ExpectedBody: "",
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
