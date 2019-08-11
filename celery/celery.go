package celery

import (
	celery "github.com/gocelery"
	"yespider-go/settings"
)
var (
	Client *celery.CeleryClient
)

func GetCeleryCli() *celery.CeleryClient{
	if Client == nil {
		Client, _ = celery.NewCeleryClient(
				celery.NewRedisCeleryBroker(settings.RedisSettings.Host),
				celery.NewRedisCeleryBackend(settings.RedisSettings.Host),
				5)

	}
	return Client
}
