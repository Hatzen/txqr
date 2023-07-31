package txqr

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"

	"github.com/Hatzen/txqr/qr"
)

// jni interface for binary data

/*
func createGifSmallAndFast(data []byte) []byte {
	createAnimatedGif(data, 300, 5, 100, qr.Medium)
}

func createGif(data []byte) []byte {
	createAnimatedGif(data, 700, 2, 500, qr.Medium)
}

func createGifLargerAndSlower(data []byte) []byte {
	createAnimatedGif(data, 1000, 1, 1000, qr.Medium)
}
*/

// jni interface for string data

func createGifSmallAndFast(data string) []byte {
	result, _ := createAnimatedGif(data, 300, 5, 100, qr.Medium)
	return result
}

func createGif(data string) []byte {
	result, _ := createAnimatedGif(data, 700, 2, 500, qr.Medium)
	return result
}

func createGifLargerAndSlower(data string) []byte {
	result, _ := createAnimatedGif(data, 1000, 1, 1000, qr.Medium)
	return result
}

// Creation Functions.

/*
func createAnimatedGif(data []byte, imgSize int, fps, size int, lvl qr.RecoveryLevel) ([]byte, error) {
	str := base64.StdEncoding.EncodeToString(data) // TODO: It might be very specific to encode and decode, maybe better do i
	return AnimatedGif(str, imgSize, fps, size, lvl)
}
*/

func createAnimatedGif(data string, imgSize int, fps, size int, lvl qr.RecoveryLevel) ([]byte, error) {
	// str := base64.StdEncoding.EncodeToString(data)
	// str := string(data[:])
	chunks, err := NewEncoder(size).Encode(data)
	if err != nil {
		return nil, fmt.Errorf("encode: %v", err)
	}

	out := &gif.GIF{
		Image: make([]*image.Paletted, len(chunks)),
		Delay: make([]int, len(chunks)),
	}
	for i, chunk := range chunks {
		qr, err := qr.Encode(chunk, imgSize, lvl)
		if err != nil {
			return nil, fmt.Errorf("QR encode: %v", err)
		}
		out.Image[i] = qr.(*image.Paletted)
		out.Delay[i] = fpsToGifDelay(fps)
	}

	var buf bytes.Buffer
	err = gif.EncodeAll(&buf, out)
	if err != nil {
		return nil, fmt.Errorf("gif create: %v", err)
	}
	return buf.Bytes(), nil
}

// fpsToGifDelay converts fps value into animated GIF
// delay value, which is in 100th of second
func fpsToGifDelay(fps int) int {
	if fps == 0 {
		return 100 // default value, 1 sec
	}
	return 100 / fps
}
