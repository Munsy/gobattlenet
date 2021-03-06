package errors

import (
	"errors"
)

// Battle.net errors for clients and various packages.
var (
	ErrNoInterfaceSupplied = errors.New("nil interface was passed to constructor")
	ErrUnsupportedArgument = errors.New("unsupported argument type was passed to constructor")
	ErrUnresolvedEndpoint  = errors.New("unresolved endpoint")
	ErrInvalidDialectForm  = errors.New("invalid dialect format")
	ErrNoTokenSupplied     = errors.New("no token supplied")
	ErrNoKeySupplied       = errors.New("no key supplied")
	ErrNoVersionSupplied   = errors.New("no client version number supplied")
)
