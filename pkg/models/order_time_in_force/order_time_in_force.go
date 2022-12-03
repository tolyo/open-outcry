package models

type OrderTimeInForce string

const (
	GTC OrderTimeInForce = "GTC"
	IOC OrderTimeInForce = "IOC"
	FOK OrderTimeInForce = "FOK"
	GTD OrderTimeInForce = "GTD"
	GTT OrderTimeInForce = "GTT"
)
