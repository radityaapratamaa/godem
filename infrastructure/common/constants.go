package common

import (
	"errors"
	"fmt"
)

var (
	ErrPatch        = errors.New("error patching")
	AccountCacheKey = fmt.Sprintf("%s:account", PrefixCacheKey)
)

const (
	ErrInputValidationCode = "INPUT_VALIDATION_ERROR"
	ErrUsecaseErrorCode    = "USECASE_ERROR"
	BucketName             = "sgrpatrol"

	EnvServerEncryptionKey = "SERVER_ENCRYPTION_KEY"
	MaxWorker              = 10
	PrefixCacheKey         = "godem"
)
