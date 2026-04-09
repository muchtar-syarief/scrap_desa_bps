package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/muchtar-syarief/scrap_desa/src/session"
)

var ClientApi *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 5,
		// Proxy: http.ProxyURL(&url.URL{
		// 	Scheme: "http",
		// 	Host:   "localhost:8888",
		// }),
	},
	Timeout: 30 * time.Second,
}

var encoder = schema.NewEncoder()

var defaultHeader = map[string]string{
	"Accept": "application/json",
	// "Accept-Encoding":    "gzip, deflate, br, zstd",
	"Accept-Language":    "en-US,en;q=0.9,id;q=0.8",
	"Cache-Control":      "no-cache",
	"Connection":         "keep-alive",
	"Host":               "sig.bps.go.id",
	"Pragma":             "no-cache",
	"Referer":            "https://sig.bps.go.id/bridging-kode/index",
	"Sec-Fetch-Dest":     "empty",
	"Sec-Fetch-Mode":     "cors",
	"Sec-Fetch-Site":     "same-origin",
	"X-Requested-With":   "XMLHttpRequest",
	"sec-ch-ua":          `"Chromium";v="146", "Not-A.Brand";v="24", "Google Chrome";v="146"`,
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": `"Windows"`,
}

type SigBpsApi interface {
	RestDropDownApi(query *RestDropDownQuery) (RestDropDownResp, error)
	RestBridgingApi(query *RestBridgingQuery) (RestBridgingResp, error)
}

type sigBpsApi struct {
	BaseUri string
	session session.Session
}

var BaseUri = "https://sig.bps.go.id"

func NewSigBpsApi(baseUri string, session session.Session) SigBpsApi {
	return &sigBpsApi{
		BaseUri: baseUri,
		session: session,
	}
}

func (api *sigBpsApi) NewRequest(method, uri string, query any, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()
		encoder.Encode(query, q)

		req.URL.RawQuery = q.Encode()
	}

	for key, header := range defaultHeader {
		req.Header.Set(key, header)
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	req.Header.Set("User-Agent", api.session.GetUserAgent())

	api.session.AddCookiesToRequest(req)

	return req, nil
}

func (api *sigBpsApi) SendRequest(req *http.Request, result any) error {
	res, err := ClientApi.Do(req)
	if err != nil {
		return err
	}

	if result != nil {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, result)
		if err != nil {
			return err
		}
	}

	return api.session.UpdateCookies(res.Cookies())
}
