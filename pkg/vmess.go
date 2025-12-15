// SPDX-License-Identifier: GPL-3.0-or-later
package pkg

import (
	"fmt"
	"strconv"
	"strings"
	"net/url"
	"encoding/json"
	"encoding/base64"
	"github.com/xtls/xray-core/infra/conf"
)

// internal
func Gen_vmess(args URLmap) (dst *conf.OutboundDetourConfig, e error) {
    map_normal (args, Vmess_Sec, "none")
    map_normal (args, ServerPort, "443")
    map_normal (args, Vmess_AlterID, "0")
    dst = &conf.OutboundDetourConfig{}
    if e = unmarshal_H (dst,
        fmt.Sprintf (
            `{
                "protocol": "%s",
                "settings": {"vnext": [
                  {
                    "address": "%s",
                    "port": %s,
                    "users": [{
                      "security": "%s",  "alterId": %s,  "id": "%s"
                    }]
                  }
                ]},
                "tag": "proxy"
             }`,
            args[Protocol], args[ServerAddress], args[ServerPort],
            args[Vmess_Sec], args[Vmess_AlterID], args[Vxess_ID],
        )); nil != e {
        // log
		return
    }
    if dst.StreamSetting, e = Gen_streamSettings (args); nil != e {
        // log
		return
    }
    return
}

// {"aid":"0", "host":"","net":"tcp","path":"/","ps":"⚡ @ViPVpn_v2ray","scy":"auto","sni":"","tls":"","type":"http","v":"2"}
// {"alpn":"h2,http/1.1","fp":"","host":"","net":"grpc","path":"","port":"2087","ps":"@freV2rayNG","scy":"auto","sni":"oko0085.shop","tls":"tls","type":"gun","v":"2"}
// {"host":"glweidf.sbs","net":"ws","path":"/linkws","scy":"auto","sni":"glweidf.sbs","tls":"tls","type":"none","v":"2"}
func Gen_vmess_URL(src *conf.OutboundDetourConfig) *url.URL {
	var vmess VmessVnext
	if e := json.Unmarshal (*src.Settings, &vmess); nil != e {
		return nil
	}
	if len(vmess.Vnext) == 0 || len(vmess.Vnext[0].Users) == 0 {
		return nil
	}
	vnext := vmess.Vnext[0]
	res := make(map[string]string, 0)
	res["add"] = vnext.Address
	res["port"] = strconv.Itoa(vnext.Port)
	res["id"] = vnext.Users[0].ID
	res["aid"] = strconv.Itoa(vnext.Users[0].AlterIds)
	res["scy"] = vnext.Users[0].Security

	stream := src.StreamSetting
	if nil != stream && nil != stream.Network {
		net := string(*stream.Network);
		res["net"] = net
		switch (net) {
		case "tcp":
			if v,e := encode_tcp_header(stream.TCPSettings.HeaderConfig); nil == e {
				res["type"] = v.Type
				res["path"] = v.Request.Path
				res["host"] = v.Request.Headers["Host"]
			}
			break;
		case "grpc":
			res["path"] = stream.GRPCSettings.ServiceName
			res["authority"] = stream.GRPCSettings.Authority
			if stream.GRPCSettings.MultiMode {
				res["type"] = "multi"
			}
			break;

		case "ws":
			res["host"] = stream.WSSettings.Host
			res["path"] = stream.WSSettings.Path
			res["type"] = "none"
			break;
		}
		if sec := stream.Security; "tls" == sec {
			res["tls"] = "tls";
			res["fp"] = stream.TLSSettings.Fingerprint
			res["allowInsecure"] = strconv.FormatBool(stream.TLSSettings.Insecure)
			res["sni"] = stream.TLSSettings.ServerName
			res["alpn"] = strings.Join(*stream.TLSSettings.ALPN, ",")
		}
	} else {
		res["net"] = "tcp";
	}

	j, err := json.Marshal(res);
	if nil != err {
		panic(err); // it's ours.
	}
	u := &url.URL{
		Scheme: "vmess",
		Host: base64.StdEncoding.EncodeToString(j),
	};
	return u
}
