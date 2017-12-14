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

import (
	C "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	C.TestingT(t)
}

func init() {
	C.Suite(GetManager())
}

func (m *Manager) TestUtils(c *C.C) {
	c.Check(hasMotionFlag(1), C.Equals, true)
	c.Check(hasMotionFlag(0), C.Equals, false)

	c.Check(hasButtonFlag(2), C.Equals, true)
	c.Check(hasButtonFlag(0), C.Equals, false)

	c.Check(hasKeyFlag(4), C.Equals, true)
	c.Check(hasKeyFlag(0), C.Equals, false)

	c.Check(keyCode2Str(25), C.Equals, "w")
}
