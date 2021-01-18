package testhelpers

import (
	"hcn/myhandlers/hcn"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllHCN bla bla...
func CasesGetAllHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/HCN/GetAllHCN",
			Function:     hcn.GetAllHCN,
			Body:         "",
			ExpectedBody: `[{"ID":1,"TeacherID":1},{"ID":2,"TeacherID":1},{"ID":3,"TeacherID":1},{"ID":4,"TeacherID":2},{"ID":5,"TeacherID":3}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetHCN bla bla...
func CasesGetHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?idddd=1",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `Url Param 'id' is missing or is invalid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetHCN?id=",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `Url Param 'id' is missing or is invalid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetHCN?id=1",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":1,"TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetHCN?id=2",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":2,"TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetHCN?id=3",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":3,"TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetHCN?id=15",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateHCN bla bla...
func CasesUpdateHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":1,"TeacherID":2}`,
			ExpectedBody: `{"ID":1,"TeacherID":2}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":1,"TeacherID":1}`,
			ExpectedBody: `{"ID":1,"TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":11,"TeacherID":2}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"TeacherID":2}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":2}`,
			ExpectedBody: `TeacherID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesCreateHCN bla bla...
func CasesCreateHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":1}`,
			ExpectedBody: `{"ID":6,"TeacherID":1}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":2}`,
			ExpectedBody: `{"ID":7,"TeacherID":2}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":AA}`,
			ExpectedBody: `TeacherID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteHCN bla bla...
func CasesDeleteHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/HCN/DeleteHCN",
			Function:     hcn.DeleteHCN,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/HCN/DeleteHCN",
			Function:     hcn.DeleteHCN,
			Body:         `{"ID":6}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/HCN/DeleteHCN",
			Function:     hcn.DeleteHCN,
			Body:         `{"ID":7}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/HCN/DeleteHCN",
			Function:     hcn.DeleteHCN,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
