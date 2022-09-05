package main

import (
	"fmt"
	"time"

	"github.com/TaynanDuarte/file-transfer/src/models"
)

func main() {

	startedAt := time.Now()

	transfer := models.TransferConstructor("../files/from", "../files/to", "json")
	transfer.Run()

	duration := time.Since(startedAt)
	fmt.Println("Duration: ", duration)

}
