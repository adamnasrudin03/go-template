package helpers

const (
	// Format date time
	FormatDateTime             = `2006-01-02 15:04:05`
	FormatDateHourMinutes      = `2006-01-02 15:04`
	FormatDate                 = `2006-01-02`
	FormatDateDDMMYYYY         = `02-01-2006`
	FormatLocalTime            = `02-Jan-2006`
	FormatLocalTimeDDsMMsYYYY  = `02 Jan 2006`
	FormatDateConcise          = `02-Jan-06`
	DateMonthFormat            = `02 January 2006`
	DateMonthFormatWithoutYear = `02 January`
	FormatTimeMinute           = `15:04`
	FormatTimeMinuteSecond     = `15:04:05`

	// Timezone country
	AsiaJakarta  = `Asia/Jakarta`
	AsiaMakassar = `Asia/Makassar`
	AsiaJayapura = `Asia/Jayapura`

	// indonesia country timezone in table city
	GMT7 = `GMT+7:00`
	GMT8 = `GMT+8:00`
	GMT9 = `GMT+9:00`

	// Month
	January   = `January`
	February  = `February`
	March     = `March`
	April     = `April`
	May       = `May`
	June      = `June`
	July      = `July`
	August    = `August`
	September = `September`
	October   = `October`
	November  = `November`
	December  = `December`
)

var (

	// loc
	loc = TimeZoneJakarta()

	// Mapping Timezone ToT imeLocation ...
	MappingTimezoneToTimeLocation = map[string]string{
		GMT7: AsiaJakarta,
		GMT8: AsiaMakassar,
		GMT9: AsiaJayapura,
	}

	// Month to Bulan
	MappingMonthToBulan = map[string]string{
		January:   `Januari`,
		February:  `Februari`,
		March:     `Maret`,
		April:     `April`,
		May:       `Mei`,
		June:      `Juni`,
		July:      `Juli`,
		August:    `Agustus`,
		September: `September`,
		October:   `Oktober`,
		November:  `November`,
		December:  `Desember`,
	}
)
