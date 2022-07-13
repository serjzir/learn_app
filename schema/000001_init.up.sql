-- пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL
);

-- задачи
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY, 
    opened BIGINT NOT NULL DEFAULT extract(epoch from now()), -- время создания задач
    closed BIGINT DEFAULT 0, 
    author_id INTEGER REFERENCES users(id) DEFAULT 0, -- автор задачи
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0, -- ответственный
    title VARCHAR(258), -- название задачи
    content TEXT  -- задачи
);

-- метки задач
CREATE TABLE labels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);

-- связь многие - ко- многим между задачами и метками
CREATE TABLE tasks_labels (
    task_id INTEGER REFERENCES tasks(id)  ON DELETE CASCADE,
    label_id INTEGER REFERENCES labels(id) 
);
-- добавить пользователя
INSERT INTO users (id, name) VALUES (0, 'default');
-- добавить задач 
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content) VALUES(12, 3, 0,0, 'Тестовая задача1', 'Тестовый текст тестовой задачи1');
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content) VALUES(4, 1, 0,0, 'Тестовая задача2', 'Тестовый текст тестовой задачи2');
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content) VALUES(1, 4, 0,0, 'Тестовая задача3', 'Тестовый текст тестовой задачи3');
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content) VALUES(4, 2, 0,0, 'Тестовая задача4', 'Тестовый текст тестовой задачи4');
-- добавить метки
INSERT INTO labels (name) VALUES('Тестовая1');
INSERT INTO labels (name) VALUES('Тестовая2');
INSERT INTO labels (name) VALUES('Тестовая3');
INSERT INTO labels (name) VALUES('Тестовая4');
-- добавить связи в таблицу
INSERT INTO tasks_labels (task_id, label_id) VALUES(1,1);
INSERT INTO tasks_labels (task_id, label_id) VALUES(2,2);
INSERT INTO tasks_labels (task_id, label_id) VALUES(3,3);
INSERT INTO tasks_labels (task_id, label_id) VALUES(4,4);