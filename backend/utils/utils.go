package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

const BUCKETURL string = "https://swoyam-pokharel-2431342-movie.s3.amazonaws.com/"
func GenerateUniqueFileName(original string) string {
	ext := filepath.Ext(original)
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%d%s", timestamp, ext)

}

func Join(strs []string, sep string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

func UploadToS3(fileHeader *multipart.FileHeader, fileName string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	url := fmt.Sprintf("https://swoyam-pokharel-2431342-movie.s3.amazonaws.com/%s", url.QueryEscape(fileName))

	req, err := http.NewRequest("PUT", url, file)
	if err != nil {
		return err
	}

	// Set required headers
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	req.Header.Set("Content-Type", contentType)

	// Add Content-Length header to avoid chunked encoding
	req.ContentLength = fileHeader.Size // This is the critical fix

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed: %d %s", resp.StatusCode, string(body))
	}

	return nil
}
