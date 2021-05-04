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
	expirationTime := time.Now().Add(120 * time.Minute)
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
	//fmt.Println("auth verify 0")
	if receivedToken != "" {
		//fmt.Println("auth verify 0.1")
		token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			fmt.Println("auth verify 1")
			return false, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Verify time expiration
			tokenTime := getTime(claims["exp"])
			actualTime := time.Now()
			if tokenTime.Sub(actualTime) < 0 {
				Db, err := config.MYSQLConnection()
				if err != nil {
					fmt.Println("auth verify 2")
					return false, err
				}
				Db.Query("SELECT DeleteToken(?)", token)
				defer Db.Close()
				//fmt.Println("auth verify 3")
				return false, fmt.Errorf("Token expired")
			}
			//fmt.Println("auth verify 3.1")
			return true, nil
		}
		//fmt.Println("auth verify 4")
		return false, err
	}
	//fmt.Println("auth verify 5")
	return false, fmt.Errorf("Token not valid:" + receivedToken)
}

// IsValid returns true if the token authenticity is ok, exists in
// db and the expiration date is valid
// Returns the validation
func IsValid(receivedToken string) (bool, error) {
	// Verify token authenticity
	ok, err := VerifyAuthenticity(receivedToken)
	if ok {
		//fmt.Println("IsValid Error 0")
		claims, err := GetTokenClaims(receivedToken)
		if err != nil {
			fmt.Println("IsValid Error 1")
			return false, err
		}
		//fmt.Println("IsValid Error 1.1")
		email := *claims.Email
		//fmt.Println("IsValid Error 1.2")
		if email != "" {
			//fmt.Println("IsValid Error 1.3")
			// Verify token in db
			Db, err := config.MYSQLConnection()
			defer Db.Close()
			if err != nil {
				fmt.Println("IsValid Error 2")
				return false, err
			}
			//fmt.Println("IsValid Error 2.1", receivedToken)
			rows, err := Db.Query("SELECT IsValidToken(?,?)", claims.ID, receivedToken)
			defer rows.Close()
			//fmt.Println("IsValid Error 2.11")
			if err == nil {
				//fmt.Println("IsValid Error 2.12")
				for rows.Next() {
					//fmt.Println("IsValid Error 2.13")
					var result string
					err = rows.Scan(&result)
					if err == nil {
						//fmt.Println("IsValid Error 2.14")
						switch result {
						case "True":
							//fmt.Println("IsValid Error 2.15")
							return true, nil
						case "False":
							//fmt.Println("IsValid Error 2.16")
							return false, errors.New("SQL function returned False")
						}
					}
				}
				return false, errors.New("SQL function returned False")
			} else {
				//fmt.Println("IsValid Error 3")
				return false, err
			}
			//fmt.Println("IsValid Error 3 final")
		} else {
			//fmt.Println("IsValid verify 3.1")
			err = errors.New("UserID does not exist")
		}
		//fmt.Println("IsValid Error 3 final final")
	}
	//fmt.Println(" IsValid Error 4")
	return false, err
}
