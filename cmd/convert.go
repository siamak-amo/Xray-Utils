// SPDX-License-Identifier: GPL-3.0-or-later
package main

import (
	"fmt"
	"encoding/json"

	pkg "github.com/siamak-amo/v2utils/pkg"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/infra/conf/serial"
	"github.com/xtls/xray-core/main/confloader"
	_ "github.com/xtls/xray-core/main/confloader/external"
)

func apply_template(dst *conf.Config, template core.ConfigSource) {
	r, err := confloader.LoadConfig(template.Name)
	if nil != err {
		fmt.Println (err);
	} else {
		c, err := serial.ReaderDecoderByFormat[template.Format](r)
		if nil != err {
			fmt.Println (err);
		} else {
			*dst = *c  // for the first time

			// TODO: maybe accept template array and merge them via:
			// dst.Override(c, file.Name)
		}
	}
}

func (opt Opt) Convert_url2json(url string) {
	cf := conf.Config{}
	if "" != opt.Template.Name {
		apply_template (&cf, opt.Template)
	}

	var e error
	var umap pkg.URLmap
	umap, e = pkg.ParseURL(url);
	if nil != e {
		return
	}
	cf.OutboundConfigs, e = pkg.Gen_outbound(umap);
	if nil != e {
		return
	}

	// TODO: skip null and empty strings
	// TODO: flag to print by indent or not
	b, err := json.Marshal(cf)
	// b, err := json.MarshalIndent (cf, "", "    ")
	if err != nil {
		fmt.Println (e)
		return
	}
	println (string(b));
}
