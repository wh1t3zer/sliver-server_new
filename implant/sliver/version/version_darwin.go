//go:build darwin

package version

import (
	"syscall"
)

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

func GetVersion() string {
	if version, err := syscall.Sysctl("kern.osproductversion"); err == nil {
		return "macOS " + version
	}
	if version, err := getKernelVersion(); err == nil {
		return "Darwin " + version
	}
	return ""
}

func getKernelVersion() (string, error) {
	return syscall.Sysctl("kern.version")
}
