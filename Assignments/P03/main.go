package main

import (
	"fmt"

	//"github.com/vizemo/imgMod/Colors"
	"github.com/vizemo/imgMod/GetPic"
	//"github.com/vizemo/imgMod/Grayscale"
	//"github.com/vizemo/imgMod/Text"
)

func main() {
	fmt.Println("Victor Marchesi")
	fmt.Println("Program 3 - Image Ascii Art")

	// Prints black space
	fmt.Print("\n")

	// Call function to get picture from URL
	GetPic.DownloadPicture()

}
