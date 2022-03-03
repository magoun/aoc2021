package helpers

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RuneSlicePop(slice []rune) (rune, []rune) {
	length := len(slice)

	if length > 0 {
		popped := slice[length-1]
		slice = slice[:length-1]
		return popped, slice
	}

	panic(slice)
}
