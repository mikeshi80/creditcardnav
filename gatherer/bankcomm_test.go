package gatherer

import (
    "testing"
    "strings"
    "io/ioutil"
)

func Test_FetchData(t *testing.T) {
    bc := BankCOMM { Fetcher : &Fetcher{}}
    body, err := bc.FetchData()

    if err != nil {
        t.Fatal("Error when fetch the data")
    }

    if len(body) == 0 {
        t.Fatal("Cannot fetch the data")
    }

    if !strings.Contains(body, "优惠细则") {
        t.Fatal("The body is wrong")
    }

}

func Test_ParseUrls(t *testing.T) {
    var (
        content []byte
        urls []string
        err error
    )
    if content, err = ioutil.ReadFile("testdata/getbankcomm.txt"); err != nil {
        t.Fatal("Cannot read the test data")
    }

    bc := BankCOMM { Fetcher : &Fetcher{}}
    if urls, err = bc.ParseUrls(string(content)); err != nil {
        t. Fatal("Cannot parse the urls")
    }

    if len(urls) != 5 {
        t.Fatalf("The urls count is not 5, is %d\n", len(urls))
    }

    if urls[0] != "http://creditcard.bankcomm.com/bcms/merc/bcms/merc2/75024.htm" {
        t.Fatalf("The first node is not 75024, is %s\n", urls[0])
    }

}


