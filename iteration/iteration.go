package iteration

const repeaCount = 5

func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < repeaCount; i++ {
		repeated += character
	}
	return repeated
}
