package parser

type KodeBpsType string

const (
	ProvinceKodeBpsType KodeBpsType = "province"
	RegionKodeBpsType   KodeBpsType = "region"
	DistrictKodeBpsType KodeBpsType = "district"
	VillageKodeBpsType  KodeBpsType = "village"
)

// 35 05 220 001
var KodeBpsMap = map[KodeBpsType]int{
	ProvinceKodeBpsType: 2,
	RegionKodeBpsType:   4,
	DistrictKodeBpsType: 7,
	VillageKodeBpsType:  10,
}

func (k KodeBpsType) IsVillage() bool {
	return len(k) == KodeBpsMap[VillageKodeBpsType]
}

func (k KodeBpsType) GetProvinceCode() string {
	if len(k) >= KodeBpsMap[ProvinceKodeBpsType] {
		result := k[:KodeBpsMap[ProvinceKodeBpsType]]
		return string(result)
	}
	return ""
}

func (k KodeBpsType) GetRegionCode() string {
	if len(k) >= KodeBpsMap[RegionKodeBpsType] {
		result := k[:KodeBpsMap[RegionKodeBpsType]]
		return string(result)
	}
	return ""
}

func (k KodeBpsType) GetDistrictCode() string {
	if len(k) >= KodeBpsMap[DistrictKodeBpsType] {
		result := k[:KodeBpsMap[DistrictKodeBpsType]]
		return string(result)
	}
	return ""
}
