package parser

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type ImageUser struct {
	ImageFile string
	UserId    string
	UserName  string
}

func ProcessImageDir(dir string, index ...int) ([]*ImageUser, error) {

	fmt.Printf("processing directory %v ...\n", dir)
	if strings.TrimSpace(dir) == "" {
		return nil, errors.New("directory name is empty")
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("empty images dir")
	}

	path := fmt.Sprintf("%v%v", filepath.Dir(dir), string(os.PathSeparator))
	fmt.Println("Path", path)
	i := 1
	if index != nil && len(index) > 0 {
		i = int(math.Max(float64(1), float64(index[0])))
	}

	imagePaths := make([]*ImageUser, len(files))
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		id := fmt.Sprintf("%05d", i)
		from := fmt.Sprintf("%s%s", path, file.Name())
		to := fmt.Sprintf("%simg.%s%s", path, id, ext)

		if from != to {
			fmt.Printf("Renaming %v to %v\n", from, to)
			err = os.Rename(from, to)
		} else {
			fmt.Printf("Not renaming - names of the files are the same %v\n", to)
		}
		if err != nil {
			return nil, err
		}
		imagePaths[i-1] = &ImageUser{
			ImageFile: to,
			UserId:    fmt.Sprintf("itu-%s", id),
			UserName:  fmt.Sprintf("ImageTestUser-%s", id),
		}
		i++
	}

	return imagePaths, nil
}

func (m *ImageUser) GetImageFile() string {
	if m != nil {
		return m.ImageFile
	}
	return ""
}

func (m *ImageUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ImageUser) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}
