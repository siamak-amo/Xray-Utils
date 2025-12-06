package pkg

import (
	"fmt"

	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
)


func Gen_main(input string) (dst *v4.Config, e error) {
	dst = &v4.Config{}
	if e = unmarshal_H (dst, input); nil != e {
		// log
	}
	return
}

func Gen_inbound(input string) (dst []v4.InboundDetourConfig, e error) {
	dst = make ([]v4.InboundDetourConfig, 0)
	if e = unmarshal_H (&dst, input); nil != e {
		// log
	}
	return
}

// internal
func Gen_outbound(args URLmap, template string) (dst []v4.OutboundDetourConfig, e error) {
	dst = make ([]v4.OutboundDetourConfig, 0)
	if e = unmarshal_H (&dst, template); nil != e {
		// log
		return
	}
	switch args[Protocol] {
	case "vless":
		v, e := Gen_vless (args)
		if e == nil {
			dst = append (dst, *v)
		} else {
			fmt.Printf("Vless Error:  %v\n", e)
		}
		break
	case "vmess":
		v, e := Gen_vmess (args)
		if e == nil {
			dst = append (dst, *v)
		} else {
			fmt.Printf("Vmess Error:  %v\n", e)
		}
		break

	default:
		return nil, not_implemented ("protocol " + args[Protocol])
	}
	return
}


