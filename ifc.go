/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package lunarcalendar

func GetLunarDateBySolar(year, month, day int32) (info CaYearInfo, isLeapMonth, ok bool) {
	isLeapMonth = false
	info, ok = getLunarDateBySolar(year, month, day)
	if !ok {
		return
	}

	leapMonth, _ := getLunarLeapYear(year)
	if leapMonth > 0 && leapMonth == info.Month {
		isLeapMonth = true
	} else if leapMonth > 0 && leapMonth > info.Month {
		info.Month += 1
	} else if leapMonth <= 0 {
		info.Month += 1
	}

	return
}

func GetSolarDateByLunar(year, month, day int32, isLeapMonth bool) (info CaYearInfo, ok bool) {
	leapMonth, _ := getLunarLeapYear(year)
	if leapMonth <= 0 {
		isLeapMonth = false
	}
	if (leapMonth > 0 && month > leapMonth) || isLeapMonth {
		month = month
	} else {
		month -= 1
	}
	info, ok = lunarToSolar(year, month, day)
	return
}

func GetLunarInfoBySolar(year, month, day int32) (info caLunarDayInfo, ok bool) {
	info, ok = solarToLunar(year, month, day)
	return
}

func GetSolarMonthCalendar(year, month int32, fill bool) (info caSolarMonthInfo, ok bool) {
	info, ok = getSolarCalendar(year, month, fill)
	return
}

func GetLunarMonthCalendar(year, month int32, fill bool) (info caLunarMonthInfo, ok bool) {
	info, ok = getLunarCalendar(year, month, fill)
	return
}
