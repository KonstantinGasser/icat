package internal

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	header = strings.NewReader("\033]1337;File=inline=1:")
	footer = strings.NewReader("\a")
)

// Show writes the image as RGB to the iTerm2
func Show(w io.Writer, path string, isURL bool) error {

	content, err := getContent(path, isURL)
	if err != nil {
		return err
	}

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()

		wc := base64.NewEncoder(base64.StdEncoding, pw)

		var reader io.Reader

		switch content.(type) {
		case *os.File:
			defer content.(*os.File).Close()
			reader = content.(*os.File)
		case io.ReadCloser:
			defer content.(io.ReadCloser).Close()
			reader = content.(io.ReadCloser)
		}

		_, err := io.Copy(wc, reader)
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
func Raw(w io.Writer, path string, isURL bool) error {
	var reader io.Reader
	content, err := getContent(path, isURL)
	if err != nil {
		return err
	}
	switch content.(type) {
	case *os.File:
		defer content.(*os.File).Close()
		reader = content.(*os.File)
	case io.ReadCloser:
		defer content.(io.ReadCloser).Close()
		reader = content.(io.ReadCloser)
	}

	encoded := getBase64(reader)

	if _, err := fmt.Fprint(w, encoded); err != nil {
		return fmt.Errorf("😱 ~ Oops I could not write the output to the os.Stdout..file a bug report 💁‍♀️\n%v", err.Error())
	}
	return nil
}

// Write writes the base64 of the image to the givin output file
func Write(from, to string, isURL bool) error {
	var reader io.Reader

	content, err := getContent(from, isURL)
	if err != nil {
		return err
	}

	switch content.(type) {
	case *os.File:
		defer content.(*os.File).Close()
		reader = content.(*os.File)
	case io.ReadCloser:
		defer content.(io.ReadCloser).Close()
		reader = content.(io.ReadCloser)
	}

	base64 := strings.NewReader(getBase64(reader))

	outf, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("😱 ~ I could not create the file %s", to)
	}
	if _, err := io.Copy(outf, base64); err != nil {
		return fmt.Errorf("😱 ~ I could not write the base64 to %s: %v", to, err.Error())
	}

	return nil
}

func open(path string) (*os.File, error) {
	return os.Open(path)
}

func getBase64(reader io.Reader) string {

	r := bufio.NewReader(reader)
	enc, _ := ioutil.ReadAll(r)
	encoded := base64.StdEncoding.EncodeToString(enc)
	return encoded
}

func getContent(path string, isURL bool) (interface{}, error) {
	time.Sleep(5 * time.Second)
	if !isURL {
		f, err := open(path)
		if err != nil {
			return nil, fmt.Errorf("😱 ~ Oops I could not open the file: %v", err.Error())
		}
		return f, nil
	}

	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("🤯 ~ Looks like something is wring with the URL or your network")
	}
	return resp.Body, nil
	return nil, nil
}
