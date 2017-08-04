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

package capture

import (
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestCaptureScreen(t *testing.T) {
	outFile := filepath.Join("_testdata", "output.png")
	err := CaptureScreen(outFile)
	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(outFile)

	if err != nil {
		t.Error(err)
	}

	c, err := png.DecodeConfig(f)

	if err != nil {
		t.Error(err)
	}

	if c.Height == 0 || c.Width == 0 {
		t.Errorf("Invalid image dimensions, %dx%d", c.Height, c.Width)
	}

	os.Remove(outFile)
}
