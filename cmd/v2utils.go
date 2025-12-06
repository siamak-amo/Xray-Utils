package main

import (
	"fmt"
	"encoding/json"

	v2 "github.com/siamak-amo/v2utils/pkg"
)

func main() {
	cfg, err := v2.Gen_main ("{}")
	// cfg, err = gen_main (
	// 	`{
    //         "log": {"loglevel": "info"},
    //         "dns": {"servers": ["1.1.1.1", "8.8.8.8"]},
    //         "routing": {
    //           "domainStrategy": "IPIfNonMatch",
    //           "rules": [{
    //             "ip": ["1.1.1.1"], "outboundTag": "proxy", "port": "53", "type": "field"
    //           }]
    //         }
    //      }`,
	// )

	// cfg.InboundConfigs, err = gen_inbound (
	// 	`[
    //         {
    //            "listen": "127.0.0.1", "port": 10808, "protocol": "socks",
    //            "settings": {"auth": "noauth", "udp": true, "userLevel": 8},
    //            "sniffing": {"destOverride": ["http", "tls"], "enabled": true},
    //            "tag": "socks"
    //         },
    //         {
    //            "listen": "127.0.0.1", "port": 10809, "protocol": "http",
    //            "settings": {"userLevel": 8},
    //            "tag": "http"
    //         }
    //      ]`,
	// )

	params := make (v2.URLmap)
	params[v2.ServerAddress] = "127.6.6.6"
	params[v2.ServerPort] = "777"
	params[v2.Security] = "tls"
	params[v2.TLS_sni] = "yourmom.xyz"
	
	params[v2.Network] = "ws"
	params[v2.WS_Path] = "/yourmom"
	// paramseaders] = `"Host": "abc",  "X-Fuck": "true"`
	
	params[v2.Protocol] = "vless"
	params[v2.Vxess_ID] = "eb46c717-3f24-4b58-a9d4-82df392e4eb6"
	params[v2.Vless_ENC] = "chert"
	
	cfg.OutboundConfigs, err = v2.Gen_outbound (params,
		`[
            {"protocol": "freedom", "settings": {}, "tag": "direct"},
            {"protocol": "blackhole", "settings": {"response": {"type": "http"}}, "tag": "block"}
         ]`,
	)

	if err != nil {
		fmt.Printf ("Error: %s\n", err)
		return
	}
	jsonStr, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(jsonStr))
}
