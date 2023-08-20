package web

import (
	"hscan/common"
	"hscan/utils/logger"
	"hscan/web/internal/common/check"
	load "hscan/web/internal/common/load"
	"hscan/web/internal/common/output"
	nuclei_parse "hscan/web/pkg/nuclei/parse"
	xray_requests "hscan/web/pkg/xray/requests"
	xray_structs "hscan/web/pkg/xray/structs"
	"strconv"
	"time"
)

func filterWebTargets() (targets []string) {
	for _, discoverResult := range common.DiscoverResults {
		if discoverResult["protocol"].(string) == "http" || discoverResult["protocol"].(string) == "https" {
			targets = append(targets, discoverResult["uri"].(string))
		}
	}

	return
}

func NucleiPoc() {
	//if !common.RunningInfo.Nuclei {
	//	return
	//}
	//if common.RunningInfo.NoPoc {
	//	return
	//}

	logger.Banner("Start Nuclei Poc Scan! ")
	targets := filterWebTargets()
	// 初始化http客户端
	xray_requests.InitHttpClient(common.RunningInfo.PocThread, common.RunningInfo.PocSocks5Proxy, time.Duration(common.RunningInfo.PocTimeout)*time.Second)

	// 初始化nuclei options
	nuclei_parse.InitExecuterOptions(common.RunningInfo.PocRate, common.RunningInfo.PocTimeout)

	common.RunningInfo.NucleiPocPath = "web/pocs/nuclei/**"

	_, nucleiPocs := load.LoadPocs(common.RunningInfo.NucleiPocPath)
	logger.Info(logger.LightGreen("Load ") +
		logger.White(strconv.Itoa(len(nucleiPocs))) +
		logger.LightGreen(" nuclei poc(s)"))
	// 初始化输出
	outputChannel, outputWg := output.InitOutput(common.RunningInfo.OutputFileName, false)

	// 初始化check
	check.InitCheck(common.RunningInfo.PocThread, common.RunningInfo.PocRate, false)
	check.Start(targets, nil, nucleiPocs, outputChannel)
	check.Wait()

	// check结束
	close(outputChannel)
	check.End()
	outputWg.Wait()

}

type XrayClients struct {
	Targets    []map[string]interface{}
	Proxy      string
	OutputFile string
	Json       bool
	Thread     int
	Rate       int
	PocTimeout int
}

func NewXrayClients(xray XrayClients) *XrayClients {
	return &XrayClients{
		Targets:    xray.Targets,
		Rate:       xray.Rate,
		Thread:     xray.Thread,
		PocTimeout: xray.PocTimeout,
		OutputFile: xray.OutputFile,
	}
}

type XrayClient struct {
	url        string
	proxy      string
	outputFile string
	json       bool
	thread     int
	pocPath    string
	rate       int
	pocTimeout int
	xrayPocs   map[string]xray_structs.Poc
}

func NewXrayClient(url string) *XrayClient {

	var pocPath = "pocs/xray2"
	rate := 20
	thread := 10
	pocTimeout := 20
	outputFile := "xray_output.txt"
	return &XrayClient{
		url:        url,
		pocPath:    pocPath,
		rate:       rate,
		thread:     thread,
		pocTimeout: pocTimeout,
		outputFile: outputFile,
	}
}

func XrayScan(xrayClient *XrayClient, allXrayPocs map[string]xray_structs.Poc) {

	// 初始化http客户端
	xray_requests.InitHttpClient(xrayClient.thread, xrayClient.proxy, time.Duration(xrayClient.pocTimeout)*time.Second)
	logger.Info("Start vulnerable Scan: " + xrayClient.url)
	// 计算xray的总发包量，初始化缓存
	xrayTotalRequests := 0
	for _, poc := range allXrayPocs {
		ruleLens := len(poc.Rules)
		// 额外需要缓存connectionID
		if poc.Transport == "tcp" || poc.Transport == "udp" {
			ruleLens += 1
		}
		xrayTotalRequests += ruleLens
	}
	//log.Println("xrayTotalRequests", xrayTotalRequests)
	if xrayTotalRequests == 0 {
		xrayTotalRequests = 1
	}
	//log.Println(xrayTotalRequests)
	xray_requests.InitCache(xrayTotalRequests)

	// 初始化输出
	outputChannel, outputWg := output.InitOutput(xrayClient.outputFile, false)

	// 初始化check
	check.InitCheck(common.RunningInfo.PocThread, common.RunningInfo.PocRate, false)

	// check开始
	check.StartXray(xrayClient.url, allXrayPocs, outputChannel)
	check.Wait()

	// check结束
	close(outputChannel)
	check.End()
	outputWg.Wait()

}

func StartXray(xrayclients *XrayClients) {
	logger.Banner("Start Xray  Scan! ")
	// 加载pocs
	allXrayPocs := initXaryPoc("pocs/xray2")
	logger.Info(logger.LightGreen("Load ") +
		logger.White(strconv.Itoa(len(allXrayPocs))) +
		logger.LightGreen(" xray poc(s)"))
	for _, discoverResult := range xrayclients.Targets {
		if discoverResult["protocol"].(string) == "http" || discoverResult["protocol"].(string) == "https" {
			url := discoverResult["uri"].(string)
			XrayScan(NewXrayClient(url), allXrayPocs)
		}
	}
}

func StartNuclei() {
	NucleiPoc()
}
