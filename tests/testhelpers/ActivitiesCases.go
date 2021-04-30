package testhelpers

import (
	"hcn/myhandlers/activities"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllActivities bla bla...
func CasesGetAllActivities() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Activities/CasesGetAllActivities",
			Function:     activities.GetAllActivities,
			Body:         "",
			ExpectedBody: `[{"ID":1,"Title":"Primera tarea, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-08 20:00:00","CourseID":1,"ClinicalCaseID":1,"HCNID":1,"Difficulty":3},{"ID":2,"Title":"Actividad de prueba","Description":"Por favor ignoren esta actividad, gracias.","Type":"Prueba","CreationDate":"2021-01-09 11:43:21","LimitDate":"2021-01-19 10:59:59","CourseID":2,"ClinicalCaseID":2,"HCNID":2,"Difficulty":1}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetActivity bla bla...
func CasesGetActivity() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Activities/GetActivity",
			Function:     activities.GetActivity,
			Body:         `{"ID":1}`,
			ExpectedBody: `{"ID":1,"Title":"Primera tarea, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-08 20:00:00","CourseID":1,"ClinicalCaseID":1,"HCNID":1,"Difficulty":3}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Activities/GetActivity",
			Function:     activities.GetActivity,
			Body:         `{"IDDD":1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Activities/GetActivity",
			Function:     activities.GetActivity,
			Body:         `{"ID":}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Activities/GetActivity",
			Function:     activities.GetActivity,
			Body:         `{"ID":2}`,
			ExpectedBody: `{"ID":2,"Title":"Actividad de prueba","Description":"Por favor ignoren esta actividad, gracias.","Type":"Prueba","CreationDate":"2021-01-09 11:43:21","LimitDate":"2021-01-19 10:59:59","CourseID":2,"ClinicalCaseID":2,"HCNID":2,"Difficulty":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Activities/GetActivity",
			Function:     activities.GetActivity,
			Body:         `{"ID":12}`,
			ExpectedBody: "(db 2) element does not exist in db",
			StatusCode:   http.StatusNotFound,
		},
	}
}

// CasesUpdateActivity bla bla...
func CasesUpdateActivity() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Type is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `LimitDate is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"Difficulty":3}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"HCNID":1,"ClinicalCaseID":3}`,
			ExpectedBody: `Difficulty is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "PUT",
			URL:          "/Activities/UpdateActivity",
			Function:     activities.UpdateActivity,
			Body:         `{"ID":1,"Title":"Primera tarea, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-08 20:00:00","CourseID":1,"ClinicalCaseID":1,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateActivity bla bla...
func CasesCreateActivity() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Creando mi primera tarea","Description":"Re difficult mai friennd","Type":"Calificable","LimitDate":"2021-01-20 23:30:00","CourseID":1,"ClinicalCaseID":2,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: "",
			StatusCode:   http.StatusCreated,
		},

		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Description":"Re easy pri, solo busquen en Google.","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `Type is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","CourseID":1,"ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `LimitDate is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","ClinicalCaseID":3,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"HCNID":1,"Difficulty":3}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"ClinicalCaseID":3,"Difficulty":3}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Activities/CreateActivity",
			Function:     activities.CreateActivity,
			Body:         `{"Title":"Primera tarea rey, matrices dispersas","Type":"Calificable","Description":"Re easy pri, solo busquen en Google.","CreationDate":"2021-01-08 12:00:00","LimitDate":"2021-01-10 12:00:00","CourseID":1,"HCNID":1,"ClinicalCaseID":3}`,
			ExpectedBody: `Difficulty is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteActivity bla bla...
func CasesDeleteActivity() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Activities/DeleteActivity",
			Function:     activities.DeleteActivity,
			Body:         `{"ID":10}`,
			ExpectedBody: "(db 2) element does not exist in db",
			StatusCode:   http.StatusNotFound,
		},
		{
			Method:       "DELETE",
			URL:          "/Activities/DeleteActivity",
			Function:     activities.DeleteActivity,
			Body:         `{"ID":3}`,
			ExpectedBody: "",
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Activities/DeleteActivity",
			Function:     activities.DeleteActivity,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
