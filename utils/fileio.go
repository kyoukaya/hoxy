// read a file from specified location
func readFileIntoList(fileloc string) []string {
	// open file from string and check for errors
	fr, err := os.Open(fileloc)
	utils.Check(err)

	// create a new buffered reader
	bufr := bufio.NewReader(fr)

	// create the finished []string
	var list []string
	// loop counter
	for {
		// initialize buffers and read a line from a file in a loop
		dat, err := bufr.ReadString('\n')
		if err == io.EOF {
			break
		}
		utils.Check(err)
		list = append(list, string(dat))
  }

  return list
}

// read a file from specified location as []byte
func readFileIntoByte(fileloc string) []byte {
	// open file from string and check for errors
	fr, err := os.Open(fileloc)
	utils.Check(err)

	// create a new buffered reader
	bufr := bufio.NewReader(fr)

	dat, err := bufr.ReadBytes(0)
	utils.Check(err)

	return dat
}
