package util

import (
	"encoding/json"
	"fmt"
	"leo/internal/pb"
	"net"

	"github.com/kakami/pkg/convert"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

func ipCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ip",
		Run: ipToLocation,
	}

	return cmd
}

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

func ipToLocation(_ *cobra.Command, args []string) {
	var err error
	if len(args) < 1 {
		pb.Wf.NewItem("> ip xxx.xxx.xxx.xxx")
		return
	}
	ip := args[0]
	if iip := net.ParseIP(ip); iip == nil {
		pb.Wf.NewItem("> ip xxx.xxx.xxx.xxx")
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
		pb.Wf.NewItem("> query failed")
		return
	}

	if resp.StatusCode() != 200 {
		pb.Wf.NewItem("> fetch error")
		return
	}

	bd, err := convert.DecodeGBK(resp.Body())
	if err != nil {
		pb.Wf.NewItem("> decode failed")
		return
	}
	// fmt.Println(bd)

	var ipJ ipJSON
	if err = json.Unmarshal(bd, &ipJ); err != nil {
		pb.Wf.NewItem("> parse json failed")
		return
	}

	if ipJ.City == "" || ipJ.Pro == "" {
		useAPI(ip)
		return
	}

	// fmt.Println(ipJ.Addr)
	pb.Wf.NewItem("> " + ipJ.Addr)
}

func useAPI(ip string) {
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
		pb.Wf.NewItem("ip-api fetch error")
		return
	}

	if resp.StatusCode() != 200 {
		pb.Wf.NewItem("> ip-api fetch error")
		return
	}

	var ipA ipAPI
	if err = json.Unmarshal(resp.Body(), &ipA); err != nil {
		pb.Wf.NewItem("> ip-api parse failed")
		return
	}

	pb.Wf.NewItem("> " + ipA.Country + "," + ipA.City)
}
