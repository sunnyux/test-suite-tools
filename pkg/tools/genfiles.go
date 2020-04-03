package tools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// generateFiles takes in a specially formatted file of all test cases of name filename
// and generates corresponding .in .arg and the suite file with provided name
// see main.go for detail
func GenerateFiles(filename string, suiteFileName string, firstDelim string, secDelim string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	rd := bufio.NewReader(f)
	count, name, in, args, ebreaker, dbreaker := 0, "", "", "", firstDelim, secDelim
	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if line == ebreaker {
			count = 0
		} else if line == dbreaker {
			count += 1
		} else {
			if count == 0 {
				name = suiteFileName
				fname := line[:len(line)-1]
				in = fname + ".in"
				args = fname + ".args"
			} else if count == 1 {
				name = args
			} else {
				name = in
			}
			f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.Write([]byte(line)); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			fmt.Print(name, ", ")
		}
	}
	return f.Close()
}
