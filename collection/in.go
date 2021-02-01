package collection

func InStr(find string, arr []string) bool {
	for _, s := range arr {
		if s == find {
			return true
		}
	}
	return false
}

func InInt(find int, arr []int) bool {
	for _, s := range arr {
		if s == find {
			return true
		}
	}
	return false
}

// 多元化判断
func InSlice(need, haystack interface{}) bool {
	switch need.(type) {
	case int:
		for _, v := range haystack.([]int) {
			if v == need {
				return true
			}
		}
	case int64:
		for _, v := range haystack.([]int64) {
			if v == need {
				return true
			}
		}
	case string:
		for _, v := range haystack.([]string) {
			if v == need {
				return true
			}
		}
	default:
		return false
	}
	return false
}
