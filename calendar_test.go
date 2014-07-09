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

import (
	"fmt"
	"testing"
)

func TestGetLunarDateBySolar(t *testing.T) {
	info, isLeapMonth, ok := GetLunarDateBySolar(2014, 8, 8)
	if !ok {
		t.Error("GetLunarDateBySolar failed")
		return
	}

	fmt.Printf("isLeapMonth: %v, info: %v\n", isLeapMonth, info)
}

func TestGetSolarDateByLunar(t *testing.T) {
	info, ok := GetSolarDateByLunar(2014, 7, 13, false)
	if !ok {
		t.Error("GetLunarDateBySolar failed")
		return
	}

	fmt.Printf("info: %v\n", info)
}

func TestGetLunarInfoBySolar(t *testing.T) {
	info, ok := GetLunarInfoBySolar(2014, 8, 8)
	if !ok {
		t.Error("GetLunarDateBySolar failed")
		return
	}

	fmt.Printf("info: %v\n", info)
}

func TestGetSolarMonthCalendar(t *testing.T) {
	info, ok := GetSolarMonthCalendar(2014, 7, true)
	if !ok {
		t.Error("GetLunarDateBySolar failed")
		return
	}

	fmt.Printf("info: %v\n", info)
}

func TestGetLunarMonthCalendar(t *testing.T) {
	info, ok := GetLunarMonthCalendar(2014, 7, true)
	if !ok {
		t.Error("GetLunarDateBySolar failed")
		return
	}

	fmt.Printf("info: %v\n", info)
}
