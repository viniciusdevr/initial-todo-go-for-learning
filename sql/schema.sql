DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
	id INT AUTO_INCREMENT NOT NULL,
	title VARCHAR(128) NOT NULL,
	description VARCHAR(255) NOT NULL,
	done BOOLEAN NOT NULL DEFAULT FALSE,
	PRIMARY KEY (id)
	);

INSERT INTO tasks
  (title, description, done)
VALUES
  ('Dish washes', 'Clean all the dishes and utensils left in the sink after meals', false),
  ('Play games', 'Spend some leisure time playing video games or board games to relax', false),
  ('Studies', 'Review class notes and complete assignments for upcoming exams', false),
  ('Read my book', 'Continue reading the current chapter of the book on the nightstand', false);