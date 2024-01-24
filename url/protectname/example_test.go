package url_test

import (
	"fmt"
	"log"
	"net/url"
)

func ExampleURL() {
	u, err := url.Parse("http://twitter.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// https://foo.com/go
}
