package test

import (
	"testing"

	"github.com/jianxin/goredis/core"
)

func TestServer(t *testing.T) {
	t.Run("SET name jianxin", func(t *testing.T) {
		success := core.Set("name", "jianxin")
		if !success {
			t.Error("SET success")
		}
	})
	t.Run("SET    ", func(t *testing.T) {
		success := core.Set("", "jianxin")
		if !success {
			t.Error("SET fail")
		}
	})
	t.Run("GET name", func(t *testing.T) {
		got, err := core.Get("name")
		want := "jianxin"
		if err != nil || got != want {
			t.Error("GET fail")
		}
	})
	t.Run("GET  ", func(t *testing.T) {
		_, err := core.Get("")
		if err == nil {
			t.Error("GET fail")
		}
	})
}
