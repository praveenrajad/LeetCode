
func Encode(strs []string) string {
	var encode strings.Builder
	for _, str := range strs {
		encode.WriteString(strconv.Itoa(len(str)))
		encode.WriteString("/")
		encode.WriteString(str)
	}

	return encode.String()
}


func Decode(str string) []string {
	i := 0
	out := []string{}
	for i < len(str) {
		j := i
		for j  < len(str) && str[j] != "/" {
			j++
		}
		length,_ := strconv.Atoi(str[i:j])
		out = append(out, str[j+1:j+1+length])
		i = j+1+length
	}
	return out
}