package testhelpers

import (
	"hcn/myhandlers/courses"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllCourses bla bla...
func CasesGetAllCourses() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Courses/GetAllCourses",
			Function:     courses.GetAllCourses,
			Body:         "",
			ExpectedBody: `[{"ID":1,"TeacherID":1,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"},{"ID":2,"TeacherID":1,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"},{"ID":3,"TeacherID":2,"Name":"Clases de piano","CreationDate":"2021-01-06 15:21:50"},{"ID":4,"TeacherID":3,"Name":"Manejando en Cali","CreationDate":"2021-01-05 11:40:12"}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetCourse bla bla...
func CasesGetCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=1",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `{"ID":1,"TeacherID":1,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=2",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `{"ID":2,"TeacherID":1,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=3",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `{"ID":3,"TeacherID":2,"Name":"Clases de piano","CreationDate":"2021-01-06 15:21:50"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=15",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateCourse bla bla...
func CasesUpdateCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":1,"TeacherID":1,"Name":"Introducción a Amongos","CreationDate":"2021-01-01 12:00:00"}`,
			ExpectedBody: `{"ID":1,"TeacherID":1,"Name":"Introducción a Amongos","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":2,"TeacherID":1,"Name":"Amongos avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			ExpectedBody: `{"ID":2,"TeacherID":1,"Name":"Amongos avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":1,"TeacherID":1,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			ExpectedBody: `{"ID":1,"TeacherID":1,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":2,"TeacherID":1,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			ExpectedBody: `{"ID":2,"TeacherID":1,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":55,"Name":"Ghost User","Email":"Ghost@email.com"}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateCourse bla bla...
func CasesCreateCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/CreateCourse",
			Function:     courses.CreateCourse,
			Body:         `{"TeacherID":1,"Name":"Apoyo moral","CreationDate":"2021-01-12 12:33:50"}`,
			ExpectedBody: `{"ID":5,"TeacherID":1,"Name":"Apoyo moral","CreationDate":"2021-01-12 12:33:50"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateCourse",
			Function:     courses.CreateCourse,
			Body:         `{"Name":"Desapoyo moral","CreationDate":"2021-01-12 15:12:21"}`,
			ExpectedBody: `Teacher is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateCourse",
			Function:     courses.CreateCourse,
			Body:         `{"TeacherID":"Pipe","Name":"Desapoyo moral","CreationDate":"2021-01-12 15:22:15"}`,
			ExpectedBody: `json: cannot unmarshal string into Go struct field Course.TeacherID of type int`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteCourse bla bla...
func CasesDeleteCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteCourse",
			Function:     courses.DeleteCourse,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteCourse",
			Function:     courses.DeleteCourse,
			Body:         `{"ID":5}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteCourse",
			Function:     courses.DeleteCourse,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
