package builder

import "supercollector/internal/config"

type Builder struct {
	conf config.Config
}

func New(conf config.Config) *Builder {
	return &Builder{
		conf: conf,
	}
}
