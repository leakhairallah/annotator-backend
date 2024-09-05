CREATE TABLE IF NOT EXISTS annotations (
     id INT AUTO_INCREMENT PRIMARY KEY,
     text TEXT NOT NULL,
     metadata JSON
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;