Simple cross-platform pure Go screenshot library. This is used by [Gauge](https://github.com/getgauge/gauge) plugins as a mean to capture screenshot, but can be used standalone nevertheless.

## Usage
`gauge_screenshot <file_name>`

## Install:
`go get github.com/getgauge/gauge_screenshot`

### Note to Windows users
In order to compile this code, `gcc` must be installed. Please ensure that you install the MinGW gcc and *not* Cygwin. Also ensure that if you have 64-bit Go installed, then MinGW/gcc is also 64-bit.

## Dependencies:
- Windows: https://github.com/AllenDang/w32
- Linux/Freebsd: https://github.com/BurntSushi/xgb
- OSx: `screencapture` :)

## Credits
This code has been forked from [vova616/screenshot](https://github.com/vova616/screenshot). The purposes have diverged slightly, the original `screenshot` package serves as a library, `gauge_screenshot` serves as a standalone utility.

A list of all dependencies can be found [here](https://github.com/getgauge/gauge_screenshot/blob/master/NOTICE.md)

## License

![GNU Public License version 3.0](http://www.gnu.org/graphics/gplv3-127x51.png)
Gauge is released under [GNU Public License version 3.0](http://www.gnu.org/licenses/gpl-3.0.txt)

## Copyright

Copyright 2015 ThoughtWorks, Inc.
