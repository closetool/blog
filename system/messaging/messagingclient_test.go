package messaging

import (
	"log"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/streadway/amqp"
)

func TestMessageHandlerLoop(t *testing.T) {

	var wg sync.WaitGroup
	var invocations = 0

	var handlerFunction = func(d amqp.Delivery) {
		log.Println("In handlerFunction")
		invocations = invocations + 1
		wg.Done()
	}

	Convey("Given", t, func() {
		var messageChannel = make(chan amqp.Delivery, 1)
		go consumeLoop(messageChannel, handlerFunction)

		Convey("When", func() {
			wg.Add(3)
			d := amqp.Delivery{Body: []byte(""), ConsumerTag: ""}
			messageChannel <- d
			messageChannel <- d
			messageChannel <- d
			wg.Wait()
			Convey("Then", func() {
				So(invocations, ShouldEqual, 3)
			})
		})
	})
}
