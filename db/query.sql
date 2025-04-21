CREATE TABLE
    IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(50)
    );

CREATE TABLE
    IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        session TEXT DEFAULT NULL
    );

CREATE TABLE
    IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(30),
        title VARCHAR(255),
        description TEXT,
        time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        postID INTEGER,
        username VARCHAR(30),
        comment TEXT,
        time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (postID) REFERENCES posts (id)
    );

CREATE TABLE
    IF NOT EXISTS likes (
        userID INTEGER,
        postID INTEGER,
        value VARCHAR(2),
        PRIMARY KEY (userID, postID),
        FOREIGN KEY (userID) REFERENCES users (id),
        FOREIGN KEY (postID) REFERENCES posts (id)
    );

CREATE TABLE
    IF NOT EXISTS categories_post (
        categoryID INTEGER,
        postID INTEGER,
        PRIMARY KEY (categoryID, postID),
        FOREIGN KEY (categoryID) REFERENCES categories (id),
        FOREIGN KEY (postID) REFERENCES posts (id)
    );



CREATE TABLE IF NOT EXISTS commentsLikes (
    userID INTEGER,
    commentID INTEGER,
    value VARCHAR(2),
    PRIMARY KEY (userID, commentID),
    FOREIGN KEY (userID) REFERENCES users (id),
    FOREIGN KEY (commentID) REFERENCES comments (id)
);
