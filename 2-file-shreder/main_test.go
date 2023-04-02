package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestShred(t *testing.T) {
	// Create a temporary file to test the Shred function
	tmpFile, err := ioutil.TempFile("", "shred_test")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some data to the temporary file
	_, err = tmpFile.Write([]byte("test data"))
	if err != nil {
		t.Fatal("Error writing to temporary file:", err)
	}

	// Call the Shred function on the temporary file
	err = Shred(tmpFile.Name())
	if err != nil {
		t.Fatal("Error shredding file:", err)
	}

	// Ensure that the file no longer exists
	_, err = os.Stat(tmpFile.Name())
	if !os.IsNotExist(err) {
		t.Error("Expected file to be deleted, but it still exists")
	}
}

func TestShredInvalidPath(t *testing.T) {
	// Call the Shred function with an invalid file path
	err := Shred("/invalid/path")
	if err == nil {
		t.Error("Expected error when shredding invalid file path, but no error was returned")
	}
}
