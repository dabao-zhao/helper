package any

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ToString 将 value 转换为字符串
func ToString(value any) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func ToInt(value any) (int, error) {
	switch value.(type) {
	case int:
		return value.(int), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		return strconv.Atoi(value.(string))
	}
	return 0, errors.New("value error")
}

func ToInt32(value any) (int32, error) {
	switch value.(type) {
	case int:
		return int32(value.(int)), nil
	case int32:
		return value.(int32), nil
	case int64:
		return int32(value.(int64)), nil
	case float64:
		return int32(value.(float64)), nil
	case float32:
		return int32(value.(float32)), nil
	case string:
		val, err := strconv.ParseInt(value.(string), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(val), nil
	}
	return 0, errors.New("value error")
}

func ToFloat64(value any) (float64, error) {
	switch value.(type) {
	case float64:
		return value.(float64), nil
	case int:
		return float64(value.(int)), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	}
	return 0, errors.New("value error")
}
