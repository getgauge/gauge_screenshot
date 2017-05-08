// Copyright 2015 ThoughtWorks, Inc.
// Copyright (C) 2012 vova616 <vova616@gmail.com>

// This file is part of Gauge-Screenshot.

// Gauge-Screenshot is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge-Screenshot is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge-Screenshot.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/getgauge/gauge_screenshot/capture"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <target filepath>\n", os.Args[0])
		os.Exit(0)
	}
	err := capture.CaptureScreen(os.Args[1])
	if err != nil {
		panic(err)
	}
}
