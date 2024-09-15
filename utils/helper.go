package utils

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot load file .env: ", err)
		panic(err)
	}

	value := GetEnvOrDefault(key, "").(string)
	return value
}

func GetEnvOrDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func WriteStringTemplate(stringTemplate string, args ...interface{}) string {
	return fmt.Sprintf(stringTemplate, args...)
}

func ExpectedInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}

func ExpectedUint(v interface{}) uint {
	var result uint
	switch v := v.(type) {
	case int:
		result = uint(v)
	case float64:
		result = uint(v)
	case string:
		convertedString, _ := strconv.ParseUint(v, 10, 32)
		result = uint(convertedString)
	case uint:
		result = v
	}
	return result
}

func ExpectedString(v interface{}) string {
	var result string
	switch v := v.(type) {
	case int, uint:
		result = fmt.Sprintf("%d", v)
	case float64:
		result = fmt.Sprintf("%f", v)
	case string:
		result = v
	}
	return result
}

func BoolPointerToBool(value *bool) bool {
	newValue := false

	if value != nil {
		newValue = *value
	}

	return newValue
}

func BoolToBoolPointer(value bool) *bool {
	newValue := new(bool)

	newValue = &value

	return newValue
}

func TimeToTimePointer(value time.Time) *time.Time {
	newValue := new(time.Time)

	newValue = &value

	return newValue
}

func IntToIntPointer(value int) *int {
	newValue := new(int)

	newValue = &value

	return newValue
}

func UintToUintPointer(value uint) *uint {
	newValue := new(uint)

	newValue = &value

	return newValue
}

func UintPointerToUint(value *uint) uint {
	return *value
}

func Uint8ToUint8Pointer(value uint8) *uint8 {
	newValue := value
	if value <= 0 {
		newValue = 0
	}

	return &newValue
}

func JSONUnmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		return err
	}
	return nil
}

func JSONMarshal(data interface{}) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ToByte(i interface{}) []byte {
	byte_, _ := JSONMarshal(i)
	return byte_
}

func Dump(i interface{}) string {
	return string(ToByte(i))
}

func MyCaller(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return details.Name()
	}

	return "failed to identify method caller"
}

// CalculatePages function that takes in the total number of items and the number of items per page,
// and calculates the number of pages required to fit all the items.
// it returns the number of pages as an integer
func CalculatePages(total, size uint) int {
	return int(math.Ceil(float64(total) / float64(size)))
}

func StringToLower(_string string) string {
	return strings.ToLower(_string)
}

func StringToUpper(s string) string {
	return strings.ToUpper(s)
}
