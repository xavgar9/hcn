package testhelpers

import (
	"hcn/myhandlers/feedbacks"
	"hcn/mymodels"
	"net/http"
)

// CasesGetAllFeedbacks bla bla...
func CasesGetAllFeedbacks() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Feedbacks/GetAllFeedbacks",
			Function:     feedbacks.GetAllFeedbacks,
			Body:         "",
			ExpectedBody: `[{"ID":1,"ActivityID":1,"StudentID":1},{"ID":2,"ActivityID":1,"StudentID":2},{"ID":3,"ActivityID":1,"StudentID":3},{"ID":4,"ActivityID":1,"StudentID":4},{"ID":5,"ActivityID":1,"StudentID":5},{"ID":6,"ActivityID":1,"StudentID":6},{"ID":7,"ActivityID":1,"StudentID":7}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetFeedback bla bla...
func CasesGetFeedback() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/Feedbacks/GetFeedback?idddd=1",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetFeedback?id=",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetFeedback?id=1",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: `{"ID":1,"ActivityID":1,"StudentID":1}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetFeedback?id=2",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: `{"ID":2,"ActivityID":1,"StudentID":2}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetFeedback?id=3",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: `{"ID":3,"ActivityID":1,"StudentID":3}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/Courses/GetFeedback?id=15",
			Function:     feedbacks.GetFeedback,
			Body:         "",
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateFeedback bla bla...
func CasesUpdateFeedback() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":6,"ActivityID":2,"StudentID":6}`,
			ExpectedBody: `{"ID":6,"ActivityID":2,"StudentID":6}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":7,"ActivityID":2,"StudentID":7}`,
			ExpectedBody: `{"ID":7,"ActivityID":2,"StudentID":7}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":6,"ActivityID":1,"StudentID":6}`,
			ExpectedBody: `{"ID":6,"ActivityID":1,"StudentID":6}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Courses/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":7,"ActivityID":1,"StudentID":7}`,
			ExpectedBody: `{"ID":7,"ActivityID":1,"StudentID":7}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":77,"ActivityID":1,"StudentID":7}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ActivityID":1,"StudentID":7}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":77,"StudentID":7}`,
			ExpectedBody: `ActivityID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/Students/UpdateFeedback",
			Function:     feedbacks.UpdateFeedback,
			Body:         `{"ID":77,"ActivityID":1}`,
			ExpectedBody: `StudentID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesCreateFeedback bla bla...
func CasesCreateFeedback() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/Courses/CreateFeedback",
			Function:     feedbacks.CreateFeedback,
			Body:         `{"ActivityID":2,"StudentID":1}`,
			ExpectedBody: `{"ID":8,"ActivityID":2,"StudentID":1}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateFeedback",
			Function:     feedbacks.CreateFeedback,
			Body:         `{"ActivityID":2,"StudentID":2}`,
			ExpectedBody: `{"ID":9,"ActivityID":2,"StudentID":2}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/Courses/CreateFeedback",
			Function:     feedbacks.CreateFeedback,
			Body:         `{"ActivityID":AA,"StudentID":2}`,
			ExpectedBody: `invalid character 'A' looking for beginning of value`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}

// CasesDeleteFeedback bla bla...
func CasesDeleteFeedback() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteFeedback",
			Function:     feedbacks.DeleteFeedback,
			Body:         `{"ID":10}`,
			ExpectedBody: `No rows deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteFeedback",
			Function:     feedbacks.DeleteFeedback,
			Body:         `{"ID":8}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteFeedback",
			Function:     feedbacks.DeleteFeedback,
			Body:         `{"ID":9}`,
			ExpectedBody: `One row deleted`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "DELETE",
			URL:          "/Courses/DeleteFeedback",
			Function:     feedbacks.DeleteFeedback,
			Body:         `{"ID":Arroz}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
	}
}
