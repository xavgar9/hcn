package mymodels

// Token bla bla...
type Token struct {
	ID             *int    `json:"ID",omitempty`
	Name           *string `json:"Name",omitempty`
	Email          *string `json:"Email",omitempty`
	Token          *string `json:"Token",omitempty`
	ExpirationDate *string `json:"ExpirationDate",omitempty`
}

// AllTokens bla bla...
type AllTokens []Token
