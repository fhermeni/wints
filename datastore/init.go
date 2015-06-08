package datastore

const (
	create = `
drop table if exists defenses;
drop table if exists defenseJuries;
drop table if exists defenses;
drop table if exists pending;
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
                        male boolean,
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
                        creation timestamp with time zone,
                        foreignCountry boolean,
                        lab boolean,
                        gratification real 
);

create table conventions(studentEmail text PRIMARY KEY,
                        male boolean,
                        creation timestamp with time zone,
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
                        skip boolean,
                        foreignCountry boolean,
                        lab boolean,
                        gratification real 
);

create table reports(student text REFERENCES users(email) on delete cascade,
                        kind text,
                        deadline timestamp with time zone,
                        delivery timestamp with time zone,
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

create table defenses(
    date timeStamp with time zone,
    room text,
    student text PRIMARY KEY REFERENCES internships(student),
    grade integer,
    private bool,
    remote bool    
);

create table juries(
    student text REFERENCES defenses(student) on delete cascade,
    jury text REFERENCES users(email)    
);

create table password_renewal(
    email text PRIMARY KEY REFERENCES users(email) on delete cascade,
    deadline timestamp,
    token text UNIQUE
);

create table pending(
    firstname text,
    lastname text,
    email text PRIMARY KEY,
    promotion text,
    major text,
    internship text,
    hidden boolean,
);`
)
