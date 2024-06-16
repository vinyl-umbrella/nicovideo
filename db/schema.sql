CREATE TABLE IF NOT EXISTS videos (
    vid INTEGER,
    title TEXT NOT NULL COLLATE utf8mb4_0900_as_ci,
    owner_nickname TEXT NOT NULL,
    description TEXT,
    duration INTEGER NOT NULL,
    registered_at DATETIME NOT NULL,
    view_count INTEGER NOT NULL DEFAULT 0,
    comment_count INTEGER NOT NULL DEFAULT 0,
    mylist_count INTEGER NOT NULL DEFAULT 0,
    like_count INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (vid),
    FULLTEXT KEY `title` (`title`) /*!50100 WITH PARSER `ngram` */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
