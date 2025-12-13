// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"os"
	"runtime"
	"syscall"
	"os/signal"

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
	if opt.Xray_instance, err = core.New(cf); nil != err {
		return err
	}
	if err = opt.Xray_instance.Start(); nil != err {
		return err
	}
	return nil
}

// Run_Xray (blocking)
func (opt *Opt) Exec_Xray() error {
	if e := opt.Run_Xray(); nil != e {
		return e
	}
	defer opt.Kill_Xray();

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	<-osSignals
	return nil
}

// Do NOT use opt.Xray_instance after this call
func (opt Opt) Kill_Xray() {
	if nil != opt.Xray_instance {
		if opt.Xray_instance.IsRunning() {
			opt.Xray_instance.Close()
			opt.Xray_instance = nil
		}
	}
}
