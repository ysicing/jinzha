package utils

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	HttpClientSkipTlsVerify *http.Client
	HttpClient              *http.Client
)

func init() {
	HttpClientSkipTlsVerify = httpClient(_httpClient{
		Transport: httpTransport(_httpTransport{
			SkipVerifySSL: true,
		}),
	})
	HttpClient = httpClient(_httpClient{
		Transport: httpTransport(_httpTransport{
			SkipVerifySSL: false,
		}),
	})

	http.DefaultClient.Timeout = time.Minute * 10
}

type _httpTransport struct {
	SkipVerifySSL bool
}

func httpTransport(conf _httpTransport) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.SkipVerifySSL},
		Proxy:           http.ProxyFromEnvironment,
	}
}

type _httpClient struct {
	Transport *http.Transport
}

func httpClient(conf _httpClient) *http.Client {
	return &http.Client{
		Transport: conf.Transport,
		Timeout:   time.Minute * 10,
	}
}
