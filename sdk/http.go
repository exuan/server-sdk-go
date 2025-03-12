package sdk

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"syscall"

	"github.com/astaxie/beego/httplib"
)

func (rc *RongCloud) do(b *httplib.BeegoHTTPRequest) (body []byte, err error) {
	return rc.httpRequest(b)
}

// Network errors that require domain switching
func isNetError(err error) bool {
	netErr, ok := err.(net.Error)
	if !ok {
		return false
	}
	// Timeout
	if netErr.Timeout() {
		return true
	}

	var opErr *net.OpError
	opErr, ok = netErr.(*net.OpError)
	if !ok {
		// URL error
		urlErr, ok := netErr.(*url.Error)
		if !ok {
			return false
		}
		opErr, ok = urlErr.Err.(*net.OpError)
		if !ok {
			return false
		}
	}

	switch t := opErr.Err.(type) {
	case *net.DNSError:
		return true
	case *os.SyscallError:
		if errno, ok := t.Err.(syscall.Errno); ok {
			switch errno {
			case syscall.ECONNREFUSED:
				return true
			case syscall.ETIMEDOUT:
				return true
			}
		}
	}

	return false
}

func (rc *RongCloud) httpRequest(b *httplib.BeegoHTTPRequest) (body []byte, err error) {
	// Use the global httpClient to avoid opening too many ports
	b.SetTransport(rc.globalTransport)
	resp, err := b.DoRequest()
	if err != nil {
		if isNetError(err) {
			rc.ChangeURI()
		}
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	rc.checkStatusCode(resp)
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err = checkHTTPResponseCode(body); err != nil {
		return nil, err
	}
	return body, err
}

// v2 api
func (rc *RongCloud) doV2(b *httplib.BeegoHTTPRequest) (body []byte, err error) {
	// Use the global httpClient to avoid opening too many ports
	b.SetTransport(rc.globalTransport)

	resp, err := b.DoRequest()
	if err != nil {
		if isNetError(err) {
			rc.ChangeURI()
		}
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	rc.checkStatusCode(resp)
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		body, err = ioutil.ReadAll(reader)
	} else {
		body, err = ioutil.ReadAll(resp.Body)
	}
	if err = checkHTTPResponseCodeV2(body); err != nil {
		return nil, err
	}
	return body, err
}

func checkHTTPResponseCode(rep []byte) error {
	code := codePool.Get().(CodeResult)
	defer codePool.Put(code)
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 200 {
		return code
	}
	return nil
}

// v2 api error
func checkHTTPResponseCodeV2(rep []byte) error {
	code := codePoolV2.Get().(CodeResultV2)
	defer codePoolV2.Put(code)
	if err := json.Unmarshal(rep, &code); err != nil {
		return err
	}
	if code.Code != 10000 && code.Code != 200 {
		return code
	}
	return nil
}
