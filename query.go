package dbquery

import (
	"database/sql"
	"errors"
	"strconv"
)

// Client - wtf is this? // TODO: add comment
type Client struct {
	sqlDB *sql.DB
}

/*
Connect - open a database connection.
dbDriver - mysql / pg / etc
*/
func (c *Client) Connect(dbDriver string, credentials DBCredentials) error {
	var err error
	dataSourceName := credentials.User + ":" + credentials.Password
	dataSourceName += "@" + credentials.Host + ":" + credentials.Port
	dataSourceName += "/" + credentials.DBName
	c.sqlDB, err = sql.Open(dbDriver, dataSourceName)
	if err != nil {
		return err
	}

	err = c.sqlDB.Ping()
	if err != nil {
		return errors.New("failed to ping db: " + err.Error())
	}
	return nil
}

func (c *Client) buildSQLQueryAdd(
	collection string, fields map[string]interface{}, returnField ...string,
) (string, []interface{}) {
	sqlQuery := "INSERT INTO " + collection + " "

	keyValues := make([]interface{}, len(fields))
	var i int = 1
	keysString := ""
	valuesString := ""
	for key, value := range fields {
		keyValues = append(keyValues, value)
		delimeter := ""
		if i > 0 {
			delimeter = ","
		}
		keysString += delimeter + key
		valuesString += delimeter + "$" + strconv.Itoa(i)
		i++
	}
	sqlQuery += "(" + keysString + ") VALUES (" + valuesString + ")"

	if len(returnField) > 0 {
		sqlQuery += " RETURNING " + returnField[0]
	}

	return sqlQuery, keyValues
}

// Add - writes a new line to the base
// fields - map with key-value
func (c *Client) Add(collection string, fields map[string]interface{}) error {
	sqlQuery, keyValues := c.buildSQLQueryAdd(collection, fields)
	return c.sqlDB.QueryRow(sqlQuery, keyValues...).Err()
}

// AddAndGet - writes a new line to the base
// fields - map with key-value
// returnField (optional) - returns the value of the specified field of the added record, if necessary
func (c *Client) AddAndGet(collection string, fieldsMap map[string]interface{}, returnField ...string) (string, error) {
	sqlQuery, keyValues := c.buildSQLQueryAdd(collection, fieldsMap)
	row := c.sqlDB.QueryRow(sqlQuery, keyValues...)
	err := row.Err()
	if err != nil {
		return "", err
	}

	var resultField string = ""
	err = row.Scan(&resultField)
	if err != nil {
		return "", err
	}

	return resultField, nil
}

// Get - request one or more values from the database.
// searchOpts must be 0 or 1 elements. You can use And, Or to combine multiple.
// returns: maps (key-value), error
func (c *Client) Get(
	collection string, fieldsMap map[string]interface{}, searchOpts ...SearchOption,
) ([]map[string]interface{}, error) {
	resultMap := []map[string]interface{}{}

	sqlQuery := ""               // TODO
	keyValues := []interface{}{} // TODO

	rows, err := c.sqlDB.Query(sqlQuery, keyValues...)
	if err != nil {
		return resultMap, err
	}
	resultArr := make([]interface{}, len(fieldsMap))
	for rows.Next() {
		err := rows.Scan(resultArr...) // test this. interface or pointer to interface?
		if err != nil {
			return resultMap, err
		}
		// TODO: add data to resultMap
	}

	return resultMap, nil
}
