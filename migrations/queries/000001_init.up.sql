CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    number CHAR(6)
);

create domain email_type as varchar(50) check(value ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+$'); /*postgres only*/

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email email_type UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    class_id INT,
    CONSTRAINT fk_class FOREIGN key (class_id) REFERENCES classes(id) ON DELETE CASCADE
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email email_type UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN DEFAULT false
);

CREATE TABLE problems (
    id SERIAL PRIMARY KEY,
    title VARCHAR(70) not null,
    description TEXT not null,
    course_id int not null,
    constraint fk_course foreign key (course_id) references courses(id) on delete cascade 
);


CREATE TABLE teachers_classes (
    teacher_id INT NOT NULL,
    class_id INT NOT NULL,
    PRIMARY KEY (teacher_id, class_id),
    CONSTRAINT fk_teacher Foreign Key (teacher_id) REFERENCES teachers(id) on delete CASCADE,
    constraint fk_class FOREIGN KEY (class_id) REFERENCES classes(id) on delete CASCADE
);



create table classes_courses (
    class_id int not null,
    course_id int not null,
    primary key (class_id, course_id),
    constraint fk_class foreign key (class_id) references classes(id) on delete cascade,
    constraint fk_course foreign key (course_id) references courses(id) on delete cascade
);


create table attempts (
    student_id int not null,
    problem_id int not null,
    primary key (student_id, problem_id),
    amount int not null default 0,
    constraint amount_gt check (amount > 0),
    done boolean default false,
    constraint fk_student foreign key (student_id) references students(id) on delete cascade,
    constraint fk_problem foreign key (problem_id) references problems(id) on delete cascade
);
