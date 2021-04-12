package tokenhelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"hcn/config"
	"hcn/mymodels"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
La cabecera (o header): que generalmente consta de dos partes: el tipo de token (JWT) y
el algoritmo con el que fue firmado (como HMAC SHA256 o RSA). Este contenido, como
podemos deducir por su nombre, está en formato JSON ({"alg": "HS256", "type": "JWT"}) y codificado en Base64Url.
El contenido (o payload): que contiene lo que se conoce como claims (generalmente, el usuario)
y otros datos adicionales. También está codificado en Base64Url.
La signatura (o signature): resultante de coger la cabecera y el contenido codificados,
un secreto, el algoritmo de firma especificado en la cabecera y firmarlo. Ésta puede ser
usada para verificar que el contenido no se modificó en el camino y, en el caso de los tokens
firmados con una clave privada, también sirve para verificar que el remitente del JWT es quién dice ser.
*/

type claims struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
	jwt.StandardClaims
}

func getTime(expirationTime interface{}) time.Time {
	var tm time.Time
	switch iat := expirationTime.(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}
	return tm
}

// CreateToken function receives a teacher model
// Returns the teacher signed token
func CreateToken(teacher mymodels.Teacher) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := claims{
		*teacher.ID,
		*teacher.Name,
		*teacher.Email,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		fmt.Println("Error creating token", err)
		return "", err
	}
	return signedToken, nil
}

// GetTokenClaims function receives a token
// Return all the data stored in claims
func GetTokenClaims(receivedToken string) (mymodels.Token, error) {
	var data mymodels.Token
	if receivedToken != "" {
		token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWTSecret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := int(claims["ID"].(float64))
			name := claims["Name"].(string)
			email := claims["Email"].(string)

			expirationDate := strings.Split(getTime(claims["exp"]).String(), " -")

			data.ID = &id
			data.Name = &name
			data.Email = &email
			data.ExpirationDate = &expirationDate[0]
			return data, nil
		}
		return data, err
	}

	return data, fmt.Errorf("Token not valid")
}

// VerifyAuthenticity function receives token and validate that
// the token was not modified or not expired
// Returns the authenticity verification
func VerifyAuthenticity(receivedToken string) (bool, error) {
	//fmt.Println("received token: ", receivedToken)
	if receivedToken != "" {
		token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			return false, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Verify time expiration
			tokenTime := getTime(claims["exp"])
			actualTime := time.Now()
			if tokenTime.Sub(actualTime) < 0 {
				Db, err := config.MYSQLConnection()
				if err != nil {
					return false, err
				}
				Db.Query("SELECT DeleteToken(?)", token)
				defer Db.Close()
				return false, fmt.Errorf("Token expired")
			}
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("Token not valid:" + receivedToken)
}

// IsValid returns true if the token authenticity is ok, exists in
// db and the expiration date is valid
// Returns the validation
func IsValid(receivedToken string) (bool, error) {
	// Verify token authenticity
	ok, err := VerifyAuthenticity(receivedToken)
	if ok {
		claims, err := GetTokenClaims(receivedToken)
		if err != nil {
			return false, err
		}
		email := *claims.Email
		if email != "" {
			// Verify token in db
			Db, err := config.MYSQLConnection()
			defer Db.Close()
			if err != nil {
				return false, err
			}
			rows, err := Db.Query("SELECT IsValidToken(?)", receivedToken)
			if err == nil {
				for rows.Next() {
					var result string
					err = rows.Scan(&result)
					if err == nil {
						switch result {
						case "True":
							return true, nil
						case "False":
							return false, errors.New("SQL function returned False")
						}
					}
				}
			} else {
				return false, err
			}
		} else {
			err = errors.New("UserID does not exist")
		}
	}
	return false, err
}
