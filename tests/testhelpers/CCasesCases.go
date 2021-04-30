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
			ExpectedBody: `[{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001},{"ID":2,"Title":"El pianista de la selva","Description":"Re La Mi Do#","Media":"../activitiesresources/img2.png","TeacherID":50002},{"ID":3,"Title":"Muerte accidental","Description":"¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","Media":"../activitiesresources/ElUltimoBeso.mp3","TeacherID":50003}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetClinicalCase bla bla...
func CasesGetClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"IDD": 1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"ID": }`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"ID": 1}`,
			ExpectedBody: `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"ID": 2}`,
			ExpectedBody: `{"ID":2,"Title":"El pianista de la selva","Description":"Re La Mi Do#","Media":"../activitiesresources/img2.png","TeacherID":50002}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"ID": 3}`,
			ExpectedBody: `{"ID":3,"Title":"Muerte accidental","Description":"¿Por qué se fue? ¿Y por qué murió? ¿Por qué el Señor me la quitó? Se ha ido al cielo y para poder ir yo...","Media":"../activitiesresources/ElUltimoBeso.mp3","TeacherID":50003}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/ClinicalCases/GetClinicalCase",
			Function:     ccases.GetClinicalCase,
			Body:         `{"ID": 12}`,
			ExpectedBody: `(db 2) element does not exist in db`,
			StatusCode:   http.StatusNotFound,
		},
	}
}

// CasesUpdateClinicalCase bla bla...
func CasesUpdateClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"le title actualizado","Description":"La descripción actualizada","Media":"../activitiesresources/updated.png","TeacherID":50001}`,
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":110,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `(db 2) element does not exist in db`,
			StatusCode:   http.StatusNotFound,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/ClinicalCases/UpdateClinicalCase",
			Function:     ccases.UpdateClinicalCase,
			Body:         `{"ID":1,"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","TeacherID":50001}`,
			ExpectedBody: `Media is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
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
			Body:         `{"Title":"El gordo Arnoldo Suárez","Description":"Estás muy gordo amigo, debes rebajar un poco de peso porque ajá hey, respeta un poquito.","Media":"../activitiesresources/img_test.png","TeacherID":50001}`,
			ExpectedBody: "",
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Description":"Benjamón era un joven con IMC PARCHADO.","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El joven parchado","Media":"../activitiesresources/img1.png","TeacherID":50001}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/CreateClinicalCase",
			Function:     ccases.CreateClinicalCase,
			Body:         `{"Title":"El joven parchado","Description":"Benjamón era un joven con IMC PARCHADO.","TeacherID":50001}`,
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
			ExpectedBody: `(db 2) element does not exist in db`,
			StatusCode:   http.StatusNotFound,
		},
		{
			Method:       "DELETE",
			URL:          "/ClinicalCases/DeleteClinicalCase",
			Function:     ccases.DeleteClinicalCase,
			Body:         `{"ID":4}`,
			ExpectedBody: "",
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
			ExpectedBody: "",
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/ClinicalCases/LinkHCN",
			Function:     ccases.LinkHCN,
			Body:         `{"ClinicalCaseID":1,"HCNID":1}`,
			ExpectedBody: `(db 2) Error 1062: Duplicate entry '1-1' for key 'uq_CCases_HCN'`,
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
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/UnlinkHCN",
			Function:     ccases.UnlinkHCN,
			Body:         `{"ClinicalCaseID":2,"HCNID":2}`,
			ExpectedBody: `(db 2) element does not exist in db`,
			StatusCode:   http.StatusNotFound,
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
