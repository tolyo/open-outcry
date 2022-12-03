package models

type CurrencyName string

type CurrencyPrecision decimal

type Currency struct {
	Name      CurrencyName
	Precision CurrencyPrecision
}

const baseQuery = `
    SELECT(
        c.name,
        c.precision
    )

    FROM currency AS c
`

func Exists(name CurrencyName) bool {
	db.QueryExists(baseQuery+"WHERE c.name = $1", name)
}
