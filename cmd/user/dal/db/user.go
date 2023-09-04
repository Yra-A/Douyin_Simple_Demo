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

package db

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/constants"
	"github.com/Yra-A/Douyin_Simple_Demo/pkg/errno"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID              int64  `gorm:"primaryKey";json:"id"`
	UserName        string `gorm:"type:varchar(255)"json:"user_name"`
	Password        string `gorm:"type:varchar(255)"json:"password"`
	Avatar          string `gorm:"type:varchar(255)"json:"avatar"`           // 用户头像 URL
	BackgroundImage string `gorm:"type:varchar(255)"json:"background_image"` // 用户背景图 URL
	Signature       string `gorm:"type:varchar(255)"json:"signature"`        // 用户个性签名
}

type Video struct {
	gorm.Model
	ID          int64
	AuthorID    int64
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
	Title       string
}

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (u User) TableName() string {
	return constants.UserTableName
}

func (u Favorite) TableName() string {
	return constants.FavoriteTableName
}

func (u Video) TableName() string {
	return constants.VideosTableName
}

// MGetUsers multiple get list of user info
//func MGetUsers(ctx context.Context, UserIds []int64) ([]*User, error) {
//  res := make([]*User, 0)
//  if len(UserIds) == 0 {
//    return res, nil
//  }
//
//  if err := DB.WithContext(ctx).Where("id in ?", UserIds).Find(&res).Error; err != nil {
//    return nil, err
//  }
//  return res, nil
//}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserByUserId get user in the database by user id
func QueryUserByUserId(user_id int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return nil, err
	}
	if user == (User{}) {
		err := errno.UserNotExistErr
		return nil, err
	}
	return &user, nil
}

func GetWorkCountByUserId(user_id int64) (int64, error) {
	var count int64
	if err := DB.Model(&Video{}).Where("author_id = ?", user_id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
