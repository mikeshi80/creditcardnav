// gatherer_test
package gatherer

import (
	"testing"
)

func Test_FetchText_1(t *testing.T) {
	g := Fetcher{}
	url := "http://www.baidu.com"
	_, err := g.FetchText(url)
	if err != nil {
		t.Log(err.Error())
		t.Fatalf("The url %s cannot be read correctly\n", url)
	}

}

func Test_FetchImage_gif(t *testing.T) {
	g := Fetcher{}
	url := "http://www.baidu.com/img/bdlogo.gif"
	img, err := g.FetchImage(url)
	if err != nil {
		t.Log(err.Error())
		t.Fatalf("The url %s cannot be read as a GIF\n", url)
	} else {
		if img == nil {
			t.Log(err.Error())
			t.Fatalf("The url %s cannot be read as a GIF\n", url)
		}
	}

}

func Test_FetchImage_png(t *testing.T) {
	g := Fetcher{}
	url := "http://soso.qstatic.com/30d/img/web/logo_index_130304.png"
	img, err := g.FetchImage(url)
	if err != nil {
		t.Log(err.Error())
		t.Fatalf("The url %s cannot be read as a PNG\n", url)
	} else {
		if img == nil {
			t.Log(err.Error())
			t.Fatalf("The url %s cannot be read as a PNG\n", url)
		}
	}

}

func Test_FetchImage_jpg(t *testing.T) {
	g := Fetcher{}
	url := "http://www.zhuanfala.com/uploadfile/201106/8/839188273.jpg"
	img, err := g.FetchImage(url)
	if err != nil {
		t.Log(err.Error())
		t.Fatalf("The url %s cannot be read as a JPG\n", url)
	} else {
		if img == nil {
			t.Log(err.Error())
			t.Fatalf("The url %s cannot be read as a JPG\n", url)
		}
	}

}
