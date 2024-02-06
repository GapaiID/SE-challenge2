package routes_test

import (
	"time"
)

type LibRedisMock struct {
	IsCalled map[string]any
}

func NewLibRedisMock() *LibRedisMock {
	return &LibRedisMock{
		IsCalled: make(map[string]interface{}),
	}
}

func (r *LibRedisMock) Set(key string, value any, expiration time.Duration) error {
	r.IsCalled["Set"] = []any{value, expiration}
	return nil
}

func (r *LibRedisMock) Get(key string, value any) error {
	r.IsCalled["Get"] = []any{key, value}
	return nil
}

func (r *LibRedisMock) Delete(keys ...string) (bool, error) {
	return false, nil
}

func (r *LibRedisMock) Check(keys ...string) (bool, error) {
	return false, nil
}
