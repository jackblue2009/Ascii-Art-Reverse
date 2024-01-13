package pkg

// By using Hex we select the color name
func HexToColor(color string) string {
	colors := map[string]string{
		"#FF0000": "red",
		"#00FF00": "green",
		"#0000FF": "blue",
		"#FFFF00": "yellow",
		"#800080": "purple",
		"#00FFFF": "cyan",
		"#808080": "gray",
		"#FFFFFF": "white",
		"#FFA500": "orange",
		"#000000": "black",
	}
	return colors[color]
}
