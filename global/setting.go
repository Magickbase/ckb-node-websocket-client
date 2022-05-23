package global

import (
	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/Magickbase/ckb-node-websocket-client/pkg/logger"
	"github.com/Magickbase/ckb-node-websocket-client/pkg/setting"
)

var (
	DatabaseSetting *setting.DatabaseSettingS
	AppSetting      *setting.AppSettingS
	Logger          *logger.Logger
	CKBClient       rpc.Client
	RPCSetting      *setting.RPCSettingS
	SystemCodeHash  *setting.SystemCodeHashS
)
