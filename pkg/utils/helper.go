package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func ConvertStringTimetoTime(t string) time.Time {
	layout := "2006-01-02T15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error - Parse time failed: %s", err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")
	return result.In(loc)
}

func ConvertStringToSqlNullString(s string) sql.NullString {
	var sqlNullString sql.NullString
	if s != "" {
		sqlNullString.String = s
		sqlNullString.Valid = true
	} else {
		sqlNullString.Valid = false
	}
	return sqlNullString
}