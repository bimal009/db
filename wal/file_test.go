package wal

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

// func TestFile(t *testing.T) {
// 	err := EnsureDir(WalDirPath)
// 	t.Log("EnsureDir:", err)

// 	files, err := GetFiles(WalDirPath)
// 	name, latest, err := LatestFile(files, WalFileExtension)
// 	Append(WalDirPath)
// 	// file, fileNumber, err := CreateWalFile(WalDirPath)
// 	// if err != nil {
// 	// 	t.Fatal("CreateWalFile failed:", err)
// 	// }
// 	// defer file.Close()

// 	t.Log("latestfile name:", name)
// 	t.Log("latestfile number:", latest)
// }

func TestAppend(t *testing.T) {

	record := &Record{
		Type:  WalPut,
		Key:   []byte("name"),
		Value: []byte("john"),
	}

	if err := Append(WalDirPath, record); err != nil {
		t.Fatalf("Append failed: %v", err)
	}

	files, err := GetFiles(WalDirPath)
	if err != nil {
		t.Fatal(err)
	}

	name, latest, err := LatestFile(files, WalFileExtension)
	if err != nil {
		t.Fatal(err)
	}

	if name == "" {
		t.Fatal("No WAL file created")
	}

	data, err := os.ReadFile(filepath.Join(WalDirPath, name))
	if err != nil {
		t.Fatal(err)
	}

	expected, err := Encode(record)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, expected) {
		t.Fatal("WAL file content does not match encoded record")
	}

	t.Log("WAL file:", name)
	t.Log("WAL file number:", latest)
	t.Logf("WAL content:\n%s", hex.Dump(data))
}
