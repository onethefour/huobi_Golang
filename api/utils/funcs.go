package utils

import (
	"fmt"
	"strconv"
)

//<code class="go hljs">
func ParseToString(val interface{}) string {
	switch t := val.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", t)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%.8f", t)
	case string:
		return t
	default:
		return "0"
		//panic(fmt.Errorf("invalid value type", t))
	}
}

//Digits 保留有效数字
func Digits(number float64, num uint64) float64 {
	number_str := fmt.Sprintf("%.8f", number)
	start := false
	n := uint64(0)
	l := len(number_str)
	posOfPoint := len(number_str)
	for i := 0; i < len(number_str); i++ {
		//log.Println(i)
		if number_str[i] == '.' {
			posOfPoint = i
			//log.Println(number_str[i], posOfPoint)
		}
		if !start {
			if number_str[i] != '0' && number_str[i] != '.' {
				start = true
			}
		}
		if start && number_str[i] != '.' {
			n++
			if n == num {
				l = i + 1
			}
		}

	}
	ret_str := number_str[0:l]
	//log.Println("ssssss", number_str, num, ret_str, posOfPoint, l)
	for i := l; i < posOfPoint; i++ {
		ret_str = ret_str + "0"
	}

	ret, _ := strconv.ParseFloat(ret_str, 64)
	//log.Println("bbbb", ret_str, ret)
	return ret
}
