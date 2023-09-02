delete from `comments`;
delete from `favorites`;
delete from `messages`;
delete from `relations`;
-- delete from `videos`;
-- delete from `users`;
-- update users set follow_count = 0, follower_count = 0, total_favorited = 0, work_count = 0, favorite_count = 0;
update users set follow_count = 0, follower_count = 0, total_favorited = 0, work_count = 0, favorite_count = 0;
update videos set favorite_count = 0, comment_count = 0;
