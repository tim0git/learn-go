package iteration

const repeatCount = 5

// Repeat returns a string of characters
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
