package gifxawssns

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	gifxaws "github.com/b2wdigital/goignite/fx/aws/v2"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {

	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			gifxaws.AWSModule(),
			fx.Provide(
				sns.NewFromConfig,
			),
		)
	})

	return options
}