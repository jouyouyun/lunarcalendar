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
	"strconv"
	"strings"
	"time"
)

type yearInfo struct {
	year  int32
	month int32
	day   int32
}

func isYearValid(year int32) bool {
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
func getLunarLeapYear(year int32) (int32, bool) {
	if !isYearValid(year) {
		return -1, false
	}

	return lunarInfos[year-MinYear].leapMonth, true
}

/**
 * 获取农历年每月的天数及一年的总天数
 */
func getLunarYearDays(year int32) (map[int32]int32, int32, bool) {
	if !isYearValid(year) {
		return nil, -1, false
	}

	info := lunarInfos[year-MinYear]
	leapMonth := info.leapMonth
	monthData := fmt.Sprintf("%b", info.lunarMonthNum)
	tmp := ""
	l := len(monthData)
	//还原数据至16位,少于16位的在前面插入0（二进制存储时前面的0被忽略)
	for i := o; i < 16-l; i++ {
		tmp += "0"
	}
	monthData = tmp + monthData

	monthNum := int32(0)
	if leapMonth {
		monthNum = 13
	} else {
		monthNum = 12
	}

	yearDays := int32(0)
	monthDayMap := make(map[int32]int32)
	for i := 0; i < monthNum; i++ {
		if monthData[i] == 0 {
			yearDays += 29
			monthDayMap[i] = 29
		} else {
			yearDays += 30
			monthDayMap[i] = 30
		}
	}

	return monthDayMap, yearDays, true
}

/**
 * 通过间隔天数查找农历日期
 */
func getLunarDateByBetween(year, between int32) (yearInfo, bool) {
	month := int32(-1)
	day := int32(-1)
	monthDayMap, yearDays, ok := getLunarYearDays(year)
	if !ok {
		fmt.Println("Get Year Days Failed For Year: ", year)
		return yearInfo{year, month, day}, false
	}

	end := int32(0)
	if between > 0 {
		end = between
	} else {
		end = yearDays + between
	}
	tmpDays := int32(0)
	for index, num := range monthDayMap {
		tmpDays += num
		if tmpDays > end {
			month = index
			tmpDays = tmpDays - num
			break
		}
	}
	day = end - tmpDays + 1

	return yearInfo{year, month, day}, true
}

/**
 * 通过公历日期获取农历日期
 */
func getLunarDateBySolar(year, month, day int32) (yearInfo, bool) {
	if !isYearValid(year) {
		return yearInfo{-1, -1, -1}, false
	}

	info := lunarInfos[year-MinYear]
	zengMonth := info.springKalendsMonth
	zengDay := info.springKalendsDay
	between, _ := getDaysBetweenSolar(year, zengMonth, zengDay,
		year, month, day)
	if between == 0 { //正月初一
		return yearInfo{year, 1, 1}, true
	} else if between < 0 {
		year -= 1
	}
	return getLunarDateByBetween(year, between)
}

/**
 * 计算两个公历日期之间的天数
 */
func getDaysBetweenSolar(year, month, day, year1, month1, day1 int32) (days, bool) {
	date := time.Date(int(year), int(month), int(day),
		0, 0, 0, 0, time.UTC).Unix()
	date1 := time.Date(int(year1), int(month1), int(day1),
		0, 0, 0, 0, time.UTC).Unix()

	return (date1 - date) / 86400, true
}

/**
 * 计算农历日期离正月初一有多少天
 */
func getDaysBetweenZheng(year, month, day int32) (int32, bool) {
	monthDayMap, yearDays, ok := getLunarYearDays(year)
	if !ok {
		fmt.Println("Get Year Days Failed For Year: ", year)
		return -1, false
	}

	days := int32(0)
	for i, d := range monthDayMap {
		if i < month {
			days += d
		} else {
			break
		}
	}

	return days + day - 1, true
}

func formatDayD4(month, day int32) string {
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
func getTermDate(year, n int32) (yearInfo, bool) {
	if !isYearValid(year) {
		return yearInfo{}, false
	}

	offset := 31556925974/1000*(year-1890) + terms[n]*60 + time.Date(1890, 0, 5, 0, 0, 0, 0, time.UTC).Unix()
	y, m, d := time.Unix(offset, 0).Date()

	return yearInfo{int32(y), int32(m), int32(d)}, true
}

/**
 * 获取公历年一年的二十四节气
 * 返回key:日期，value:节气中文名
 */
func getYearTerm(year) {
}

/**
 * 获取生肖
 * year: 干支所在年(默认以立春前的公历年作为基数)
 */
func getYearZodiac(year int32) (string, bool) {
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
func cyclical(num int32) (string, bool) {
	return lunarData["heavenlyStems"][num%10] + lunarData["earthlyBranches"][num%12], true
}

/**
 * 获取干支纪年
 * year 干支所在年
 * offset 偏移量，默认为0，便于查询一个年跨两个干支纪年(以立春为分界线)
 */
func getLunarYearName(year, offset int32) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	offset = offset || 0
	return cyclical(year - 1890 + 25 + offset)
}

/**
 * 获取干支纪月
 * year,month 公历年，干支所在月
 * offset 偏移量，默认为0，便于查询一个年跨两个干支纪年(以立春为分界线)
 */
func getLunarMonthName(year, month, offset int32) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	offset = offset || 0
	return cyclical((year-1890)*12 + month + 12 + offset)
}

/**
 * 获取干支纪日
 * year,month,day 公历年，月，日
 */
func getLunarDayName(year, month, day int32) (string, bool) {
	if !isYearValid(year) {
		return "", false
	}

	//当日与1890/1/1 相差天数
	//1890/1/1与 1970/1/1 相差29219日, 1890/1/1 日柱为壬午日(60进制18)
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
	dayCyclical := date/86400 + 29219 + 18
	return cyclical(dayCyclical)
}

/**
 * 获取公历月份的天数
 */
func getSolarMonthDays(year, month int32) (int32, bool) {
	if !isYearValid(year) {
		return -1, false
	}

	monthDays := []int32{}
	if isLeapYear(year) {
		monthDays = []int32{31, 29, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	} else {
		monthDays = []int32{31, 28, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}

	return monthDays[month], true
}

func isLeapYear(year int32) bool {
	return (year%4 == 0 && year%100 == 0) || year%400 == 0
}

/**
 * 统一日期输入参数(输入月份从1开始，内部月份统一从0开始)
 */
func formatDate(year, month, day int32) {
}

/**
 * 将农历转换为公历
 * year,month,day 农历年，月(1-13，有闰月)，日
 */
func lunarToSolar(year, month, day int32) (int32, int32, int32, bool) {
	if !isYearValid(year) {
		return -1, -1, -1, false
	}

	between := getDaysBetweenZheng(year, month, day)
	info := lunarInfos[year-MinYear]
	zengMonth := info.springKalendsMonth
	zengDay := info.springKalendsDay

	offDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix() + between*86400
	newDate := time.Unix(offDate, 0)
	y, m, d := newDate.Date()

	return int32(y), int32(m), int32(d), true
}

/**
 * 将公历转换为农历
 */
func solarToLunar(year, month, day int32) (int32, int32, int32, bool) {
	if !isYearValid(year) {
		return -1, -1, -1, false
	}

	cacheObj.setCurrent(year)
}
