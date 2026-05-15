package utils

import "github.com/Raghart/traveling_planer/travinfo/internal/types"

func LoadLocations() map[string]types.CountryLocation {
	mapLocations := map[string]types.CountryLocation{}
	mapLocations["Argentina"] = types.CountryLocation{
		Latitude:  -38.416097,
		Longitude: -63.616672,
	}
	mapLocations["Bolivia"] = types.CountryLocation{
		Latitude:  -16.290154,
		Longitude: -63.588653,
	}
	mapLocations["Brazil"] = types.CountryLocation{
		Latitude:  -14.235004,
		Longitude: -51.92528,
	}
	mapLocations["Canada"] = types.CountryLocation{
		Latitude:  56.130366,
		Longitude: -106.346771,
	}
	mapLocations["Chile"] = types.CountryLocation{
		Latitude:  -35.675147,
		Longitude: -71.542969,
	}
	mapLocations["Colombia"] = types.CountryLocation{
		Latitude:  4.570868,
		Longitude: -74.297333,
	}
	mapLocations["Costa Rica"] = types.CountryLocation{
		Latitude:  9.748917,
		Longitude: -83.753428,
	}
	mapLocations["Cuba"] = types.CountryLocation{
		Latitude:  21.521757,
		Longitude: -77.781167,
	}
	mapLocations["Dominica"] = types.CountryLocation{
		Latitude:  15.414999,
		Longitude: -61.370976,
	}
	mapLocations["Dominican Republic"] = types.CountryLocation{
		Latitude:  18.735693,
		Longitude: -70.162651,
	}
	mapLocations["Ecuador"] = types.CountryLocation{
		Latitude:  -1.831239,
		Longitude: -78.183406,
	}
	mapLocations["Grenada"] = types.CountryLocation{
		Latitude:  12.262776,
		Longitude: -61.604171,
	}
	mapLocations["French Guiana"] = types.CountryLocation{
		Latitude:  3.933889,
		Longitude: -53.125782,
	}
	mapLocations["Guyana"] = types.CountryLocation{
		Latitude:  4.860416,
		Longitude: -58.93018,
	}
	mapLocations["Honduras"] = types.CountryLocation{
		Latitude:  15.199999,
		Longitude: -86.241905,
	}
	mapLocations["Saint Lucia"] = types.CountryLocation{
		Latitude:  13.909444,
		Longitude: -60.978893,
	}
	mapLocations["Mexico"] = types.CountryLocation{
		Latitude:  23.634501,
		Longitude: -102.552784,
	}
	mapLocations["Nicaragua"] = types.CountryLocation{
		Latitude:  12.865416,
		Longitude: -85.207229,
	}
	mapLocations["Panama"] = types.CountryLocation{
		Latitude:  8.537981,
		Longitude: -80.782127,
	}
	mapLocations["Peru"] = types.CountryLocation{
		Latitude:  -9.189967,
		Longitude: -75.015152,
	}
	mapLocations["Puerto Rico"] = types.CountryLocation{
		Latitude:  18.220833,
		Longitude: -66.590149,
	}
	mapLocations["Paraguay"] = types.CountryLocation{
		Latitude:  -23.442503,
		Longitude: -58.443832,
	}
	mapLocations["Suriname"] = types.CountryLocation{
		Latitude:  3.919305,
		Longitude: -56.027783,
	}
	mapLocations["El Salvador"] = types.CountryLocation{
		Latitude:  13.794185,
		Longitude: -88.89653,
	}
	mapLocations["Trinidad and Tobago"] = types.CountryLocation{
		Latitude:  10.691803,
		Longitude: -61.222503,
	}
	mapLocations["United States"] = types.CountryLocation{
		Latitude:  37.09024,
		Longitude: -95.712891,
	}
	mapLocations["Uruguay"] = types.CountryLocation{
		Latitude:  -32.522779,
		Longitude: -55.765835,
	}
	mapLocations["Venezuela"] = types.CountryLocation{
		Latitude:  6.42375,
		Longitude: -66.58973,
	}
	mapLocations["Guatemala"] = types.CountryLocation{
		Latitude:  15.783471,
		Longitude: -90.230759,
	}
	mapLocations["Belize"] = types.CountryLocation{
		Latitude:  17.189877,
		Longitude: -88.49765,
	}
	mapLocations["Jamaica"] = types.CountryLocation{
		Latitude:  18.109581,
		Longitude: -77.297508,
	}
	mapLocations["Haiti"] = types.CountryLocation{
		Latitude:  18.971187,
		Longitude: -72.285215,
	}
	mapLocations["Bahamas"] = types.CountryLocation{
		Latitude:  25.03428,
		Longitude: -77.39628,
	}
	mapLocations["Barbados"] = types.CountryLocation{
		Latitude:  13.193887,
		Longitude: -59.543198,
	}
	mapLocations["Saint Kitts and Nevis"] = types.CountryLocation{
		Latitude:  17.357822,
		Longitude: -62.782998,
	}
	mapLocations["Antigua and Barbuda"] = types.CountryLocation{
		Latitude:  17.0608,
		Longitude: -61.7964,
	}

	return mapLocations
}
