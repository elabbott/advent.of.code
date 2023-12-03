package global

func StringToArray(input string) []string {
	arrayResult := []string{}
	for _, char := range input {
		arrayResult = append(arrayResult, string(char))
	}
	return arrayResult
}
