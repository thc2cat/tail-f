package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Linux tail is nice, don't reinvent...
// Windows reopen fails, loop it!

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage : tail.exe fichier")
	}

	if len(os.Args) > 2 {
		// write a line avery secs
		go appendstring(os.Args[1], 50*time.Millisecond)
	}
	//tail part
	var thissize int64
	for {
		thissize = checkfile(os.Args[1], thissize)
		time.Sleep(1 * time.Second)
	}
}

func checkfile(filename string, size int64) int64 {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if size != fi.Size() {
		filedumplast(filename, size)
		size = fi.Size()
	}
	return size
}

func filedumplast(fname string, start int64) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 1024*1024) //so short!
	taille, err := file.ReadAt(buf, start)
	if err != io.EOF {
		log.Fatal(err)
	}
	fmt.Print((string)(buf[0:taille]))
}

// Truncate/Appends  strings to this file
func appendstring(filename string, sleeptime time.Duration) {
	//	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	for counter := 0; ; counter++ {
		if _, err = f.WriteString(fmt.Sprintf("line (%v)\n", counter)); err != nil {
			log.Fatal(err)
		}
		time.Sleep(sleeptime)
	}
}
