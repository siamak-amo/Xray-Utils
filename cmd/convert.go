// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
)

// Converts proxy URL @url to template @opt.Template
// returns error Only on fatal failures
func (opt Opt) Convert_url2json(url string) (error) {
	if "" != opt.Template.Name {
		opt.Apply_template(&opt.CFG)
	}
	if e := opt.Init_Outbound_byURL(url); nil != e {
		return nil // not fatal
	}

	if err := opt.CFG_Out(); nil != err {
		return err // IO error is fatal (invalid paths / broken pipes)
	}
	return nil
}
