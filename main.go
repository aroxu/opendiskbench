package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// const testFileSize = 4096 // 4MB in KB
	// const testFileSize = 10485760 // 10GB in KB
	const testFileSize = 1048576 // 1GB in KB
	const defaultBufSize = 4096

	fSize := int64(testFileSize) * (1024)
	err := writeFile(fSize, defaultBufSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, fSize, err)
	}
}

func writeFile(fSize int64, bufferSize int) error {
	fName := "test.file.dummy" // test file
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Fatalln("Failed to delete file.")
			return
		}
	}(fName)
	f, errCreateFile := os.Create(fName)
	if errCreateFile != nil {
		return errCreateFile
	}
	buffer := make([]byte, bufferSize)
	buffer[len(buffer)-1] = '\n'
	writer := bufio.NewWriterSize(f, len(buffer))

	written := int64(0)
	startTime := time.Now()
	for i := int64(0); i < fSize; i += int64(len(buffer)) {
		nn, err := writer.Write(buffer)
		written += int64(nn)
		if err != nil {
			return err
		}
	}
	errFlush := writer.Flush()
	since := time.Since(startTime)

	if errFlush != nil {
		return errFlush
	}
	errCloseFile := f.Close()
	if errCloseFile != nil {
		return errCloseFile
	}
	fmt.Printf("written: %dB %dns %.2fGB %.2fs %.2fMB/s\n",
		written, since,
		float64(written)/(1024*1024*1024), float64(since)/float64(time.Second),
		(float64(written)/(1024*1024))/(float64(since)/float64(time.Second)),
	)
	return nil
}
