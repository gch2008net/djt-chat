package main

import (
	"fmt"
	"strings"

	"github.com/openimsdk/gomake/mageutil"
)

func main() {
	fmt.Println("mian.................11111111111111")
	Build()

}

var Default = Build

func Build() {

	//制定系统架构进行build
	platforms := "linux_amd64"

	for _, platform := range strings.Split(platforms, " ") {
		mageutil.CompileForPlatform(platform)
	}

	mageutil.PrintGreen("All binaries under cmd and tools were successfully compiled.")
}

func Start() {
	mageutil.InitForSSC()
	// err := setMaxOpenFiles()
	// if err != nil {
	// 	mageutil.PrintRed("setMaxOpenFiles failed " + err.Error())
	// 	os.Exit(1)
	// }
	mageutil.StartToolsAndServices()
}

func Stop() {
	mageutil.StopAndCheckBinaries()
}

func Check() {
	mageutil.CheckAndReportBinariesStatus()
}
