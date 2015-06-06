package hcutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/proxy"
)

var (
	sp = fmt.Sprintf
)

func NewClientSocks5(proxyAddr string) (*http.Client, error) {
	p, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		return nil, makeErr(err, "proxy error")
	}
	client := &http.Client{
		Transport: &http.Transport{
			Dial: p.Dial,
		},
	}
	return client, nil
}

var (
	DefaultRetryCount    = 8
	DefaultRetryInterval = time.Second * 3
)

func GetBytes(client *http.Client, url string) ([]byte, error) {
	retryCount := DefaultRetryCount
retry:
	resp, err := client.Get(url)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("get %s", url))
		}
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("get content of %s", url))
		}
	}
	return content, nil
}

func DoBytes(client *http.Client, req *http.Request) ([]byte, error) {
	retryCount := DefaultRetryCount
retry:
	resp, err := client.Do(req)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("do %v", req))
		}
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("get content of %v", req))
		}
	}
	return content, nil
}

func GetGoqueryDoc(client *http.Client, url string) (*goquery.Document, error) {
	retryCount := DefaultRetryCount
retry:
	resp, err := client.Get(url)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("get %s", url))
		}
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		if retryCount > 0 {
			retryCount--
			time.Sleep(DefaultRetryInterval)
			goto retry
		} else {
			return nil, makeErr(err, sp("get doc of %s", url))
		}
	}
	return doc, nil
}
