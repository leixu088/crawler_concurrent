package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	cookie1 := "sid=fdd788d1-a787-4498-a918-c5a73b770345; ec=tib25Qzn-1627802546653-4b4144f4574fb278069660; FSSBBIl1UgzbN7NO=5BpJ8gSAzwwVfeXLvxtsr1qC0ZtIxEEbQ00Q2nemwc.BteZhi_NNNVyxxI0tEB_qv6S.rQ6bPP6Z82EbiONjeIa; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1627802552; _exid=Y5uc55UlusWihxVniDvlL7oVhhnebdWJQ9l8xic%2BckCtkpMIt9ooz5bq5DmQ2Zz4IQ4mUOzVUNa2Z6dbBWOV8w%3D%3D; _efmdata=Gx5ZoINlpAoAYTozpdbPaAy0rbRYChiTwuo7LVGkRJ2Guq4mWGh1xJMYuW07rVUPsU84GU4RtZeM9U%2BQoqEY466av56ZkIGWr5K%2FhvTRr04%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1627809291; FSSBBIl1UgzbN7NP=53ARzyKlDM8lqqqm4qpTx8Go0FBrg_w1P8lgiGNl76aWMJprBkyoZ2odYyt1mcX.fpabQiX2ZNJjM52i4avuL1Eh4gW3vaDRbZJXLbBnnfedfgOtaxMVYdA0ChgrP35gAZO8xeaIoj7GFmJP8Polol9eZxX3IOwEOIj6m48VMcVSh1RYqdkcXPvKCA1dhIRZT_iJF_7GJdwoAg8gzkjWmyLmAsBl_Nz4sHJlJabO7StvNHEQ0pilTMxFJjvDV3VIXYjTbhE067Y927GEmdAt9z2"
	req.Header.Add("cookie", cookie1)
	resp, err := client.Do(req)
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// 把网页转为utf-8编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
