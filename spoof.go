package main

import (
	"fmt"
	"time"

	"github.com/Oakes6/guids"
	"github.com/fatih/color"
	"golang.org/x/sys/windows/registry"
)

func main() {

	var res string
	fmt.Println("Запустить цикл? [y/n]")
	fmt.Scanln(&res)
	if res == "y" {
		for {
			spoof()
			time.Sleep(20 * time.Millisecond)
		}
	}
	
}
func spoof() {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	guid := guids.GUIDV4()
	key.SetStringValue("machineGuid", guid)

	
	key, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\IDConfigDB\Hardware Profiles\0001`, registry.WRITE)
	if err != nil {
		color.Red("[-] Ошибка")
	}
	defer key.Close()
	guid = guids.GUIDV4()
	key.SetStringValue("HwProfileGuid", "{" + guid + "}")
	color.Green(time.Now().Format("[ 15:04:05 ] ") + "Успешно установлен " + guid)

}