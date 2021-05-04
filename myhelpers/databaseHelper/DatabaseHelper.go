package databaseHelper

import (
	"database/sql"
	"errors"
	"hcn/config"
	"reflect"
	"strconv"

	helper "hcn/myhelpers/structValidationHelper"
)

// -------------------------------------------------------
// CRUD FUNCTIONS
// -------------------------------------------------------

// Insert executes a insert query without returning any rows.
func Insert(model interface{}) (int, error) {

	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return 0, errors.New("(db 1) " + err.Error())
	}

	// Prepare query
	query := createInsertQuery(model)
	//fmt.Println("create query:", query)

	// Execute query
	rows, err := db.Exec(query)
	if err != nil {
		return 0, errors.New("(db 2) " + err.Error())
	}

	// Handle results
	lastInsertedID, err := rows.LastInsertId()
	if err != nil {
		return 0, errors.New("(db 3) " + err.Error())
	}

	return int(lastInsertedID), nil
}

// GetAll executes a select query returning multiples rows.
func GetAll(model interface{}) (*sql.Rows, error) {

	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return nil, errors.New("(db 1) " + err.Error())
	}

	// Prepare query
	query := createGetAllQuery(model)
	//fmt.Println("getAll query:", query)

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("(db 2) " + err.Error())
	}

	return rows, nil
}

// Get executes a select query returning one row.
func Get(model interface{}) (*sql.Rows, error) {

	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return nil, errors.New("(db 1) " + err.Error())
	}

	// Check if the model exists in db
	if !exists(model) {
		return nil, errors.New("(db 2) element does not exist in db")
	}

	// Prepare query
	query := createGetQuery(model)
	//fmt.Println("GET query", query)

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("(db 3) " + err.Error())
	}

	return rows, nil
}

// Update executes a update query returning check variable.
func Update(model interface{}) (bool, error) {

	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return false, errors.New("(db 1) " + err.Error())
	}

	// Check if the model exists in db
	if !exists(model) {
		return false, errors.New("(db 2) element does not exist in db")
	}

	// Prepare query
	query := createUpdateQuery(model)
	//fmt.Println("update query:", query)

	// Execute query
	result, err := db.Exec(query)
	if err != nil {
		return false, errors.New("(db 3) " + err.Error())
	}

	// Handle results
	cnt, err := result.RowsAffected()
	if err != nil {
		return false, errors.New("(db 4) " + err.Error())
	}
	if cnt == 1 {
		return true, nil
	} else {
		return false, errors.New("(db 5) Expected to affect 1 row, affected " + strconv.Itoa(int(cnt)) + " rows")
	}
}

// Delete executes a delete query returning check variable.
func Delete(model interface{}) (bool, error) {
	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return false, errors.New("(db 1) " + err.Error())
	}

	// Check if the model exists in db
	if !exists(model) {
		return false, errors.New("(db 2) element does not exist in db")
	}

	// Prepare query
	query := createDeleteQuery(model)
	//fmt.Println("delete query:", query)

	// Execute query
	result, err := db.Exec(query)
	if err != nil {
		return false, errors.New("(db 3) " + err.Error())
	}

	// Handle results
	cnt, err := result.RowsAffected()
	if err != nil {
		return false, errors.New("(db 4) " + err.Error())
	}
	if cnt == 1 {
		return true, nil
	}
	modelName, _, _, _ := helper.GetFields(model)
	if modelName == "SolvedHCN" {
		return true, nil
	}
	return false, errors.New("(db 5) Expected to affect 1 row, affected " + strconv.Itoa(int(cnt)) + " rows")
}

// -------------------------------------------------------
// AUXILIAR FUNCTIONS
// -------------------------------------------------------

// Some models does not have the same name his table in db.
// So, we have to change the variable value base in the model name.
func tableName(modelName string) string {
	var tableName string
	dbTablesNames := map[string]string{
		"Activity":           "Activities",
		"Announcement":       "Announcements",
		"ClinicalCase":       "Clinical_Cases",
		"Course":             "Courses",
		"HCN":                "HCN",
		"SolvedHCN":          "Solved_HCN",
		"Teacher":            "Teachers",
		"Student":            "Students",
		"StudentTuition":     "Students_Courses",
		"HCNVinculation":     "CCases_HCN",
		"CourseClinicalCase": "Courses_CCases",
		"CourseHCN":          "Courses_HCN",
	}

	if _, ok := dbTablesNames[modelName]; ok {
		tableName = dbTablesNames[modelName]
	} else {
		tableName = modelName
	}
	return tableName
}

// Some models are the representation of an intermediate
// table in db. The filter field is not the ID PK.
func filterField(tableName string) string {
	dbIntermediateTables := map[string]string{
		"Students_Courses": "CourseID",
		"Courses_HCN":      "CourseID",
		"Courses_CCases":   "CourseID",
		"CCases_HCN":       "CourseID",
		"Solved_HCN":       "ActivityID",
	}

	if _, ok := dbIntermediateTables[tableName]; ok {
		return dbIntermediateTables[tableName]
	}
	return "ID"
}

