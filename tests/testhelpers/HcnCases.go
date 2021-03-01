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
			ExpectedBody: `[{"ID":1,"TeacherID":50001,"MongoID":"60346574367b678c2e13c072"},{"ID":2,"TeacherID":50001,"MongoID":"5tgdfdghghdfdbc1295d976d"},{"ID":3,"TeacherID":50001,"MongoID":"dfvdbcd2437sfbc1295d54de"},{"ID":4,"TeacherID":50002,"MongoID":"nhddbcd2437d1bc8fvewigd3"},{"ID":5,"TeacherID":50003,"MongoID":"knv949f95dvreij9mvo00dmd"}]`,
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
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?id=",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?id=1",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":1,"TeacherID":50001,"MongoID":"60346574367b678c2e13c072"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?id=2",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":2,"TeacherID":50001,"MongoID":"5tgdfdghghdfdbc1295d976d"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?id=3",
			Function:     hcn.GetHCN,
			Body:         "",
			ExpectedBody: `{"ID":3,"TeacherID":50001,"MongoID":"dfvdbcd2437sfbc1295d54de"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "GET",
			URL:          "/HCN/GetHCN?id=15",
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
			URL:          "/HCN/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":1,"TeacherID":50002,"MongoID":"EstaEsun4Prueb4"}`,
			ExpectedBody: `{"ID":1,"TeacherID":50002,"MongoID":"EstaEsun4Prueb4"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":1,"TeacherID":50001,"MongoID":"EstaEsNoun4Prueb4"}`,
			ExpectedBody: `{"ID":1,"TeacherID":50001,"MongoID":"EstaEsNoun4Prueb4"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":1,"TeacherID":50001,"MongoID":"60346574367b678c2e13c072"}`,
			ExpectedBody: `{"ID":1,"TeacherID":50001,"MongoID":"60346574367b678c2e13c072"}`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"ID":11,"TeacherID":50002,"MongoID":"EstaEsNoun4Prueb4"}`,
			ExpectedBody: `No rows updated`,
			StatusCode:   http.StatusOK,
		},
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCN",
			Function:     hcn.UpdateHCN,
			Body:         `{"TeacherID":50002}`,
			ExpectedBody: `ID is empty or not valid`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCN",
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
			URL:          "/HCN/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":50001,"MongoID": "EstaEsun4Prueb4"}`,
			ExpectedBody: `{"ID":6,"TeacherID":50001,"MongoID":"EstaEsun4Prueb4"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/HCN/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":50002,"MongoID":"EstaEsun4Prueb4"}`,
			ExpectedBody: `{"ID":7,"TeacherID":50002,"MongoID":"EstaEsun4Prueb4"}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:       "POST",
			URL:          "/HCN/CreateHCN",
			Function:     hcn.CreateHCN,
			Body:         `{"TeacherID":5000AA}`,
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

// CasesDeleteAllHCNMongo bla bla...
func CasesDeleteAllHCNMongo() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "DELETE",
			URL:          "/HCN/DeleteAllHCNMongo",
			Function:     hcn.DeleteAllHCNMongo,
			Body:         ``,
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetAllHCNMongo1 bla bla...
func CasesGetAllHCNMongo1() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/HCN/GetAllHCNMongo",
			Function:     hcn.GetAllHCNMongo,
			Body:         "",
			ExpectedBody: "{}",
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesCreateHCNMongo bla bla...
func CasesCreateHCNMongo() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:   "GET",
			URL:      "/HCN/CreateHCNMongo",
			Function: hcn.CreateHCNMongo,
			Body: `{
				"GeneralData": {
					"ValorationDate": "2020/02/66",
					"HCNNumber": "526",
					"AdmissionDate": "2020/02/06",
					"Room": "26"
				},
				"PatientData": {
					"FullName": "Benito Antonio Martínez",
					"Birthdate": "526",
					"Gender": "2020/02/06",
					"Sex": "M",
					"Age": 26, 
					"EPS": "PR Salud", 
					"Telephone": "31658245", 
					"Occupation": "Singer",
					"CivilStatus": "Single",        
					"Feedback": "Información completa"
				},
				"ConsultationReason": "fatiga",
				"Anthropometry": {
					"Weight": {
						"Actual": 63.2,
						"Usual": 65,
						"Reference": 64,
						"ChangeWeight": 0.05,
						"Interpretation": "El paciente está bien de peso según los indicadores",
						"Feedback": "Buen trabajo"
					},
					"TricipitalFold": {
						"Value": 12.6,	
						"Interpretation": "Interpretación tricipal",
						"Feedback": "Buen trabajo"
					},
					"BrachialPerimeter": {
						"Value": 12.6,	
						"Interpretation": "Interpretación braquial",
						"Feedback": "Buen trabajo"
					},
					"AbdominalPerimeter": {
						"Value": 12.6,	
						"Interpretation": "Interpretación abdominal",
						"Feedback": "Buen trabajo"
					},
					"SubscapularFold": {
						"Value": 12.6,	
						"Interpretation": "Interpretación subescapular",
						"Feedback": "Buen trabajo"
					},
			
					"Height": {
						"Value": 12.6,	
						"Interpretation": "Interpretación altura",
						"Feedback": "Buen trabajo"
					},
					"Structure": {
						"Value": 12.6,	
						"Interpretation": "Interpretación estructura",
						"Feedback": "Buen trabajo"
					},
					"BMI": {
						"Value": 12.6,	
						"Interpretation": "Interpretación IMC",
						"Feedback": "Buen trabajo"
					}
				},
				"Biochemistry": [
					{
						"Date": "2020/01/02",
						"Parameter": "El parmaétro",
						"Value": "26.2",
						"ReferenceValue":"25",
						"Interpretation": "Está un toque mal "
					},
					{
						"Date": "2020/01/02",
						"Parameter": "El parmaétro dos",
						"Value": "26.2",
						"ReferenceValue":"25",
						"Interpretation": "Está un toque mal "
					}
				]
			}`,
			ExpectedBody: `{"GeneralData":{"ValorationDate":"2020/02/66","HCNNumber":"526","AdmissionDate":"2020/02/06","Room":"26"},"ConsultationReason":"fatiga","Anthropometry":{"Weight":{"Actual":"","Usual":"","Reference":"","ChangeWeight":"","Interpretation":"El paciente está bien de peso según los indicadores","Feedback":"Buen trabajo"},"TricipitalFold":{"Value":"","Interpretation":"Interpretación tricipal","Feedback":"Buen trabajo"},"BrachialPerimeter":{"Value":"","Interpretation":"Interpretación braquial","Feedback":"Buen trabajo"},"AbdominalPerimeter":{"Value":"","Interpretation":"Interpretación abdominal","Feedback":"Buen trabajo"},"SubscapularFold":{"Value":"","Interpretation":"Interpretación subescapular","Feedback":"Buen trabajo"},"Height":{"Value":"","Interpretation":"Interpretación altura","Feedback":"Buen trabajo"},"Structure":{"Value":"","Interpretation":"Interpretación estructura","Feedback":"Buen trabajo"},"BMI":{"Value":"","Interpretation":"Interpretación IMC","Feedback":"Buen trabajo"}},"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "},{"Date":"2020/01/02","parameter":"El parmaétro dos","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]}`,
			StatusCode:   http.StatusCreated,
		},
		{
			Method:   "GET",
			URL:      "/HCN/CreateHCNMongo",
			Function: hcn.CreateHCNMongo,
			Body: `{
					"PatientData": {
						"FullName": "Benito Antonio Martínez",
						"Birthdate": "526",
						"Gender": "2020/02/06",
						"Sex": "M",
						"Age": 26,
						"EPS": "PR Salud",
						"Telephone": "31658245",
						"Occupation": "Singer",
						"CivilStatus": "Single",
						"Feedback": "Información completa"
					},
					"ConsultationReason": "fatiga",
					"Anthropometry": {
						"Weight": {
							"Actual": 63.2,
							"Usual": 65,
							"Reference": 64,
							"ChangeWeight": 0.05,
							"Interpretation": "El paciente está bien de peso según los indicadores",
							"Feedback": "Buen trabajo"
						},
						"TricipitalFold": {
							"Value": 12.6,
							"Interpretation": "Interpretación tricipal",
							"Feedback": "Buen trabajo"
						},
						"BrachialPerimeter": {
							"Value": 12.6,
							"Interpretation": "Interpretación braquial",
							"Feedback": "Buen trabajo"
						},
						"AbdominalPerimeter": {
							"Value": 12.6,
							"Interpretation": "Interpretación abdominal",
							"Feedback": "Buen trabajo"
						},
						"SubscapularFold": {
							"Value": 12.6,
							"Interpretation": "Interpretación subescapular",
							"Feedback": "Buen trabajo"
						},

						"Height": {
							"Value": 12.6,
							"Interpretation": "Interpretación altura",
							"Feedback": "Buen trabajo"
						},
						"Structure": {
							"Value": 12.6,
							"Interpretation": "Interpretación estructura",
							"Feedback": "Buen trabajo"
						},
						"BMI": {
							"Value": 12.6,
							"Interpretation": "Interpretación IMC",
							"Feedback": "Buen trabajo"
						}
					},
					"Biochemistry": [
						{
							"Date": "2020/01/02",
							"Parameter": "El parmaétro",
							"Value": "26.2",
							"ReferenceValue":"25",
							"Interpretation": "Está un toque mal "
						},
						{
							"Date": "2020/01/02",
							"Parameter": "El parmaétro dos",
							"Value": "26.2",
							"ReferenceValue":"25",
							"Interpretation": "Está un toque mal "
						}
					]
				}`,
			ExpectedBody: `{"ConsultationReason":"fatiga","Anthropometry":{"Weight":{"Actual":"","Usual":"","Reference":"","ChangeWeight":"","Interpretation":"El paciente está bien de peso según los indicadores","Feedback":"Buen trabajo"},"TricipitalFold":{"Value":"","Interpretation":"Interpretación tricipal","Feedback":"Buen trabajo"},"BrachialPerimeter":{"Value":"","Interpretation":"Interpretación braquial","Feedback":"Buen trabajo"},"AbdominalPerimeter":{"Value":"","Interpretation":"Interpretación abdominal","Feedback":"Buen trabajo"},"SubscapularFold":{"Value":"","Interpretation":"Interpretación subescapular","Feedback":"Buen trabajo"},"Height":{"Value":"","Interpretation":"Interpretación altura","Feedback":"Buen trabajo"},"Structure":{"Value":"","Interpretation":"Interpretación estructura","Feedback":"Buen trabajo"},"BMI":{"Value":"","Interpretation":"Interpretación IMC","Feedback":"Buen trabajo"}},"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "},{"Date":"2020/01/02","parameter":"El parmaétro dos","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]}`,
			StatusCode:   http.StatusCreated,
		},

		{
			Method:       "GET",
			URL:          "/HCN/CreateHCNMongo",
			Function:     hcn.CreateHCNMongo,
			Body:         `{ "ConsultationReason": "fatiga", }`,
			ExpectedBody: `{"ConsultationReason":"fatiga"}`,
			StatusCode:   http.StatusCreated,
		},

		{
			Method:   "GET",
			URL:      "/HCN/CreateHCNMongo",
			Function: hcn.CreateHCNMongo,
			Body: `{ "Biochemistry": [
						{
							"Date": "2020/01/02",
							"Parameter": "El parmaétro",
							"Value": "26.2",
							"ReferenceValue":"25",
							"Interpretation": "Está un toque mal "
						}				
					]
				}`,
			ExpectedBody: `{"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]}`,
			StatusCode:   http.StatusCreated,
		},

		{
			Method:       "GET",
			URL:          "/HCN/CreateHCNMongo",
			Function:     hcn.CreateHCNMongo,
			Body:         `{}`,
			ExpectedBody: ``,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesGetAllHCNMongo2 bla bla...
func CasesGetAllHCNMongo2() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "GET",
			URL:          "/HCN/GetAllHCNMongo",
			Function:     hcn.GetAllHCNMongo,
			Body:         "",
			ExpectedBody: `[{"GeneralData":{"ValorationDate":"2020/02/66","HCNNumber":"526","AdmissionDate":"2020/02/06","Room":"26"},"ConsultationReason":"fatiga","Anthropometry":{"Weight":{"Actual":"","Usual":"","Reference":"","ChangeWeight":"","Interpretation":"El paciente está bien de peso según los indicadores","Feedback":"Buen trabajo"},"TricipitalFold":{"Value":"","Interpretation":"Interpretación tricipal","Feedback":"Buen trabajo"},"BrachialPerimeter":{"Value":"","Interpretation":"Interpretación braquial","Feedback":"Buen trabajo"},"AbdominalPerimeter":{"Value":"","Interpretation":"Interpretación abdominal","Feedback":"Buen trabajo"},"SubscapularFold":{"Value":"","Interpretation":"Interpretación subescapular","Feedback":"Buen trabajo"},"Height":{"Value":"","Interpretation":"Interpretación altura","Feedback":"Buen trabajo"},"Structure":{"Value":"","Interpretation":"Interpretación estructura","Feedback":"Buen trabajo"},"BMI":{"Value":"","Interpretation":"Interpretación IMC","Feedback":"Buen trabajo"}},"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "},{"Date":"2020/01/02","parameter":"El parmaétro dos","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]},{"ConsultationReason":"fatiga","Anthropometry":{"Weight":{"Actual":"","Usual":"","Reference":"","ChangeWeight":"","Interpretation":"El paciente está bien de peso según los indicadores","Feedback":"Buen trabajo"},"TricipitalFold":{"Value":"","Interpretation":"Interpretación tricipal","Feedback":"Buen trabajo"},"BrachialPerimeter":{"Value":"","Interpretation":"Interpretación braquial","Feedback":"Buen trabajo"},"AbdominalPerimeter":{"Value":"","Interpretation":"Interpretación abdominal","Feedback":"Buen trabajo"},"SubscapularFold":{"Value":"","Interpretation":"Interpretación subescapular","Feedback":"Buen trabajo"},"Height":{"Value":"","Interpretation":"Interpretación altura","Feedback":"Buen trabajo"},"Structure":{"Value":"","Interpretation":"Interpretación estructura","Feedback":"Buen trabajo"},"BMI":{"Value":"","Interpretation":"Interpretación IMC","Feedback":"Buen trabajo"}},"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "},{"Date":"2020/01/02","parameter":"El parmaétro dos","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]},{"ConsultationReason":"fatiga"},{"Biochemistry":[{"Date":"2020/01/02","parameter":"El parmaétro","Value":"26.2","ReferenceValue":"25","Interpretation":"Está un toque mal "}]}]`,
			StatusCode:   http.StatusOK,
		},
	}
}

// CasesUpdateHCNMongo bla bla...
func CasesUpdateHCNMongo() mymodels.AllTest {
	return mymodels.AllTest{
		{
			Method:       "POST",
			URL:          "/HCN/UpdateHCNMongo",
			Function:     hcn.UpdateHCNMongo,
			Body:         "",
			ExpectedBody: "{}",
			StatusCode:   http.StatusOK,
		},
	}
}
