package helper

func MonthConverter(month string) string {
	var result string

	switch month {
	case "January":
		result = "1"

	case "February":
		result = "2"

	case "March":
		result = "3"

	case "April":
		result = "4"

	case "May":
		result = "5"
	
	case "June":
		result = "6"

	case "July":
		result = "7"

	case "August":
		result = "8"

	case "September":
		result = "9"

	case "October":
		result = "10"

	case "November":
		result = "11"

	case "December":
		result = "12"
	}

	return result
}
