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

package main

import (
	"fmt"
	"time"
)

type yearInfo struct {
	year  int
	month int
	day   int
}

func isYearValid(year int) bool {
	if year > MaxYear || year < MinYear {
		fmt.Printf("Invalid Year: %d. Year Range(%d - %d)\n",
			year, MinYear, MaxYear)
		return false
	}

	return true
}

/**
 * 判断农历年闰月数
 * @param {Number} year 农历年
 * return 闰月数 （月份从1开始）
 */
func getLunarLeapYear(year int) (int, bool) {
	if !isYearValid(year) {
		return -1, false
	}

	info := lunarInfos[year-MinYear]
	//fmt.Println("LeapMonth: ", info.leapMonth)
	//fmt.Println("MonthNum: ", info.lunarMonthNum)
	//fmt.Println("ZhengMonth: ", info.zhengMonth)
	//fmt.Println("ZhengDay: ", info.zhengDay)
	return info.leapMonth, true
}

/**
 * 获取农历年每月的天数及一年的总天数
 */
func getLunarYearDays(year int) ([]dateDayInfo, int, bool) {
	if !isYearValid(year) {
		return nil, -1, false
	}

	info := lunarInfos[year-MinYear]
	leapMonth := info.leapMonth
	monthData := fmt.Sprintf("%b", info.lunarMonthNum)
	//fmt.Println("Month Bianry before insert: ", monthData)
	tmp := ""
	l := len(monthData)
	//还原数据至16位,少于16位的在前面插入0（二进制存储时前面的0被忽略)
	for i := 0; i < 16-l; i++ {
		tmp += "0"
	}
	monthData = tmp + monthData
	//fmt.Println("Month Bianry after insert: ", monthData)

	monthNum := 0
	if leapMonth > 0 {
		monthNum = 13
	} else {
		monthNum = 12
	}

	yearDays := 0
	monthDayInfos := []dateDayInfo{}
	for i := 0; i < monthNum; i++ {
		tmp := dateDayInfo{}
		if monthData[i] == '0' {
			yearDays += 29
			tmp.days = 29
		} else {
			yearDays += 30
			tmp.days = 30
		}
		// 让月份从1开始，不从0开始
		t := i + 1
		// 处理闰月
		if i >= leapMonth {
			t -= 1
		}
		tmp.index = t
		monthDayInfos = append(monthDayInfos, tmp)
	}

	return monthDayInfos, yearDays, true
}

/**
 * 通过间隔天数查找农历日期
 */
func getLunarDateByBetween(year, between int) (yearInfo, bool) {
	month := int(-1)
	day := int(-1)
	monthDayInfos, yearDays, ok := getLunarYearDays(year)
	if !ok {
		fmt.Println("Get Year Days Failed For Year: ", year)
		return yearInfo{year, month, day}, false
	}

	end := int(0)
	if between > 0 {
		end = between
	} else {
		end = yearDays - between
	}
	fmt.Println("Between: ", end)
	tmpDays := int(0)
	for _, info := range monthDayInfos {
		tmpDays += info.days
		//fmt.Println("\tTmp: ", tmpDays)
		if tmpDays > end {
			month = info.index
			tmpDays = tmpDays - info.days
			break
		}
	}
	day = end - tmpDays + 1

	return yearInfo{year, month, day}, true
}

/**
 * 通过公历日期获取农历日期
 */
func getLunarDateBySolar(year, month, day int) (yearInfo, bool) {
	if !isYearValid(year) {
		return yearInfo{-1, -1, -1}, false
	}

	info := lunarInfos[year-MinYear]
	zengMonth := info.zhengMonth
	zengDay := info.zhengDay
	between, _ := getDaysBetweenSolar(year, zengMonth, zengDay,
		year, month, day)
	if between == 0 { //正月初一
		return yearInfo{year, 1, 1}, true
	} else if between < 0 {
		year -= 1
	}
	return getLunarDateByBetween(year, int(between))
}

/**
 * 计算两个公历日期之间的天数
 */
func getDaysBetweenSolar(year, month, day, year1, month1, day1 int) (int64, bool) {
	date := time.Date(int(year), time.Month(month), int(day),
		0, 0, 0, 0, time.UTC).Unix()
	date1 := time.Date(int(year1), time.Month(month1), int(day1),
		0, 0, 0, 0, time.UTC).Unix()

	return (date1 - date) / 86400, true
}

/**
 * 计算农历日期离正月初一有多少天
 */
