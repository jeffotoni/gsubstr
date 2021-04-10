package gsubstr

// substr (text, 0, 4)
// substr (text, 2, 2)
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
		lenf = nums[0]
		if lenf < 0 && leni == 0 {
			tmp := -lenf
			lenf = len(value) - tmp
		} else if lenf < 0 && leni > 0 {
			tmp := -lenf
			lenf = len(value) - tmp - leni
		} else if lenf < 0 && leni < 0 {
			tmpf := -lenf
			tmpi := -leni
			//tmp3 := 0

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

			// if tmp > tmp2 {
			// 	tmp3 = tmp - tmp2
			// } else {
			// 	tmp3 = tmp2 - tmp
			// }

			lenf = tmp3
			println("lenf:", lenf, " - tmpi:", tmpi, "tmpf:", tmpf)
		}
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

	// if leni < 0 || lenf < 0 {
	// 	return ""
	// }

	println(":::", leni, " ::::", lenf)
	if len(value) < leni {
		return ""
	}

	leny := lenf
	soma := leni + lenf
	if soma > len(value) {
		soma = len(value)
	}
	leny = soma
	//println(value[leni:leny])
	return value[leni:leny]
}
