package db

func CreateSeeds() {

	db.Execute(`
        INSERT INTO currency(
          name,
          precision
        )
        VALUES (
          'BTC',
          5
        );
      `)

	db.Execute(`
        SELECT create_payment_account(
          'MASTER',
          'EUR'
        )
      `)

	db.Execute(`
        SELECT create_payment_account(
          'MASTER',
          'BTC'
        )
      `)

	models.CreateFxInstrument("BTC_EUR", "BTC", "EUR")
	models.CreateInstrument("SPX", "EUR")

}
