package global

import (
	"cnhtc/gin-vue-admin/AppServer/config"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *sqlx.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)
