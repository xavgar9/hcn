package mymodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -------------------------------------------
// MySQL Models
// -------------------------------------------

// HCN is the assessment in MySQL (Historia Clínica Nutricional)...
type HCN struct {
	ID        *int `json:"ID"`
	TeacherID *int `json:"TeacherID"`
}

// HCNVinculation bla bla...
type HCNVinculation struct {
	ID             *int `json:"ID"`
	ClinicalCaseID *int `json:"ClinicalCaseID"`
	HCNID          *int `json:"HCNID"`
}

// AllHCN bla bla...
type AllHCN []HCN

// -------------------------------------------
// Mongo Models
// -------------------------------------------

// HCNmongo contains all sections of the nutritional assessment
type HCNmongo struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GeneralDatamongo   GeneralDatamongo   `json:"GeneralDatamongo,omitempty" bson:"GeneralDatamongo,omitempty"`
	ConsultationReason *string            `json:"ConsultationReason,omitempty" bson:"ConsultationReason,omitempty"`
}

// GeneralDatamongo contains all the general data of the patient
type GeneralDatamongo struct {
	ValorationDate *string `json:"ValorationDate,omitempty" bson:"ValorationDate,omitempty"`
	HCNNumber      *int    `json:"HCNNumber,omitempty" bson:"HCNNumber,omitempty"`
	AdmissionDate  *string `json:"AdmissionDate,omitempty" bson:"AdmissionDate,omitempty"`
	Room           *string `json:"Room,omitempty" bson:"Room,omitempty"`
}

// Anthropometry contains all the anthropometry data of the patient
type Anthropometry struct {
	Weight weight `json:"Weight,omitempty" bson:"Weight,omitempty"`

	TricipitalFold               *int `json:"TricipitalFold,omitempty" bson:"TricipitalFold,omitempty"`
	TricipitalFoldInterpretation *int `json:"TricipitalFoldInterpretation,omitempty" bson:"TricipitalFoldInterpretation,omitempty"`

	BrachialPerimeter               *string `json:"BrachialPerimeter,omitempty" bson:"BrachialPerimeter,omitempty"`
	BrachialPerimeterInterpretation *string `json:"BrachialPerimeterInterpretation,omitempty" bson:"BrachialPerimeterInterpretation,omitempty"`
}

type weight struct {
	Actual               *int     `json:"Actual,omitempty" bson:"Actual,omitempty"`
	Usual                *int     `json:"Usual,omitempty" bson:"Usual,omitempty"`
	Reference            *int     `json:"Reference,omitempty" bson:"Reference,omitempty"`
	ChangeWeight         *float32 `json:"ChangeWeight,omitempty" bson:"ChangeWeight,omitempty"`
	WeightInterpretation *string  `json:"WeightInterpretation,omitempty" bson:"WeightInterpretation,omitempty"`
}

// AllHCNmongo bla bla...
type AllHCNmongo []HCNmongo
