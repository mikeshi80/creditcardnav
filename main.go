// CreditCardNav project main.go
package main

import (
	"creditcardnav/gatherer"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	g := gatherer.Fetcher{}
	g.FetchText("http://baidu.com")
}
