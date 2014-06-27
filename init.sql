drop table if exists internships;
drop table if exists pending_internships;
drop table if exists interviews;
drop table if exists applications;
drop table if exists students;
drop table if exists sessions;
drop table if exists roles;
drop table if exists logs;
drop type if exists app_status;
drop table if exists users;
drop type if exists role;
-- user
create table users(email text UNIQUE ON UPDATE UPDATE,
				  username text,
				  firstname text,
				  lastname text,
				  tel text,
				  password text
				  );

-- students
create table students(email text PRIMARY KEY REFERENCES users(email),
					  major text,
					  promotion text,
					  last_action timestamp
);

-- major_adm
create type role as enum('student', 'tutor', 'admin', 'root', 'al','ihm','vim','ubinet','kis','cssr','imafa');
create table roles(email text,
					   role text,
					   constraint pk_major_adm PRIMARY KEY (email, role)
);

-- applications
create type app_status as enum('rejected', 'open', 'granted', 'validating', 'validated');
create table applications(aid serial  PRIMARY KEY,
						  date timestamp,
						  email text REFERENCES students(email),
						  company text,
						  note text,
						  status app_status);

-- interviews
create table interviews(aid integer REFERENCES applications(aid),
						date date,
						constraint pk_interviews PRIMARY KEY(aid, date)
);

--logs
create table logs(email text,
				  date timestamp,
				  message text);

-- sessions
create table sessions(email text REFERENCES users(email),
		      token text,
		      last timestamp,
		      constraint pk_uid PRIMARY KEY(email)
		      );

create table internships(student text PRIMARY KEY REFERENCES students(email),
                        startTime timestamp,
                        endTime timeStamp,
                        tutor text REFERENCES users(email),
                        midTermDeadline timestamp,
                        company text,
                        companyWWW text,
                        supervisorFn text,
                        supervisorLn text,
                        supervisorEmail text,
                        supervisorTel text
);

create table pending_internships(student text PRIMARY KEY REFERENCES students(email),
                        startTime timestamp,
                        endTime timeStamp,
                        tutorFn text,
                        tutorLn text,
                        tutorEmail text,
                        tutorTel text,
                        midTermDeadline timestamp,
                        company text,
                        companyWWW text,
                        supervisorFn text,
                        supervisorLn text,
                        supervisorEmail text,
                        supervisorTel text
)