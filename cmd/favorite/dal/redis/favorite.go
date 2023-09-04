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

package redis

import (
	"strconv"
)

const (
	likeSuffix  = ":like"
	likedSuffix = ":liked"
)

type (
	Favorite struct{}
)

// AddLike 往集合 user_id:like 中添加 video_id
func (f Favorite) AddLike(user_id, video_id int64) {
	add(rdb, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

// AddLiked 往集合 video_id:liked 中添加 user_id
func (f Favorite) AddLiked(user_id, video_id int64) {
	add(rdb, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

// DelLike 从集合 user_id:like 中删除 video_id
func (f Favorite) DelLike(user_id, video_id int64) {
	del(rdb, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

// DelLiked 从集合 video_id:liked 中删除 user_id
func (f Favorite) DelLiked(user_id, video_id int64) {
	del(rdb, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

// CheckLike 检查 user_id:like 是否存在
func (f Favorite) CheckLike(user_id int64) bool {
	return check(rdb, strconv.FormatInt(user_id, 10)+likeSuffix)
}

// CheckLiked 检查 video_id:liked 是否存在
func (f Favorite) CheckLiked(video_id int64) bool {
	return check(rdb, strconv.FormatInt(video_id, 10)+likedSuffix)
}

// ExistLike 检查 user_id:like 中是否存在 video_id
func (f Favorite) ExistLike(user_id, video_id int64) bool {
	return exist(rdb, strconv.FormatInt(user_id, 10)+likeSuffix, video_id)
}

// ExistLiked 检查 video_id:liked 中是否存在 user_id
func (f Favorite) ExistLiked(user_id, video_id int64) bool {
	return exist(rdb, strconv.FormatInt(video_id, 10)+likedSuffix, user_id)
}

// CountLike 统计 user_id:like 中的数量
func (f Favorite) CountLike(user_id int64) (int64, error) {
	return count(rdb, strconv.FormatInt(user_id, 10)+likeSuffix)
}

// CountLiked 统计 video_id:liked 中的数量
func (f Favorite) CountLiked(video_id int64) (int64, error) {
	return count(rdb, strconv.FormatInt(video_id, 10)+likedSuffix)
}

// GetLike 获取 user_id:like 中的所有 video_id
func (f Favorite) GetLike(user_id int64) []int64 {
	return get(rdb, strconv.FormatInt(user_id, 10)+likeSuffix)
}

// GetLiked 获取 video_id:liked 中的所有 user_id
func (f Favorite) GetLiked(video_id int64) []int64 {
	return get(rdb, strconv.FormatInt(video_id, 10)+likedSuffix)
}
