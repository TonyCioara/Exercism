package series

func All(num int, s string) []string {

	runes := []rune(s)
	returnStrings := []string{}

	for index := 0; index <= len(runes)-num; index++ {
		slice := runes[index : index+num]
		newString := string(slice)
		returnStrings = append(returnStrings, newString)
	}

	return returnStrings
}

func UnsafeFirst(num int, s string) string {

	runes := []rune(s)

	if num > len(runes) {
		return ""
	}

	slice := runes[0:num]
	newString := string(slice)

	return newString
}
