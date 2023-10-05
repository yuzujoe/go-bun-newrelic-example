DROP TABLE IF EXISTS employees;

CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    department VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO employees (name, age, department) VALUES
    ('Alice', 30, 'Engineering'),
    ('Bob', 35, 'Marketing'),
    ('Charlie', 25, 'Sales'),
    ('David', 45, 'HR'),
    ('Eve', 40, 'Finance');
