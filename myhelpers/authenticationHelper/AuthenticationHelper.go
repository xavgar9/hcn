package authenticationhelper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hcn/config"
	token "hcn/myhelpers/tokenHelper"
	"hcn/mymodels"
)

// isTeacherDataComplete checks if all the teacher data is in
// the struct
func isTeacherDataComplete(teacher mymodels.Teacher) bool {
	switch {
	case (teacher.ID == nil) || (*teacher.ID*1 <= 0):
		return false
	case (teacher.Name == nil) || (len(*teacher.Name) == 0):
		return false
	case (teacher.Email == nil) || (len(*teacher.Email) == 0):
		return false
	default:
		return true
	}
}

// AreCredentialsValid check the credentials in db
func AreCredentialsValid(teacher mymodels.Teacher) bool {
	Db, err := config.MYSQLConnection()
	defer Db.Close()
	if err != nil {
		return false
	}
	hasher := md5.New()
	hasher.Write([]byte(*teacher.Password))
	md5Password := hex.EncodeToString(hasher.Sum(nil))

	if err != nil {
		fmt.Println("(MD5) ", err.Error())
		return false
	}
	rows, err := Db.Query("SELECT AreCredentialsValid(?,?)", teacher.Email, md5Password)
	defer rows.Close()
	if err != nil {
		fmt.Println("(SQL) ", err.Error())
		return false
	}
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			fmt.Println("(SQL) ", err.Error())
			return false
		}
		if result == "True" {
			return true
		}
	}
	return false
}

// UserAuthentication function receives a teachers model and verifies the
// users credentials
// Returns token
func UserAuthentication(teacher mymodels.Teacher) (mymodels.Token, error) {
	var finalToken mymodels.Token
	var id int
	var name string
	var Db, _ = config.MYSQLConnection()
	err := Db.QueryRow("SELECT ID, Name FROM Teachers WHERE Email=?", teacher.Email).Scan(&id, &name)
	defer Db.Close()
	if err != nil {
		return finalToken, err
	}
	teacher.ID = &id
	teacher.Name = &name

	if isTeacherDataComplete(teacher) {
		newToken, err := token.CreateToken(teacher)
		if err != nil {
			return finalToken, err
		}
		//fmt.Println("user 2", newToken)
		ok, err := token.VerifyAuthenticity(newToken)
		if err != nil {
			return finalToken, err
		}
		//fmt.Println("user 3", ok)
		//	fmt.Println("user 3", err)
		// Save the token in bd for further verifications
		if ok {
			if AreCredentialsValid(teacher) {
				claims, err := token.GetTokenClaims(newToken)
				//fmt.Println("user 4", claims)
				if err != nil {
					return finalToken, err
				}

				Db, err := config.MYSQLConnection()
				defer Db.Close()
				//fmt.Println("user 5", err)
				if err != nil {
					return finalToken, err
				}
				//fmt.Println("user 5.1", *claims.ExpirationDate)
				rows, err := Db.Query("SELECT SaveToken(?,?,?)", claims.Email, newToken, claims.ExpirationDate)
				defer rows.Close()
				//fmt.Println("user 6", err)
				if err != nil {
					fmt.Println("(SQL) ", err.Error())
					return finalToken, err
				}
				for rows.Next() {
					var result string
					//fmt.Println("user 6.3", err)
					err := rows.Scan(&result)
					//fmt.Println("user 6.5", err)
					if err != nil {
						fmt.Println("(SQL) ", err.Error())
						return finalToken, err
					}
					//fmt.Println("user 7", result)
					id := claims.ID
					name := claims.Name
					email := claims.Email
					exp := claims.ExpirationDate
					if result == "Insert" || result == "Updated" {
						finalToken.ID = id
						finalToken.Name = name
						finalToken.Email = email
						finalToken.ExpirationDate = exp
						finalToken.Token = &newToken
						return finalToken, nil
					}
					return finalToken, fmt.Errorf("Teacher does not exist")
				}
				//fmt.Println("errors", rows.Err())
			}
			return finalToken, fmt.Errorf("Credentials invalid")
		}
	}
	return finalToken, fmt.Errorf("Teacher data is not complete")
}
