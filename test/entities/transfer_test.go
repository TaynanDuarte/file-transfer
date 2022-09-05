package transfer_test

import (
	"testing"

	"github.com/TaynanDuarte/file-transfer/src/models"
)

func BenchmarkTransfer(b *testing.B) {

	transfer := models.TransferConstructor("../test-files/from", "../test-files/to", "txt")
	for i := 0; i < b.N; i++ {
		transfer.Run()
	}

}
