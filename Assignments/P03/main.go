package main

import (
	"fmt"

	"github.com/Vizemo/imgMod/Colors"
	"github.com/Vizemo/imgMod/GetPic"
	"github.com/Vizemo/imgMod/Grayscale"
	"github.com/Vizemo/imgMod/Text"
)

func main() {
	fmt.Println("Victor Marchesi")
	fmt.Println("Program 3 - Image Ascii Art")

	// Prints black space
	fmt.Print("\n")

	// Call function to get picture from URL
	GetPic.GetPicture()

	Colors.GetColors()

	Grayscale.Grayscale()

	Text.MakePic()

}
