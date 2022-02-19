package utils

import (
	"bytes"
	"os"
	"time"
	"unicode"
)

func PString(str string) *string {
	return &str
}

func PFloat64(f float64) *float64 {
	return &f
}

func PInt(i int) *int {
	return &i
}

func PTime(t time.Time) *time.Time {
	return &t
}

func PBool(b bool) *bool {
	return &b
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func OnlyDigits(str string) string {
	buf := bytes.NewBufferString("")
	for _, r := range str {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}
