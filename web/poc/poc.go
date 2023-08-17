package poc

import (
	"hscan/common"
	"hscan/utils/logger"
	"hscan/web/poc/internal/common/check"
	utils "hscan/web/poc/internal/common/load"
	"hscan/web/poc/internal/common/output"
	nuclei_parse "hscan/web/poc/pkg/nuclei/parse"
	xray_requests "hscan/web/poc/pkg/xray/requests"
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

	var file = ""
	var json = false
	var xrayProxy = ""

	// 初始化dnslog平台
	//structs.InitReversePlatform(common.RunningInfo.CeyeApi, common.RunningInfo.CeyeDomain, time.Duration(common.RunningInfo.PocTimeout)*time.Second)
	//if common.RunningInfo.ReversePlatformType != xray_structs.ReverseType_Ceye {
	//	logger.Warning("No Ceye api, use dnslog.cn")
	//}

	// 初始化http客户端
	xray_requests.InitHttpClient(common.RunningInfo.PocThread, xrayProxy, time.Duration(common.RunningInfo.PocTimeout)*time.Second)

	// 初始化nuclei options
	nuclei_parse.InitExecuterOptions(common.RunningInfo.PocRate, common.RunningInfo.PocTimeout)

	common.RunningInfo.NucleiPocPath = "web/pocs/nuclei/**"
	//common.RunningInfo.NucleiPocPath = "web/pocs/xrayv2/*"
	xrayPocs, nucleiPocs := utils.LoadPocs(common.RunningInfo.NucleiPocPath)

	logger.Info(logger.LightGreen("Load ") +
		// logger.White(strconv.Itoa(len(xrayPocMap))) +
		// logger.LightGreen(" xray poc(s), ") +
		logger.White(strconv.Itoa(len(nucleiPocs))) +
		logger.LightGreen(" nuclei poc(s)"))

	// 计算xray的总发包量，初始化缓存
	xrayTotalRequests := 0
	totalTargets := len(targets)
	for _, poc := range xrayPocs {
		ruleLens := len(poc.Rules)
		// 额外需要缓存connectionID
		if poc.Transport == "tcp" || poc.Transport == "udp" {
			ruleLens += 1
		}
		xrayTotalRequests += totalTargets * ruleLens
	}
	if xrayTotalRequests == 0 {
		xrayTotalRequests = 1
	}
	xray_requests.InitCache(xrayTotalRequests)

	// 初始化输出
	outputChannel, outputWg := output.InitOutput(file, json)

	// 初始化check
	check.InitCheck(common.RunningInfo.PocThread, common.RunningInfo.PocRate, false)

	// check开始
	check.Start(targets, xrayPocs, nucleiPocs, outputChannel)
	check.Wait()

	// check结束
	close(outputChannel)
	check.End()
	outputWg.Wait()

}

func XrayPoc() {
	//if !common.RunningInfo.Nuclei {
	//	return
	//}
	//if common.RunningInfo.NoPoc {
	//	return
	//}
	logger.Banner("Start Xray Scan! (Poc version 2)")
	targets := filterWebTargets()

	var file = ""
	var json = false
	var xrayProxy = ""

	// 初始化dnslog平台
	//structs.InitReversePlatform(common.RunningInfo.CeyeApi, common.RunningInfo.CeyeDomain, time.Duration(common.RunningInfo.PocTimeout)*time.Second)
	//if common.RunningInfo.ReversePlatformType != xray_structs.ReverseType_Ceye {
	//	logger.Warning("No Ceye api, use dnslog.cn")
	//}

	// 初始化http客户端
	xray_requests.InitHttpClient(common.RunningInfo.PocThread, xrayProxy, time.Duration(common.RunningInfo.PocTimeout)*time.Second)

	// 初始化nuclei options
	nuclei_parse.InitExecuterOptions(common.RunningInfo.PocRate, common.RunningInfo.PocTimeout)

	common.RunningInfo.NucleiPocPath = "web/pocs/xrayv2/*"
	xrayPocs, nucleiPocs := utils.LoadPocs(common.RunningInfo.NucleiPocPath)
	logger.Info(logger.LightGreen("Load ") +
		// logger.White(strconv.Itoa(len(xrayPocMap))) +
		// logger.LightGreen(" xray poc(s), ") +
		logger.White(strconv.Itoa(len(xrayPocs))) +
		logger.LightGreen(" xray poc(s)"))

	// 计算xray的总发包量，初始化缓存
	xrayTotalRequests := 0
	totalTargets := len(targets)
	for _, poc := range xrayPocs {
		ruleLens := len(poc.Rules)
		// 额外需要缓存connectionID
		if poc.Transport == "tcp" || poc.Transport == "udp" {
			ruleLens += 1
		}
		xrayTotalRequests += totalTargets * ruleLens
	}
	if xrayTotalRequests == 0 {
		xrayTotalRequests = 1
	}
	xray_requests.InitCache(xrayTotalRequests)

	// 初始化输出
	outputChannel, outputWg := output.InitOutput(file, json)

	// 初始化check
	check.InitCheck(common.RunningInfo.PocThread, common.RunningInfo.PocRate, false)

	// check开始
	check.Start(targets, xrayPocs, nucleiPocs, outputChannel)
	check.Wait()

	// check结束
	close(outputChannel)
	check.End()
	outputWg.Wait()

}

