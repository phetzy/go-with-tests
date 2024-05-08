package iteration

func Repeat(character string) string {
	var repeated string

	/*
		NOTE: For loop pre 1.22
		for i := 0; i < 5; i++ {
			repeated += character
		}
	*/

	// For loop 1.22 and later
	for i := range 5 {
		repeated += character
		i++
	}

	return repeated
}