// RowsToStruct scans rows to the slice pointed to by dest.
// The slice elements must be pointers to structs with exported
// fields corresponding to the the columns in the result set.
//
// The function panics if dest is not as described above.
// https://stackoverflow.com/questions/62240553/generalizing-sql-rows-scan-in-go
func RowsToStruct(rows *sql.Rows, dest interface{}) error {

	// 1. Create a slice of structs from the passed struct type of model
	//
	// Not needed, the caller passes pointer to destination slice.
	// Elem() dereferences the pointer.
	//
	// If you do need to create the slice in this function
	// instead of using the argument, then use
	// destv := reflect.MakeSlice(reflect.TypeOf(model).

	destv := reflect.ValueOf(dest).Elem()

	// Allocate argument slice once before the loop.

	args := make([]interface{}, destv.Type().Elem().NumField())

	// 2. Loop through each row

	for rows.Next() {

		// 3. Create a struct of passed mode interface{} type
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		// 4. Scan the row results to a slice of interface{}
		// 5. Set the field values of struct created in step 3 using the slice in step 4
		//
		// Scan directly to the struct fields so the database
		// package handles the conversion from database
		// types to a Go types.
		//
		// The slice args is filled with pointers to struct fields.

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		// 6. Add the struct created in step 3 to slice created in step 1

		destv.Set(reflect.Append(destv, rowv))

	}
	return nil
}

// exists veryfies if the model exists in db
func exists(model interface{}) bool {

	// Create connection
	db, err := config.MYSQLConnection()
	defer db.Close()
	if err != nil {
		return false
	}

	// Prepare query
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	query := createExistsQuery(modelName, fieldsNames, fieldsValues)
	//fmt.Println("exists query:", query)

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		return false
	}

	// Handle results
	for rows.Next() {
		return true
	}
	return false
}

func createExistsQuery(modelName string, fieldsNames []string, fieldsValues []string) string {
	tableName := tableName(modelName)
	filterField := filterField(tableName)
	query := "SELECT ID"
	if filterField == "ID" {
		query = query + " FROM " + tableName + " WHERE ID=" + fieldsValues[0]
	} else {
		if tableName == "Courses_CCases" || tableName == "Courses_HCN" || tableName == "Solved_HCN" {
			query = query + " FROM " + tableName + " WHERE " + fieldsNames[1] + "=" + fieldsValues[1]
		} else {
			query = query + " FROM " + tableName + " WHERE " + fieldsNames[1] + "=" + fieldsValues[1] + " AND " + fieldsNames[2] + "=" + fieldsValues[2]
		}
	}
	return query
}

func createInsertQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	query := "INSERT INTO " + tableName + "("
	for i, fieldName := range fieldsNames {
		if i != 0 {
			if i == len(fieldsNames)-1 {
				query = query + fieldName
			} else {
				query = query + fieldName + ", "
			}
		} else {
			if modelName == "Student" || modelName == "Teacher" {
				query = query + fieldName + ", "
			}
		}
	}

	query = query + ") VALUES ("
	for i, fieldValue := range fieldsValues {
		if i != 0 {
			if i == len(fieldsNames)-1 {
				query = query + "'" + fieldValue + "'"
			} else {
				query = query + "'" + fieldValue + "', "
			}
		} else {
			if modelName == "Student" || modelName == "Teacher" {
				query = query + fieldValue + ", "
			}
		}
	}
	query = query + ")"
	return query
}

func createGetQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	query := "SELECT "

	for i, fieldName := range fieldsNames {
		if i == len(fieldsNames)-1 {
			query = query + fieldName
		} else {
			query = query + fieldName + ", "
		}
	}

	filterField := filterField(tableName)
	if filterField == "ID" {
		query = query + " FROM " + tableName + " WHERE ID=" + fieldsValues[0]
	} else {
		query = query + " FROM " + tableName + " WHERE " + filterField + "=" + fieldsValues[1]
	}

	return query
}

func createGetAllQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	query := "SELECT "

	if tableName == "Students_Courses" {
		query = query + "ID, Name, Email FROM Students WHERE ID IN (SELECT StudentID FROM Students_Courses WHERE CourseID=" + fieldsValues[1] + ")" //no siempre puede ser la posición 1
	} else {
		for i, fieldName := range fieldsNames {
			if i == len(fieldsNames)-1 {
				query = query + fieldName
			} else {
				query = query + fieldName + ", "
			}
		}
		query = query + " FROM " + tableName
	}

	return query
}

func createUpdateQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	query := "UPDATE " + tableName + " SET "

	if tableName == "Courses_CCases" || tableName == "Courses_HCN" {
		return createUpdateVisibilityQuery(model)
	} else if tableName == "Solved_HCN" {
		return query + fieldsNames[6] + "=" + fieldsValues[6] + " WHERE ActivityID=" + fieldsValues[1] + " AND Solver=" + fieldsValues[2]
	}

	for i, fieldName := range fieldsNames {
		if i == len(fieldsNames)-1 {
			query = query + fieldName + "='" + fieldsValues[i] + "' "
		} else if i != 0 {
			query = query + fieldName + "='" + fieldsValues[i] + "', "
		}
	}
	query = query + "WHERE ID=" + fieldsValues[0]
	return query
}

func createDeleteQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	filterField := filterField(tableName)
	query := "DELETE"
	if tableName == "Solved_HCN" {
		query = query + " FROM " + tableName + " WHERE " + fieldsNames[1] + "=" + fieldsValues[1]
	} else if filterField == "ID" {
		query = query + " FROM " + tableName + " WHERE ID=" + fieldsValues[0]
	} else {
		query = query + " FROM " + tableName + " WHERE " + fieldsNames[1] + "=" + fieldsValues[1] + " AND " + fieldsNames[2] + "=" + fieldsValues[2]
	}
	return query
}

func createUpdateVisibilityQuery(model interface{}) string {
	modelName, fieldsNames, fieldsValues, _ := helper.GetFields(model)
	tableName := tableName(modelName)
	query := "UPDATE " + tableName + " SET Displayable=" + fieldsValues[3] + " WHERE " + fieldsNames[1] + "=" + fieldsValues[1] + " AND " + fieldsNames[2] + "=" + fieldsValues[2]
	return query
}
