package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

func ConvertIntToSqlNullInt32(s int) sql.NullInt32 {
	var sqlNullInt32 sql.NullInt32
	if s != 0 {
		sqlNullInt32.Int32 = int32(s)
		sqlNullInt32.Valid = true
	} else {
		sqlNullInt32.Valid = false
	}
	return sqlNullInt32
}
func ConvertInt32ToSqlNullInt32(s int32) sql.NullInt32 {
	var sqlNullInt32 sql.NullInt32
	if s != 0 {
		sqlNullInt32.Int32 = s
		sqlNullInt32.Valid = true
	} else {
		sqlNullInt32.Valid = false
	}
	return sqlNullInt32
}
func GetFileExtension(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		return ("." + parts[len(parts)-1])
	}
	return ""
}