func getDaysBetweenZheng(year, month, day int) (int, bool) {
	monthDayInfos, _, ok := getLunarYearDays(year)
	if !ok {
		fmt.Println("Get Year Days Failed For Year: ", year)
		return -1, false
	}

	days := int(0)
	for _, info := range monthDayInfos {
		if info.index < month {
			days += info.days
		} else {
			break
		}
	}

	return days + day - 1, true
}

func formatDayD4(month, day int) string {
	monStr := ""
	dayStr := ""
	if month < 10 {
		monStr = fmt.Sprintf("0%d", month)
	} else {
		monStr = fmt.Sprintf("%d", month)
	}

	if day < 10 {
		dayStr = fmt.Sprintf("0%d", day)
	} else {
		dayStr = fmt.Sprintf("%d", day)
	}

	return fmt.Sprintf("d%s%s", monStr, dayStr)
}

/**
 * 某年的第n个节气为几日
 * 31556925974.7为地球公转周期，是毫秒
 * 1890年的正小寒点：01-05 16:02:31，1890年为基准点
 * year 公历年
 * n 第几个节气，从0小寒起算
 * 由于农历24节气交节时刻采用近似算法，可能存在少量误差(30分钟内)
 */
func getTermDate(year, n int) (yearInfo, bool) {
	if !isYearValid(year) {
		return yearInfo{}, false
	}

	offset := 31556925974/1000*(int64(year)-1890) + int64(termInfo[n])*60 + time.Date(1890, 1, 5, 16, 2, 31, 0, time.UTC).Unix()
	y, m, d := time.Unix(offset, 0).Date()

	return yearInfo{y, int(m), d}, true
}

/**
 * 获取公历年一年的二十四节气
 * 返回key:日期，value:节气中文名
 */
func getYearTerm(year int) map[string]string {
	res := make(map[string]string)
	month := 0
	for i := 0; i < 24; i++ {
		if info, ok := getTermDate(year, i); !ok {
			continue
		} else {
			// 每个月中有两个节气
			if i < 2 {
				month = 1
			} else {
				month = i/2 + 1
			}
			res[formatDayD4(month, info.day)] = lunarData["solarTerm"][i]
		}
	}

	return res
}

/**
 * 获取生肖
 * year: 干支所在年(默认以立春前的公历年作为基数)
 */
func getYearZodiac(year int) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	num := year - 1890 + 25 //参考干支纪年的计算，生肖对应地支
	return lunarData["zodiac"][num%12], true
}

/**
 * 计算天干地支
 * num 60进制中的位置(把60个天干地支，当成一个60进制的数)
 */
func cyclical(num int) (string, bool) {
	return lunarData["heavenlyStems"][num%10] + lunarData["earthlyBranches"][num%12], true
}

/**
 * 获取干支纪年
 * year 干支所在年
 * offset 偏移量，默认为0，便于查询一个年跨两个干支纪年(以立春为分界线)
 */
func getLunarYearName(year, offset int) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	offset = offset | 0
	return cyclical(year - 1890 + 25 + offset)
}

/**
 * 获取干支纪月
 * year,month 公历年，干支所在月
 * offset 偏移量，默认为0，便于查询一个年跨两个干支纪年(以立春为分界线)
 */
func getLunarMonthName(year, month, offset int) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	offset = offset | 0
	return cyclical((year-1890)*12 + month + 12 + offset)
}

/**
 * 获取干支纪日
 * year,month,day 公历年，月，日
 */
func getLunarDayName(year, month, day int) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	//当日与1890/1/1 相差天数
	//1890/1/1与 1970/1/1 相差29219日, 1890/1/1 日柱为壬午日(60进制18)
	date := time.Date(int(year), time.Month(month), int(day),
		0, 0, 0, 0, time.UTC).Unix()
	dayCyclical := date/86400 + 29219 + 18
	return cyclical(int(dayCyclical))
}

/**
 * 获取公历月份的天数
 */
