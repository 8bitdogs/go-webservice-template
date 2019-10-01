package main

import "github.com/antonmashko/envconf"

type config struct {
	// TODO: add required config
	// see: https://github.com/antonmashko/envconf#how-it-works
}

func parse() (*config, error) {
	var cfg config
	err := envconf.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
