drop table if exists finalreport;
drop table if exists midtermreport;
drop table if exists defenses;
drop table if exists reports;
drop table if exists internships;
drop table if exists pending_internships;
drop table if exists interviews;
drop table if exists applications;
drop table if exists students;
drop table if exists sessions;
drop table if exists logs;
drop type if exists app_status;
drop table if exists password_renewal;

drop table if exists users;
-- user
create table users(email text PRIMARY KEY,
				  firstname text,
				  lastname text,
				  tel text,
				  password text,
				  role text
				  );

-- students
create table students(email text PRIMARY KEY REFERENCES users(email),
					  major text,
					  promotion text,
					  last_action timestamp
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
                        supervisorTel text,
                        title text
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
                        supervisorTel text,
                        title text
);

create table reports(student text REFERENCES students(email),
                        kind text,
                        deadline timestamp,
                        grade integer,
                         cnt bytea,
                         constraint pk_reports PRIMARY KEY(student, kind)
                        );

create table defenses(id text,
                      content text);

create table password_renewal(
    email text PRIMARY KEY REFERENCES users(email) on delete cascade,
    deadline timestamp,
    token text UNIQUE
)