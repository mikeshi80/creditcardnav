package gatherer

/*
The data gatherer for Bank COMM 交通银行
*/
import (
	"net/url"
    "strings"
	ds "creditcardnav/data"
	"github.com/PuerkitoBio/goquery"
    "code.google.com/p/go.net/html"
)

type BankCOMM struct {
	*Fetcher
}

func (this *BankCOMM) ParseUrls(body string) (urls []string, err error) {
    node, err := html.Parse(strings.NewReader(body))
    doc := goquery.NewDocumentFromNode(node)

    urls = make([]string, 0, 20)
    doc.Find("div.ml-pic>a").Each(func(i int, s *goquery.Selection) {
        attr, exists := s.Attr("href")
        if exists {
            url := "http://creditcard.bankcomm.com/bcms/merc" + attr
            urls = append(urls, url)
        }
    })
    return urls, nil
}

func (this *BankCOMM) FetchData() (body string, err error) {
	values := make(url.Values, 10)
	values.Set("cityId", "2c92de8f2ee5d079012ee5d50dda0003")
	values.Set("cityName", "上海")
	values.Set("circleId", "0")
	values.Set("catalog", "0")
	values.Set("cardType", "")
	values.Set("keyWord", "")
	values.Set("pageSize", "20")
	values.Set("pageNo", "1")
	values.Set("orderBy", "0")
	values.Set("isPage", "true")

	body, err = this.FetchTextByPost("http://creditcard.bankcomm.com/bcms/front/merchant/ajax/search.do", values)

	if err != nil {
		return "", err
	}

	return body, err
}

func (this *BankCOMM) ClawerList() (events []ds.Event, err error) {
	return nil, nil
}
