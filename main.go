package main

import (
	"context"
	"eventmapper/configs"
	"eventmapper/middlewares"
	"flag"
	"github.com/WajoxSoftware/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
)

var configfile = flag.String("configfile", "config.yml", "load config from `file`")
var cpuprofile = flag.String("cpuprofile", "pprof/cpu.pprof", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "pprof/mem.mprof", "write memory profile to `file`")

func main() {
	closeCh := make(chan bool)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	log.Printf("[x] Starting application...")

	log.Printf("[x] Load config")

	if *configfile == "" {
		log.Fatal("could not read config file")
	}

	config := configs.LoadConfigFile(*configfile)

	log.Printf("[x] Create router")
	routerHandler := CreateNewRouterHandler(config)

	log.Printf("[x] Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.HttpAuthType, config.HttpAuthParams))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(routerHandler)

	if config.DisableHandlers {
		log.Printf("[*] Handlers are disabled")
	} else {
		log.Printf("[x] Start events listener")
		BindEventsHandlers(config, closeCh)
	}

	if config.DisableGrpc {
		log.Printf("[*] GRPC is disabled")
	} else {
		log.Printf("[x] Start grpc server")
		StartGrpc(config)
	}

	var httpServer *http.Server
	if config.DisableHttp {
		log.Printf("[*] Http is disabled")
	} else {
		log.Printf("[x] Start http server")
		httpServer = StartHttpsServer(mware, config)
	}

	<-stop

	closeCh <- true

	if httpServer != nil {
		httpServer.Shutdown(context.Background())
	}
}
