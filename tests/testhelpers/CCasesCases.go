package testhelpers

import (
	"hcn/myhandlers/ccases"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllClinicalCases bla bla...
func CasesGetAllClinicalCases() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetAllClinicalCases",
			Function:     ccases.GetAllClinicalCases,
			Body:         "",
			ExpectedBody: `[{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1},{"ID":2,"Title":"El pianista de la selva","Description":"Re La Mi Do#","Media":"../activitiesresources/img2.png","TeacherID":2},{"ID":3,"Title":"Muerte accidental","Description":"¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","Media":"../activitiesresources/ElUltimoBeso.mp3","TeacherID":3}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetClinicalCase bla bla...
func CasesGetClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?idddd=1",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?id=",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?id=1",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?id=2",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: `{"ID":2,"Title":"El pianista de la selva","Description":"Re La Mi Do#","Media":"../activitiesresources/img2.png","TeacherID":2}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?id=3",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: `{"ID":3,"Title":"Muerte accidental","Description":"¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","Media":"../activitiesresources/ElUltimoBeso.mp3","TeacherID":3}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase?id=15",
			Function:     ccases.GetClinicalCase,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateClinicalCase bla bla...
func CasesUpdateClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"le title actualizado","Description":"La descripción actualizada","Media":"../activitiesresources/updated.png","TeacherID":1}`,
			ExpectedBody: `{"ID":1,"Title":"le title actualizado","Description":"La descripción actualizada","Media":"../activitiesresources/updated.png","TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":110,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","TeacherID":1}`,
			ExpectedBody: `Media is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png"}`,
			ExpectedBody: `TeacherID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesCreateClinicalCase bla bla...
func CasesCreateClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El gordo Arnoldo Suárez","Description":"Estás muy gordo amigo, debes rebajar un poco de peso porque ajá hey, respeta un poquito.","Media":"../activitiesresources/img_test.png","TeacherID":1}`,
			ExpectedBody: `{"ID":4,"Title":"El gordo Arnoldo Suárez","Description":"Estás muy gordo amigo, debes rebajar un poco de peso porque ajá hey, respeta un poquito.","Media":"../activitiesresources/img_test.png","TeacherID":1}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El joven parchado","Media":"../activitiesresources/img1.png","TeacherID":1}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","TeacherID":1}`,
			ExpectedBody: `Media is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png"}`,
			ExpectedBody: `TeacherID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteClinicalCase bla bla...
func CasesDeleteClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/ClinicalCases/DeleteClinicalCase",
			Function:     ccases.DeleteClinicalCase,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/ClinicalCases/DeleteClinicalCase",
			Function:     ccases.DeleteClinicalCase,
			Body:         `{"ID":4}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/ClinicalCases/DeleteClinicalCase",
			Function:     ccases.DeleteClinicalCase,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesLinkHCN vinculates vinculates one HCN with one Clinical Case
func CasesLinkHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":2,"HCNID":2}`,
			ExpectedBody: `HCN added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":3,"HCNID":3}`,
			ExpectedBody: `HCN added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":1,"HCNID":1}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1-1' for key 'uq_CCases_HCN'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"HCNID":4}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":2}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":3,"HCNID":5}`,
			ExpectedBody: `HCN added`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUnlinkHCN bla bla...
func CasesUnlinkHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"ClinicalCaseID":2,"HCNID":2}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"ClinicalCaseID":2,"HCNID":2}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"ClinicalCaseID":3,"HCNID":3}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"ClinicalCaseID":2}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"HCNID":4}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
