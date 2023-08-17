package judge

import (
	"bytes"
	"encoding/hex"
	"hscan/utils/logger"

	"hscan/discover/proxy"
)

func TcpRDP(result map[string]interface{}, Args map[string]interface{}) bool {
	timeout := Args["Timeout"].(int)
	host := result["host"].(string)
	port := result["port"].(int)

	conn, err := proxy.ConnProxyTcp(host, port, timeout)
	if logger.DebugError(err) {
		return false
	}

	//msg1 := "\x03\x00\x00\x13\x0e\xe0\x00\x00\x00\x00\x00\x01\x00\x08\x00\x03\x00\x00\x00"
	msg2 := "\x03\x00\x00\x2b\x26\xe0\x00\x00\x00\x00\x00\x43\x6f\x6f\x6b\x69\x65\x3a\x20\x6d\x73\x74\x73\x68\x61\x73\x68\x3d\x75\x73\x65\x72\x30\x0d\x0a\x01\x00\x08\x00\x00\x00\x00\x00"
	_, err = conn.Write([]byte(msg2))
	if logger.DebugError(err) {
		return false
	}

	reply := make([]byte, 256)
	_, _ = conn.Read(reply)
	if conn != nil {
		_ = conn.Close()
	}

	var buffer [256]byte
	if bytes.Equal(reply[:], buffer[:]) {
		return false
	} else if hex.EncodeToString(reply[0:8]) != "030000130ed00000" {
		return false
	} else {
		result["protocol"] = "rdp"
	}

	os := map[string]string{}
	/*** msg1 os finger ***
	os["030000130ed000001234000209080002000000"]="Windows 7/Windows Server 2008"
	os["030000130ed00000123400021f080002000000"]="Windows 10/Windows Server 2019"
	os["030000130ed00000123400020f080002000000"]="Windows 8.1/Windows Server 2012 R2"
	*/
	os["030000130ed000001234000209080000000000"] = "Windows 7/Windows Server 2008 R2"
	os["030000130ed000001234000200080000000000"] = "Windows 7/Windows Server 2008"
	os["030000130ed000001234000201080000000000"] = "Windows Server 2008 R2"
	os["030000130ed000001234000207080000000000"] = "Windows 8/Windows server 2012"
	os["030000130ed00000123400020f080000000000"] = "Windows 8.1/Windows Server 2012 R2"
	os["030000130ed000001234000300080001000000"] = "Windows 10/Windows Server 2016"
	os["030000130ed000001234000300080005000000"] = "Windows 10/Windows 11/Windows Server 2019"

	for k, v := range os {
		if k == hex.EncodeToString(reply[0:19]) {
			result["banner.string"] = v
			return true
		}
	}
	result["banner.string"] = hex.EncodeToString(reply[0:19])
	result["banner.byte"] = reply
	return true
}
