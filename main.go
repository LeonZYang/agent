package main

import (
	"flag"
	"fmt"
	"github.com/LeonZYang/agent/cron"
	"github.com/LeonZYang/agent/funcs"
	"github.com/LeonZYang/agent/g"
	"os"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	g.InitRootDir()
	g.InitLocalIps()
	g.InitRpcClients()

	funcs.BuildMappers()

	go cron.InitDataHistory()

	cron.ReportAgentStatus()
	cron.SyncBuiltinMetrics()
	cron.SyncTrustableIps()
	cron.Collect()

	select {}

}
