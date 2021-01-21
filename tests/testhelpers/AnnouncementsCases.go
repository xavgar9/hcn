package testhelpers

import (
	"hcn/myhandlers/announcements"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllAnnouncements bla bla...
func CasesGetAllAnnouncements() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Announcements/GetAllAnnouncements",
			Function:     announcements.GetAllAnnouncements,
			Body:         "",
			ExpectedBody: `[{"ID":1,"CourseID":1,"Title":"¡Bienvenidos al curso!","Description":"Este es un curso básico de Matlab. LOS AMO.","CreationDate":"2021-01-17 13:34:28"},{"ID":2,"CourseID":1,"Title":"¡Primera tarea!","Description":"Resuelvan una matriz dispersa 100x100.","CreationDate":"2021-01-17 13:34:28"},{"ID":3,"CourseID":1,"Title":"Hola a todos","Description":"Hola muchachos, los quiero mucho. Estudien bye!","CreationDate":"2021-01-17 13:34:28"},{"ID":4,"CourseID":1,"Title":"Material guía","Description":"Busquen en Youtube. \"Accidentes de tránsito graves sin censura.\"","CreationDate":"2021-01-17 13:34:28"}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetAnnouncement bla bla...
func CasesGetAnnouncement() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?id=1",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: `{"ID":1,"CourseID":1,"Title":"¡Bienvenidos al curso!","Description":"Este es un curso básico de Matlab. LOS AMO.","CreationDate":"2021-01-17 13:34:28"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?iddd=1",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?id=",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?id=2",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: `{"ID":2,"CourseID":1,"Title":"¡Primera tarea!","Description":"Resuelvan una matriz dispersa 100x100.","CreationDate":"2021-01-17 13:34:28"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?id=3",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: `{"ID":3,"CourseID":1,"Title":"Hola a todos","Description":"Hola muchachos, los quiero mucho. Estudien bye!","CreationDate":"2021-01-17 13:34:28"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Announcements/GetAnnouncement?id=15",
			Function:     announcements.GetAnnouncement,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateAnnouncement bla bla...
func CasesUpdateAnnouncement() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Announcements/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"ID":1,"CourseID":1,"Title":"Anuncio 1","Description":"Este es el anuncio 1","CreationDate":"2021-01-09 00:10:30"}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"Title":"Anuncio 1","Description":"Este es el anuncio 1","CreationDate":"2021-01-09 00:10:30"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"Title":1,"Description":"Amongos para el lunes","CreationDate":"2021-01-01 12:20:08"}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"ID":1,"Description":"Este es el anuncio 1","CreationDate":"2021-01-09 00:10:30"}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"ID":1,"Title":"Anuncio 1","CreationDate":"2021-01-09 00:10:30"}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"ID":1,"Title":"Anuncio 1","Description":"Este es el anuncio 1"}`,
			ExpectedBody: `CreationDate is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/UpdateAnnouncement",
			Function:     announcements.UpdateAnnouncement,
			Body:         `{"ID":1,"CourseID":1,"Title":"¡Bienvenidos al curso!","Description":"Este es un curso básico de Matlab. LOS AMO.","CreationDate":"2021-01-17 13:34:28"}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"Title":"¡Bienvenidos al curso!","Description":"Este es un curso básico de Matlab. LOS AMO.","CreationDate":"2021-01-17 13:34:28"}`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateAnnouncement bla bla...
func CasesCreateAnnouncement() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Announcements/CreateAnnouncement",
			Function:     announcements.CreateAnnouncement,
			Body:         `{"CourseID":1,"Title":"Título de prueba 1","Description":"Descripción de prueba 1"}`,
			ExpectedBody: `{"ID":5,"CourseID":1,"Title":"Título de prueba 1","Description":"Descripción de prueba 1","CreationDate":null}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/CreateAnnouncement",
			Function:     announcements.CreateAnnouncement,
			Body:         `{"Title":"Título de prueba 1","Description": "Descripción de prueba 1"}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/CreateAnnouncement",
			Function:     announcements.CreateAnnouncement,
			Body:         `{"CourseID":1,"Description": "Descripción de prueba 1"}`,
			ExpectedBody: `Title is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Announcements/CreateAnnouncement",
			Function:     announcements.CreateAnnouncement,
			Body:         `{"CourseID":1,"Title": "Título de prueba 1"}`,
			ExpectedBody: `Description is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteAnnouncement bla bla...
func CasesDeleteAnnouncement() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Announcements/DeleteAnnouncement",
			Function:     announcements.DeleteAnnouncement,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Announcements/DeleteAnnouncement",
			Function:     announcements.DeleteAnnouncement,
			Body:         `{"ID":5}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Announcements/DeleteAnnouncement",
			Function:     announcements.DeleteAnnouncement,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
