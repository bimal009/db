package wal

import "testing"

func TestFile(t *testing.T) {
	err := EnsureDir(WalDirPath)
	t.Log("EnsureDir:", err)

	file, fileNumber, err := CreateWalFile(WalDirPath)
	if err != nil {
		t.Fatal("CreateWalFile failed:", err)
	}
	defer file.Close()

	t.Log("CreateWalFile:", file.Name())
	t.Log("File number:", fileNumber)
}
