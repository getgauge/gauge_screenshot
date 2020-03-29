Simple cross-platform pure Go screenshot library. This is used by [Gauge](https://github.com/getgauge/gauge) plugins as a mean to capture screenshot, but can be used standalone nevertheless.

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Usage
`gauge_screenshot <file_name>`

## Install:
`go get github.com/getgauge/gauge_screenshot`

### Offline installation

Download the plugin from [Releases](https://github.com/getgauge/gauge_screenshot/releases)

### Note to Windows users
In order to compile this code, `gcc` must be installed. Please ensure that you install the MinGW gcc and *not* Cygwin. Also ensure that if you have 64-bit Go installed, then MinGW/gcc is also 64-bit.

One way to install the required way of `gcc` on Windows is to use the [mingw-w64-install.exe](https://sourceforge.net/projects/mingw-w64/files/Toolchains%20targetting%20Win32/Personal%20Builds/mingw-builds/installer/mingw-w64-install.exe/download) which can fetch and install the desired version of `gcc`.

## Dependencies:
- Windows: https://github.com/TheTitanrain/w32
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
