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
)

func test() {
	n, _ := getLunarLeapYear(2014)
	fmt.Printf("2014 闰月： %d\n", n)
	//monthDays, yearDays, _ := getLunarYearDays(2014)
	//fmt.Printf("2014 days: %d\n", yearDays)
	//for _, info := range monthDays {
	//fmt.Printf("\t%d 月: %d\n", info.index, info.days)
	//}

	//fmt.Println("\n通过间隔天数查找农历日期")
	//info, _ := getLunarDateByBetween(2014, 100)
	//fmt.Println("year: 2014, between: 100")
	//fmt.Printf("\t%d - %d - %d\n", info.year, info.month, info.day)

	//fmt.Println("\n通过公历日期获取农历日期")
	//fmt.Println("Date: 2014 - 4 - 29")
	//info, _ = getLunarDateBySolar(2014, 4, 29)
	//fmt.Printf("\t%d - %d - %d\n", info.year, info.month, info.day)

	//fmt.Println("\n计算两个公历日期之间的天数")
	//num, _ := getDaysBetweenSolar(2014, 2, 1, 2014, 10, 1)
	//fmt.Println("2014 - 2 - 3, 2014 - 10 -3")
	//fmt.Printf("\tnum: %d\n", num)

	//fmt.Println("\n计算农历日期离正月初一有多少天")
	//days, _ := getDaysBetweenZheng(2014, 4, 13)
	//fmt.Println("\tdays: ", days)

	//fmt.Println("\n第n个节气")
	//for i := 0; i < 24; i++ {
	//info, _ = getTermDate(2014, i)
	//fmt.Printf("\t%s: %v - %v - %v\n",
	//lunarData["solarTerm"][i], info.year, info.month, info.day)
	//}

	//fmt.Println("\n获取公历年一年的二十四节气")
	//resMap := getYearTerm(2014)
	//for k, v := range resMap {
	//fmt.Printf("\t%s: %s\n", k, v)
	//}

	//fmt.Println("\n获取生肖")
	//ret, _ := getYearZodiac(2014)
	//fmt.Printf("\t2014 Zodiac: %s\n", ret)

	//fmt.Println("\n获取干支纪年")
	//ret, _ := getLunarYearName(2014, 0)
	//fmt.Printf("\tGanZhiYear: %s\n", ret)
	//fmt.Println("获取干支纪月")
	//ret, _ = getLunarMonthName(2014, 5, 0)
	//fmt.Printf("\tGanZhiMonth: %s\n", ret)
	//fmt.Println("获取干支纪日")
	//ret, _ = getLunarDayName(2014, 5, 4)
	//fmt.Printf("\tGanZhiDay: %s\n", ret)

	//fmt.Println("\n获取公历月份的天数")
	//ret, _ := getSolarMonthDays(2014, 6)
	//fmt.Printf("\t2014-6 days: %d\n", ret)

	//fmt.Println("\n将农历转换为公历")
	//ret, _ := lunarToSolar(2014, 3, 22)
	//fmt.Printf("\t2014 - 3 - 22 date: %v - %v - %v\n",
	//ret.year, ret.month, ret.day)

	fmt.Println("\n将公历转换为农历")
	ret, _ := solarToLunar(2014, 4, 5)
	fmt.Printf("\tlunarYear: %v\n", ret.lunarYear)
	fmt.Printf("\tlunarMonth: %v\n", ret.lunarMonth)
	fmt.Printf("\tlunarDay: %v\n", ret.lunarDay)
	fmt.Printf("\tlunarLeapMonth: %v\n", ret.lunarLeapMonth)
	fmt.Printf("\tlunarMonthName: %v\n", ret.lunarMonthName)
	fmt.Printf("\tlunarDayName: %v\n", ret.lunarDayName)
	fmt.Printf("\tganZhiYear: %v\n", ret.ganZhiYear)
	fmt.Printf("\tganZhiMonth: %v\n", ret.ganZhiMonth)
	fmt.Printf("\tganZhiDay: %v\n", ret.ganZhiDay)
	fmt.Printf("\tzodiac: %v\n", ret.zodiac)
	fmt.Printf("\tterm: %v\n", ret.term)
	fmt.Printf("\tsolarFestival: %v\n", ret.solarFestival)
	fmt.Printf("\tlunarFestival: %v\n", ret.lunarFestival)
}

func main() {
	test()
}
