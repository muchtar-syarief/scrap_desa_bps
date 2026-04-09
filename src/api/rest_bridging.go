package api

import (
	"net/url"
)

type RestBridgingQuery struct {
	Level        LevelType `schema:"level"`
	Parent       string    `schema:"parent"`
	PeriodeMerge string    `schema:"periode_merge"`
}

type RestBridging struct {
	KodeBps   string `json:"kode_bps"`
	NamaBps   string `json:"nama_bps"`
	KodeDagri string `json:"kode_dagri"`
	NamaDagri string `json:"nama_dagri"`
}

type RestBridgingResp []*RestBridging

func (api *sigBpsApi) RestBridgingApi(query *RestBridgingQuery) (RestBridgingResp, error) {
	uri := "/rest-bridging/getwilayah"
	endpoint, err := url.JoinPath(api.BaseUri, uri)
	if err != nil {
		return nil, err
	}

	req, err := api.NewRequest("GET", endpoint, query, nil, nil)
	if err != nil {
		return nil, err
	}

	var result RestBridgingResp
	err = api.SendRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
