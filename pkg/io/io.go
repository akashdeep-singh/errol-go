package io

//go:generate mockgen -source io.go -destination ../mocks/mock_io.go -package mocks

type IO interface {
	Write(writer interface{}, data []byte) error
	Read(reader interface{}) ([]byte, error)
}
