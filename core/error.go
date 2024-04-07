package core

// Pf: panic forever
func Pf(err error) {
	if err != nil {
		panic(err)
	}
}
