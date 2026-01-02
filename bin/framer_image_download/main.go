package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	pageURL := "http://localhost:8000/x.html"
	outputDir := "images"
	os.MkdirAll(outputDir, 0755)

	// Create HTTP client
	client := &http.Client{}

	// Fetch page with User-Agent
	req, _ := http.NewRequest("GET", pageURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/115.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	imgs := extractImgSrc(doc)
	if len(imgs) == 0 {
		log.Println("No images found on the page")
		return
	}

	for _, src := range imgs {
		imgURL := resolveURL(pageURL, src)
		downloadImage(imgURL, outputDir, client)
	}

	if err := zipFolder(outputDir, "images.zip"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done! Created images.zip")
}

func extractImgSrc(n *html.Node) []string {
	var urls []string
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			switch a.Key {
			case "src", "data-src":
				if a.Val != "" {
					urls = append(urls, a.Val)
				}
			case "srcset":
				parts := strings.Split(a.Val, ",")
				for _, p := range parts {
					urlPart := strings.Fields(p)[0]
					if urlPart != "" {
						urls = append(urls, urlPart)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		urls = append(urls, extractImgSrc(c)...)
	}
	return urls
}

func resolveURL(base, href string) string {
	if strings.HasPrefix(href, "//") {
		return "https:" + href
	}
	u, err := url.Parse(href)
	if err != nil || u.IsAbs() {
		return href
	}
	baseU, _ := url.Parse(base)
	return baseU.ResolveReference(u).String()
}

func downloadImage(imgURL, folder string, client *http.Client) {
	req, _ := http.NewRequest("GET", imgURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/115.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to download:", imgURL, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Failed to download:", imgURL, resp.Status)
		return
	}

	name := filepath.Base(imgURL)
	if strings.Contains(name, "?") {
		name = strings.Split(name, "?")[0]
	}
	outPath := filepath.Join(folder, name)

	out, err := os.Create(outPath)
	if err != nil {
		fmt.Println("Failed to create file:", outPath, err)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		fmt.Println("Failed to save file:", outPath, err)
		return
	}
	fmt.Println("Downloaded:", name)
}

func zipFolder(folder, zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	return filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Base(path)
		header.Method = zip.Deflate

		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, f)
		return err
	})
}
