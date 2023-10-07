package main

import (
	"fmt"

	"github.com/Vizemo/img_mod/Colors"
	"github.com/Vizemo/img_mod/Getpic"
	"github.com/Vizemo/img_mod/Grayscale"
	"github.com/Vizemo/img_mod/Text"
)

func main() {
	fmt.Println("Victor Marchesi")
	fmt.Println("Program 3 - Image Ascii Art")

	// Prints blank space
	fmt.Print("\n")

	// Call function to get picture from URL
	Getpic.GetPicture()

	Colors.GetColor()

	Grayscale.Grayscale()

	Text.ColorText()

}
