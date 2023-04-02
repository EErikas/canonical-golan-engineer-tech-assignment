package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

const NumberOfRewrites = 3

func main() {
	// Check that a file path was provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: shred <file_path>")
		os.Exit(1)
	}

	// Get the file path from the command-line arguments
	filePath := os.Args[1]

	// Call the Shred function to overwrite and delete the file
	err := Shred(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func Shred(path string) error {
	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("File does not exist: %s", path)
	}

	// Open the file
	fmt.Printf("Preparing to shred file %s\n", path)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return err
	}
	fileSize := fileInfo.Size()

	// Overwrite the file 3 times with random data
	for i := 0; i < NumberOfRewrites; i++ {
		fmt.Printf("\t -> Writing random data. Pass %d out of %d\n", i+1, NumberOfRewrites)

		// Seek to the beginning of the file
		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Println("Error seeking to beginning of file:", err)
			return err
		}

		// Generate random data to overwrite the file
		data := make([]byte, fileSize)
		_, err = rand.Read(data)
		if err != nil {
			fmt.Println("Error generating random data:", err)
			return err
		}

		// Write the random data to the file
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing random data to file:", err)
			return err
		}
	}

	// Close and remove the file
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return err
	}
	err = os.Remove(path)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return err
	}

	fmt.Println("File shredded successfully:", path)
	return nil
}
