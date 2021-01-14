package mymodels

// HCN (Historia Clínica Nutricional)...
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
