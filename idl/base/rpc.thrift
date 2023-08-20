namespace go rpc

// 视频信息
struct VideoInfo {
	1: i64 id
	2: i64 author_id
	3: string play_url
	4: string cover_url
	5: i64 favorite_count
	6: i64 comment_count
	7: string title
}

// 用户信息
struct UserInfo {
  1: i64 id
  2: string name
  3: i64 follow_count
  4: i64 follower_count
  5: bool is_follow
  6: string avatar
  7: string background_image
  8: string signature
  9: i64 total_favorited
  10: i64 work_count
  11: i64 favorite_count
}

