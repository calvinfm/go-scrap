package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

type Phones struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price string `json:"price"`
	Image string `json:"image"`
	Rating string `json:"rating"`
	StoreName string `json:"store_name"`
	TotalRate string `json:"total_rate"`
}

func main() {

	var phones [][]string

	k := 1
	for true {
		var err error
		var client = &http.Client{}

		request, err := http.NewRequest("GET", fmt.Sprint("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=",k), nil)
		if err != nil {
			log.Println(err.Error())
		}
		k++

		//request.AddCookie(&http.Cookie{
		//	Name:       "_abck",
		//	Value:      "958254095A0E6B18C234733F951B5D4C~-1~YAAQQf5ebyMv4Z1/AQAA+g+dqwfNkPAoLQS1QwGnsd/xS9x/3xsv20OwE+IRc9wjCmPqAAX+2aWFBz3g7ERppG91C1PoHAmLCiYhh5c8JLs76zvLCysmID2BItcaBWfz6aMK03i7G1eYK0xQRB/Vc6yI9rsGrpKAyubF2Th7qGDcWs3iWJx/tq14mQwCUcubU87P5ipVnlmu/TKbk/+ObQ7oa6hagmG4ZJ6pUMYwtP3HEo1+Y1PT6yCSVlB9YOCJCA/WPpUlwn4QwkKFru0t8QjKby5opqHjEEBJC3v3u9AGzXDFtKnZweccgFQzdVhNoxnvqexM293yi+pO5umdvwCyuzllat4K79obmVRM3Ob2/zMgtZB5luz1CChhjs5eNz1lxPs20NbzhEm6IeWYbaIOJ3RDQ3YUH9bMSQ==~-1~-1~-1",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2023,3,21, 8,34,31,0,time.Local),
		//	RawExpires: "Mon, 21 Mar 2023 08:34:31 GMT",
		//	MaxAge:     31536000,
		//	Secure:     true,
		//	HttpOnly:   false,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "ak_bmsc",
		//	Value:      "40809A19BFB3D640CD4A30249A9500BB~000000000000000000000000000000~YAAQZAkFcst2Ipt/AQAADMIlrA8LsemewAC/TRv97cdodq1jTh/hQ90BaU8Kp2kZRNLG+4WZvkISsW8IhCoibPgaX+Z2yz9NqGIJvfQ6PdaOolCGNX6MjaGZhDuyLMAEfBV8ZyttCB8+XVBYOZBNr2uGAXZb3Jkexny236xCfF+wnKHOW6nQQsIfwXv3oTbdzPjvAlBKc+U3mIsqcsVz/x/LHNhLa6GPiDds5dznVQIbTi64Xv48aHhWBLt7vfVyjkCJpAUeDKPyCmeHC2UD0XEo+talN+2EWP0TnQM1WmA2SlKxGLPKMsfkfJRqeDNAbFKK0QLYRdnmwOUAud7SnoWArIG5PtOgmh337woWP2FEsldynMLibOuOiIsNpnCjMrHsIrr2LQiX+JeM/A==",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2022,3,21, 11,03,50,0,time.Local),
		//	RawExpires: "Sun, 21 Mar 2022 11:03:50 GMT",
		//	MaxAge:     0,
		//	Secure:     false,
		//	HttpOnly:   true,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "bm_sz",
		//	Value:      "693AC9BA0F88B09FCDBCCFC314755BC5~YAAQQf5ebzRS4Z1/AQAAJR6eqw8dSOljW1GRIhp4VMbB9RkbaQDraRwNvqj1M/lsHyLu27UPpCQa1LtZ3GcZPDXrELAGdiu3Wo1uh3aqmvXgG6lNyqnbIard6R5DFMf2XFYWWg33YsRk9my8RdgGR5HrQ7xh2+AL2syznZBMGYXGZl2Zh2KwX1tkMlEm4mCpZs8SeNXUFZetZggBYj7p55qxBTBPEWIW53fxBerZEpoUIPdvnnuDp8p0fY7h31w6O22bID7LUZNcvLv7ZXMRXj0P87DZRIatHEf+5cfSB6aDJnuzAVw=~4337713~3293749",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2022,3,21, 12,35,40,0,time.Local),
		//	RawExpires: "Sun, 21 Mar 2022 12:35:40 GMT",
		//	MaxAge:     14400,
		//	Secure:     false,
		//	HttpOnly:   false,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "Cookie_1",
		//	Value:      "value",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2022,3,20, 15,07,29,0,time.Local),
		//	RawExpires: "Mon, 20 Mar 2023 15:07:29 GMT",
		//	MaxAge:     0,
		//	Secure:     false,
		//	HttpOnly:   false,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "_UUID_NONLOGIN_",
		//	Value:      "13dd3a9a3287907619f27b86565bc3bf",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2070,12,31, 0,0,0,0,time.Local),
		//	RawExpires: "Mon, 31 Dec 2070 00:00:00 GMT",
		//	MaxAge:     0,
		//	Secure:     false,
		//	HttpOnly:   false,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "_UUID_NONLOGIN_.sig",
		//	Value:      "Yt2Di47l6RszEDaiqa9qFnaaWmw",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2070,12,31, 0,0,0,0,time.Local),
		//	RawExpires: "Mon, 31 Dec 2070 00:00:00 GMT",
		//	MaxAge:     0,
		//	Secure:     false,
		//	HttpOnly:   false,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "bm_mi",
		//	Value:      "C26602036C934F2115BCE024D591CBD6~AGnZXtjyW+YurjJvbQXcNi68vCf2C3xsLRGhfdFz9uorVHdbDOJ/Zu46/p6viEEhk+LZgCM47fa4la2E6MpGLnuPke/zo0hDgoU2FT9Izbf6yIgDLNDBvWsQI3hJy8MaIZdk2gt2N4QfySRMgiTaWkk4qAr48sqGp0pCHVL85nydG1UEzVMS8mhJtJbThS+POBrJptZpLKus4gg/lVMRiekTgiAz6HyZYVkNtXIqguvYgX9NLbjDGb0Pu17cfUpGUJHVc83AC4qtx6RHH64QaA==",
		//	Path:       "/",
		//	Domain:     ".tokopedia.com",
		//	Expires:    time.Date(2070,12,31, 0,0,0,0,time.Local),
		//	RawExpires: "Mon, 31 Dec 2070 00:00:00 GMT",
		//	MaxAge:     6773,
		//	Secure:     false,
		//	HttpOnly:   true,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		//request.AddCookie(&http.Cookie{
		//	Name:       "serverECT",
		//	Value:      "4g",
		//	Path:       "/",
		//	Domain:     "",
		//	Expires:    time.Date(2022,3,20, 15,07,29,0,time.Local),
		//	RawExpires: "",
		//	MaxAge:     0,
		//	Secure:     true,
		//	HttpOnly:   true,
		//	SameSite:   0,
		//	Raw:        "",
		//	Unparsed:   nil,
		//})

		request.Header.Set("accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Header.Set("accept-encoding","gz?ip, deflate, br")
		request.Header.Set("accept-language","en-US,en;q=0.9")
		request.Header.Set("cache-control","max-age=0")
		request.Header.Set("ect","4g")
		request.Header.Set("Cache-Control","no-cache")
		request.Header.Set("Connection","keep-alive")
		request.Header.Set("Host","443")
		request.Header.Set("sec-ch-ua",` Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99`)
		request.Header.Set("sec-ch-ua-mobile","?0")
		request.Header.Set("sec-ch-ua-platform","Windows")
		request.Header.Set("sec-fetch-dest","document")
		request.Header.Set("sec-fetch-mode","navigate")
		request.Header.Set("sec-fetch-site","cross-site")
		request.Header.Set("sec-fetch-user","?1")
		request.Header.Set("upgrade-insecure-requests","1")
		request.Header.Set("user-agent",`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36`)
		request.Header.Add("cookie",`hfv_banner=true; g_state={"i_p":1647820662498,"i_l":2};`)

		res, err := client.Do(request)
		if err != nil {
			log.Println(err.Error())
		}
		defer res.Body.Close()

		//res, err := http.Get("https://www.tokopedia.com/p/handphone-tablet/handphone?ob=5&page=1")

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		fmt.Printf("%+v\n",res)
		fmt.Println(res.Header.Get("Set-Cookie"))
		temp := strings.Split(res.Header.Get("Set-Cookie"),"; ")
		for _, s := range temp {
			fmt.Println(s)
			temps := strings.SplitN(s,"=",2)
			for _, s2 := range temps {
				fmt.Println(s2)
			}
		}
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".css-13l3l78 .css-bk6tzz .css-89jnbj .css-16vw0vn ").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find(".css-1bjwylw").Text()
			fmt.Printf("Name %d: %s\n", i, title)

			price := s.Find(".css-o5uqvq").Text()
			fmt.Printf("Price %d: %s\n", i, price)

			img,_ := s.Find(".css-79elbk").Children().Children().Children().Attr("src")

			fmt.Printf("img %d: %s\n", i, img)

			store := ""
			s.Find(".css-1kr22w3").Each(func(i int, s *goquery.Selection) {

				if i == 1 {
					fmt.Printf("Store %s\n",  s.Text())
					store = s.Text()
				}
			})

			var rate int
			s.Find(".css-177n1u3").Each(func(i int, s *goquery.Selection) {
				rate++
			})

			if rate == 0 {
				fmt.Printf("No Rating")
			}else {
				fmt.Printf("Rating %d Star \n", rate)
			}

			totalRate := s.Find(".css-153qjw7").Children().Find("span").Text()
			fmt.Printf("Total Rating %s\n",  totalRate)

			phones = append(phones,[]string{title,title,price,img,fmt.Sprintf("Rating %d Star", rate),store,fmt.Sprint("Total Rating ",  totalRate)})

		})

		if len(phones) >= 100 {
			break
		}

	}

	f, err := os.Create("users.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(phones) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}

}
