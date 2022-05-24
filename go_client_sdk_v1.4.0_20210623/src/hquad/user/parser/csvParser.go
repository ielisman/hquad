package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ParseUsersFile(fileName string, imageDir string) (map[string]map[string]string, error) {

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(imageDir)) == 0 {
		imageDir = filepath.Dir(fileName)
		fmt.Printf("Image Dir with user images '%s' is set from users file '%s'\n", imageDir, fileName)
	}

	recs, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nProcessing users file '%s'\n", fileName)
	headerMap := make(map[int]string)
	recordsMap := make(map[string]map[string]string)
	for i, r := range recs {
		if i == 0 { // header row
			for j, c := range r {
				headerMap[j] = c // headerMap[1] = "stu-715"
			}
		} else { // rows of records
			doNotSkip := true
			rowMap := make(map[string]string)
			var recordId string
			for j, c := range r {
				if headerMap[j] == "user_id" {
					recordId = c // "stu-715"
				}
				if strings.Index(headerMap[j], "face_image_file") >= 0 { // verify if file exists
					imgPath := fmt.Sprintf("%s%c%s", imageDir, os.PathSeparator, c)
					// err == nil // errors.Is(err, fs.ErrNotExist) // os.IsNotExist(err)
					if _, err := os.Stat(imgPath); errors.Is(err, fs.ErrNotExist) {
						fmt.Printf(" Image file %s does not exist. skipping record %s\n", imgPath, r)
						doNotSkip = false
						break
					}
				}
				// rowMap{user_id} = "stu-715"
				rowMap[headerMap[j]] = c
			}
			// recordsMap{user_id} = rowMap{ name:HQTest-User-715, user_id:stu-715, face_image_file1:715.jpg, ...}
			if doNotSkip {
				recordsMap[recordId] = rowMap
			}

		}
	}
	fmt.Printf("Finished processing users file '%s' with %d records\n\n", fileName, len(recs)-1)

	return recordsMap, err

}

/*recordsMap, err := parser.ParseUsersFile(*usersFile, *imageDir)
if err != nil {
	fmt.Printf("Error retrieving users file: %v\n", err)
	os.Exit(0)
}

if len(recordsMap) < 0 {
	for k, m := range recordsMap {
		fmt.Printf("Working on record with user id %s\n", k)
		for c, v := range m {
			fmt.Printf(" %-30s : %-30s\n", c, v)
		}
	}
}*/
