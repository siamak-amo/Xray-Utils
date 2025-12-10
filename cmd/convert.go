// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"encoding/json"

	log "github.com/siamak-amo/v2utils/log"
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

	// TODO: skip null and empty strings
	// TODO: flag to print by indent or compact
	b, err := json.Marshal(opt.CFG)
	// b, err := json.MarshalIndent (cf, "", "    ")
	if err != nil {
		log.Errorf ("json.Marshal failed - %v\n", err);
		return nil // not fatal
	}

	if err := opt.Out(b); nil != err {
		return err // IO error is fatal (invalid paths / broken pipes)
	}
	return nil
}
