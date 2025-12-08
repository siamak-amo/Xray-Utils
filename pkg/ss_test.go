// SPDX-License-Identifier: GPL-3.0-or-later
package pkg

import (
	"testing"
)


func TestGen_ss (t *testing.T) {
	tc := TestCase[ServerCFG] {T: t,
		Input: map[URLMapper]string {
	    	    Protocol:       "shadowsocks",
				ServerAddress:	"vpn.net",
				ServerPort:		"1234",
				SS_Method:      "chacha20-ietf-poly1305",
				SS_Password:    "p@ssw0rd",
			},
		Output: ServerCFG{},
	}

	v, e := Gen_ss (tc.Input)
	if nil != e {
		t.Fatalf ("gen_ss failed: %v\n", e)
		return
	}
	tc.Do(v);

	ss := tc.Output.Settings.Servers[0];
	tc.Assert (tc.Output.Protocol,    tc.Input[Protocol])
	tc.Assert (ss.Address,            tc.Input[ServerAddress])
	tc.Assert (ss.Port,               tc.Input[ServerPort])
	tc.Assert (ss.Method,             tc.Input[SS_Method])
	tc.Assert (ss.Password,           tc.Input[SS_Password])
}
