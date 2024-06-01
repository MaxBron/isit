

CREATE TABLE faculties
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id INTEGER,
	Name Text PRIMARY KEY,
	
	Catalog TEXT
);

CREATE TABLE Student_groups
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id INTEGER,
	Number INTEGER PRIMARY KEY,
	faculty_name Text,
	FOREIGN KEY (faculty_name) REFERENCES Faculties(Name)
	
);

CREATE TABLE Students
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	Id INTEGER PRIMARY KEY,
	first_name CHARACTER VARYING(30),
    last_name CHARACTER VARYING(30),
	sur_name CHARACTER VARYING(30),
	course_number INTEGER,
	Location TEXT,
	group_number INTEGER,
	FOREIGN KEY (group_number) REFERENCES Student_groups(Number)
);

CREATE TABLE dormitories
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id INTEGER,
	room_number INTEGER PRIMARY KEY,
	student_id INTEGER,
	FOREIGN KEY (student_id) REFERENCES Students(Id)
);

CREATE TABLE Scholarships
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id INTEGER,
	Type TEXT,
	Student_id INTEGER,
	Amount INTEGER,
	FOREIGN KEY (Student_Id) REFERENCES Students(Id),
	PRIMARY KEY (Type, Student_Id)
);

CREATE TABLE users
(
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	id INTEGER,
	login TEXT PRIMARY KEY,
	password TEXT,
	role TEXT
);

INSERT INTO USERS (login, password, role) VALUES ('qwerty', 'qwerty', 'admin');
INSERT INTO USERS (login, password, role) VALUES ('aaa', 'aaa', 'user');