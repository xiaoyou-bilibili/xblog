// Package db 数据库测试脚本
// @Description
// @Author 小游
// @Date 2021/05/25
package db

import "testing"

// go测试单个文件 go test -V internal/db/db_test.go run TestDatabaseBackup

func TestDatabaseBackup(t *testing.T) {
	DatabaseBackup()
}

func TestDatabaseRestore(t *testing.T) {
	err := DatabaseRestore()
	if err != nil {
		return
	}
}
