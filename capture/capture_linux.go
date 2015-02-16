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
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"image"
	"image/png"
	"os"
)

func ScreenRect() (image.Rectangle, error) {
	c, err := xgb.NewConn()
	if err != nil {
		return image.Rectangle{}, err
	}
	defer c.Close()

	screen := xproto.Setup(c).DefaultScreen(c)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func CaptureScreen(filename string) error {
	r, e := ScreenRect()
	if e != nil {
		return e
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	img, err := CaptureRect(r)
	if err != nil {
		return err
	}

	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}

func CaptureRect(rect image.Rectangle) (*image.RGBA, error) {
	c, err := xgb.NewConn()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	screen := xproto.Setup(c).DefaultScreen(c)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(c, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}
