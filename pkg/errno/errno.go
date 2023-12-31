// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                  = 0
	FalseCode                    = 10000
	ServiceErrCode               = 10001
	ParamErrCode                 = 10002
	UserAlreadyExistErrCode      = 10003
	AuthorizationFailedErrCode   = 10004
	UsernameOrPasswordNilErrCode = 10005
	UserNotExistErrCode          = 10006
	PasswordIncorrectErrCode     = 10007
	TokenEmptyErrCode            = 10008
	LoginFailedErrCode           = 10009
	VideoExceedMaxSizeErrCode    = 10010
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                  = NewErrNo(SuccessCode, "Success")
	ServiceErr               = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr                 = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr      = NewErrNo(UserAlreadyExistErrCode, "用户已存在")
	AuthorizationFailedErr   = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UsernameOrPasswordNilErr = NewErrNo(UsernameOrPasswordNilErrCode, "用户名或密码不能为空")
	UserNotExistErr          = NewErrNo(UserNotExistErrCode, "该用户不存在")
	PasswordIncorrectErr     = NewErrNo(PasswordIncorrectErrCode, "密码不正确")
	LoginFailedErr           = NewErrNo(LoginFailedErrCode, "用户名或密码不正确")
	TokenEmptyErr            = NewErrNo(TokenEmptyErrCode, "token 为空")
	VideoExceedMaxSizeErr    = NewErrNo(VideoExceedMaxSizeErrCode, "单个视频不得超过 128 MB")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
