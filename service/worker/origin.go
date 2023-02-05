package worker

import (
	"log"
	"time"

	"github.com/kardianos/service"
	"github.com/spf13/viper"

	"tdp-cloud/module/worker"
)

type origin struct{}

func (p *origin) Start(s service.Service) error {

	log.Println("service start")

	go p.run()
	return nil

}

func (p *origin) Stop(s service.Service) error {

	log.Println("service stop")
	return nil

}

func (p *origin) run() {

	defer p.timer()

	remote := viper.GetString("worker.remote")

	if err := worker.Daemon(remote); err != nil {
		log.Println(err)
	}

}

func (p *origin) timer() {

	log.Println("连接已断开，将在5秒后重试")
	time.Sleep(time.Second * 5)
	p.run()

}