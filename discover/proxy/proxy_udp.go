package proxy

import (
	"hscan/common"
	"net"
	"strconv"
	"time"

	"hscan/utils/logger"

	"github.com/txthinking/socks5"
)

func ConnProxyUdp(host string, port int, timeout int) (net.Conn, error) {
	target := net.JoinHostPort(host, strconv.Itoa(port))
	scheme, address, proxyUri, err := common.RunningInfo.ProxySchema, common.RunningInfo.ProxyHost, common.RunningInfo.ProxyProxy, common.RunningInfo.ProxyErr
	if logger.DebugError(err) {
		return nil, err
	}
	var conn net.Conn

	if proxyUri == "" {
		conn, err = net.DialTimeout("udp", target, time.Duration(timeout)*time.Second)
		if logger.DebugError(err) {
			return nil, err
		}
	}

	if scheme == "http" {
		target := net.JoinHostPort(host, strconv.Itoa(port))
		conn, err = net.DialTimeout("udp", target, time.Duration(timeout)*time.Second)
		if logger.DebugError(err) {
			return nil, err
		}
	}

	if scheme == "socks5" {
		dialer, err := socks5.NewClient(address, "", "", timeout, timeout)
		if logger.DebugError(err) {
			logger.Error("Cannot initialize socks5 proxy")
			return nil, err
		}
		conn, err = dialer.Dial("udp", target)
		if logger.DebugError(err) {
			return nil, err
		}
	}

	err = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	if logger.DebugError(err) {
		if conn != nil {
			_ = conn.Close()
		}
		return nil, err
	}
	return conn, nil
}
