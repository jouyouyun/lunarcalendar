LunarCalendar
============

农历（阴历）万年历，支持农历与公历之间相互转换，含有二十四节气，天干地支纪年纪月纪日，生肖属相及农历传统节假日信息等功能。带有黄历数据，支持1891-2100年。

此库的实现是参考 [zzyss86](https://github.com/zzyss86/LunarCalendar) 来实现的。

Install
=========

`go get github.com/jouyouyun/lunarcalendar`

`API` 介绍
=========

使用到的结构体：

```go
type CaYearInfo struct {
	Year  int32
	Month int32
	Day   int32
}

type caLunarDayInfo struct {
	LunarLeapMonth int32
	LunarMonthName string
	LunarDayName   string
	GanZhiYear     string
	GanZhiMonth    string
	GanZhiDay      string
	Zodiac         string
	Term           string
	SolarFestival  string
	LunarFestival  string
	Worktime       int32
}

type caSolarMonthInfo struct {
	FirstDayWeek int32
	Days         int32
	Datas        []CaYearInfo
}

type caLunarMonthInfo struct {
	FirstDayWeek int32
	Days         int32
	Datas        []caLunarDayInfo
}
```

`func GetLunarDateBySolar(year, month, day int32) (info CaYearInfo, isLeapMonth,ok bool)`

通过公历日期获得农历日期

`func GetSolarDateByLunar(year, month, day int32, isLeapMonth bool) (info CaYearInfo, ok bool)`

通过农历日期获得公历日期

`func GetLunarInfoBySolar(year, month, day int32) (info caLunarDayInfo, ok bool))`

通过公历日期获得农历日期的详细信息，包括黄历、节气、假日。

`func GetSolarMonthCalendar(year, month int32, fill bool) (info caSolarMonthInfo, ok bool)`

获取以整个月的公历信息，`fill` 表示是否 `7x6` 的结构。

`func GetLunarMonthCalendar(year, month int32, fill bool) (info caLunarMonthInfook bool)`

获取以整个月的农历信息，`fill` 表示是否 `7x6` 的结构。输入的是公历的年月日期。
