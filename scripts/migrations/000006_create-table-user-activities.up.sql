CREATE TABLE IF NOT EXISTS user_activities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    post_id INT NOT NULL,
    is_liked BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by LONGTEXT NOT NULL,
    created_by LONGTEXT NOT NULL,
    CONSTRAINT fk_user_id_user_activities FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_post_id_user_activities FOREIGN KEY (post_id) REFERENCES posts(id)
);