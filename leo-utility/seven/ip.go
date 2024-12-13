package seven

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/kakami/pkg/convert"
	"github.com/valyala/fasthttp"

	aw "github.com/deanishe/awgo"
)

type ipJSON struct {
	IP         string `json:"ip"`
	Pro        string `json:"pro"`
	ProCode    string `json:"proCode"`
	City       string `json:"city"`
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
	Err        string `json:"err"`
}

type ipAPI struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

var client *fasthttp.Client

// BuildIPQueryFeedback ...
func BuildIPQueryFeedback(wf *aw.Workflow, ip string) {
	var err error
	if iip := net.ParseIP(ip); iip == nil {
		wf.NewItem("> ip xxx.xxx.xxx.xxx")
		return
	}

	// http://whois.pconline.com.cn/ipJson.jsp?ip=58.216.12.57&json=true
	if client == nil {
		client = &fasthttp.Client{}
	}
	urlString := fmt.Sprintf("http://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true", ip)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(urlString)
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	if err = client.Do(req, resp); err != nil {
		wf.NewItem("> query failed")
		return
	}

	if resp.StatusCode() != 200 {
		wf.NewItem("> fetch error")
		return
	}

	bd, err := convert.DecodeGBK(resp.Body())
	if err != nil {
		wf.NewItem("> decode failed")
		return
	}
	// fmt.Println(bd)

	var ipJ ipJSON
	if err = json.Unmarshal(bd, &ipJ); err != nil {
		wf.NewItem("> parse json failed")
		return
	}

	if ipJ.City == "" || ipJ.Pro == "" {
		buildIPAPIFeedback(wf, ip)
		return
	}

	// fmt.Println(ipJ.Addr)
	wf.NewItem("> " + ipJ.Addr)
}

func buildIPAPIFeedback(wf *aw.Workflow, ip string) {
	var err error
	// http://ip-api.com/json/58.216.12.57?lang=zh-CN
	urlString := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(urlString)
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	if err = client.Do(req, resp); err != nil {
		wf.NewItem("ip-api fetch error")
		return
	}

	if resp.StatusCode() != 200 {
		wf.NewItem("> ip-api fetch error")
		return
	}

	var ipA ipAPI
	if err = json.Unmarshal(resp.Body(), &ipA); err != nil {
		wf.NewItem("> ip-api parse failed")
		return
	}

	wf.NewItem("> " + ipA.Country + "," + ipA.City)
}
