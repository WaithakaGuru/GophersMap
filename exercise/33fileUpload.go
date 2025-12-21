/* Exercise Thirty three
33. Create an endpoint to upload a file and save it to disk.
*/

package exercise

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var PORT = 3040

const (
	MaxUploadSize = 10 << 20
	UploadPath    = "./uploads"
)

func hanldeUpload(w http.ResponseWriter, r *http.Request) {
	// limit request body size
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	// check whether the upload exceeds maxLimit
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// retrieve file from form field "file"
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// create uploads directory if not exists
	if err := os.MkdirAll(UploadPath, os.ModePerm); err != nil {
		http.Error(w, "Unable to create upload directory", http.StatusInternalServerError)
		return
	}

	// sanitize filename
	filename := filepath.Base(header.Filename)
	filePath := filepath.Join(UploadPath, filename)

	// create destination file
	dstFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dstFile.Close()

	// copy file content to the destination file
	if _, err := io.Copy(dstFile, file); err != nil {
		http.Error(w, "File upload failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	successStmt := `\n File uploaded successfully %s`
	fmt.Fprint(w, successStmt, filename)
}

func StartUploadHandlerServer() {
	http.HandleFunc("/upload", hanldeUpload)

	log.Printf("\n Upload Serve running on port: %d", PORT)
	log.Fatal(http.ListenAndServe(":3040", nil))
}
