package main

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"path"
	"strings"

	// "Underscore Import" the package "image/jpeg" for the side effect of
	// format registration. Will allow "jpeg" format to be decoded as
	// an unknown image format in the convertToPNG function
	_ "image/jpeg"
	// Enable two lines below to allow gif and png files to also be
	// decoded as unknown image formats in the convertToPNG function
	// _ "image/gif"
	// _ "image/png"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	// Decode decodes an image that has been encoded in a registered format.
	// The string returned is the format name used during format registration.
	// Format registration is typically done by an init function in the codec-
	// specific package.
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	// Encode the image to png and write it to w
	return png.Encode(w, img)
}

func main() {
	filename := "image.jpg"

	// Open an image file
	img, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Open file error:", err)
	}
	defer img.Close()

	// Get the filename's extension
	extension := path.Ext(filename) // .jpg

	// Create a filename to save the png to
	pfilename := strings.TrimSuffix(filename, extension) + ".png"

	var bb bytes.Buffer
	if err := convertToPNG(&bb, img); err != nil {
		log.Fatalln(err)
	}

	pf, err := os.Create(pfilename)
	if err != nil {
		log.Fatalln("Create file error:", err)
	}
	defer pf.Close()

	_, err = pf.Write(bb.Bytes())
	if err != nil {
		log.Fatalln(err)
	}
}
