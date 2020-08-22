package main

import (
	"flag"
	"fmt"
	"os"
)

// flags:
// mode -d (display) | -w (write: string path) | -r (raw)

func main() {

	// flag mode
	mode := flag.NewFlagSet("mode", flag.ExitOnError)
	display := mode.Bool("d", true, "If set img is printed in command-line (works for iTerm2 only)")
	write := mode.String("w", "nil", "write base64 encoded img to given file")
	raw := mode.Bool("r", false, "If set base64 encoded img printed to command-line")
	flag.Parse()

	if err := iCat(*display, *raw, *write); err != nil {
		fmt.Fprintf(os.Stderr, "😕: Clound not display img: %v|n", err)
	}
}

func iCat(display, raw bool, write string) error {
	return nil
}

// package main

// import (
// 	"encoding/base64"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// var (
// 	header = strings.NewReader("\033]1337;File=inline=1:")
// 	footer = strings.NewReader("\a")
// 	usage  = func() {
// 		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
// 		flag.PrintDefaults()
// 	}
// )

// func main() {
// 	path := flag.String("path", "nil", "path to image to display")
// 	raw := flag.NewFlagSet("raw", flag.ExitOnError)
// 	stdout := raw.Bool("print", false, "Print img in base64 to Stdout")
// 	saveAt := raw.String("path", "nil", "save base64 in givien path as file")
// 	flag.Parse()
// 	fmt.Fprintf(os.Stderr, "%v, %s, %s\n", *stdout, *saveAt, *path)
// 	if *path == "nil" {
// 		fmt.Println("🤬: Dude I need a file!?")
// 		os.Exit(1)
// 	}
// 	if err := displayIMG(*path, *stdout, *saveAt); err != nil {
// 		fmt.Fprintf(os.Stderr, "😕: Clound not display img: %v|n", err)
// 		os.Exit(1)
// 	}
// }

// func displayIMG(path string, raw bool, saveAt string) error {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	if err := copy(os.Stdout, f, raw, saveAt); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func copy(w io.Writer, r io.Reader, raw bool, saveAt string) error {

// 	pr, pw := io.Pipe()
// 	go func() {
// 		defer pw.Close()
// 		wc := base64.NewEncoder(base64.StdEncoding, pw)
// 		_, err := io.Copy(wc, r)
// 		if err != nil {
// 			pw.CloseWithError(err)
// 			return
// 		}

// 		if err := wc.Close(); err != nil {
// 			pw.CloseWithError(err)
// 			return
// 		}

// 	}()

// 	nf, err := os.Create("./t.txt")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "😕: Cloud not write base64 encoded img to %s:\n%v", saveAt, err.Error())
// 	}
// 	if _, err := io.Copy(nf, pr); err != nil {
// 		fmt.Fprintf(os.Stderr, "😕: Clound not write base64 encoded img to %s:\n%v", saveAt, err.Error())
// 	}
// 	if !raw {
// 		if err := ioCopy(w, io.MultiReader(header, pr, footer)); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	// if saveAt != "nil" {

// 	// }
// 	if err := ioCopy(w, pr); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func ioCopy(w io.Writer, r io.Reader) error {
// 	_, err := io.Copy(w, r)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
