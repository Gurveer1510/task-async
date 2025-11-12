package sqlparser

func ParseSqlFile(fileContent string) []string {
	var result []string
	temp := ""
	for i := 0; i < len(fileContent); i++ {
		char := string(fileContent[i])
		if char == ";"{
			result = append(result, temp)
			temp = ""
		} else {
			temp += char
		}
	}

	if temp != "" {
		result = append(result, temp)
	}
	return result
}
