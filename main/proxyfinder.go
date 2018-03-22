package main

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"

	"os"

	"fmt"
	"github.com/richardsang2008/proxy_finder/model"
	"net"
	"net/url"

	"net/http"
	"time"

	"github.com/richardsang2008/proxy_finder/controller"
)

//"https://www.us-proxy.org/","https://free-proxy-list.net/"
var us_proxy_url = [1]string{
	"https://www.us-proxy.org/"}
var niaticurl = "https://club.pokemon.com/us/pokemon-trainer-club/sign-up/"
var hidemyname_urls = [5]string{
	"https://hidemy.name/en/proxy-list/?country=US#list",
	"https://hidemy.name/en/proxy-list/?country=US&start=64#list",
	"https://hidemy.name/en/proxy-list/?country=US&start=128#list",
	"https://hidemy.name/en/proxy-list/?country=US&start=192#list",
	"https://hidemy.name/en/proxy-list/?country=US&start=256#list"}
var proxiesChannel = make(chan model.ProxyRecord)

func CheckProxyForCreate(record model.ProxyRecord) {

	proxy := fmt.Sprintf("http://%s:%d", record.IP, record.Port)
	fmt.Printf("processing %s   ==> \n",record.ToString())
	//only for testing
	//proxy ="http://138.197.192.64:65000"

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		record.IsCreateAccountOk = false
	}
	log.Println(fmt.Sprintf("call using %s", proxy))
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 20 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		Proxy:               http.ProxyURL(proxyUrl),
	}
	netClient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: netTransport,
	}

	resp, err1 := netClient.Get(niaticurl)

	if err1 != nil {
		record.IsCreateAccountOk = false
		return
	} else {

		if resp.StatusCode == 200 {
			record.IsCreateAccountOk = true
			log.Println("success", proxy)
			fmt.Printf("insert %s \n",record.ToString())
			proxiesChannel <- record
		} else {
			record.IsCreateAccountOk = false
		}
		return
	}
	defer resp.Body.Close()
}
func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/proxy_finder.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	proxies, err := controller.ScrapeFreeProxyListNet(us_proxy_url[0])
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("out/proxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, item := range *proxies {
		go CheckProxyForCreate(item)

	}
	//<-proxiesChannel
	i :=1
	for proxy := range proxiesChannel {
		fmt.Printf("/v",i)
		i++
		if proxy.IsCreateAccountOk {
			proxystr := proxy.ToString() + ","
			fmt.Fprintf(f, proxystr)
		}
	}
	close(proxiesChannel)
}
