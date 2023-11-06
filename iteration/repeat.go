package iteration

// Repeat generates a string comprised of 'character' repeated 'count' times
func Repeat(character string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
