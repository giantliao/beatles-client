// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/giantliao/beatles-client-lib/app/cmd"
	"github.com/giantliao/beatles-client-lib/config"
	"github.com/giantliao/beatles-client-lib/webmain"
	"github.com/giantliao/beatles-mac-client/setting"
	"github.com/sevlyar/go-daemon"
	"log"
	"net/http"
	"path"
)

var (
	Version   string
	Build     string
	BuildTime string
)

func main() {
	cmd.CmdVersion = Version
	cmd.CmdBuild = Build
	cmd.CmdBuildTime = BuildTime


	if _, err := http.Get("http://127.0.0.1:50211"); err == nil {
		webmain.OpenBrowser("http://127.0.0.1:50211")
		return
	}

	cmd.InitCfg()
	cfg := config.GetCBtlc()
	cfg.Save()

	daemondir := config.GetBtlcHomeDir()
	cntxt := daemon.Context{
		PidFileName: path.Join(daemondir, "beetle.pid"),
		PidFilePerm: 0644,
		LogFileName: path.Join(daemondir, "beetle.log"),
		LogFilePerm: 0640,
		WorkDir:     daemondir,
		Umask:       027,
		Args:        []string{},
	}
	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		log.Println("beetle client starting, please check log at:", path.Join(daemondir, "beetle.log"))

		return
	}
	defer cntxt.Release()

	webmain.StartWEBService(&setting.MacSetting{})
}