func getSolarMonthDays(year, month int) (int, bool) {
	if !isYearValid(year) {
		return -1, false
	}

	monthDays := []int{}
	if isLeapYear(year) {
		monthDays = []int{31, 29, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	} else {
		monthDays = []int{31, 28, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}

	return monthDays[month], true
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 == 0) || year%400 == 0
}

/**
 * 统一日期输入参数(输入月份从1开始，内部月份统一从0开始)
 */
func formatDate(year, month, day int) {
}

/**
 * 将农历转换为公历
 * year,month,day 农历年，月(1-13，有闰月)，日
 */
func lunarToSolar(year, month, day int) (yearInfo, bool) {
	if !isYearValid(year) {
		return yearInfo{-1, -1, -1}, false
	}

	between, _ := getDaysBetweenZheng(year, month, day)
	//info := lunarInfos[year-MinYear]
	//zengMonth := info.zhengMonth
	//zengDay := info.zhengDay

	offDate := time.Date(int(year), time.Month(month), int(day),
		0, 0, 0, 0, time.UTC).Unix() + int64(between)*86400
	newDate := time.Unix(offDate, 0)
	y, m, d := newDate.Date()

	return yearInfo{int(y), int(m), int(d)}, true
}

/**
 * 将公历转换为农历
 */
func solarToLunar(year, month, day int) (map[string]interface{}, bool) {
	if !isYearValid(year) {
		return nil, false
	}

	cacheObj.setCurrent(year)
	// 立春日期
	v, ok := cacheObj.getCache("term2")
	if !ok {
		info, _ := getTermDate(year, 2)
		v = cacheObj.setCache("term2", info)
	}
	term2 := v.(yearInfo)

	// 二十四节气
	v, ok = cacheObj.getCache("termList")
	if !ok {
		list := getYearTerm(year)
		v = cacheObj.setCache("termList", list)
	}
	termList := v.(map[string]string)

	//某月第一个节气开始日期
	firstTerm, _ := getTermDate(year, month*2)
	//干支所在年份
	ganZhiYear := int(0)
	if month > 1 || (month == 1 && day >= term2.day) {
		ganZhiYear = year + 1
	} else {
		ganZhiYear = year
	}
	//干支所在月份（以节气为界）
	ganZhiMonth := int(0)
	if day >= firstTerm.day {
		ganZhiMonth = month + 1
	} else {
		ganZhiMonth = month
	}

	lunarDate, _ := getLunarDateBySolar(year, month, day)
	lunarLeapMonth, _ := getLunarLeapYear(lunarDate.year)
	lunarMonthName := ""
	if lunarLeapMonth > 0 && lunarLeapMonth == lunarDate.month {
		lunarMonthName = "闰" + lunarData["monthCn"][lunarDate.month-1] + "月"
	} else if lunarLeapMonth > 0 && lunarLeapMonth > lunarDate.month {
		lunarMonthName = lunarData["monthCn"][lunarDate.month-1] + "月"
	} else {
		lunarMonthName = lunarData["monthCn"][lunarDate.month] + "月"
	}

	//农历节日判断
	lunarFtv := ""
	lunarMonthInfos, _, _ := getLunarYearDays(lunarDate.year)
	lunarMonthLen := int(len(lunarMonthInfos))
	//除夕
	if lunarDate.month == (lunarMonthLen-1) && lunarDate.day == lunarMonthInfos[lunarMonthLen-1].days {
		lunarFtv = lunarFestival["d0100"]
	} else if lunarLeapMonth > 0 && lunarDate.month > lunarLeapMonth {
		lunarFtv = lunarFestival[formatDayD4(lunarDate.month-1, lunarDate.day)]
	} else {
		lunarFtv = lunarFestival[formatDayD4(lunarDate.month, lunarDate.day)]
	}

	// 返回结果
	resMap := make(map[string]interface{})
	zodiac, _ := getYearZodiac(ganZhiYear)
	resMap["zodiac"] = zodiac
	yearName, _ := getLunarYearName(ganZhiYear, 0)
	resMap["GanZhiYear"] = yearName
	monthName, _ := getLunarMonthName(year, ganZhiMonth, 0)
	resMap["GanZhiMonth"] = monthName
	dayName, _ := getLunarDayName(year, month, day)
	resMap["GanZhiDay"] = dayName
	resMap["term"] = termList[formatDayD4(month, day)]
	resMap["lunarYear"] = lunarDate.year
	resMap["lunarMonth"] = lunarDate.month
	resMap["lunarDay"] = lunarDate.day
	resMap["lunarMonthName"] = lunarMonthName
	resMap["lunarDayName"] = lunarData["dateCn"][lunarDate.day-1]
	resMap["lunarLeapMonth"] = lunarLeapMonth
	resMap["solarFestival"] = solarFestival[formatDayD4(month, day)]
	resMap["lunarFestival"] = lunarFtv

	return resMap, true
}
