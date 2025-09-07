package hello

// Hello returns a greeting to the provided name
func Hello(name string) string {
	if name == "Oliver" {
		return "Hello, Oli"
	}
	return "Hello, " + name
}
