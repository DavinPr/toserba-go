package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func Init(filename, configType string) {
	viper.AutomaticEnv()
	viper.SetConfigName(filename)
	viper.SetConfigType(configType)
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.ReadInConfig()
}

func MustGetInt(key string) int {
	v, err := strconv.Atoi(MustGetString(key))
	mustParseKey(err, key)
	return v
}

func MustGetInt64(key string) int64 {
	v, err := strconv.Atoi(MustGetString(key))
	mustParseKey(err, key)
	return int64(v)
}

func MustGetFloat32(key string) float32 {
	v, err := strconv.ParseFloat(MustGetString(key), 32)
	mustParseKey(err, key)
	return float32(v)
}

func MustGetFloat64(key string) float64 {
	v, err := strconv.ParseFloat(MustGetString(key), 64)
	mustParseKey(err, key)
	return v
}

func MustGetString(key string) string {
	mustHaveKey(key)
	return GetString(key)
}

func MustGetUint(key string) uint {
	v, err := strconv.ParseUint(MustGetString(key), 10, 0)
	mustParseKey(err, key)
	return uint(v)
}

func MustGetJSON(key string, val interface{}) error {
	strValue := MustGetString(key)
	err := json.Unmarshal([]byte(strValue), val)
	return err
}

func MustGetDuration(key string) time.Duration {
	d, err := cast.ToDurationE(MustGetString(key))
	mustParseKey(err, key)
	return d
}

func GetInt(key string) int {
	v, err := strconv.Atoi(GetString(key))
	if err != nil {
		return 0
	}
	return v
}

func GetInt64(key string) int64 {
	v, err := strconv.ParseInt(GetString(key), 10, 64)
	if err != nil {
		return int64(0)
	}
	return v
}

func GetString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func GetStringSlice(key string) []string {
	value := GetString(key)
	if value == "" {
		return []string{}
	}
	return strings.Split(value, ",")
}

func GetIntSlice(key string) []int {
	strValues := GetStringSlice(key)
	if len(strValues) == 0 {
		return []int{}
	}

	var retValues []int
	for _, str := range strValues {
		if val, err := strconv.Atoi(str); err != nil {
			return []int{}
		} else {
			retValues = append(retValues, val)
		}
	}
	return retValues
}

func GetInt64Slice(key string) []int64 {
	strValues := GetStringSlice(key)
	if len(strValues) == 0 {
		return []int64{}
	}

	var retValues []int64
	for _, val := range strValues {
		if v, err := strconv.ParseInt(val, 10, 64); err != nil {
			return []int64{}
		} else {
			retValues = append(retValues, v)
		}

	}
	return retValues
}

func GetFloat32Slice(key string) []float32 {
	strValues := GetStringSlice(key)
	if len(strValues) == 0 {
		return []float32{}
	}

	var retValues []float32
	for _, val := range strValues {
		if v, err := strconv.ParseFloat(val, 32); err != nil {
			return []float32{}
		} else {
			retValues = append(retValues, float32(v))
		}

	}
	return retValues
}

func GetFloat64Slice(key string) []float64 {
	strValues := GetStringSlice(key)
	if len(strValues) == 0 {
		return []float64{}
	}

	var retValues []float64
	for _, val := range strValues {
		if v, err := strconv.ParseFloat(val, 64); err != nil {
			return []float64{}
		} else {
			retValues = append(retValues, v)
		}

	}
	return retValues
}

func GetFeature(key string) bool {
	value := GetString(key)
	if value == "" {
		return false
	}

	boolVal, err := strconv.ParseBool(value)
	mustParseKey(err, key)
	return boolVal
}

func GetUint(key string) uint {
	value, err := strconv.ParseUint(GetString(key), 10, 0)
	if err != nil {
		return 0
	}
	return uint(value)
}

func GetJSON(key string, val interface{}) error {
	strValue := GetString(key)
	if strValue == "" {
		return nil
	}

	err := json.Unmarshal([]byte(strValue), val)
	return err
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(GetString(key))
}

func mustHaveKey(key string) {
	if !hasKey(key) {
		log.Fatalf("%s key is not set", key)
	}
}

func mustParseKey(err error, key string) {
	if err != nil {
		log.Fatalf("Could not parse key: %s, error: %s", key, err)
	}
}

func hasKey(key string) bool {
	return viper.IsSet(key) || os.Getenv(key) != ""
}
