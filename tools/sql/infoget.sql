DELIMITER //
CREATE PROCEDURE favinfo()
BEGIN
	select id, author_id, favorite_count, comment_count, title from videos;
	select * from users;
	select * from favorites;
END //

DELIMITER ;

