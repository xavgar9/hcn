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
			ExpectedBody: `[{"ID":1,"TeacherID":50001,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"},{"ID":2,"TeacherID":50001,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"},{"ID":3,"TeacherID":50002,"Name":"Clases de piano","CreationDate":"2021-01-06 15:21:50"},{"ID":4,"TeacherID":50003,"Name":"Manejando en Cali","CreationDate":"2021-01-05 11:40:12"}]`,
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
			ExpectedBody: `{"ID":1,"TeacherID":50001,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=2",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `{"ID":2,"TeacherID":50001,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetCourse?id=3",
			Function:     courses.GetCourse,
			Body:         "",
			ExpectedBody: `{"ID":3,"TeacherID":50002,"Name":"Clases de piano","CreationDate":"2021-01-06 15:21:50"}`,
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
			Body:         `{"ID":1,"TeacherID":50001,"Name":"Introducción a Amongos","CreationDate":"2021-01-01 12:00:00"}`,
			ExpectedBody: `{"ID":1,"TeacherID":50001,"Name":"Introducción a Amongos","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":2,"TeacherID":50001,"Name":"Amongos avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			ExpectedBody: `{"ID":2,"TeacherID":50001,"Name":"Amongos avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":1,"TeacherID":50001,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			ExpectedBody: `{"ID":1,"TeacherID":50001,"Name":"Introducción a Matlab","CreationDate":"2021-01-01 12:00:00"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateCourse",
			Function:     courses.UpdateCourse,
			Body:         `{"ID":2,"TeacherID":50001,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
			ExpectedBody: `{"ID":2,"TeacherID":50001,"Name":"Matlab avanzado","CreationDate":"2021-01-01 12:20:08"}`,
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
			Body:         `{"TeacherID":50001,"Name":"Apoyo moral","CreationDate":"2021-01-12 12:33:50"}`,
			ExpectedBody: `{"ID":5,"TeacherID":50001,"Name":"Apoyo moral","CreationDate":"2021-01-12 12:33:50"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateCourse",
			Function:     courses.CreateCourse,
			Body:         `{"Name":"Desapoyo moral","CreationDate":"2021-01-12 15:12:21"}`,
			ExpectedBody: `TeacherID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateCourse",
			Function:     courses.CreateCourse,
			Body:         `{"TeacherID":"Pipe","Name":"Desapoyo moral","CreationDate":"2021-01-12 15:22:15"}`,
			ExpectedBody: `TeacherID is empty or not valid`,
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
			Body:         `{"CourseID":1,"StudentID":10005}`,
			ExpectedBody: `Student added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":1,"StudentID":10006}`,
			ExpectedBody: `Student added`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"CourseID":1,"StudentID":10004}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1-10004' for key 'uq_Students_Courses'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Students/AddStudent",
			Function:     courses.AddStudent,
			Body:         `{"StudentID":10006}`,
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
			Body:         `{"CourseID":Arroz,"StudentID":10002}`,
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
			ExpectedBody: `[{"ID":10001,"Name":"Daniel Gómez Sermeño","Email":"goma@email.com"},{"ID":10002,"Name":"Xavier Garzón López","Email":"xavgar9@email.com"},{"ID":10003,"Name":"Juan F. Gil","Email":"transfer10@email.com"},{"ID":10004,"Name":"Edgar Silva","Email":"ednosil@email.com"}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Students/GetAllStudentsCourse?id=2",
			Function:     courses.GetAllStudentsCourse,
			Body:         ``,
			ExpectedBody: `[{"ID":10005,"Name":"Juanita María Parra Villamíl","Email":"juanitamariap@email.com"},{"ID":10006,"Name":"Sebastián Rodríguez Osorio Silva","Email":"sebasrosorio98@email.com"},{"ID":10007,"Name":"Andrés Felipe Garcés","Email":"andylukast@email.com"}]`,
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
			Body:         `{"CourseID":1, "StudentID":10004}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":10005}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":10006}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveStudent",
			Function:     courses.RemoveStudent,
			Body:         `{"CourseID":1, "StudentID":10007}`,
			ExpectedBody: `No rows deleted`,
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

// CasesAddHCN bla bla...
func CasesAddHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":1,"HCNID":4,"Displayable":1}`,
			ExpectedBody: `HCN added to course`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":2,"HCNID":5,"Displayable":0}`,
			ExpectedBody: `HCN added to course`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":1,"HCNID":1,"Displayable":0}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1-1' for key 'uq_Courses_HCN'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"HCNID":4,"Displayable":1}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":2,"Displayable":0}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":2,"HCNID":5}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":2,"HCNID":5,"Displayable":-1}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":2,"HCNID":5,"Displayable":2}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddHCN",
			Function:     courses.AddHCN,
			Body:         `{"CourseID":"Arroz","StudentID":10005,"Displayable":0}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesGetAllHCNCourse bla bla...
func CasesGetAllHCNCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllHCNCourse?iddddd=1",
			Function:     courses.GetAllHCNCourse,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllHCNCourse?id=1",
			Function:     courses.GetAllHCNCourse,
			Body:         ``,
			ExpectedBody: `[{"ID":1,"CourseID":1,"HCNID":1,"Displayable":1},{"ID":2,"CourseID":1,"HCNID":2,"Displayable":0},{"ID":4,"CourseID":1,"HCNID":4,"Displayable":1}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllHCNCourse?id=2",
			Function:     courses.GetAllHCNCourse,
			Body:         ``,
			ExpectedBody: `[{"ID":3,"CourseID":2,"HCNID":3,"Displayable":1},{"ID":5,"CourseID":2,"HCNID":5,"Displayable":0}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllHCNCourse",
			Function:     courses.GetAllHCNCourse,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesRemoveHCN bla bla...
func CasesRemoveHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveHCN",
			Function:     courses.RemoveHCN,
			Body:         `{"CourseID":1, "HCNID":4}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveHCN",
			Function:     courses.RemoveHCN,
			Body:         `{"CourseID":1, "HCNID":4}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveHCN",
			Function:     courses.RemoveHCN,
			Body:         `{"CourseID":2, "HCNID":5}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveHCN",
			Function:     courses.RemoveHCN,
			Body:         ``,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesVisibilityHCN bla bla...
func CasesVisibilityHCN() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"ID":1,"CourseID":1,"HCNID":1,"Displayable":0}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"HCNID":1,"Displayable":0}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"ID":1,"CourseID":1,"HCNID":1,"Displayable":1}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"HCNID":1,"Displayable":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"CourseID",1 "HCNID":4,"Displayable":1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"ID":1,"HCNID":4,"Displayable":1}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"ID":1,"CourseID":2,"Displayable":0}`,
			ExpectedBody: `HCNID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityHCN",
			Function:     courses.VisibilityHCN,
			Body:         `{"ID":1,"CourseID":2,"HCNID":5}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesAddClinicalCase bla bla...
func CasesAddClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":1,"ClinicalCaseID":2,"Displayable":1}`,
			ExpectedBody: `Clinical Case added to course`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":2,"ClinicalCaseID":3,"Displayable":0}`,
			ExpectedBody: `Clinical Case added to course`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":1,"ClinicalCaseID":1,"Displayable":0}`,
			ExpectedBody: `(SQL) Error 1062: Duplicate entry '1-1' for key 'uq_Courses_CCases'`,
			StatusCode:   http.StatusConflict,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"ClinicalCaseID":4,"Displayable":1}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":2,"Displayable":0}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":2,"ClinicalCaseID":5}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":2,"ClinicalCaseID":5,"Displayable":-1}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":2,"ClinicalCaseID":5,"Displayable":2}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/AddClinicalCase",
			Function:     courses.AddClinicalCase,
			Body:         `{"CourseID":"Arroz","StudentID":10005,"Displayable":0}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesGetAllClinicalCasesCourse bla bla...
func CasesGetAllClinicalCasesCourse() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllClinicalCases?iddddd=1",
			Function:     courses.GetAllClinicalCases,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllClinicalCases?id=1",
			Function:     courses.GetAllClinicalCases,
			Body:         ``,
			ExpectedBody: `[{"ID":1,"CourseID":1,"ClinicalCaseID":1,"Displayable":1},{"ID":4,"CourseID":1,"ClinicalCaseID":2,"Displayable":1}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllClinicalCases?id=2",
			Function:     courses.GetAllClinicalCases,
			Body:         ``,
			ExpectedBody: `[{"ID":2,"CourseID":2,"ClinicalCaseID":2,"Displayable":1},{"ID":5,"CourseID":2,"ClinicalCaseID":3,"Displayable":0}]`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/GetAllClinicalCases",
			Function:     courses.GetAllClinicalCases,
			Body:         ``,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesRemoveClinicalCase bla bla...
func CasesRemoveClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveClinicalCase",
			Function:     courses.RemoveClinicalCase,
			Body:         `{"CourseID":1, "ClinicalCaseID":2}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveClinicalCase",
			Function:     courses.RemoveClinicalCase,
			Body:         `{"CourseID":1, "ClinicalCaseID":2}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveClinicalCase",
			Function:     courses.RemoveClinicalCase,
			Body:         `{"CourseID":2, "ClinicalCaseID":3}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/RemoveClinicalCase",
			Function:     courses.RemoveClinicalCase,
			Body:         ``,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesVisibilityClinicalCase bla bla...
func CasesVisibilityClinicalCase() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"ID":1,"CourseID":1,"ClinicalCaseID":1,"Displayable":0}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"ClinicalCaseID":1,"Displayable":0}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"ID":1,"CourseID":1,"ClinicalCaseID":1,"Displayable":1}`,
			ExpectedBody: `{"ID":1,"CourseID":1,"ClinicalCaseID":1,"Displayable":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"CourseID",1 "ClinicalCaseID":4,"Displayable":1}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"ID":1,"ClinicalCaseID":4,"Displayable":1}`,
			ExpectedBody: `CourseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"ID":1,"CourseID":2,"Displayable":0}`,
			ExpectedBody: `ClinicalCaseID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Courses/VisibilityClinicalCase",
			Function:     courses.VisibilityClinicalCase,
			Body:         `{"ID":1,"CourseID":2,"ClinicalCaseID":5}`,
			ExpectedBody: `Displayable is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
