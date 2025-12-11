// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"runtime"

	// pkg "github.com/siamak-amo/v2utils/pkg"
	// log "github.com/siamak-amo/v2utils/log"
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/app/proxyman/inbound"
)

// Runs xray-core instance (none blocking)
func (opt *Opt) Run_Xray() error {
	var err error
	var cf *core.Config

	// Cleanup sh** we have done so far to make the config
	runtime.GC()

	cf, err = opt.CFG.Build()
	if nil != err {
		return err
	}
	if opt.client, err = core.New(cf); nil != err {
		return err
	}
	if err = opt.client.Start(); nil != err {
		return err
	}
	return nil
}

// Do NOT use opt.client after this call
func (opt Opt) Kill_Xray() {
	if nil != opt.client {
		if opt.client.IsRunning() {
			opt.client.Close()
			opt.client = nil
		}
	}
}
