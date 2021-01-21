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
			URL:          "/Courses/GetCourse?idddd=1",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
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

// CasesAddStudent bla bla...
func CasesAddStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":1,"StudentID":5}`,
			ExpectedBody: `Student added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":1,"StudentID":6}`,
			ExpectedBody: `Student added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":1,"StudentID":5}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1-5' for key 'uq_Students_Courses'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"StudentID":6}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":6}`,
			ExpectedBody: `StudentID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":Arroz,"StudentID":2}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesGetAllStudentsCourse bla bla...
func CasesGetAllStudentsCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Students/GetAllStudentsCourse?iddddd=1",
			Function:     courses.GetAllStudentsCourse,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/GetAllStudentsCourse?id=1",
			Function:     courses.GetAllStudentsCourse,
			Body:         ``,
			ExpectedBody: `[{"ID":1,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"},{"ID":2,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"},{"ID":3,"Name":"Juan F. Gil","Email":"transfer10@email.com"},{"ID":4,"Name":"Edgar Silva","Email":"ednosil@email.com"}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/GetAllStudentsCourse?id=2",
			Function:     courses.GetAllStudentsCourse,
			Body:         ``,
			ExpectedBody: `[{"ID":5,"Name":"Juanita María Parra Villamíl","Email":"juanitamariap@email.com"},{"ID":6,"Name":"Sebastián Rodríguez Osorio Silva","Email":"sebasrosorio98@email.com"},{"ID":7,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/GetAllStudentsCourse",
			Function:     courses.GetAllStudentsCourse,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesRemoveStudent bla bla...
func CasesRemoveStudent() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":5}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":5}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":6}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         ``,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
