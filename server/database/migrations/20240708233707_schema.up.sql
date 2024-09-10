CREATE TABLE House (
    house_id INT AUTO_INCREMENT PRIMARY KEY,
    house_name VARCHAR(50) NOT NULL,
    address VARCHAR(255) NOT NULL
);


CREATE TABLE User(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
)


CREATE TABLE Message (
    id INT AUTOINCREMENT PRIMARY KEY ,
    sender VARCHAR(50) NOT NULL,
    reciever VARCHAR(50) NOT NULL
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- FOREIGN KEY (user_id) REFERENCES User(user_id)
);
