package iteration

func Repeat(char string, times int) string {
	var res string
	for i := 0; i < times; i++ {
		res += char
	}
	return res
}
