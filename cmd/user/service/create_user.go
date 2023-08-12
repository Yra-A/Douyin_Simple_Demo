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

package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
	"io"

	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/user/dal/db"
)

type CreateUserService struct {
  ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
  return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.UserRegisterRequest) error {
  users, err := db.QueryUser(s.ctx, req.Username)
  if err != nil {
    return err
  }
  if len(users) != 0 {
    return errno.UserAlreadyExistErr
  }

  h := md5.New()
  if _, err = io.WriteString(h, req.Password); err != nil {
    return err
  }
  passWord := fmt.Sprintf("%x", h.Sum(nil))
  return db.CreateUser(s.ctx, []*db.User{{
    UserName: req.Username,
    Password: passWord,
  }})
}
