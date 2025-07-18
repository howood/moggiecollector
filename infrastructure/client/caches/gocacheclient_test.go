package caches_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/howood/moggiecollector/infrastructure/client/caches"
)

func Test_GoCacheClient(t *testing.T) {
	t.Parallel()

	setkey := "testkey"
	setdata := "setdata"
	ctx := context.Background()
	client := caches.NewGoCacheClient()
	//nolint:errcheck
	client.Set(ctx, setkey, setdata, 60*time.Second)
	getdata, ok, err := client.Get(ctx, setkey)
	if err != nil {
		t.Fatalf("failed to get cache: %v", err)
	}
	if !ok {
		t.Fatalf("failed to get cache")
	}
	//nolint:forcetypeassert
	if reflect.DeepEqual(getdata.(string), setdata) == false {
		t.Fatalf("failed compare cache data ")
	}
	//nolint:forcetypeassert
	t.Log(getdata.(string))
	t.Log("success GoCacheClient")
}
