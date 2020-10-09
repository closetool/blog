package exit

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Listen(handleExit func()) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-s
		handleExit()
		logrus.Infof("service shutdown")
		os.Exit(0)
	}()
}
