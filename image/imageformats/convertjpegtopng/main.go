// The standard package library supports a number of common image formats,
// such as GIF, JPEG and PNG. If you know the format of a source image
// file, you can decode from an io.Reader directly.
package main

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

// ConvertJPEGToPNG converts from JPEG to PNG.
func convertJPEGToPNG(w io.Writer, r io.Reader) error {
	img, err := jpeg.Decode(r)
	if err != nil {
		return err
	}

	// Encode the image to png and write it to w
	return png.Encode(w, img)
}

func main() {
	jfilename := "image.jpg"

	// Open a jpeg image file
	jf, err := os.Open(jfilename)
	if err != nil {
		log.Fatalln("Open file error:", err)
	}
	defer jf.Close()

	// Get the filename's extension
	extension := path.Ext(jfilename) // .jpg

	// Create a filename to save the png to
	pfilename := strings.TrimSuffix(jfilename, extension) + ".png"

	var bb bytes.Buffer
	if err := convertJPEGToPNG(&bb, jf); err != nil {
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
