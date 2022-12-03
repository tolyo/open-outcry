package models

type OrderFill string

const (
	Full    OrderFill = "FULL"
	Partial OrderFill = "PARTIAL"
	None    OrderFill = "NONE"
)
