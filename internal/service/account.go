package service

//Fiat Currency Account Types 
const USDACCOUNTTYPE = "USD"
const NAIRAACCOUNTTYPE = "NAIRA"

type Account struct{
	Type string		`json:"type"`
}

