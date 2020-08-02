package converter

import (
	"fmt"
	"strconv"
	"strings"
)

// Converter defines y = f(x)
type Converter func(value interface{}) interface{}


// SplitWith ...
func SplitWith(sep string) Converter {
	return func(value interface{}) interface{} {
		if valueStr, ok := value.(string); ok {
			return strings.Split(valueStr, sep)
		} else {
			return []string{}
		}
	}
}
// Convert2Bool ...
func Convert2Bool(next Converter) Converter {
	return func(value interface{}) interface{} {
		if next != nil {
			value = next(value)
		}
		result := 0
		if v, ok := value.(int); ok {
			result = v
		} else if v, ok := value.([]int); ok {
			if len(v) >= 1 {
				result = v[0]
			}
		} else if v, ok := value.(string); ok {
			vInt, err := strconv.Atoi(v)
			if err != nil && len(v) > 0 {
				return false
			}
			result = vInt
		}

		return result > 0
	}
}

// Convert2Int ...
func Convert2Int(next Converter) Converter {
	return func(value interface{}) interface{} {
		// check nesting converter
		if next != nil {
			value = next(value)
		}
		if vs, ok := value.(string); ok {
			i, err := strconv.Atoi(vs)
			if nil != err && len(vs) > 0 {
				return 0
			}
			return i
		}
		return 0
	}
}

// Convert2Set ...
func Convert2StringSet(next Converter) Converter {
	return func(value interface{}) interface{} {
		result := make(map[string]byte)
		// check nesting converter
		if next != nil {
			value = next(value)
		}
		fmt.Println(value)
		if nil != value {
			if vs, ok := value.([]int32) ; ok {
				for _, v := range vs {
					result[strconv.Itoa(int(v))] = 0
				}
			} else if vs, ok := value.([]int64) ; ok {
				for _, v := range vs {
					result[strconv.FormatInt(v, 10)] = 0
				}
			} else if vs, ok := value.([]string) ; ok {
				for _, v := range vs {
					if len(v)>0 {
						result[v] = 0
					}
				}
			}
		}
		return result
	}
}

