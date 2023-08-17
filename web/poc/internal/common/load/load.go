package utils

import (
	"hscan/utils/logger"
	"hscan/web/poc/utils"
	"path/filepath"
	"strings"

	"github.com/yargevad/filepathx"

	nuclei_parse "hscan/web/poc/pkg/nuclei/parse"
	nuclei_structs "hscan/web/poc/pkg/nuclei/structs"

	xray_parse "hscan/web/poc/pkg/xray/parse"
	xray_structs "hscan/web/poc/pkg/xray/structs"
)

// 读取pocs
func LoadPocs(pocPath string) (map[string]xray_structs.Poc, map[string]nuclei_structs.Poc) {
	xrayPocMap := make(map[string]xray_structs.Poc)
	nucleiPocMap := make(map[string]nuclei_structs.Poc)

	// 加载poc函数
	LoadPoc := func(pocFile string) {
		if utils.Exists(pocFile) && utils.IsFile(pocFile) {
			pocPath, err := filepath.Abs(pocFile)
			if err != nil {
				logger.Error("Get poc filepath error: " + pocFile)
			}
			utils.DebugF("Load poc file: %v", pocFile)

			xrayPoc, err := xray_parse.ParsePoc(pocPath)
			if err == nil {
				xrayPocMap[pocPath] = *xrayPoc
				return
			}
			nucleiPoc, err := nuclei_parse.ParsePoc(pocPath)

			if err == nil {
				nucleiPocMap[pocPath] = *nucleiPoc
				return
			}

			if err != nil {
				utils.WarningF("Poc[%s] Parse error", pocFile)
			}

		} else {
			utils.WarningF("Poc file not found: '%v'", pocFile)
		}
	}

	utils.InfoF("Load from poc path: %v", pocPath)

	pocFiles, err := filepathx.Glob(pocPath)
	if err != nil {
		logger.Error("Path glob match error: " + err.Error())
	}
	for _, pocFile := range pocFiles {
		// 只解析yml或yaml文件
		if strings.HasSuffix(pocFile, ".yml") || strings.HasSuffix(pocFile, ".yaml") {
			LoadPoc(pocFile)
		}
	}

	return xrayPocMap, nucleiPocMap
}
