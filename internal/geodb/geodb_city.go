package geodb

import (
	"embed"

	"github.com/oschwald/geoip2-golang"
	"go.uber.org/fx"
)

type GeoDBParams struct {
	fx.In
}

type GeoDBResult struct {
	fx.Out
	SearchGeoIPReaderCity *geoip2.Reader `name:"geoip_city"`
	SearchGeoIPReaderASN  *geoip2.Reader `name:"geoip_asn"`
}

//go:embed GeoLite2-City.mmdb
//go:embed GeoLite2-ASN.mmdb
var mmdbData embed.FS

func NewGeoDB(params GeoDBParams) GeoDBResult {
	mmdbBytes, err := mmdbData.ReadFile("GeoLite2-City.mmdb")
	db_city, err := geoip2.FromBytes(mmdbBytes)
	if err != nil {
		panic("GeoIP2-City.mmdb open failed!")
	}
	mmdbBytes, err = mmdbData.ReadFile("GeoLite2-ASN.mmdb")
	db_ASN, err := geoip2.FromBytes(mmdbBytes)
	if err != nil {
		panic("GeoIP2-ASN.mmdb open failed!")
	}
	return GeoDBResult{
		SearchGeoIPReaderCity: db_city,
		SearchGeoIPReaderASN:  db_ASN,
	}
}
