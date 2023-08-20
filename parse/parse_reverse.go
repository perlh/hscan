package parse

import (
	"hscan/common"
	xray_structs "hscan/web/pkg/xray/structs"
	"net/http"
	"strings"
)

func ParseReversePlatform(api, domain string) {
	// 选择平台
	if api != "" && domain != "" && strings.HasSuffix(domain, ".ceye.io") {
		common.RunningInfo.CeyeApi = api
		common.RunningInfo.CeyeDomain = domain
		common.RunningInfo.ReversePlatformType = xray_structs.ReverseType_Ceye
	} else {
		common.RunningInfo.ReversePlatformType = xray_structs.ReverseType_DnslogCN

		// 设置请求相关参数
		common.RunningInfo.DnslogCNGetDomainRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getdomain.php", nil)
		common.RunningInfo.DnslogCNGetRecordRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getrecords.php", nil)

	}
}
