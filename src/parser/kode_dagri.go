package parser

import (
	"strings"
)

type KodeDagriType string

const (
	ProvinceKodeDagriType KodeDagriType = "province"
	RegionKodeDagriType   KodeDagriType = "region"
	DistrictKodeDagriType KodeDagriType = "district"
	VillageKodeDagriType  KodeDagriType = "village"
)

var KodeDagriMap = map[KodeDagriType]int{
	ProvinceKodeDagriType: 1,
	RegionKodeDagriType:   2,
	DistrictKodeDagriType: 3,
	VillageKodeDagriType:  4,
}

func (k KodeDagriType) GetProvinceCode() string {
	if k == "0" {
		return "0"
	}

	data := strings.Split(string(k), ".")
	if len(data) >= 1 {
		return data[0]
	}
	return ""
}

func (k KodeDagriType) GetRegionCode() string {
	if k == "0" {
		return "0"
	}

	data := strings.Split(string(k), ".")
	if len(data) >= 2 {
		return strings.Join(data[:2], ".")
	}
	return ""
}

func (k KodeDagriType) GetDistrictCode() string {
	if k == "0" {
		return "0"
	}

	data := strings.Split(string(k), ".")
	if len(data) >= 3 {
		return strings.Join(data[:3], ".")
	}
	return ""
}
