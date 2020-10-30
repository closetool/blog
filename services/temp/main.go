package main

import (
	"bytes"
	"flag"
	"os"
	"time"

	"github.com/closetool/blog/system/messaging"
	"github.com/sirupsen/logrus"
)

func main() {
	data := flag.String("d", "", "data")
	queueName := flag.String("q", "", "queue name")
	flag.Parse()
	bts := []byte(*data)
	messaging.Client = &messaging.MessagingClient{}
	messaging.Client.ConnectToBroker("amqp://guest:guest@localhost:5672/")
	go func() {
		timer := time.NewTimer(3 * time.Second)
		select {
		case <-timer.C:
			logrus.Errorf("time out, check if function was called in main method")
			os.Exit(1)
		}
	}()
	rpl, err := messaging.Client.PublishOnQueueWaitReply(bts, *queueName)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	if !bytes.Contains(rpl, []byte("00000")) {
		logrus.Errorln(string(rpl))
		return
	}

	logrus.Infoln(string(rpl))
}
