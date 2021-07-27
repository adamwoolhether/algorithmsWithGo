package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"

// My method. Passes, unless special chars are used like 中文
// Use the runes solution below to handlrunses.
/*
func Reverse(word string) string {
	var a strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		a.WriteString(string(word[i]))
	}
	return a.String()
}
*/

/*
// My way seems better than this one:
func Reverse(word string) string {
	var rev string
	for i := len(word) - 1; i >= 0; i-- {
		rev = rev + string(word[i])
	}
	return rev
}
*/
// Same as above, but build backwards
/*
func Reverse(word string) string {
var rev string
for i := 0; i < len(word); i++ {
	res = string(word[i]) + rev
}
return rev
*/


// Build with runes (I didn't think of this!)
// This handles more non-US charsets
func Reverse(word string) string {
	var rev string
	for _, v := range word {
		rev = string(v) + rev
	}
	return rev
}

