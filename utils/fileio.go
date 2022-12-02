package utils

import (
	"bufio"
	"io"
	"os"
)

// read a file from specified location
func ReadFileIntoList(fileloc string, delim byte) []string {
	// open file from string and check for errors
	fr, err := os.Open(fileloc)
	Check(err)

	// create a new buffered reader
	bufr := bufio.NewReader(fr)

	// create the finished []string
	var list []string
	// loop counter
	for {
		// initialize buffers and read a line from a file in a loop
		dat, err := bufr.ReadString(delim)
		if err == io.EOF {
			break
		}
		Check(err)
		list = append(list, dat)
	}

	return list
}
