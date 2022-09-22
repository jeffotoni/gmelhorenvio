package fmts

import (
	gcat "github.com/jeffotoni/gconcat"
)

//Concat contaquena
func Concat(strs ...interface{}) string {
	return gcat.Concat(strs...)
}

//Concat contaquena
func ConcatStr(strs ...string) string {
	return gcat.ConcatStr(strs...)
}
