/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

// #cgo CFLAGS: -g -Wall
// #cgo pkg-config: x11 xi
// #include "xinput.h"
import "C"

func init() {
	go C.start_listen()
}

//export go_handle_raw_event
func go_handle_raw_event(evt_type int, detail int32, x, y, mask int32) {
	switch evt_type {
	case C.XI_RawKeyPress:
		GetManager().handleKeyboardEvent(detail, true, x, y)
	case C.XI_RawKeyRelease:
		GetManager().handleKeyboardEvent(detail, false, x, y)
	case C.XI_RawTouchBegin:
		GetManager().handleButtonEvent(1, true, x, y)
	case C.XI_RawButtonPress:
		GetManager().handleButtonEvent(detail, true, x, y)
	case C.XI_RawTouchEnd:
		GetManager().handleButtonEvent(1, false, x, y)
	case C.XI_RawButtonRelease:
		GetManager().handleButtonEvent(detail, false, x, y)

	case C.XI_RawTouchUpdate:
		GetManager().handleCursorEvent(x, y, false)
	case C.XI_RawMotion:
		/**
		* mouse left press: mask = 256
		* mouse right press: mask = 512
		* mouse middle press: mask = 1024
		**/
		if mask >= 256 {
			GetManager().handleCursorEvent(x, y, true)
		} else {
			GetManager().handleCursorEvent(x, y, false)
		}
	}
}

func getButtonState(event *C.XIDeviceEvent) []int {
	var buttons []int
	for i := 0; i < int(event.buttons.mask_len)*8; i++ {
		if C.xi_mask_is_set(event.buttons.mask, C.char(i)) != 0 {
			buttons = append(buttons, i)
		}
	}
	return buttons
}
