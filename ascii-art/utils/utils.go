package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	hashStandard   = "a51f800619146db0c42d26db3114c99f"
	hashThinkertoy = "8efd138877a4b281312f6dd1cbe84add"
	hashShadow     = "d44671e556d138171774efbababfc135"
)

func AsciiConverter(w http.ResponseWriter, text string, name string) (string, int) {
	file := "ascii-art/files/" + name + ".txt"

	bb, err := readTxt(file)
	if err != nil {
		log.Println(err)
		return "", http.StatusInternalServerError
	}

	err = hashSum(file, name)
	if err != nil {
		log.Println(err)
		return "", http.StatusInternalServerError
	}

	splittedTxt := splitTxt(bb)

	txtWithNewlines := checkNewLine(text)

	checkedText := checkErr(txtWithNewlines)
	if checkedText != nil {
		return "", http.StatusBadRequest
	}

	return textToAscii(splittedTxt, txtWithNewlines), 0
}

func readTxt(s string) ([]byte, error) {
	text, err := os.ReadFile(s)
	if err != nil {
		return nil, errors.New("unable to open file/file does not exist")
	}
	return text, nil
}

func hashSum(path string, name string) error {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		return errors.New("unable to open file/file does not exist")
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		return err
	}
	hashSum := fmt.Sprintf("%x", h.Sum(nil))
	switch name {
	case "standard":
		if hashSum != hashStandard {
			return errors.New("font file is currupted")
		}
	case "thinkertoy":
		if hashSum != hashThinkertoy {
			return errors.New("font file is currupted")
		}
	case "shadow":
		if hashSum != hashShadow {
			return errors.New("font file is currupted")
		}
	}

	return nil
}

func splitTxt(b []byte) [][]string {
	text := [][]string{{}}
	element := strings.Split(string(b), "\n")
	count, j := 0, 0
	for _, word := range element {

		if count == 9 {
			count = 0
			text = append(text, []string{})
			j++
		}
		if len(word) != 0 {
			text[j] = append(text[j], word)
		}
		count++
	}
	return text
}

func checkErr(s []string) error {
	for _, h := range s[0] {
		if (h < 32 || h > 126) && h != 10 {
			return errors.New("non-printable argument")
		}
	}

	for _, j := range s {
		if j == "" && len(s) == 1 {
			return nil
		}
	}
	return nil
}

func checkNewLine(s string) []string {
	element := strings.ReplaceAll(s, "\r\n", "\n")
	str := strings.Split(element, "\n")

	return str
}

func textToAscii(bigtext [][]string, arg []string) string {
	var isAllNewLine bool
	s := ""
	for _, word := range arg {
		if word != "" {
			isAllNewLine = true
		}
	}

	count := 0
	if !isAllNewLine {
		count = 1
	}

	for i := 0; i < len(arg)-count; i++ {
		if arg[i] == "" {
			s = s + "\n"
		} else {
			for j := 0; j < 8; j++ {
				for _, g := range arg[i] {
					s = s + (bigtext[int(g)-32][j])
				}

				s = s + "\n"
			}
		}
	}

	return s
}
