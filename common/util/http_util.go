package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	u "net/url"
	"strings"
)

var methodError = errors.New("请求方式错误,暂时只支持GET和post")

func Request[I any, R any](i *I, url string, method string, headers map[string]string) (*R, error) {
	var r R
	var err error
	var resp *http.Response

	switch strings.ToUpper(method) {
	default:
		return nil, methodError
	case "GET":
		data := u.Values{} // url encode（携带get请求参数）

		if i != nil {
			marshal, _ := json.Marshal(i)

			var temp map[string]string
			err = json.Unmarshal(marshal, &temp)
			if err != nil {
				return nil, err
			}

			for k, v := range temp {
				data.Set(k, v)
			}

		}

		encode := data.Encode()
		parse, _ := u.Parse(url)
		parse.RawPath = encode
		var request *http.Request
		request, err = http.NewRequest("GET", parse.String(), nil)
		if err != nil {
			//common.LoggerIns.Error("构建请求失败")
			return nil, err
		}

		//设置头
		setHeaders(request, headers)

		resp, err = http.DefaultClient.Do(request)

	case "POST":
		var bs []byte

		if i != nil {
			bs, err = json.Marshal(i)
			if err != nil {
				return nil, err
			}
		}

		reader := bytes.NewReader(bs)
		var request *http.Request
		request, err = http.NewRequest("POST", url, reader)
		if err != nil {
			//common.LoggerIns.Error("构建请求失败")
			return nil, err
		}

		//设置头
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
		setHeaders(request, headers)

		client := http.Client{}
		// 返回服务端的响应数据
		resp, err = client.Do(request)
	}

	if err != nil {
		//common.LoggerIns.Errorf("请求返回失败: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，解析body错误: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，body反序列化错误: %s", err.Error())
		return nil, err
	}

	return &r, nil
}

func Post[I any, R any](i *I, url string, headers map[string]string) (*R, error) {

	var resp *http.Response
	var err error

	var bs []byte
	if i != nil {
		bs, err = json.Marshal(i)
		if err != nil {
			return nil, err
		}
	}

	reader := bytes.NewReader(bs)
	var request *http.Request
	request, err = http.NewRequest("POST", url, reader)
	if err != nil {
		//common.LoggerIns.Error("构建请求失败")
		return nil, err
	}

	//common.LoggerIns.Infof("POST请求:URL=[%s],headers=[%#v]", url, headers)

	//设置头
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	setHeaders(request, headers)

	client := http.Client{}
	// 返回服务端的响应数据
	resp, err = client.Do(request)
	if err != nil {
		//common.LoggerIns.Error("请求失败:%s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，解析body错误: %s", err.Error())
		return nil, err
	}

	var r R
	err = json.Unmarshal(body, &r)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，body反序列化错误: %s", err.Error())
		return nil, err
	}

	return &r, nil
}

func Get[I any, R any](i *I, url string, headers map[string]string) (*R, error) {
	data := u.Values{} // url encode（携带get请求参数）

	var resp *http.Response
	var err error

	if i != nil {
		marshal, _ := json.Marshal(i)
		var temp map[string]string
		err = json.Unmarshal(marshal, &temp)
		if err != nil {
			return nil, err
		}
		for k, v := range temp {
			data.Set(k, v)
		}
	}

	encode := data.Encode()
	parse, _ := u.Parse(url)
	parse.RawPath = encode

	var request *http.Request
	request, err = http.NewRequest("GET", parse.String(), nil)
	if err != nil {
		//common.LoggerIns.Error("构建请求失败")
		return nil, err
	}

	//设置头
	setHeaders(request, headers)

	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回失败: %s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，解析body错误: %s", err.Error())
		return nil, err
	}

	var r R
	err = json.Unmarshal(body, &r)
	if err != nil {
		//common.LoggerIns.Errorf("请求返回成功，body反序列化错误: %s", err.Error())
		return nil, err
	}

	return &r, nil
}

func setHeaders(request *http.Request, m map[string]string) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			request.Header.Set(k, v)
		}
	}
}
