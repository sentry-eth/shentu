package utils

import "errors"

var (
	// ErrBaseAccount is the error for BaseAccount assertion
	ErrBaseAccount = errors.New("the account is not a BaseAccount")
	// ErrEmptyCVMCode is the error when the query of the CVM code is empty
	ErrEmptyCVMCode = errors.New("the cvm code is empty")
	// ErrEmptyCVMAbi is the error when the query of the CVM abi is empty
	ErrEmptyCVMAbi = errors.New("the cvm abi is empty")
)
