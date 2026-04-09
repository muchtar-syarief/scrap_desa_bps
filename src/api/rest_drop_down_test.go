package api_test

import (
	"testing"

	"github.com/muchtar-syarief/scrap_desa/src/api"
	"github.com/muchtar-syarief/scrap_desa/src/session"
	"github.com/stretchr/testify/assert"
)

func TestRestDropDownApi(t *testing.T) {
	session := session.NewSession()
	bpsApi := api.NewSigBpsApi(api.BaseUri, session)

	t.Run("test province", func(t *testing.T) {
		query := api.RestDropDownQuery{
			Level:        api.ProvinceLevelType,
			Parent:       "2025_1.2025",
			PeriodeMerge: "2025_1.2025",
		}
		response, err := bpsApi.RestDropDownApi(&query)
		assert.Nil(t, err)
		assert.NotEmpty(t, response)

		// raw, err := json.MarshalIndent(response, "", "	")
		// assert.Nil(t, err)
		// log.Println(string(raw))
	})

	t.Run("test region", func(t *testing.T) {
		query := api.RestDropDownQuery{
			Level:        api.RegionLevelType,
			Parent:       "35",
			PeriodeMerge: "2025_1.2025",
		}
		response, err := bpsApi.RestDropDownApi(&query)
		assert.Nil(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("test district", func(t *testing.T) {
		query := api.RestDropDownQuery{
			Level:        api.DistrictLevelType,
			Parent:       "3505",
			PeriodeMerge: "2025_1.2025",
		}
		response, err := bpsApi.RestDropDownApi(&query)
		assert.Nil(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("test village", func(t *testing.T) {
		query := api.RestDropDownQuery{
			Level:        api.VillageLevelType,
			Parent:       "3505220",
			PeriodeMerge: "2025_1.2025",
		}
		response, err := bpsApi.RestDropDownApi(&query)
		assert.Nil(t, err)
		assert.NotEmpty(t, response)
	})

	// raw, err := json.MarshalIndent(response, "", "	")
	// assert.Nil(t, err)
	// log.Println(string(raw))
}
