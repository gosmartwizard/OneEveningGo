package main

func Merge(a []string, b []string) []string {
	a = append(a, b...)

	return a
}

func Remove(a []string, index int) []string {

	len := len(a)

	if index == len-1 {
		return a[:index]
	} else if index == 0 {
		return a[1:]
	}

	a = append(a[:index], a[index+1:]...)

	return a
}

func RemoveLast(a []string) []string {
	return Remove(a, len(a)-1)
}
