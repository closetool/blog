package amqp

import (
	"sync"

	"github.com/closetool/blog/system/messaging"
	"github.com/closetool/blog/system/reply"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var ConfigCache *sync.Map

func GetConfig() {
	messaging.Client.SubscribeToQueueAndReply("config.getConfig", "config.getConfig", func(d amqp.Delivery) []byte {
		logrus.Debugln(string(d.Body))
		keys := []string{}
		if err := jsoniter.Unmarshal(d.Body, &keys); err != nil {
			logrus.Debugln(err)
			return reply.ErrorBytes(reply.ParamError)
		}
		res := map[string]string{}
		for _, key := range keys {
			valueInt, ok := ConfigCache.Load(key)
			if ok {
				value, ok := valueInt.(string)
				if ok {
					res[key] = value
				}
			}
		}
		return reply.ModelBytes(res)
	})
}
