package types

type currType string

type Currency struct {
	From currType
	To   currType
}

const (
	CanadianDollars   currType = "CAD"
	Dollars           currType = "USD"
	MexicanPeso       currType = "MXN"
	GuatemaQuetzal    currType = "GTQ"
	BelizeDollars     currType = "BZD"
	HondurasCurrency  currType = "HNL"
	NicaraguaCurrency currType = "NIO"
	CostaCurrency     currType = "CRC"
	CubaCurrency      currType = "CUP"
	JamaicanDollars   currType = "JMD"
	DominicanCurrency currType = "DOP"
	BahamasDollars    currType = "BSD"
	BarbadosDollars   currType = "BBD"
	TrinidadDollars   currType = "TTD"
	AntiguaDollars    currType = "XCD"
	ColombianPeso     currType = "COP"
	GuyanaDollars     currType = "GYD"
	SurinameDollars   currType = "SRD"
	Euros             currType = "EUR"
	PeruvianCurrency  currType = "PEN"
	BrazilianCurrency currType = "BRL"
	BolivianCurrency  currType = "BOB"
	ParaguayCurrency  currType = "PYG"
	ChileCurrency     currType = "CLP"
	ArgentinaPeso     currType = "ARG"
	UruguayCurrency   currType = "UYU"
)
