package stringscmp

import (
	"reflect"
)

func containsStringArrays(small []string, large []string) (eq bool) {
	if len(large) == 0 {
		return
	}

	eq = true
	for i := range small {
		var eq1 bool
		for j := range large {
			eq1 = eq1 || reflect.DeepEqual(small[i], large[j])
		}
		eq = eq && eq1
	}

	return
}

func compareStringArrays(a []string, b []string) (eq bool) {
	if len(a) != len(b) {
		return
	}

	for i := range b {
		for j := range a {
			eq = reflect.DeepEqual(b[i], a[j])
			if eq {
				break
			}
		}
		if !eq {
			break
		}
	}
	return
}
