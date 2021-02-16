package mymodels

// --------------------------------------------------------------------------------------
// MySQL Models
// --------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------
// Mongo Models
// --------------------------------------------------------------------------------------

// HCNmongo contains all sections of the nutritional assessment
type HCNmongo struct {
	ID                 *string        `json:"_id,omitempty" bson:"_id,omitempty"`
	GeneralData        GeneralData    `json:"GeneralData,omitempty" bson:"GeneralData,omitempty"`
	PatientData        PatientData    `json:"PatientData,omitempty" bson:"PatientData,omitempty"`
	ConsultationReason *string        `json:"ConsultationReason,omitempty" bson:"ConsultationReason,omitempty"`
	Anthropometry      Anthropometry  `json:"Anthropometry,omitempty" bson:"Anthropometry,omitempty"`
	Biochemistry       []Biochemistry `json:"Biochemistry,omitempty" bson:"Biochemistry,omitempty"`
	Interpretation     *string        `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback           *string        `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

// GeneralData contains all the general data of the nutriotional assessement
type GeneralData struct {
	ValorationDate *string `json:"ValorationDate,omitempty" bson:"ValorationDate,omitempty"`
	HCNNumber      *string `json:"HCNNumber,omitempty" bson:"HCNNumber,omitempty"`
	AdmissionDate  *string `json:"AdmissionDate,omitempty" bson:"AdmissionDate,omitempty"`
	Room           *string `json:"Room,omitempty" bson:"Room,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

// PatientData contains all the general data of the patient
type PatientData struct {
	FullName       *string `json:"FullName,omitempty" bson:"FullName,omitempty"`
	Birthdate      *string `json:"Birthdate,omitempty" bson:"Birthdate,omitempty"`
	Gender         *string `json:"Gender,omitempty" bson:"Gender,omitempty"`
	Sex            *string `json:"Sex,omitempty" bson:"Sex,omitempty"`
	Age            *int    `json:"Age,omitempty" bson:"Age,omitempty"`
	EPS            *string `json:"EPS,omitempty" bson:"EPS,omitempty"`
	Telephone      *string `json:"Telephone,omitempty" bson:"Telephone,omitempty"`
	Occupation     *string `json:"Occupation,omitempty" bson:"Occupation,omitempty"`
	CivilStatus    *string `json:"CivilStatus,omitempty" bson:"CivilStatus,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

// Anthropometry contains all the anthropometry data of the patient
type Anthropometry struct {
	Weight             weight             `json:"Weight,omitempty" bson:"Weight,omitempty"`
	TricipitalFold     tricipitalFold     `json:"TricipitalFold,omitempty" bson:"TricipitalFold,omitempty"`
	BrachialPerimeter  brachialPerimeter  `json:"BrachialPerimeter,omitempty" bson:"BrachialPerimeter,omitempty"`
	AbdominalPerimeter abdominalPerimeter `json:"AbdominalPerimeter,omitempty" bson:"AbdominalPerimeter,omitempty"`
	SubscapularFold    subscapularFold    `json:"SubscapularFold,omitempty" bson:"SubscapularFold,omitempty"`
	Height             height             `json:"Height,omitempty" bson:"Height,omitempty"`
	Structure          structure          `json:"Structure,omitempty" bson:"Structure,omitempty"`
	BMI                bmi                `json:"BMI,omitempty" bson:"BMI,omitempty"`
	Interpretation     *string            `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback           *string            `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type weight struct {
	Actual         *string `json:"Actual,omitempty" bson:"Actual,omitempty"`
	Usual          *string `json:"Usual,omitempty" bson:"Usual,omitempty"`
	Reference      *string `json:"Reference,omitempty" bson:"Reference,omitempty"`
	ChangeWeight   *string `json:"ChangeWeight,omitempty" bson:"ChangeWeight,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type tricipitalFold struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type brachialPerimeter struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type abdominalPerimeter struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type subscapularFold struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type height struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type structure struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

type bmi struct {
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

// Biochemistry contains all the anthropometry data of the patient
type Biochemistry struct {
	Date           *string `json:"Date,omitempty" bson:"Date,omitempty"`
	Parameter      *string `json:"parameter,omitempty" bson:"parameter,omitempty"`
	Value          *string `json:"Value,omitempty" bson:"Value,omitempty"`
	ReferenceValue *string `json:"ReferenceValue,omitempty" bson:"ReferenceValue,omitempty"`
	Interpretation *string `json:"Interpretation,omitempty" bson:"Interpretation,omitempty"`
	Feedback       *string `json:"Feedback,omitempty" bson:"Feedback,omitempty"`
}

// AllHCNmongo bla bla...
type AllHCNmongo []HCNmongo
