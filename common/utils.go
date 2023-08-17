package common

import (
	"bytes"
	"fmt"
	"github.com/dimiro1/banner"
	"os"
)

func createBanner(text string) {
	isEnabled := true
	isColorEnabled := true
	banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
	//return banner
}

func PrintBanner() {
	//bannerText := "hscan" + version
	//createBanner(bannerText)
	//fmt.Println(banner)
	banner := `                                        
 #    #   ####    ####     ##    #    # 
 #    #  #       #    #   #  #   ##   # 
 ######   ####   #       #    #  # #  # 
 #    #       #  #       ######  #  # # 
 #    #  #    #  #    #  #    #  #   ## 
 #    #   ####    ####   #    #  #    #   
Version: ` + version
	fmt.Println(banner)
}
