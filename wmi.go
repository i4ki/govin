package govin

import "errors"

type (
	wmiRegionRange struct {
		low1, low2         rune
		high1, high2       rune
		continent, country string
	}

	WMIManufacturer struct {
		code string
		name string
	}

	WMI struct {
		code, continent, country string
		manufacturer             WMIManufacturer
	}

	WMIMap map[string]WMI
)

var (
	codes = []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N',
		'P', 'R', 'S', 'T', 'U', 'V', 'X', 'Y', 'Z', '1', '2', '3', '4',
		'5', '6', '7', '8', '9', '0',
	}

	countryRanges = append(africaRanges, samericaRanges...)

	wmiMap WMIMap
)

func init() {
	var err error

	wmiMap, err = genWMIMap()

	if err != nil {
		panic(err)
	}
}

func toCode(r rune) int {
	for i := 0; i < len(codes); i++ {
		if codes[i] == r {
			return i
		}
	}

	return -1
}

func genWMIMap() (WMIMap, error) {
	return genWMIFromRanges(countryRanges)
}

func genWMIFromRanges(cranges []wmiRegionRange) (WMIMap, error) {
	wmiMap := make(WMIMap)

	if len(cranges) == 0 {
		return nil, errors.New("No range")
	}

	for r := 0; r < len(cranges); r++ {
		rng := cranges[r]

		for i := toCode(rng.low1); i <= toCode(rng.high1); i++ {
			for j := toCode(rng.low2); j <= toCode(rng.high2); j++ {
				wmi := WMI{}
				ci := codes[i]
				cj := codes[j]

				wmi.code = string([]rune{ci, cj})

				wmi.continent = rng.continent
				wmi.country = rng.country

				wmiMap[wmi.code] = wmi
			}
		}
	}

	return wmiMap, nil
}

func GetWMIMap() WMIMap { return wmiMap }
