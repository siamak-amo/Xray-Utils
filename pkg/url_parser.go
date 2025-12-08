package pkg

import (
	"fmt"
	"errors"
	"net/url"
	"encoding/base64"
)

func ParseURL(link string) (URLmap, error) {
	u, e := url.Parse(link)
	if nil != e {
		return nil, e
	}

	switch (u.Scheme) {
	case "vless":
		return parse_vless_url (u), nil
	case "vmess":
		return nil, nil

	case "trojan","ss":
		return nil, not_implemented (u.Scheme)

	default:
		return nil, errors.New ("Invalid URL scheme")
	}
}

// 	url: "vless://uuid@address:port?key=val..."
func parse_vless_url (u *url.URL) (URLmap) {
	res := make (URLmap, 0)
	params := Str2Strr(u.Query())

	res[Protocol] = "vless"
	res[ServerPort] = u.Port()
	res[Vxess_ID] = u.User.Username()
	res[ServerAddress] = u.Hostname()
	res[Security] = params.Pop ("security")
	res[Vless_ENC] = params.Pop ("encryption")

	res[Network] = params.Pop ("type")
	switch (res[Network]) {
	case "ws":
		res[WS_Path] = params.Pop ("path")
		res[WS_Headers] = params.Pop ("host")
		break;

	case "tcp":
		res[TCP_HTTP_Host] = params.Pop ("host")
		res[TCP_HTTP_Path] = params.Pop ("path")
		res[TCP_HeaderType] = params.Pop ("headerType")
		break;

	case "grpc":
		res[GRPC_Mode] = params.Pop ("mode")
		res[GRPC_MultiMode] = params.Pop ("multiMode")
		res[GRPC_ServiceName] = params.Pop ("serviceName")
		break;
	default:
		break;
	}

	switch (res[Security]) {
	case "tls":
		res[TLS_fp] = params.Pop ("fp")
		res[TLS_sni] = params.Pop ("sni")
		res[TLS_ALPN] = params.Pop ("alpn")
		break;

	case "reality":
		res[REALITY_fp] = params.Pop ("fp")
		res[REALITY_sni] = params.Pop ("sni")
		res[REALITY_ShortID] = params.Pop ("sid")
		res[REALITY_SpiderX] = params.Pop ("spx")
		res[REALITY_PublicKey] = params.Pop ("pbk")
		break;

	default:
		break;
	}

	for key, v := range params {
		if len(v) >= 1 {
			fmt.Printf ("parse_vless_url: parameter '%v' was ignored.\n", key)
		}
	}
	return res
}
