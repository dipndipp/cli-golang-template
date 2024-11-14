package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Validasi jumlah argumen CLI
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing arguments.\nUsage: go-template-cli create <project-name>")
		return
	}

	// Validasi command "create"
	command := os.Args[1]
	if command != "create" {
		fmt.Printf("Error: Unknown command '%s'.\nUsage: go-template-cli create <project-name>\n", command)
		return
	}

	// Ambil nama proyek dari argumen
	projectName := os.Args[2]

	// Dapatkan lokasi executable
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	baseDir := filepath.Dir(exePath) // Direktori executable
	templatePath := filepath.Join(baseDir, "template") // Path template
	destinationPath := filepath.Join(".", projectName) // Path proyek baru

	// Pastikan folder template ada
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		fmt.Printf("Error: Template folder not found at '%s'.\n", templatePath)
		return
	}

	// Pastikan proyek baru tidak tumpang tindih
	if _, err := os.Stat(destinationPath); !os.IsNotExist(err) {
		fmt.Printf("Error: Project folder '%s' already exists.\n", projectName)
		return
	}

	// Salin template ke proyek baru
	fmt.Printf("Creating project '%s'...\n", projectName)
	err = copyDir(templatePath, destinationPath)
	if err != nil {
		fmt.Printf("Error generating project: %v\n", err)
		return
	}

	fmt.Println("Project generated successfully!")
}

// Fungsi untuk menyalin folder
func copyDir(src string, dst string) error {
	// Iterasi melalui folder sumber
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Hitung path relatif
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// Buat direktori di tujuan
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// Salin file
		return copyFile(path, targetPath)
	})
}

// Fungsi untuk menyalin file
func copyFile(src string, dst string) error {
	// Buat direktori tujuan jika belum ada
	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	// Buka file sumber
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Buat file tujuan
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Salin isi file
	_, err = io.Copy(destFile, sourceFile)
	return err
}
