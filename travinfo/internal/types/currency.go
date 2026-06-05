package types

type CurrType string

type Currency struct {
	From         string
	FromCurrency CurrType
	To           string
	ToCurrency   CurrType
	ExchangeRate float64
	InverseRate  float64
}

const (
	CanadianDollars   CurrType = "CAD"
	Dollars           CurrType = "USD"
	MexicanPeso       CurrType = "MXN"
	GuatemaQuetzal    CurrType = "GTQ"
	BelizeDollars     CurrType = "BZD"
	HondurasCurrency  CurrType = "HNL"
	NicaraguaCurrency CurrType = "NIO"
	CostaCurrency     CurrType = "CRC"
	CubaCurrency      CurrType = "CUP"
	JamaicanDollars   CurrType = "JMD"
	DominicanCurrency CurrType = "DOP"
	BahamasDollars    CurrType = "BSD"
	BarbadosDollars   CurrType = "BBD"
	TrinidadDollars   CurrType = "TTD"
	AntiguaDollars    CurrType = "XCD"
	ColombianPeso     CurrType = "COP"
	GuyanaDollars     CurrType = "GYD"
	SurinameDollars   CurrType = "SRD"
	Euros             CurrType = "EUR"
	PeruvianCurrency  CurrType = "PEN"
	BrazilianCurrency CurrType = "BRL"
	BolivianCurrency  CurrType = "BOB"
	ParaguayCurrency  CurrType = "PYG"
	ChileCurrency     CurrType = "CLP"
	ArgentinaPeso     CurrType = "ARS"
	UruguayCurrency   CurrType = "UYU"
)
