package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"hscan/common"
	"hscan/discover"
	"hscan/nonweb"
	"hscan/parse"
	"hscan/probe"
	"hscan/utils/logger"
	"hscan/web"
	"io/ioutil"
	"time"
)

func loadConfig(filename string) common.InputInfoStruct {
	data, err := ioutil.ReadFile(filename)
	if err != nil {

		fmt.Println("config file: ", filename, "[Error]")
		fmt.Println(err)
		return common.InputInfoStruct{}
	}
	//log.Println(string(data))
	var config common.InputInfoStruct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		//log.Println(err)
		fmt.Println("config file: ", filename, "[Error]")
		fmt.Println(err)
		return common.InputInfoStruct{}
	}
	fmt.Println("config file: ", filename, "[Yes]")
	return config
}

func main() {
	filename := "./config.yaml"
	startTime := time.Now()
	common.PrintBanner()
	common.InputInfo = loadConfig(filename)

	parse.Flag(&common.InputInfo)

	parse.Parse(&common.InputInfo, &common.RunningInfo)
	probe.Probe()
	// 主机发现
	discover.Discover()
	// xray扫描
	xray()
	// nuclei扫描
	if common.RunningInfo.Nuclei {
		web.StartNuclei()
	}

	// 服务扫描
	nonweb.Service()
	logger.Info(fmt.Sprintf("Task finish, consumption of time: %s", time.Now().Sub(startTime)))
}

func xray() {
	x := web.XrayClients{}
	x.Targets = common.DiscoverResults
	x.Proxy = common.RunningInfo.ProxyProxy
	x.OutputFile = common.RunningInfo.OutputFileName
	x.Rate = common.RunningInfo.PocRate
	x.Thread = common.RunningInfo.Thread
	xray := web.NewXrayClients(x)
	web.StartXray(xray)
}
