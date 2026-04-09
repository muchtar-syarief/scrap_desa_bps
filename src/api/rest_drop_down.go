package api

import (
	"net/url"
)

type LevelType string

const (
	ProvinceLevelType LevelType = "provinsi"
	RegionLevelType   LevelType = "kabupaten"
	DistrictLevelType LevelType = "kecamatan"
	VillageLevelType  LevelType = "desa"
)

type RestDropDownQuery struct {
	Level        LevelType `schema:"level"`
	Parent       string    `schema:"parent"`
	PeriodeMerge string    `schema:"periode_merge"`
}

type RestDropDown struct {
	Code string `json:"kode"`
	Name string `json:"nama"`
}

type RestDropDownResp []*RestDropDown

func (api *sigBpsApi) RestDropDownApi(query *RestDropDownQuery) (RestDropDownResp, error) {
	uri := "/rest-drop-down/getwilayah"
	endpoint, err := url.JoinPath(api.BaseUri, uri)
	if err != nil {
		return nil, err
	}

	req, err := api.NewRequest("GET", endpoint, query, nil, nil)
	if err != nil {
		return nil, err
	}

	var result RestDropDownResp
	err = api.SendRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
