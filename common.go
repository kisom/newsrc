package main

import "errors"

var (
	errUnknownLicense = errors.New(`newsrc: unknown license`)
	errFileExists     = errors.New(`newsrc: file exists`)
)

type opts struct {
	path    string
	license string
	force   bool
	body    string
}
