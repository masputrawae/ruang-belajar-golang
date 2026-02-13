DROP DATABASE IF EXISTS belajar;

CREATE DATABASE belajar;

USE belajar;

CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(500) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB;

CREATE TABLE statuses (
  id VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(20) NOT NULL
) ENGINE=InnoDB;

CREATE TABLE priorities (
  id VARCHAR(20) NOT NULL UNIQUE,
  name VARCHAR(20) NOT NULL
) ENGINE=InnoDB;

-- META
INSERT INTO statuses (id, name)
VALUES ('PLANNING', 'Planning'), 
('ACTIVE', 'Active'), 
('IN_PROGRESS', 'In Progress'),
('DONE', 'Done'),
('ARCHIVE', 'Archive');

INSERT INTO priorities (id, name)
VALUES ('HIGHEST', 'Highest'), 
('HIGH', 'High'), 
('MEDIUM', 'Medium'),
('LOW', 'Low'),
('LOWEST', 'Lowest');

CREATE TABLE todos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  task TEXT NOT NULL,
  status_id VARCHAR(20) DEFAULT "PLANNING",
  priority_id VARCHAR(20) DEFAULT "MEDIUM",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  FOREIGN KEY (user_id) REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  
  FOREIGN KEY (status_id) REFERENCES statuses(id)
    ON DELETE RESTRICT,
  FOREIGN KEY (priority_id) REFERENCES priorities(id)
    ON DELETE RESTRICT,

  CONSTRAINT 
    chk_status 
  CHECK (
    status_id 
    IN 
    ('PLANNING', 'ACTIVE', 'IN_PROGRESS', 'DONE', 'ARCHIVE')
  ),

  CONSTRAINT 
    chk_priority 
  CHECK (
    priority_id 
    IN ('HIGHEST', 'HIGH', 'MEDIUM', 'LOW', 'LOWEST')
  )
) ENGINE=InnoDB;

SELECT * FROM users;
