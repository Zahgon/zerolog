package zerolog

import (
	"context"
)

func isNilValue(e error) bool { _ = "STUB: not implemented"; return false }

func appendFields(dst []byte, fields interface{}, stack bool, ctx context.Context, hooks []Hook) []byte {
	_ = "STUB: not implemented"
	return nil
}

func appendObject(dst []byte, obj LogObjectMarshaler, stack bool, ctx context.Context, hooks []Hook) []byte {
	_ = "STUB: not implemented"
	return nil
}

func appendFieldList(dst []byte, kvList []interface{}, stack bool, ctx context.Context, hooks []Hook) []byte {
	_ = "STUB: not implemented"
	return nil
}
