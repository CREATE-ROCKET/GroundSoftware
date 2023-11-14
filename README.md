# GroundSoftware
### Serial Communication Native Application
[![Supports Linux](https://img.shields.io/badge/Supports-Linux-blue)](https://github.com/CREATE-ROCKET/GroundSoftware)
[![Supports Windows](https://img.shields.io/badge/Supports-Windows-blue)](https://github.com/CREATE-ROCKET/GroundSoftware)
[![License](https://img.shields.io/github/license/CREATE-ROCKET/GroundSoftware)](https://github.com/CREATE-ROCKET/GroundSoftware/blob/main/LICENSE)

<p align="center">
  <img width=80% src="https://github.com/CREATE-ROCKET/GroundSoftware/assets/105796502/3d06c227-d81a-41fa-992b-77fa518e4c38" alt="Logo">
</p>

![Actions Status](https://github.com/CREATE-ROCKET/GroundSoftware/actions/workflows/wails_build.yml/badge.svg?branch=main)
[![Release Version](https://img.shields.io/github/v/release/CREATE-ROCKET/GroundSoftware?include_prereleases)](https://github.com/CREATE-ROCKET/GroundSoftware/releases)
[![Release Date](https://img.shields.io/github/release-date/CREATE-ROCKET/GroundSoftware)](https://github.com/CREATE-ROCKET/GroundSoftware/releases)
[![Build Status](https://travis-ci.com/CREATE-ROCKET/GroundSoftware.svg?branch=main)](https://travis-ci.com/CREATE-ROCKET/GroundSoftware)

This native application enables serial communication. It is specifically designed for use with the NEC920 module communication. It operates seamlessly on both Windows and Linux platforms. Communication is reflected in real-time on the screen.

## Used by CREATE-ROCKET

This application is actively used by the [CREATE-ROCKET](https://github.com/CREATE-ROCKET) organization for their serial communication needs.

## Installation

To install, download the application from [Release](https://github.com/CREATE-ROCKET/GroundSoftware/releases/latest)

- `GroundSoftware-amd64-installer.exe` is for Windows.
- `wailsTest` is for Linux.

If you create a 'log' directory in the same directory as the application, log files will be generated within it. Each log file is unique to each application launch.

## How to Use

1. Download the application from [Release](https://github.com/CREATE-ROCKET/GroundSoftware/releases/latest)
2. Install the application on your system.
3. Open the application.
4. Create a directory named 'log' in the same directory as the application.
5. Click on 'Port Select' and select the desired port.
6. Click on 'Serial Start' to initiate the serial communication. If any data is received, it will be displayed.
7. If you are using the NEC920 module,
   1. Input Dst and Src.
   2. Set DstId and SrcId.
8. Enter the message.
9. Click 'Send' to transmit the message.
10. Click 'Send Module Command' to send a message as a command of the NEC920 module.

## Technology Stack

This application is built using [Wails](https://github.com/wailsapp/wails), incorporating Go and Vue.js.

![wails](https://img.shields.io/badge/Wails-DF0000.svg?style=for-the-badge&logo=Wails&logoColor=white)
![Golang](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=Go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-4FC08D.svg?style=for-the-badge&logo=vuedotjs&logoColor=white)

If you tag your commits with versions such as `v1.0.0`, a new release will be created.

## Command
```bash
xxd -p -c 1 2023-11-12-17-59-29/raw_2023-11-12-17-59-29.txt | tr -d '\n' | sed 's/\(..\)/\1 /g' | tr -s ' '
xxd -p -c 1 2023-11-12-17-59-29/raw_2023-11-12-17-59-29.txt | tr -d '\n' | sed 's/\(..\)/\1 /g' | tr -s ' ' > output.txt
xxd -p -c 1 raw_*.txt | tr -d '\n' | sed 's/\(..\)/\1 /g' | tr -s ' ' > output.txt
```

## License

See the [LICENSE](./LICENSE) file for licensing details.

## Author

[Luftalian](https://github.com/Luftalian)
