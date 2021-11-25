package util

import "time"

func GetMillsecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetSecond() int64 {
	return time.Now().UnixNano() / 1e9
}

func GetTotalDayByMonth(year uint64, month uint64) (bool, uint64) {
	if month < 1 || month > 13 || year <= 0 {
		return false, uint64(0)
	}

	if month == 13 {
		month = 1
		year = year + 1
	}

	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		return true, 31
	}

	if month == 2 {
		if IsLeapYear(year) {
			return true, 29
		} else {
			return true, 28
		}
	}

	return true, 30
}

func IsLeapYear(year uint64) bool {
	return (year%100 != 0 && year%4 == 0) || (year%400 == 0)
}
