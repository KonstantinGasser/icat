package internal

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	header = strings.NewReader("\033]1337;File=inline=1:")
	footer = strings.NewReader("\a")
)

func open(path string) (*os.File, error) {
	return os.Open(path)
}

// Show writes the image as RGB to the iTerm2
func Show(w io.Writer, path string) error {

	f, err := open(path)
	if err != nil {
		return fmt.Errorf("😱 ~ Oops I could not open the file: %v", err.Error())
	}
	defer f.Close()

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()

		wc := base64.NewEncoder(base64.StdEncoding, pw)

		_, err := io.Copy(wc, f)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("😱 ~ Oops I could not encode the image to base64: %v", err.Error()))
			return
		}
		if err := wc.Close(); err != nil {
			pw.CloseWithError(fmt.Errorf("😱 ~ Oops I could not close the Writer..file a bug report 💁‍♀️\n%v", err.Error()))
			return
		}
	}()

	if _, err := io.Copy(w, io.MultiReader(header, pr, footer)); err != nil {
		return fmt.Errorf("😱 ~ Oops I could not write the output to the os.Stdout..file a bug report 💁‍♀️\n%v", err.Error())
	}
	return nil
}

// Raw outputs the base64 encoded image the the os.Stdout
func Raw(w io.Writer, path string) error {

	f, err := open(path)
	if err != nil {
		return fmt.Errorf("😱 ~ Oops I could not open the file: %v", err.Error())
	}
	defer f.Close()

	encoded := getBase64(f)

	if _, err := fmt.Fprint(w, encoded); err != nil {
		return fmt.Errorf("😱 ~ Oops I could not write the output to the os.Stdout..file a bug report 💁‍♀️\n%v", err.Error())
	}
	return nil
}

// Write writes the base64 of the image to the givin output file
func Write(from, to string) error {

	f, err := open(from)
	if err != nil {
		return fmt.Errorf("😱 ~ Oops I could not open the file: %v", err.Error())
	}
	defer f.Close()

	base64 := strings.NewReader(getBase64(f))

	outf, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("😱 ~ I could not create the file %s", to)
	}
	if _, err := io.Copy(outf, base64); err != nil {
		return fmt.Errorf("😱 ~ I could not write the base64 to %s: %v", to, err.Error())
	}

	return nil
}

func getBase64(f *os.File) string {

	r := bufio.NewReader(f)
	enc, _ := ioutil.ReadAll(r)
	encoded := base64.StdEncoding.EncodeToString(enc)
	return encoded
}
