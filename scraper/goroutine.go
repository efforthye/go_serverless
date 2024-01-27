package main

// import (
// 	"fmt"
// 	"go/doc"
// 	"net/http"
// 	"strings"
// )

// // 레이트레이팅 -> 타임슬립
// // 네이버 지도

// func scrape(endpoint string) (content string, err error) {
// 	resp, err := http.Get(endpoint)
// 	if err != nil {
// 		return content: "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return content: "", error
// 	}

// 	// resp.Body를 이용해서 필요한 내용물 분리
// 	return content: "", errerr: nil
// }

// func main() {
// 	// 엔드포인트들 세팅
// 	// 사이트1: https://kubernetespodcast.com/
// 	// 사이트2: https://www.heroku.com/

// 	endpoints := []string{"https://kubernetespodcast.com/"}

// 	for _, endpoint := range endpoints {
// 		content, err := scrape(endpoint)
// 		if err != nil {
// 			// do something

// 		}
// 		fmt.Println(content)
// 	}

// }

// doc, err: = goquery.NewDocumentFromReader(resp.Body)
// if err != nil {
// 	return
// }

// doc.Find(selector: "div.episode h3").Each(func(i int, s *goquery.Selection){
// 	if  strings.Contains(strings.ToLower(s.Text()), substr) {

// 	}
// })

func main() {
}
