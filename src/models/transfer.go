package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type transfer struct {
	files         []File
	from, to      string
	fileExtension string
}

// Return a instance of Transfer
func TransferConstructor(from string, to string, fileExtention string) transfer {
	transfer := transfer{from: from, to: to, fileExtension: fileExtention}
	transfer.PrepareFiles()

	return transfer
}

// Receive a file path and return only the file name
func GetFileNameFromFilePath(filePath string) string {

	if strings.Contains(filePath, "/") {
		lastSlashIndex := strings.LastIndex(filePath, "/")
		return filePath[lastSlashIndex+1:]
	}

	return filePath
}

// Load files references and crete readers and writers
func (t *transfer) PrepareFiles() {

	filesInfo, err := os.ReadDir(t.from)

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range filesInfo {

		hasTheCorrectExtension := strings.Contains(fileInfo.Name(), t.fileExtension)
		if hasTheCorrectExtension {
			fmt.Println("Loading: " + fileInfo.Name())

			originFilePath := t.from + "/" + GetFileNameFromFilePath(fileInfo.Name())
			originFile, _ := os.Open(originFilePath)
			reader := bufio.NewReader(originFile)

			outFilePath := t.to + "/" + GetFileNameFromFilePath(fileInfo.Name())
			outFile, _ := os.Create(outFilePath)
			writer := bufio.NewWriter(outFile)

			file := File{reader: reader, writer: writer, originFileReference: originFile}
			t.files = append(t.files, file)
		}

	}

}

// Start the transfer
func (t *transfer) Run() {

	var wg sync.WaitGroup
	wg.Add(len(t.files))

	for _, file := range t.files {
		go func(_file File) {

			fmt.Println("Transferring: ", _file.FileName())

			buffer := make([]byte, 1024)
			for {
				numberOfBytesRead, err := _file.reader.Read(buffer)

				if err != nil {
					fmt.Println("Finished: " + _file.originFileReference.Name())
					break
				}

				buffer = buffer[:numberOfBytesRead] // removing zeros

				_file.writer.Write(buffer)
				_file.writer.Flush()

			}

			wg.Done()

			_file.originFileReference.Close()
			fmt.Println("Closed: " + _file.FileName())
		}(file)
	}

	wg.Wait()

}