//
//func Poc2() {
//	var target = "http://127.0.0.1:8080"
//
//	var (
//		targetFiles = make([]string, 0)
//		poc         = "web/pocs/xrayv2/"
//		pocPath     = make([]string, 0)
//		apiKey      = "ceye.io api key"
//		domain      = "ceye.io subdomain"
//		tags        = "filter poc by tag"
//		file        = "Result file to write"
//		json        = false
//		success     = false
//		proxy       = ""
//		threads     = 10
//		timeout     = 20
//		rate        = 100
//		debug       = false
//		verbose     = false
//	)
//	timeoutSecond := time.Duration(timeout) * time.Second
//	// 初始化http客户端
//	xray_requests.InitHttpClient(threads, proxy, timeoutSecond)
//	xrayPocs = utils.LoadPocs("web/pocs/xrayv2/")
//
//	// 加载目标
//	//targets := LoadTargets(target, targetFiles)
//}

//// 读取pocs
//func LoadPocs(pocs *[]string, pocPaths *[]string) (map[string]xray_structs.Poc, map[string]nuclei_structs.Poc) {
//	xrayPocMap := make(map[string]xray_structs.Poc)
//	nucleiPocMap := make(map[string]nuclei_structs.Poc)
//
//	// 加载poc函数
//	LoadPoc := func(pocFile string) {
//		if utils.Exists(pocFile) && utils.IsFile(pocFile) {
//			pocPath, err := filepath.Abs(pocFile)
//			if err != nil {
//				log.Printf("Get poc filepath error: "+pocFile, 4)
//			}
//			utils.DebugF("Load poc file: %v", pocFile)
//
//			xrayPoc, err := xray_parse.ParsePoc(pocPath)
//			if err == nil {
//				xrayPocMap[pocPath] = *xrayPoc
//				return
//			}
//			nucleiPoc, err := nuclei_parse.ParsePoc(pocPath)
//
//			if err == nil {
//				nucleiPocMap[pocPath] = *nucleiPoc
//				return
//			}
//
//			if err != nil {
//				utils.WarningF("Poc[%s] Parse error", pocFile)
//			}
//
//		} else {
//			utils.WarningF("Poc file not found: '%v'", pocFile)
//		}
//	}
//
//	for _, pocFile := range *pocs {
//		LoadPoc(pocFile)
//	}
//	for _, pocPath := range *pocPaths {
//		utils.DebugF("Load from poc path: %v", pocPath)
//
//		pocFiles, err := filepath.Glob(pocPath)
//		if err != nil {
//			log.Printf("Path glob match error: "+err.Error(), 6)
//		}
//		for _, pocFile := range pocFiles {
//			// 只解析yml或yaml文件
//			if strings.HasSuffix(pocFile, ".yml") || strings.HasSuffix(pocFile, ".yaml") {
//				LoadPoc(pocFile)
//			}
//		}
//
//	}
//
//	utils.InfoF("Load [%d] xray poc(s), [%d] nuclei poc(s)", len(xrayPocMap), len(nucleiPocMap))
//
//	return xrayPocMap, nucleiPocMap
//}
