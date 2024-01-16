CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    IF NOT EXISTS Boards (
        id UUID PRIMARY KEY,
        title VARCHAR(50) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    );

CREATE TABLE
    IF NOT EXISTS Sections (
        id UUID PRIMARY KEY,
        title VARCHAR(50) NOT NULL,
        board_id UUID NOT NULL,
        FOREIGN KEY (board_id) REFERENCES Boards (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    );

CREATE TABLE
    IF NOT EXISTS Tasks (
        id UUID PRIMARY KEY,
        title VARCHAR(50) NOT NULL,
        section_id UUID NOT NULL,
        FOREIGN KEY (section_id) REFERENCES Sections (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    );

CREATE TABLE
    IF NOT EXISTS Sub_tasks (
        id UUID PRIMARY KEY,
        title VARCHAR(50) NOT NULL,
        task_id UUID NOT NULL,
        FOREIGN KEY (task_id) REFERENCES Tasks (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    );

CREATE TABLE
    IF NOT EXISTS Tags (
        id UUID PRIMARY KEY,
        title VARCHAR(50) NOT NULL,
        board_id UUID NOT NULL,
        FOREIGN KEY (board_id) REFERENCES Boards (id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    );

CREATE TABLE
    IF NOT EXISTS Tag_Tasks (
        tag_id UUID NOT NULL,
        task_id UUID NOT NULL,
        FOREIGN KEY (tag_id) REFERENCES Tags (id),
        FOREIGN KEY (task_id) REFERENCES Tasks (id)
    );