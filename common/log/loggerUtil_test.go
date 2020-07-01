package clog

import "testing"

func TestInfo(t *testing.T) {
	Info("测试一下")
}

func TestError(t *testing.T) {
	Error("测试一下")
}

func TestWarning(t *testing.T) {
	Warning("测试一下")
}
