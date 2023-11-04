-- Create the habits table
CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    goal INT NOT NULL
);

-- Create the habit_entries table
CREATE TABLE IF NOT EXISTS habit_entries (
    id SERIAL PRIMARY KEY,
    habit_id INT NOT NULL REFERENCES habits(id),
    date DATE NOT NULL,
    status bool NOT NULL
);

-- Insert some data into the habits table
INSERT INTO habits (name, description, goal)
VALUES ('Exercise', 'Exercise for at least 30 minutes 3 days a week.', 3),
('Eat healthy', 'Eat at least 5 servings of fruits and vegetables per day.', 5),
('Read', 'Read for at least 30 minutes per day.', 30),
('Meditate', 'Meditate for at least 10 minutes per day.', 10);

-- Insert some data into the habit_entries table
INSERT INTO habit_entries (habit_id, date, status)
VALUES (1, '2023-11-05', TRUE),
(1, '2023-11-06', FALSE),
(1, '2023-11-07', TRUE),
(1, '2023-11-08', TRUE);

INSERT INTO habit_entries (habit_id, date, status)
VALUES (2, '2023-11-05', TRUE),
(2, '2023-11-06', TRUE),
(2, '2023-11-07', FALSE),
(2, '2023-11-08', TRUE);

INSERT INTO habit_entries (habit_id, date, status)
VALUES (3, '2023-11-05', TRUE),
(3, '2023-11-06', TRUE),
(3, '2023-11-07', TRUE),
(3, '2023-11-08', FALSE);

INSERT INTO habit_entries (habit_id, date, status)
VALUES (4, '2023-11-05', TRUE),
(4, '2023-11-06', FALSE),
(4, '2023-11-07', TRUE),
(4, '2023-11-08', TRUE);