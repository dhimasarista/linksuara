package utility

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"

	"github.com/nfnt/resize"
)

func CompressFirst(filepath string) ([]byte, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Membaca file sebagai gambar
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	// Menyesuaikan Ukuran Gambar
	img = resize.Resize(1000, 0, img, resize.Lanczos3)

	// Mengonversi ke buffer
	var buffer bytes.Buffer
	err = jpeg.Encode(&buffer, img, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
