/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"github.com/Yra-A/Douyin_Simple_Demo/cmd/user/rpc"
	"github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/favorite"
	"sync"

	"github.com/Yra-A/Douyin_Simple_Demo/cmd/user/dal/db"
	user "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/user"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUser(user_id int64) (*user.User, error) {
	u := &user.User{}
	// TODO 修改 errChan 长度
	errChan := make(chan error, 4)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		dbUser, err := db.QueryUserByUserId(user_id)
		if err != nil {
			errChan <- err
		} else {
			u.Id = dbUser.ID
			u.Name = dbUser.UserName
			u.Avatar = dbUser.Avatar
			u.BackgroundImage = dbUser.BackgroundImage
			u.Signature = dbUser.Signature
		}
		wg.Done()
	}()

	go func() {
		WorkCount, err := db.GetWorkCountByUserId(user_id)
		if err != nil {
			errChan <- err
		} else {
			u.WorkCount = WorkCount
		}
		wg.Done()
	}()

	go func() {
		FavoriteCount, err := rpc.FavoriteCountByUserId(s.ctx, &favorite.FavoriteCountByUserIDRequest{UserId: user_id})
		if err != nil {
			errChan <- err
		} else {
			u.FavoriteCount = FavoriteCount.FavoriteCount
		}
		wg.Done()
	}()

	go func() {
		TotalFavorited, err := rpc.TotalFavoritedByAuthorId(s.ctx, &favorite.TotalFavoritedByAuthorIDRequest{AuthorId: user_id})
		if err != nil {
			errChan <- err
		} else {
			u.TotalFavorited = TotalFavorited.TotalFavorited
		}
		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return &user.User{}, result
	default:
	}
	u.Id = user_id
	return u, nil
}
