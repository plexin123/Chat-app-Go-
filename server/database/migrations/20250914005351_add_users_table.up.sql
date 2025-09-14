CREATE TABLE House (
    house_id SERIAL PRIMARY KEY,
    house_name VARCHAR(50) NOT NULL,
    address VARCHAR(255) NOT NULL,
);


CREATE TABLE "Users"(
    "id" BIGSERIAL PRIMARY KEY,
    "username" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
)


CREATE TABLE Message (
    id SERIAL PRIMARY KEY ,
    sender VARCHAR(50) NOT NULL,
    receiver VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- FOREIGN KEY (user_id) REFERENCES User(user_id)
);

