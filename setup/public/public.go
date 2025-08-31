package public

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"log"
)

const (
	downloadURL = "https://drive.google.com/uc?export=download&id=YOUR_FILE_ID"
	archiveName = "public.zip"
)

func FetchPublicFolder() {
	resp, err := http.Get(downloadURL)
	if err != nil {
		log.Fatalf("ERROR: Failed to download file: %v", err)
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("ERROR: Failed to download file: %s", resp.Status)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		log.Fatalf("ERROR: Failed to read response body: %v", err)
	}

	zipFile, err := os.Create(archiveName)
	if err != nil {
		log.Fatalf("ERROR: Failed to create zip file: %v", err)
	}
	
	defer zipFile.Close()

	if _, err := zipFile.Write(buf.Bytes()); err != nil {
		log.Fatalf("ERROR: Failed to write zip file: %v", err)
	}

	if err := unzip(archiveName, "public"); err != nil {
		log.Fatalf("ERROR: Failed to unzip file: %v", err)
	}

	if err := os.Remove(archiveName); err != nil {
		log.Fatalf("ERROR: Failed to delete zip file: %v", err)
	}
}

func unzip(src, dest string) error {
	zipFile, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to open zip file: %w", err)
	}

	defer zipFile.Close()

	for _, file := range zipFile.File {
		if err := extractFile(file, dest); err != nil {
			return fmt.Errorf("ERROR: Failed to extract file %s: %w", file.Name, err)
		}
	}

	return nil
}

func extractFile(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("ERROR: Failed to open zip file entry %s: %w", f.Name, err)
	}

	defer rc.Close()

	outPath := filepath.Join(dest, f.Name)
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("ERROR: Failed to create directory for %s: %w", outPath, err)
	}

	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to create file %s: %w", outPath, err)
	}

	defer outFile.Close()

	if _, err := io.Copy(outFile, rc); err != nil {
		return fmt.Errorf("ERROR: Failed to extract file %s: %w", outPath, err)
	}

	return nil
}
