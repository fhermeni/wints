-- reset if needed then create the tables required to store internship stuff
drop table if exists pending;
drop table if exists reports;
drop table if exists internships;
drop table if exists conventions;
drop table if exists sessions;
drop table if exists password_renewal;
drop table if exists defenses;
drop table if exists surveys;
drop table if exists users;

-- user
create table users(email text PRIMARY KEY,
				  firstname text,
				  lastname text,
				  tel text,
				  password text,
				  role integer
				  );

-- sessions
create table sessions(email text REFERENCES users(email) on delete cascade,
		      token text,
		      last timestamp,
		      constraint pk_uid PRIMARY KEY(email)
		      );

create table internships(student text PRIMARY KEY REFERENCES users(email) on delete cascade,
                        startTime timestamp,
                        endTime timeStamp,
                        tutor text REFERENCES users(email), --prevent to delete the tutor if he is tutoring someone
                        company text,
                        companyWWW text,
                        supervisorFn text,
                        supervisorLn text,
                        supervisorEmail text,
                        supervisorTel text,
                        title text
);

create table conventions(student text PRIMARY KEY REFERENCES users(email) on delete cascade,
                        startTime timestamp,
                        endTime timeStamp,
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
                        title text
);

create table reports(student text REFERENCES users(email) on delete cascade,
                        kind text,
                        deadline timestamp,
                        grade integer,
                        cnt bytea,
                        constraint pk_reports PRIMARY KEY(student, kind)
                        );

create table password_renewal(
    email text PRIMARY KEY REFERENCES users(email) on delete cascade,
    deadline timestamp,
    token text UNIQUE
)

create table pending(
    firstname text,
    lastname text,
    email text PRIMARY KEY,
    promotion text,
    major text,
    internship text
)

create table defenseSessions(
    date timestamp with time zone,
    room text,
    pause integer,
    constraint pk_unique PRIMARY KEY(date, room)
);

create table defenseJuries(
    date timestamp with time zone,
    room text,
    jury text REFERENCES users(email) on delete cascade,
    constraint pk_jury PRIMARY KEY(date, room, jury),
    constraint fk_session FOREIGN KEY(date, room) REFERENCES defenseSessions(date, room) on delete cascade    
);

create table defenses(
    date timeStamp with time zone,
    room text,
    student text PRIMARY KEY REFERENCES internships(student),
    grade integer,
    private bool,
    remote bool,
    constraint fk_session FOREIGN KEY(date, room) REFERENCES defenseSessions(date, room) on delete cascade
);