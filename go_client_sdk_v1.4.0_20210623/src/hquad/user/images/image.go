package images

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

type FaceImage struct {
	ImgFile string
	ImgData []byte
}

func (m *FaceImage) GetImgFile() string {
	if m != nil {
		return m.ImgFile
	}
	return ""
}

func (m *FaceImage) GetImgData() []byte {
	if m != nil {
		return m.ImgData
	}
	return nil
}

func GetFaceImage(imgFile string) (*FaceImage, error) {

	fmt.Printf("getting images data from images %v ...\n", imgFile)
	reader, err := os.Open(imgFile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})
	jpeg.Encode(buf, img, nil) //buf := new(bytes.Buffer)
	return &FaceImage{imgFile, buf.Bytes()}, nil
}

func GetPNGImage(imgFile string) ([]byte, error) {

	fmt.Printf("getting images data from images %v ...\n", imgFile)
	reader, err := os.Open(imgFile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})
	jpeg.Encode(buf, img, nil) //buf := new(bytes.Buffer)
	return buf.Bytes(), nil
}
