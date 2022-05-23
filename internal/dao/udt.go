package dao

import "github.com/Magickbase/ckb-node-websocket-client/internal/model"

func (d Dao) GetUdtByTypeHash(typeHash string) (*model.Udt, error) {
	udt := model.Udt{
		TypeHash: typeHash,
	}
	return udt.GetByTypeHash(d.engine)
}
