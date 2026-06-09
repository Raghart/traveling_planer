package types

type CountryFestivities struct {
	Country     string
	Festivities []FestivityData
}

type CountryDescription struct {
	Name        string
	Description string
}

type CountryAsciiImg struct {
	Name  string
	Image string
}

type FestivityData struct {
	Date        string
	LocalName   string
	Name        string
	CountryCode string
	Fixed       bool
	Global      bool
	Types       []string
}

type CountryJsonFestivities []struct {
	Date        string   `json:"date"`
	LocalName   string   `json:"localName"`
	Name        string   `json:"name"`
	CountryCode string   `json:"countryCode"`
	Fixed       bool     `json:"fixed"`
	Global      bool     `json:"global"`
	Types       []string `json:"types"`
}

type WikipediaJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		Normalized []struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"normalized"`
		Pages map[string]struct {
			Pageid  int    `json:"pageid"`
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

type PhotoJsonURL struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		ID             string `json:"id"`
		Slug           string `json:"slug"`
		Width          int    `json:"width"`
		Height         int    `json:"height"`
		Color          string `json:"color"`
		Description    string `json:"description"`
		AltDescription string `json:"alt_description"`
		Urls           struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
			SmallS3 string `json:"small_s3"`
		} `json:"urls"`
		Likes int `json:"likes"`
	} `json:"results"`
}

type CurrencyJSON struct {
	Result          string `json:"result"`
	BaseCode        string `json:"base_code"`
	ConversionRates struct {
		USD float64 `json:"USD"`
		AED float64 `json:"AED"`
		AFN float64 `json:"AFN"`
		ALL float64 `json:"ALL"`
		AMD float64 `json:"AMD"`
		ANG float64 `json:"ANG"`
		AOA float64 `json:"AOA"`
		ARS float64 `json:"ARS"`
		AUD float64 `json:"AUD"`
		AWG float64 `json:"AWG"`
		AZN float64 `json:"AZN"`
		BAM float64 `json:"BAM"`
		BBD float64 `json:"BBD"`
		BDT float64 `json:"BDT"`
		BGN float64 `json:"BGN"`
		BHD float64 `json:"BHD"`
		BIF float64 `json:"BIF"`
		BMD float64 `json:"BMD"`
		BND float64 `json:"BND"`
		BOB float64 `json:"BOB"`
		BRL float64 `json:"BRL"`
		BSD float64 `json:"BSD"`
		BTN float64 `json:"BTN"`
		BWP float64 `json:"BWP"`
		BYN float64 `json:"BYN"`
		BZD float64 `json:"BZD"`
		CAD float64 `json:"CAD"`
		CDF float64 `json:"CDF"`
		CHF float64 `json:"CHF"`
		CLF float64 `json:"CLF"`
		CLP float64 `json:"CLP"`
		CNH float64 `json:"CNH"`
		CNY float64 `json:"CNY"`
		COP float64 `json:"COP"`
		CRC float64 `json:"CRC"`
		CUP float64 `json:"CUP"`
		CVE float64 `json:"CVE"`
		CZK float64 `json:"CZK"`
		DJF float64 `json:"DJF"`
		DKK float64 `json:"DKK"`
		DOP float64 `json:"DOP"`
		DZD float64 `json:"DZD"`
		EGP float64 `json:"EGP"`
		ERN float64 `json:"ERN"`
		ETB float64 `json:"ETB"`
		EUR float64 `json:"EUR"`
		FJD float64 `json:"FJD"`
		FKP float64 `json:"FKP"`
		FOK float64 `json:"FOK"`
		GBP float64 `json:"GBP"`
		GEL float64 `json:"GEL"`
		GGP float64 `json:"GGP"`
		GHS float64 `json:"GHS"`
		GIP float64 `json:"GIP"`
		GMD float64 `json:"GMD"`
		GNF float64 `json:"GNF"`
		GTQ float64 `json:"GTQ"`
		GYD float64 `json:"GYD"`
		HKD float64 `json:"HKD"`
		HNL float64 `json:"HNL"`
		HRK float64 `json:"HRK"`
		HTG float64 `json:"HTG"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		IMP float64 `json:"IMP"`
		INR float64 `json:"INR"`
		IQD float64 `json:"IQD"`
		IRR float64 `json:"IRR"`
		ISK float64 `json:"ISK"`
		JEP float64 `json:"JEP"`
		JMD float64 `json:"JMD"`
		JOD float64 `json:"JOD"`
		JPY float64 `json:"JPY"`
		KES float64 `json:"KES"`
		KGS float64 `json:"KGS"`
		KHR float64 `json:"KHR"`
		KID float64 `json:"KID"`
		KMF float64 `json:"KMF"`
		KRW float64 `json:"KRW"`
		KWD float64 `json:"KWD"`
		KYD float64 `json:"KYD"`
		KZT float64 `json:"KZT"`
		LAK float64 `json:"LAK"`
		LBP float64 `json:"LBP"`
		LKR float64 `json:"LKR"`
		LRD float64 `json:"LRD"`
		LSL float64 `json:"LSL"`
		LYD float64 `json:"LYD"`
		MAD float64 `json:"MAD"`
		MDL float64 `json:"MDL"`
		MGA float64 `json:"MGA"`
		MKD float64 `json:"MKD"`
		MMK float64 `json:"MMK"`
		MNT float64 `json:"MNT"`
		MOP float64 `json:"MOP"`
		MRU float64 `json:"MRU"`
		MUR float64 `json:"MUR"`
		MVR float64 `json:"MVR"`
		MWK float64 `json:"MWK"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		MZN float64 `json:"MZN"`
		NAD float64 `json:"NAD"`
		NGN float64 `json:"NGN"`
		NIO float64 `json:"NIO"`
		NOK float64 `json:"NOK"`
		NPR float64 `json:"NPR"`
		NZD float64 `json:"NZD"`
		OMR float64 `json:"OMR"`
		PAB float64 `json:"PAB"`
		PEN float64 `json:"PEN"`
		PGK float64 `json:"PGK"`
		PHP float64 `json:"PHP"`
		PKR float64 `json:"PKR"`
		PLN float64 `json:"PLN"`
		PYG float64 `json:"PYG"`
		QAR float64 `json:"QAR"`
		RON float64 `json:"RON"`
		RSD float64 `json:"RSD"`
		RUB float64 `json:"RUB"`
		RWF float64 `json:"RWF"`
		SAR float64 `json:"SAR"`
		SBD float64 `json:"SBD"`
		SCR float64 `json:"SCR"`
		SDG float64 `json:"SDG"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		SHP float64 `json:"SHP"`
		SLE float64 `json:"SLE"`
		SLL float64 `json:"SLL"`
		SOS float64 `json:"SOS"`
		SRD float64 `json:"SRD"`
		SSP float64 `json:"SSP"`
		STN float64 `json:"STN"`
		SYP float64 `json:"SYP"`
		SZL float64 `json:"SZL"`
		THB float64 `json:"THB"`
		TJS float64 `json:"TJS"`
		TMT float64 `json:"TMT"`
		TND float64 `json:"TND"`
		TOP float64 `json:"TOP"`
		TRY float64 `json:"TRY"`
		TTD float64 `json:"TTD"`
		TVD float64 `json:"TVD"`
		TWD float64 `json:"TWD"`
		TZS float64 `json:"TZS"`
		UAH float64 `json:"UAH"`
		UGX float64 `json:"UGX"`
		UYU float64 `json:"UYU"`
		UZS float64 `json:"UZS"`
		VES float64 `json:"VES"`
		VND float64 `json:"VND"`
		VUV float64 `json:"VUV"`
		WST float64 `json:"WST"`
		XAF float64 `json:"XAF"`
		XCD float64 `json:"XCD"`
		XCG float64 `json:"XCG"`
		XDR float64 `json:"XDR"`
		XOF float64 `json:"XOF"`
		XPF float64 `json:"XPF"`
		YER float64 `json:"YER"`
		ZAR float64 `json:"ZAR"`
		ZMW float64 `json:"ZMW"`
		ZWG float64 `json:"ZWG"`
		ZWL float64 `json:"ZWL"`
	} `json:"conversion_rates"`
}
