package httpresponse

import "errors"

var (
	ErrDeviceIDRequired = errors.New("device-id is required")

	ErrSourceRequired = errors.New("source is required")
	ErrSourceNotValid = errors.New("source not valid")

	ErrSignatureRequired = errors.New("signature is required")
	ErrSignatureNotValid = errors.New("signature not valid")

	ErrEpochRequired = errors.New("epoch is required")
)
