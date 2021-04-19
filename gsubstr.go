// Package gsubstr responsible for taking parts of the string as needed
// A very simple lib to use and very useful in our day to day
//
// By @jeffotoni
package gsubstr

// Substr function responsible for slicing a string according to its limits
func Substr(value string, leni int, nums ...int) string {
	var lenf int
	if len(nums) > 1 {
		return ""
	}
	if len(nums) == 0 {
		if leni < 0 {
			lenf = -(leni)
			if lenf > len(value) {
				return ""
			}
		} else {
			lenf = len(value)
		}
	} else {
		lenf = lenfLast(value, nums[0], leni)
	}

	if leni < 0 {
		tmp0 := -leni
		if tmp0 > len(value) {
			leni = 0
		} else {
			tmp := len(value) + leni
			leni = tmp
		}
	}

	if len(value) < leni {
		return ""
	}

	soma := leni + lenf
	if soma > len(value) {
		soma = len(value)
	}
	return value[leni:soma]
}

func lenfLast(value string, lenf, leni int) int {
	if lenf < 0 && leni == 0 {
		tmp := -lenf
		lenf = len(value) - tmp
	} else if lenf < 0 && leni > 0 {
		tmp := -lenf
		lenf = len(value) - tmp - leni
	} else if lenf < 0 && leni < 0 {
		tmpf := -lenf
		tmpi := -leni
		if tmpi >= len(value) {
			tmpi = 0
		} else {
			tmpi = len(value) - tmpi
		}
		if tmpf > len(value) {
			tmpf = len(value) - tmpf
		}
		tmp3 := 0
		soma := tmpf + tmpi
		if soma > len(value) {
			tmp3 = 0
		} else {
			tmp3 = len(value) - tmpf - tmpi
		}
		lenf = tmp3
	}
	return lenf
}
