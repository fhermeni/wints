package datastore

const (
	create = `
drop table if exists reports;
drop table if exists surveys;
drop table if exists internships;
drop table if exists conventions;
drop table if exists sessions;
drop table if exists password_renewal;
drop table if exists defenses;
drop table if exists users;
    create table users(email text PRIMARY KEY,
				  firstname text,
				  lastname text,
				  tel text,
				  password text,
				  role integer
				  );

create table sessions(email text REFERENCES users(email) on delete cascade,
		      token text,
		      last timestamp with time zone,
		      constraint pk_uid PRIMARY KEY(email)
		      );

create table internships(student text PRIMARY KEY REFERENCES users(email) on delete cascade,
                        startTime timestamp with time zone,
                        endTime timeStamp with time zone,
                        tutor text REFERENCES users(email), --prevent to delete the tutor if he is tutoring someone
                        promotion text,
                        major text,
                        company text,
                        companyWWW text,
                        supervisorFn text,
                        supervisorLn text,
                        supervisorEmail text,
                        supervisorTel text,
                        title text,
                        nextPosition integer,
                        nextContact text,
                        researchLab boolean,
                        foreignCountry boolean,
                        gratification real
);

create table conventions(studentEmail text PRIMARY KEY,
                        studentFn text,
                        studentLn text,
                        studentTel text,
                        promotion text,
                        startTime timestamp with time zone,
                        endTime timeStamp with time zone,
                        tutorFn text,
                        tutorLn text,
                        tutorEmail text,
                        tutorTel text,                        
                        company text,
                        companyWWW text,
                        supervisorFn text,
                        supervisorLn text,
                        supervisorEmail text,
                        supervisorTel text,
                        title text,
                        researchLab boolean,
                        foreignCountry boolean,
                        gratification real,
                        skip boolean
);

create table reports(student text REFERENCES users(email) on delete cascade,
                        kind text,
                        deadline timestamp with time zone,
                        grade integer,
                        comment text,
                        private boolean,
                        cnt bytea,
                        toGrade bool,
                        constraint pk_reports PRIMARY KEY(student, kind)
                        );

create table surveys(student text REFERENCES users(email) on delete cascade,
                        kind text,
                        deadline timestamp with time zone,
                        timestamp timestamp with time zone,
                        answers json,
                        token text UNIQUE,
                        constraint pk_surveys PRIMARY KEY(student, kind)
                        );

create table defenses(id text,
                      content text);

create table password_renewal(
    email text PRIMARY KEY REFERENCES users(email) on delete cascade,
    deadline timestamp,
    token text UNIQUE
)`
)
