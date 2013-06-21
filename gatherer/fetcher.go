// Gatherer project Gatherer.go
package gatherer

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Fetcher struct {
}

type ImageFormatError struct {
	reason string
}

func (this *ImageFormatError) Error() string {
	return this.reason
}

func (this *Fetcher) FetchText(url string, args ...interface{}) (body string, err error) {
	url = fmt.Sprintf(url, args...)

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("This url [%s] cannot be read correctly", url)
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	return string(content), err
}

func (this *Fetcher) FetchTextByPost(url string, args url.Values) (body string, err error) {
	resp, err := http.PostForm(url, args)

	if err != nil {
		log.Printf("This post url [%s] cannot be read correctly", url)
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	return string(content), err
}

func (this *Fetcher) FetchImage(url string) (img image.Image, err error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("This url [%s] cannot be read correctly", url)
		return nil, err
	}

	defer resp.Body.Close()

	switch strings.ToLower(path.Ext(url)) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(resp.Body)
	case ".png":
		img, err = png.Decode(resp.Body)
	case ".gif":
		img, err = gif.Decode(resp.Body)
	default:
		err = &ImageFormatError{"Unknown image format"}
	}
	return
}
