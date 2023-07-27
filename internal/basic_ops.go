package internal

import (
	"github.com/pieterclaerhout/go-log"
	"reflect"
	"sync"
)

func SetKey(key string, val string, store *sync.Map) (any, bool) {
	value, isExists := store.Load(key)
	if isExists {
		if reflect.TypeOf(val) != reflect.TypeOf(value) {
			return "Type Mismatch for the key", false
		}
		store.Store(key, val)
		return "OK", true
	} else {
		store.Store(key, val)
		return "OK", true
	}
}

func GetKey(key string, store *sync.Map) (any, bool) {
	value, isExists := store.Load(key)
	log.Info(value, isExists)
	if isExists {
		return value, true
	} else {
		return "Nil", false
	}
}

func DelKey(key string, store *sync.Map) (any, bool) {
	_, keyExists := store.Load(key)
	if keyExists {
		store.Delete(key)
		return "OK", true
	} else {
		return "Key not Found", false
	}

}
