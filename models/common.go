package models

import (
	"fmt"
)

type Resp struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GetModelByFields 通过表字段获取1条数据
func GetModelByFields[I any](tableName string, fieldMap map[string]interface{}) (i *I, err error) {
	session := Engine.Table(tableName)
	for field, val := range fieldMap {
		session.Where(fmt.Sprintf("%s = ?", field), val)
	}
	session.Limit(1)

	i = new(I)
	_, err = session.Get(i)
	return
}
