package module

import (
	rendering "github.com/adm87/finch-rendering/module"
	resources "github.com/adm87/finch-resources/module"
	tiled "github.com/adm87/finch-tiled/module"
)

func RegisterModule() error {
	if err := rendering.RegisterModule(); err != nil {
		return err
	}
	if err := resources.RegisterModule(); err != nil {
		return err
	}
	if err := tiled.RegisterModule(); err != nil {
		return err
	}
	return nil
}
