namespace go favorite

struct FavoriteActionRequest {
    1: string token   (api.query="token")    // 用户鉴权token
    2: i64 video_id   (api.query="video_id")       // 视频id
    3: i32 action_type   (api.query="action_type")      // 1-点赞，2-取消点赞
    4: i64 user_id       // 用户id
}

struct FavoriteActionResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg // 返回状态描述
}

struct FavoriteListRequest {
    1: i64 user_id          // 用户id
    2: string token     (api.query="token")       // 用户鉴权token
}

struct FavoriteListResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: list<Video> video_list  // 用户点赞视频列表
}

struct Video {
    1: i64 id,                 // 视频唯一标识
    2: User author,            // 视频作者信息
    3: string play_url,        // 视频播放地址
    4: string cover_url,       // 视频封面地址
    5: i64 favorite_count,     // 视频的点赞总数
    6: i64 comment_count,      // 视频的评论总数
    7: bool is_favorite,       // true-已点赞，false-未点赞
    8: string title            // 视频标题
}

struct User {
    1: i64 id,                 // 用户id
    2: string name,            // 用户名称
    3: i64 follow_count, // 关注总数
    4: i64 follower_count, // 粉丝总数
    5: bool is_follow,         // true-已关注，false-未关注
    6: string avatar, // 用户头像
    7: string background_image, // 用户个人页顶部大图
    8: string signature,  // 个人简介
    9: i64 total_favorited, // 获赞数量
    10: i64 work_count,   // 作品数量
    11: i64 favorite_count // 点赞数量
}
//is_favorite   ueser_id like video_id
struct IsFavoriteRequest {
    1: i64 user_id (api.query="user_id")
    2: i64 video_id  (api.query="video_id")
}

struct IsFavoriteResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: bool  is_favorite
}

struct FavoriteCountRequest {
    1: i64 video_id  (api.query="video_id")
}
struct FavoriteCountResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: i64 favorite_count
}

struct FavoriteCountByUserIDRequest {
    1: i64 user_id  (api.query="user_id")
}

struct FavoriteCountByUserIDResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: i64 favorite_count
}

struct TotalFavoritedByAuthorIDRequest {
    1: i64 author_id  (api.query="author_id")
}

struct TotalFavoritedByAuthorIDResponse {
    1: i32 status_code,        // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
    3: i64 total_favorited
}

service FavoriteService {
    // 点赞操作
    FavoriteActionResponse FavoriteAction(1:required FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    // 获取喜欢列表
    FavoriteListResponse FavoriteList(1:required FavoriteListRequest req) (api.get="/douyin/favorite/list/")
    // 获取视频获赞数量
    FavoriteCountResponse FavoriteCount(1:required FavoriteCountRequest req)
    // 获取用户点赞过的视频数量
    FavoriteCountByUserIDResponse FavoriteCountByUserID(1:FavoriteCountByUserIDRequest req)// 获取喜欢计数
    // 获取作者的获赞数
    TotalFavoritedByAuthorIDResponse TotalFavoritedByAuthorID(1:TotalFavoritedByAuthorIDRequest req)
    // 是否喜欢
    IsFavoriteResponse IsFavorite(1:IsFavoriteRequest req)
}

