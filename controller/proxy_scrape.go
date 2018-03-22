package controller

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"

	"errors"

	"github.com/richardsang2008/proxy_finder/model"
)

func ScapeHideMyName(url string, tbmapper map[string]int) (*[]model.ProxyRecord, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	var proxies []model.ProxyRecord
	//use CSS selector found with the browser
	doc.Find("body tbody tr ").Each(func(index int, item *goquery.Selection) {
		var data map[int]string
		data = make(map[int]string)
		item.Find("td").Each(func(index2 int, item2 *goquery.Selection) {
			data[index2] = item2.Text()
		})
		record := model.ProxyRecord{}
		if len(data) == 8 {
			record.IP = data[tbmapper["IP"]]
			record.Port, _ = strconv.Atoi(data[tbmapper["PORT"]])
			record.Country = data[2]
			if data[4] == "anonymous" {
				record.Type = model.Anoymous
			} else if data[4] == "elite proxy" {
				record.Type = model.Eliteproxy
			} else if data[4] == "transparent" {
				record.Type = model.Transparent
			} else {
				record.Type = model.Unknown
			}
			if data[6] == "yes" {
				record.IsHttps = true
			} else {
				record.IsHttps = false
			}
			proxies = append(proxies, record)

		} else {
			log.Fatal("url Data format has changed")
		}

	})
	if len(proxies) > 0 {
		return &proxies, nil
	}
	return nil, errors.New("Data format has changed")
}

func ScrapeFreeProxyListNet(url string) (*[]model.ProxyRecord, error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	var proxies []model.ProxyRecord
	//use CSS selector found with the browser
	doc.Find("body tbody tr ").Each(func(index int, item *goquery.Selection) {
		var data map[int]string
		data = make(map[int]string)
		item.Find("td").Each(func(index2 int, item2 *goquery.Selection) {
			data[index2] = item2.Text()
		})
		record := model.ProxyRecord{}
		if len(data) == 8 {
			record.IP = data[0]
			record.Port, _ = strconv.Atoi(data[1])
			record.CountryCode = data[2]
			record.Country = data[3]
			if data[4] == "anonymous" {
				record.Type = model.Anoymous
			} else if data[4] == "elite proxy" {
				record.Type = model.Eliteproxy
			} else if data[4] == "transparent" {
				record.Type = model.Transparent
			} else {
				record.Type = model.Unknown
			}
			if data[6] == "yes" {
				record.IsHttps = true
			} else {
				record.IsHttps = false
			}
			proxies = append(proxies, record)

		} else {
			log.Fatal("url Data format has changed")
		}
	})
	if len(proxies) > 0 {
		return &proxies, nil
	}
	return nil, errors.New("Data format has changed")
}
