package gatherer

/*
The data gatherer for Bank COMM 交通银行
*/
import (
	ds "creditcardnav/data"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

type BankCOMM struct {
	*Fetcher
}

func (this *BankCOMM) GenerateURLs() (urls []string, err error) {
	values := make(url.Values, 10)
	values["cityId"] = "2c92de8f2ee5d079012ee5d50dda0003"
	values["cityName"] = "上海"
	values["circleId"] = "0"
	values["catalog"] = "0"
	values["cardType"] = ""
	values["keyWord"] = ""
	values["pageSize"] = "20"
	values["pageNo"] = "1"
	values["orderBy"] = "0"
	values["isPage"] = "true"

	body, err := this.FetchTextByPost("http://creditcard.bankcomm.com/bcms/front/merchant/ajax/search.do", values)

	if err != nil {
		return nil, err
	}

	return body, err
}

func (this *BankCOMM) ClawerList() (events []ds.Event, err error) {
	return nil, nil
}
