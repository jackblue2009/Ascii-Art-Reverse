package pkg

/*
Get the coloring code to use it in the PrintAsciiArtWord function
*/

func GetColor(color string) string {
	colors := map[string]string{
		"default": "\033[0m",
		"red":     "\033[1;31m",
		"green":   "\033[0;32m",
		"yellow":  "\033[1;33m",
		"blue":    "\033[0;34m",
		"purple":  "\033[0;35m",
		"cyan":    "\033[0;36m",
		"gray":    "\033[0;37m",
		"white":   "\033[1;37m",
		"orange":  "\033[0;38;5;208m",
		"black":   "\033[0;30m",
	}
	return colors[color]
}
