package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
// MY Solution:
func DecToBase(dec, base int) string {
	var values = "0123456789ABCDEF"
	var res string

	for dec != 0 {
		m := dec % base
		dec = dec / base
		res = string(values[m]) + res //instructor also suggested: res = fmt.Sprintf("%X%s", m, res)
	}
	return res
}

/*
// My Psuedo-code:
if dec / base == 0 {
    return dec
}
for d = dec; d != 0 {
    d % base = n
    d / base = d
    result = string(n) + result
}
overBaseTen := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

if base > 10 {
    for d = dec; d != 0 {
        d % base = n
        d / base = d
        result = overBaseTen[n] + result
    }
}
*/

/*
// I used this some Printf statements
// in GoPlayground for clarity:
func DecToBase(dec, base int) string {
	var res string
	d := dec
	for d != 0 {
		n := d % base
		fmt.Printf("n: %d\n", n)
		d = d / base
		fmt.Printf("d: %d\n\n", d)
		res = fmt.Sprint(n) + res
	}
	return res
}

*/
