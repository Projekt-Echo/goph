package log

import "go.uber.org/zap"

var Zapper *zap.Logger

func init() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.DisableStacktrace = true
	Zapper, _ = zapCfg.Build()
	Zapper.Info("Zapper Init")
}
