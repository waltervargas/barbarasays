package colombian

var (
	// decir is not exported, is private to the colombian package.
	decir string = "Malpariculilambido!"
	// decir2 string = "Arepa con Quesito!"
)

// this is an exported function Dice because
// starts with a Capital letter.
func Dice() string {
	return decir
}
