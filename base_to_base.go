package module01

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
// This one was easy, just a matter of combining the two previous
func BaseToBase(value string, base, newBase int) string {
	//var res string
	//dec := 0
	//multiplier := 1
	//
	//for i := len(value) - 1; i >= 0; i-- {
	//	var val int
	//	fmt.Sscanf(string(value[i]), "%X", &val)
	//	dec += multiplier * val
	//	multiplier *= base
	//}
	//
	//for dec != 0 {
	//	m := dec % newBase
	//	dec = dec / newBase
	//	res = fmt.Sprintf("%X%s", m, res)
	//}
	//
	//return res

	//Or, just simple call the other functions:
	return DecToBase(BaseToDec(value, base), newBase)
}
