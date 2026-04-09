package api_test

import (
	"testing"

	"github.com/muchtar-syarief/scrap_desa/src/api"
	"github.com/muchtar-syarief/scrap_desa/src/session"
	"github.com/stretchr/testify/assert"
)

func TestRestBridgingApi(t *testing.T) {
	session := session.NewSession()
	bpsApi := api.NewSigBpsApi(api.BaseUri, session)

	query := api.RestBridgingQuery{
		Level:        api.RegionLevelType,
		Parent:       "35",
		PeriodeMerge: "2025_1.2025",
	}
	response, err := bpsApi.RestBridgingApi(&query)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)

	// raw, err := json.MarshalIndent(response, "", "	")
	// assert.Nil(t, err)
	// log.Println(string(raw))
}
