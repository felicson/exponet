//Package exponet for extract exhibitions from exponet.ru
//It get exhibitions list and parse each item from list
package exponet

import (
	"bytes"
	"exponet/expo"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html/charset"
)

const (
	domain = "exponet.ru"
)

var (
	clearDateRg  = regexp.MustCompile(`[^\d|\.]`)
	removeTagsRg = regexp.MustCompile("</?[div|a].*?>")
	rg           = regexp.MustCompile(`by-id/.*index\.ru\.html$`)
)

//GetExhibitions main extract func
func GetExhibitions(url string) ([]expo.Expo, error) {

	var (
		exhs  []expo.Expo
		links []string
		err   error
	)

	data, err := getHTML(url)
	if err != nil {
		return nil, err
	}
	rdr := bytes.NewReader(data)
	if links, err = getIndex(rdr); err != nil {
		return nil, err
	}

	for _, l := range links {
		data, err := getHTML(l)
		if err != nil {
			return nil, err
		}
		ex, err := parseExpo(bytes.NewReader(data))
		if err != nil {
			continue
		}
		if ex.Valid() {
			exhs = append(exhs, ex)
		}
	}
	return exhs, nil
}

//getIndex parse exhibitions urls
func getIndex(rdr io.Reader) ([]string, error) {

	root, err := htmlquery.Parse(rdr)
	if err != nil {
		return nil, err
	}
	var links []string
	sel, err := htmlquery.QueryAll(root, `//div[@id="maincontent"]//*/a[@href]`)
	if err != nil {
		return nil, err
	}
	for _, i := range sel {
		val := htmlquery.SelectAttr(i, "href")
		if rg.MatchString(val) {
			links = append(links, fmt.Sprintf("https://%s%s", domain, val))
		}
	}
	return links, nil
}

//parseExpo parse exhibition item
func parseExpo(rdr io.Reader) (expo.Expo, error) {

	root, err := htmlquery.Parse(rdr)
	if err != nil {
		return expo.Expo{}, err
	}
	s, err := htmlquery.Query(root, `//div[@class="exhibition"]/div[@class="row"]/div[@class="col-xs-exh-header"]`)
	if err != nil {
		return expo.Expo{}, err
	}
	exhName, _ := htmlquery.Query(s, "//h1")
	announce, _ := htmlquery.Query(s, "//p[1]")
	dates, _ := htmlquery.Query(s, "//p[2]/b")
	city, _ := htmlquery.Query(s, "//p[3]/b")
	datesA := strings.Split(htmlquery.InnerText(dates), "-")

	desc, _ := htmlquery.Query(root, `//div[@class="article"]/div[@class="content"]`)

	html := htmlquery.OutputHTML(desc, false)
	html = strings.ReplaceAll(html, "<noindex>", "")
	html = strings.ReplaceAll(html, "</noindex>", "")
	html = removeTagsRg.ReplaceAllString(html, "")

	dateStart, err := parseTime(datesA[0])
	if err != nil {
		return expo.Expo{}, err
	}
	dateEnd, err := parseTime(datesA[1])
	if err != nil {
		return expo.Expo{}, err
	}

	return expo.Expo{
		DateStart:   dateStart,
		DateEnd:     dateEnd,
		City:        htmlquery.InnerText(city),
		Announce:    htmlquery.InnerText(announce),
		Name:        htmlquery.InnerText(exhName),
		Description: html,
	}, nil
}

func getHTML(url string) ([]byte, error) {

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var b bytes.Buffer

	rdr, err := charset.NewReader(r.Body, "windows-1251")
	if err != nil {
		return nil, err
	}
	if _, err := b.ReadFrom(rdr); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func parseTime(str string) (time.Time, error) {
	t, err := time.Parse("02.01.2006", clearDateRg.ReplaceAllString(str, ""))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
