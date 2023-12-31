package mtproto

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"open.chat/pkg/log"
)

// FILE_MIGRATE_X = 303000;
// PHONE_MIGRATE_X = 303001;
// NETWORK_MIGRATE_X = 303002;
// USER_MIGRATE_X = 303003;
//
// ERROR_SEE_OTHER code has _X is dc number, We use custom NewXXXX()
func NewFileMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("FILE_MIGRATE_%d: %s", dc, message),
	}}
}

func NewFileMigrateX2(dc int) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("FILE_MIGRATE_%d", dc),
	}}
}

func NewPhoneMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("PHONE_MIGRATE_%d: %s", dc, message),
	}}
}

func NewPhoneMigrateX2(dc int) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("PHONE_MIGRATE_%d", dc),
	}}
}

func NewNetworkMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("NETWORK_MIGRATE_%d: %s", dc, message),
	}}
}

func NewNetworkMigrateX2(dc int) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("NETWORK_MIGRATE_%d", dc),
	}}
}

func NewUserMigrateX(dc int32, message string) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("USER_MIGRATE_%d: %s", dc, message),
	}}
}

func NewUserMigrateX2(dc int32) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_ERROR_SEE_OTHER),
		ErrorMessage: fmt.Sprintf("USER_MIGRATE_%d", dc),
	}}
}

// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
//
func NewFloodWaitX(second int32, message string) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_FLOOD),
		ErrorMessage: fmt.Sprintf("FLOOD_WAIT_%d: %s", second, message),
	}}
}

func NewFloodWaitX2(second int) *TLRpcError {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(TLRpcErrorCodes_FLOOD),
		ErrorMessage: fmt.Sprintf("FLOOD_WAIT_%d", second),
	}}
}

// normal code NewXXX
func NewRpcError(code int32, message string) (err *TLRpcError) {
	if name, ok := TLRpcErrorCodes_name[int32(code)]; ok {
		if code <= int32(TLRpcErrorCodes_OTHER2) {
			err = &TLRpcError{Data2: &RpcError{
				ErrorCode:    code,
				ErrorMessage: fmt.Sprintf("%s: %s", name, message),
			}}
		} else {
			switch code {
			// Not
			case int32(TLRpcErrorCodes_FILE_MIGRATE_X),
				int32(TLRpcErrorCodes_NETWORK_MIGRATE_X),
				int32(TLRpcErrorCodes_PHONE_MIGRATE_X),
				int32(TLRpcErrorCodes_USER_MIGRATE_X):
				err = &TLRpcError{Data2: &RpcError{
					ErrorCode:    int32(TLRpcErrorCodes_OTHER2),
					ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use New%s(dc, %s), ", name, name, message),
				}}
				log.Error(err.Error())

			case int32(TLRpcErrorCodes_FLOOD_WAIT_X):
				err = &TLRpcError{Data2: &RpcError{
					ErrorCode:    int32(TLRpcErrorCodes_FLOOD),
					ErrorMessage: fmt.Sprintf("FLOOD_WAIT_%s: %s", name, name),
				}}
				log.Error(err.Error())
			default:
				code2 := code / 1000
				if code2 == 0 {
					code2 = code
				}

				err = &TLRpcError{Data2: &RpcError{
					// subcode = code * 1000 + i
					ErrorCode:    int32(code2),
					ErrorMessage: name,
				}}
			}
		}
	} else {
		err = &TLRpcError{Data2: &RpcError{
			// subcode = code * 10000 + i
			ErrorCode:    int32(TLRpcErrorCodes_INTERNAL),
			ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: code = %d, message = %s", code, message),
		}}
	}

	return
}

// normal code NewXXX
func NewRpcError2(code TLRpcErrorCodes) (err *TLRpcError) {
	if name, ok := TLRpcErrorCodes_name[int32(code)]; ok {
		if code <= TLRpcErrorCodes_OTHER2 {
			err = &TLRpcError{Data2: &RpcError{
				ErrorCode:    int32(code),
				ErrorMessage: name,
			}}
		} else {
			switch code {
			// Not
			case TLRpcErrorCodes_FILE_MIGRATE_X,
				TLRpcErrorCodes_NETWORK_MIGRATE_X,
				TLRpcErrorCodes_PHONE_MIGRATE_X,
				TLRpcErrorCodes_USER_MIGRATE_X:
				err = &TLRpcError{Data2: &RpcError{
					ErrorCode:    int32(TLRpcErrorCodes_OTHER2),
					ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use New%s(dc), ", name, name),
				}}
				log.Errorf(err.Error())
			case TLRpcErrorCodes_FLOOD_WAIT_X:
				err = &TLRpcError{Data2: &RpcError{
					ErrorCode:    int32(TLRpcErrorCodes_FLOOD),
					ErrorMessage: fmt.Sprintf("INTERNAL_SERVER_ERROR: Not invoke NewRpcError(%s), please use NewFloodWaitX2(seconds), ", name),
				}}
				log.Error(err.Error())
			default:
				code2 := code / 1000
				if code2 == 0 {
					code2 = code
				}

				err = &TLRpcError{Data2: &RpcError{
					// subcode = code * 1000 + i
					ErrorCode:    int32(code2),
					ErrorMessage: name,
				}}
			}
		}
	} else {
		err = &TLRpcError{Data2: &RpcError{
			// subcode = code * 10000 + i
			ErrorCode:    int32(TLRpcErrorCodes_INTERNAL),
			ErrorMessage: "INTERNAL_SERVER_ERROR",
		}}
	}

	return
}

// Impl error interface
func (e *TLRpcError) IsOK() bool {
	if e == nil {
		return true
	}
	return e.GetErrorCode() == int32(TLRpcErrorCodes_ERROR_CODE_OK)
}

func (e *TLRpcError) Error() string {
	return e.DebugString()
}

func (e *TLRpcError) Code() int {
	return int(e.GetErrorCode())
}

func (e *TLRpcError) Message() string {
	return e.GetErrorMessage()
}

func (e *TLRpcError) Details() []interface{} {
	return nil
}

// Impl error interface
func (e *TLRpcError) ToGrpcStatus() *status.Status {
	return status.New(codes.Internal, e.Error())
}

func FromGRPCStatus(status *status.Status) *TLRpcError {
	return NewRpcError(int32(status.Code()), status.Message())
}

/*
// Impl error interface
func (e *TLRpcError) ToMetadata() (metadata.MD) {
	// return status.New(codes.Internal, e.Error())
	if name2, ok := TLRpcErrorCodes_name[e.ErrorCode]; ok {
		return metadata.Pairs(
			"rpc_error_code", name2,
			"rpc_error_message", e.ErrorMessage)
	}

	return metadata.Pairs(
		"rpc_error_code", "OTHER2",
		"rpc_error_message", fmt.Sprintf("INTERNAL_SERVER_ERROR: %s", e.ErrorMessage))
}

func NewRpcErrorFromMetadata(md metadata.MD) (*TLRpcError, error) {
	e := &TLRpcError{}

	if v, ok := getFirstKeyVal(md, "rpc_error_code"); ok {
		if code, ok := TLRpcErrorCodes_value[v]; !ok {
			return nil, fmt.Errorf("Invalid rpc_error_code: %s", v)
		} else {
			e.ErrorCode = code
		}
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_code")
	}

	if v, ok := getFirstKeyVal(md, "rpc_error_message"); !ok {
		e.ErrorMessage = v
	} else {
		return nil, fmt.Errorf("Not found metadata's key: rpc_error_message")
	}

	return e, nil
}
*/
