package main

import (
	"ascii-art-reverse/pkg"
	"os"
	"strings"
	"testing"
)

func TestArgsList(t *testing.T) {
	testStrings := []string{"--align=justify", "hello", "shadow"}
	align, outputfile, reverse, color, strToBeColored, str, banner, err := pkg.ArgsList(testStrings)
	if align != "justify" || outputfile != "" || reverse != "" || color != nil || strToBeColored != nil || str != "hello" || banner != "shadow" || err != nil {
		t.Error("Error: there is problem in ArgsList function!")
	}
}
func TestCheckColor(t *testing.T) {
	testStrings := []string{"red", "blue", "yellow"}
	if pkg.CheckColor(testStrings) != true {
		t.Error("Error: there is problem in CheckColor function!")
	}
}
func TestCheckString(t *testing.T) {
	if pkg.CheckString("h") != true {
		t.Error("Error: there is problem in CheckString function!")
	}
}
func TestColorName(t *testing.T) {
	testStrings := "--color=red"
	res, err := pkg.ColorName(testStrings)
	if res != "red" || err != nil {
		t.Error("Error: there is problem in ColorName function!")
	}
}
func TestColorStringMap(t *testing.T) {
	color := []string{"red", "blue", "yellow"}
	strcol := []string{"l", "r", "n"}
	res := pkg.ColorStringMap(strcol, color)
	if res["l"] != "red" || res["r"] != "blue" || res["n"] != "yellow" {
		t.Error("Error: there is problem in ColorStringMap function!")
	}
}
func TestGetAscii(t *testing.T) {
	lines := strings.Split(pkg.GetBannerFile("banners/standard.txt"), "\n")
	if pkg.GetAscii('h', 1, lines) != " _      " {
		t.Error("Error: there is problem in GetAscii function!")
	}
}
func TestGetBannerFile(t *testing.T) {
	file, err := os.ReadFile("banners/standard.txt")
	text := string(file)
	if pkg.GetBannerFile("banners/standard.txt") != text || err != nil {
		t.Error("Error: there is problem in GetBannerFile function!")
	}
}
func TestGetColor(t *testing.T) {
	if pkg.GetColor("red") != "\033[1;31m" {
		t.Error("Error: there is problem in GetColor function!")
	}
}
func TestConverter(t *testing.T) {
	h, s, l := pkg.StringsToFloat64("hsl(0, 100%, 50%)")
	r, g, b := pkg.HSLToRGB(h, s, l)
	hexcode := pkg.RGBToHex(int(r), int(g), int(b))
	if pkg.HexToColor(hexcode) != "red" {
		t.Error("Error: there is problem in Converter functions!")
	}
}
func TestStringsToInt(t *testing.T) {
	testStrings := "rgb(255, 0, 0)"
	r, g, b := pkg.StringsToInt(testStrings)
	if r != 255 || g != 0 || b != 0 {
		t.Error("Error: there is problem in StringsToInt function!")
	}
}

func TestLenOfStr(t *testing.T) {
	lines := strings.Split(pkg.GetBannerFile("banners/standard.txt"), "\n")
	text := "h"
	if pkg.LenOfStr(text, lines) != 8 {
		t.Error("Error: there is problem in LenOfStr function!")
	}
}
