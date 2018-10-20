package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	//<-rateLimiter

	// 解决403的问题
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	encoding := determineEncoding(bufio.NewReader(resp.Body))
	utf8reader := transform.NewReader(resp.Body, encoding.NewDecoder())
	return ioutil.ReadAll(utf8reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
