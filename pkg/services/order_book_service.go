package services

import (
	log "github.com/sirupsen/logrus"
	"open-outcry/pkg/db"
	"open-outcry/pkg/models"
)

type PriceVolume struct {
	Price  float64
	Volume float64
}

func GetVolumeAtPrice(instrumentName models.InstrumentName, side models.OrderSide, price models.OrderPrice) float64 {
	res := db.QueryVal[float64](`
	 SELECT volume
	 FROM price_level
	 WHERE side = $2
	   AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
	   AND price =  $3
   `, instrumentName, side, price)
	return res
}

func GetVolumes(instrumentName models.InstrumentName, side models.OrderSide) []PriceVolume {
	var orderBy string
	switch side {
	case models.Sell:
		orderBy = "ASC"
	case models.Buy:
		orderBy = "DESC"
	}
	rows, err := db.Instance().Query(`
		SELECT price, volume
		FROM price_level
		WHERE side = $2
		AND price > 0
		AND instrument_id = (SELECT id FROM instrument WHERE name = $1)
		ORDER BY price `+orderBy,
		instrumentName,
		side,
	)
	if err != nil {
		log.Fatal(err)
	}

	res := make([]PriceVolume, 0)
	for rows.Next() {
		var priceVolume PriceVolume
		rows.Scan(&priceVolume.Price, &priceVolume.Volume)
		res = append(res, priceVolume)
	}
	return res
}
