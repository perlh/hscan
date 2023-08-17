package get

import (
	"bytes"
	"hscan/discover/proxy"
	"hscan/utils/logger"
	"time"
)

func TcpProtocol(host string, port int, timeout int) ([]byte, error) {
	conn, err := proxy.ConnProxyTcp(host, port, timeout)
	if logger.DebugError(err) {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(time.Duration(2) * time.Second))
	reply := make([]byte, 256)
	_, err = conn.Read(reply)

	var buffer [256]byte
	if err == nil && bytes.Equal(reply[:], buffer[:]) == false {
		if conn != nil {
			_ = conn.Close()
		}
		return reply, nil

	}
	conn, err = proxy.ConnProxyTcp(host, port, timeout)
	if logger.DebugError(err) {
		return nil, err
	}
	msg := "GET /test HTTP/1.1\r\n\r\n"
	_, err = conn.Write([]byte(msg))
	if logger.DebugError(err) {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(time.Duration(2) * time.Second))
	reply = make([]byte, 256)
	_, _ = conn.Read(reply)
	if conn != nil {
		_ = conn.Close()
	}
	return reply, nil
}
