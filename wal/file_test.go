package wal

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/bimal009/db/pkg/constants"
	"github.com/bimal009/db/pkg/file"
)

func TestFile(t *testing.T) {
	err := file.EnsureDir(constants.WalDirPath)
	t.Log("EnsureDir:", err)

	files, err := file.GetFiles(constants.WalDirPath)
	name, latest, err := file.LatestFile(files, constants.WalFileExtension)
	Append(constants.WalDirPath, &Record{
		Type:  constants.Put,
		Key:   []byte("name"),
		Value: []byte("john"),
	})
	// file, fileNumber, err := CreateWalFile(constants.WalDirPath)
	// if err != nil {
	// 	t.Fatal("CreateWalFile failed:", err)
	// }
	// defer file.Close()

	t.Log("latestfile name:", name)
	t.Log("latestfile number:", latest)
}

func TestAppend(t *testing.T) {

	record := &Record{
		Type:  constants.Put,
		Key:   []byte("name"),
		Value: []byte("john"),
	}

	if err := Append(constants.WalDirPath, record); err != nil {
		t.Fatalf("Append failed: %v", err)
	}

	files, err := file.GetFiles(constants.WalDirPath)
	if err != nil {
		t.Fatal(err)
	}

	name, latest, err := file.LatestFile(files, constants.WalFileExtension)
	if err != nil {
		t.Fatal(err)
	}

	if name == "" {
		t.Fatal("No WAL file created")
	}

	data, err := os.ReadFile(filepath.Join(constants.WalDirPath, name))
	if err != nil {
		t.Fatal(err)
	}

	expected, err := EncodeWalRecord(record)
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
