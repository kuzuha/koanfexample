package config

import (
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/samber/do"
	"github.com/samber/mo"
	"koanfexample/echo"
	"strings"
)

type (
	Config struct {
		Greeting Greeting `koanf:"greeting"`
	}

	Greeting struct {
		Message string `koanf:"message"`
	}
)

func NewConfig(_ *do.Injector) (Config, error) {
	kf := koanf.New(".")
	mr := mo.EmptyableToOption(kf.Load(file.Provider("config.yaml"), yaml.Parser())).
		FlatMap(func(error) mo.Option[error] {
			return mo.EmptyableToOption(kf.Load(file.Provider("config.json"), json.Parser()))
		})

	mr.MapNone(func() (error, bool) {
		err := kf.Load(env.Provider("KOANF", "_", func(s string) string {
			return strings.ReplaceAll(strings.ToLower(
				strings.TrimPrefix(s, "KOANF_")), "_", ".")
		}), yaml.Parser())
		return err, err != nil
	})

	c := Config{}
	mr.MapNone(func() (error, bool) {
		err := kf.Unmarshal("", &c)
		return err, err != nil
	})
	return c, mr.OrEmpty()
}

func Watch(in *echo.InjectorMiddleware) {
	err := file.Provider("config.yaml").Watch(func(event interface{}, err error) {
		_, err = NewConfig(nil)
		if err != nil {
			println(err.Error())
			return
		}
		in.Apply(func(in *do.Injector) {
			// do nothing, only refresh injector
		})
	})
	if err != nil {
		panic(err)
	}
}
