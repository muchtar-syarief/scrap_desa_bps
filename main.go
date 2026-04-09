package main

import (
	"log"
	"sync/atomic"

	"github.com/muchtar-syarief/scrap_desa/src/api"
	"github.com/muchtar-syarief/scrap_desa/src/helper"
	"github.com/muchtar-syarief/scrap_desa/src/parser"
	"github.com/muchtar-syarief/scrap_desa/src/session"
)

type Result struct {
	KodeProvinceDagri string `csv:"kode_provinsi_dagri"`
	KodeProvinceBps   string `csv:"kode_provinsi_bps"`

	KodeRegionDagri string `csv:"kode_kabkota_dagri"`
	KodeRegionBps   string `csv:"kode_kabkota_bps"`

	KodeDistrictDagri string `csv:"kode_kecamatan_dagri"`
	KodeDistrictBps   string `csv:"kode_kecamatan_bps"`

	KodeVillageDagri string `csv:"kode_desa_kelurahan_dagri"`
	KodeVillageBps   string `csv:"kode_desa_kelurahan_bps"`

	VillageNameDagri string `csv:"nama_desa_kelurahan_dagri"`
	VillageNameBps   string `csv:"nama_desa_kelurahan_bps"`
}

var PeriodeData = "2025_1.2025"

func main() {
	log.Println("[ START ] : Starting Process Data..... ")
	sess := session.NewSession()
	bpsApi := api.NewSigBpsApi(api.BaseUri, sess)

	var results = make([]*Result, 0)

	var counter atomic.Uint64

	provinceFilter, err := bpsApi.RestDropDownApi(&api.RestDropDownQuery{
		Level:        api.ProvinceLevelType,
		Parent:       PeriodeData,
		PeriodeMerge: PeriodeData,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, province := range provinceFilter {
		query := api.RestDropDownQuery{
			Level:        api.RegionLevelType,
			Parent:       province.Code,
			PeriodeMerge: PeriodeData,
		}
		regionFilter, err := bpsApi.RestDropDownApi(&query)
		if err != nil {
			log.Fatal(err)
			return
		}

		for _, region := range regionFilter {
			query := api.RestDropDownQuery{
				Level:        api.DistrictLevelType,
				Parent:       region.Code,
				PeriodeMerge: PeriodeData,
			}
			districtFilter, err := bpsApi.RestDropDownApi(&query)
			if err != nil {
				log.Fatal(err)
				return
			}

			for _, district := range districtFilter {
				query := api.RestBridgingQuery{
					Level:        api.VillageLevelType,
					Parent:       district.Code,
					PeriodeMerge: PeriodeData,
				}
				data, err := bpsApi.RestBridgingApi(&query)
				if err != nil {
					log.Fatal(err)
					return
				}

				for _, item := range data {
					kodeDagri := parser.KodeDagriType(item.KodeDagri)
					kodeBps := parser.KodeBpsType(item.KodeBps)

					result := &Result{
						KodeProvinceDagri: kodeDagri.GetProvinceCode(),
						KodeProvinceBps:   kodeBps.GetProvinceCode(),
						KodeRegionDagri:   kodeDagri.GetRegionCode(),
						KodeRegionBps:     kodeBps.GetRegionCode(),
						KodeDistrictDagri: kodeDagri.GetDistrictCode(),
						KodeDistrictBps:   kodeBps.GetDistrictCode(),
						KodeVillageDagri:  item.KodeDagri,
						KodeVillageBps:    item.KodeBps,
						VillageNameDagri:  item.NamaDagri,
						VillageNameBps:    item.NamaBps,
					}

					results = append(results, result)

					counter.Add(1)
					log.Printf("[ PROCESS ] : %d Process %s %s %s  Data........", counter.Load(), province.Code, region.Code, district.Code)
				}
			}
		}
	}

	log.Println("[ FINISH ] : Saving Data.......")
	file := "./data_desa_bps.csv"
	err = helper.SaveCSVFile(file, results)
	if err != nil {
		log.Fatal(err)
		return
	}
}
