package test

import (
	"github.com/cholthi/gotrue/conf"
	"github.com/cholthi/gotrue/storage"
)

func SetupDBConnection(globalConfig *conf.GlobalConfiguration) (*storage.Connection, error) {
	return storage.Dial(globalConfig)
}
