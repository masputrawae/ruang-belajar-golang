DROP DATABASE IF EXISTS belajar;
CREATE DATABASE belajar;

USE belajar;

CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(300) NOT NULL
) ENGINE=InnoDB;

DESC users;

CREATE TABLE statuses(
  id VARCHAR(15) PRIMARY KEY UNIQUE,
  name VARCHAR(15) NOT NULL
) ENGINE=InnoDB;

CREATE TABLE priorities(
  id VARCHAR(15) PRIMARY KEY UNIQUE,
  name VARCHAR(15) NOT NULL
) ENGINE=InnoDB;

-- CONSTANT META (can not update or delete)
INSERT INTO
  statuses(id, name)
VALUES
  ("PLANNING", "Planning"),
  ("ACTIVE", "Active"),
  ("IN_PROGRESS", "In Progress"),
  ("DONE", "Done"),
  ("CANCELLED", "Cancelled"),
  ("ARCHIVE", "Archive");

INSERT INTO
  priorities(id, name)
VALUES
  ("HIGHEST", "Highest"),
  ("HIGH", "High"),
  ("MEDIUM", "Medium"),
  ("LOW", "Low"),
  ("LOWEST", "Lowest");

SELECT * FROM statuses;
SELECT * FROM priorities;

CREATE TABLE todos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  task TEXT NOT NULL,
  status_id VARCHAR(15) DEFAULT 'PLANNING',
  priority_id VARCHAR(15) DEFAULT 'MEDIUM',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  
  FOREIGN KEY (status_id) REFERENCES statuses(id)
    ON DELETE RESTRICT,
  
  FOREIGN KEY (priority_id) REFERENCES priorities(id)
    ON DELETE RESTRICT
) ENGINE=InnoDB;