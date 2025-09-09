//go:build windows

package version

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"golang.org/x/sys/windows"
)

const VER_NT_WORKSTATION = 0x0000001

func getOSVersion() string {
	osVersion := windows.RtlGetVersion()

	osName := getWindowsName(osVersion)

	var servicePack string
	if osVersion.ServicePackMajor != 0 {
		servicePack = fmt.Sprintf(" Service Pack %d", osVersion.ServicePackMajor)
	}

	return fmt.Sprintf("%s%s build %d", osName, servicePack, osVersion.BuildNumber)
}

func getWindowsName(osVersion *windows.OsVersionInfoEx) string {
	if osVersion.MajorVersion >= 10 {
		if osVersion.ProductType == VER_NT_WORKSTATION {
			if osVersion.BuildNumber >= 22000 {
				return "11"
			} else {
				return "10"
			}
		} else {
			// Windows Server
			switch osVersion.BuildNumber {
			case 14393:
				return "Server 2016"
			case 17763:
				return "Server 2019"
			case 20348:
				return "Server 2022"
			default:
				if osVersion.BuildNumber >= 20348 {
					return "Server 2022+"
				} else {
					return "Server 2016+"
				}
			}
		}
	}

	if osVersion.MajorVersion == 6 {
		switch osVersion.MinorVersion {
		case 0:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				return "Vista"
			} else {
				return "Server 2008"
			}
		case 1:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				return "7"
			} else {
				return "Server 2008 R2"
			}
		case 2:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				return "8"
			} else {
				return "Server 2012"
			}
		case 3:
			if osVersion.ProductType == VER_NT_WORKSTATION {
				return "8.1"
			} else {
				return "Server 2012 R2"
			}
		}
	}

	return "Unknown Windows"
}

// GetVersion returns the os version information
func GetVersion() string {
	return getOSVersion()
}
