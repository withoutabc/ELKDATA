USE movie;

CREATE TABLE movie_ratings (

     movie_ratings_id INT AUTO_INCREMENT PRIMARY KEY,
      user_id INT,
       movie_id INT,
       rating DECIMAL(3, 2),
      timestamp TIMESTAMP
);


INSERT INTO movie_ratings (user_id, movie_id, rating, timestamp)
VALUES
    (1, 101, 4.5, '2023-09-01 19:15:00'),
    (1, 102, 3.0, '2023-09-01 20:30:00'),
    (2, 101, 4.0, '2023-09-01 21:45:00'),
    (2, 102, 3.5, '2023-09-02 10:20:00'),
    (3, 103, 5.0, '2023-09-02 14:55:00'),
    (3, 104, 4.0, '2023-09-02 16:10:00'),
    (4, 101, 4.0, '2023-09-03 08:30:00'),
    (4, 103, 4.5, '2023-09-03 09:45:00'),
    (5, 102, 3.5, '2023-09-03 12:15:00'),
    (5, 104, 4.0, '2023-09-03 13:30:00');