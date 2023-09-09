package models

import (
	"fmt"
)

type Resp struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetModelByFields[I any](tableName string, fieldMap map[string]interface{}) (i *I, err error) {
	session := Engine.Table(tableName)
	for field, val := range fieldMap {
		session.Where(fmt.Sprintf("%s = ?", field), val)
	}

	i = new(I)
	_, err = session.Get(i)
	return
}
