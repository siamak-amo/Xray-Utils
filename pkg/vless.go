// SPDX-License-Identifier: GPL-3.0-or-later
package pkg

import (
	"fmt"
	"github.com/xtls/xray-core/infra/conf"
)

// internal
func Gen_vless(args URLmap) (dst *conf.OutboundDetourConfig, e error) {
    map_normal (args, Vless_ENC, "none")
    map_normal (args, ServerPort, "443")
    map_normal (args, Vless_Level, "0")
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
                      "encryption": "%s",  "level": %s,  "id": "%s"
                    }]
                  }
                ]},
                "tag": "proxy"
             }`,
            args[Protocol], args[ServerAddress], args[ServerPort],
            args[Vless_ENC], args[Vless_Level], args[Vxess_ID],
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
