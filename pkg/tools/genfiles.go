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
	count, name, in, args, content, prevline, write := 0, "", "", "", "", "", false
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			if name != "" {
				f, err := os.OpenFile(in, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				if _, err := f.Write([]byte(content)); err != nil {
					log.Fatal(err)
				}
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
				fmt.Print(in, ", ")
			}
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line != firstDelim && line != secDelim && !write {
			content += line
		} else if line == secDelim {
			count += 1
			if count == 1 {
				fname := prevline[:len(prevline)-1]
				in = fname + ".in"
				args = fname + ".args"
				name = suiteFileName
			} else if count == 2 {
				name = args
			}
			write = true
		} else if line == firstDelim && name != "" {
			count = 0
			name = in
			write = true
		}
		if write {
			f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.Write([]byte(content)); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			fmt.Print(name, ", ")
			content = ""
			write = false
		}
		prevline = line
	}
	return f.Close()
}
