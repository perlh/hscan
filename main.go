package main

import (
	"fmt"
	"hscan/common"
	"hscan/discover"
	"hscan/nonweb"
	"hscan/parse"
	"hscan/probe"
	"hscan/utils/logger"
	"hscan/web/fscan"
	"hscan/web/poc"
	"time"
)

func main() {
	startTime := time.Now()
	common.PrintBanner()
	parse.Flag(&common.InputInfo)
	parse.Parse(&common.InputInfo, &common.RunningInfo)
	probe.Probe()

	discover.Discover()
	// Xray v1 Scan
	fscan.Poc()
	//logger.Banner("x")

	poc.XrayPoc()
	poc.NucleiPoc()

	// Xray v2 Scan
	// TODO Xray v2
	// Nuclei Scan
	nonweb.Service()
	logger.Info(fmt.Sprintf("Task finish, consumption of time: %s", time.Now().Sub(startTime)))
}
