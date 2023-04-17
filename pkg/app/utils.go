package app

func IsStringEmpty(s string) bool {
	return len(s) == 0
}

func GetAQIDescription(aqi int) string {
	switch aqi {
	case 1:
		return "Good"
	case 2:
		return "Fair"
	case 3:
		return "Moderate"
	case 4:
		return "Poor"
	case 5:
		return "Very poor"
	}

	return ""
}
