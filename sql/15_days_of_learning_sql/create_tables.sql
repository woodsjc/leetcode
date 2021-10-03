CREATE DATABASE test;
CREATE SCHEMA test;

USE test.test;

CREATE TABLE hackers (
    hacker_id int primary key identity(1,1)
    , name varchar(255)
);

CREATE TABLE submissions (
    submission_date date
    , submission_id int primary key identity(1,1)
    , hacker_id int 
    , score int
    , foreign key (hacker_id) references hackers(hacker_id)
    	ON DELETE SET NULL
    	ON UPDATE CASCADE
);

INSERT INTO hackers (name) VALUES 
    ('Rose')
    , ('Angela')
    , ('Frank')
    , ('Patrick')
    , ('Lisa')
    , ('Kimberly')
    , ('Bonnie')
    , ('Michael')
;

WITH vals (submission_date, name, score) as (
    SELECT '2016-03-01', 'Angela', 0
    UNION
    SELECT '2016-03-01', 'Kimberly', 15
    UNION
    SELECT '2016-03-01', 'Michael', 60
    UNION
    SELECT '2016-03-01', 'Frank', 70
    UNION
    SELECT '2016-03-02', 'Angela', 0
    UNION
    SELECT '2016-03-02', 'Rose', 60
    UNION
    SELECT '2016-03-02', 'Michael', 25
    UNION
    SELECT '2016-03-02', 'Michael', 0
    UNION
    SELECT '2016-03-03', 'Angela', 0
    UNION
    SELECT '2016-03-03', 'Frank', 70
)
INSERT INTO submissions (submission_date, hacker_id, score)
SELECT 
    v.submission_date
    , h.hacker_id
    , v.score
FROM vals v
JOIN hackers h
    ON v.name = h.name
;
