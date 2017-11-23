package util

import (
	"io"
	"net/http"
	"os"
)

func main() {
	file, url := os.Args[1], os.Args[2]
	Download(file, url)
}

// Download this downloads a thing to a file
func Download(file, url string) {

	out, _ := os.Create(file)
	defer out.Close()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
}
