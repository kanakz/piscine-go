package piscine

func ArrLen(args []string) int {
	k := 0
	for range args {
		k++
	}
	return k
}

func ConcatParams(args []string) string {
	ans := ""
	for index, arg := range args {
		ans += arg
		if index != ArrLen(args)-1 {
			ans += "\n"
		}
	}
	return ans
}
