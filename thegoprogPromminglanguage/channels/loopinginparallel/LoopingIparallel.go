package main

import (
	"gopl.io/ch8/thumbnail"
	"log"
)

func main() {
	filenames := []string{
		"image1.jpg",
		"image2.png",
		"image3.jpeg",
	}
	makeThumbnails(filenames)

}
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Print(err)
		}
	}
}
