package sqlstore

const (
	create = `drop table if exists users cascade;
drop table if exists sessions cascade;
drop table if exists password_renewal cascade;
drop table if exists students cascade;
drop table if exists conventions cascade;
drop table if exists reports cascade;
drop table if exists surveys cascade;
drop table if exists defenseSessions cascade;
drop table if exists defenses cascade;
drop table if exists defenseJuries cascade;

create table users(
    email text,
    firstname text,
    lastname text,
    tel text,
    password text,
    role integer,
    lastVisit timestamp with time zone,
    constraint pk_email PRIMARY KEY(email)
);

create table sessions(
    email text,
    token text,
    expire timestamp with time zone,
    constraint pk_sessions_email PRIMARY KEY(email),
    constraint fk_sessions_email FOREIGN KEY(email) REFERENCES users(email) on delete cascade on update cascade
);


create table password_renewal(
    email text,    
    token text unique,
    constraint pk_password_renewal_email PRIMARY KEY(email),
    constraint fk_password_renewal_email FOREIGN KEY(email) REFERENCES users(email) on delete cascade on update cascade
);

create table students(
    email text,
    major text,
    promotion text,
    nextPosition text,
    nextFrance boolean,
    nextSameCompany boolean,
    nextPermanent boolean,
    nextContact text,
    skip boolean,
    male boolean,
    constraint pk_students_email PRIMARY KEY(email),
    constraint fk_students_email FOREIGN KEY (email) REFERENCES users(email) on delete cascade on update cascade
);

create table conventions(
    student text,    
    startTime timestamp with time zone,
    endTime timestamp with time zone,
    tutor text,
    companyName text,
    companyWWW text,
    supervisorFn text,
    supervisorLn text,
    supervisorEmail text,
    supervisorTel text,
    title text,
    creation timestamp with time zone,
    foreignCountry boolean,
    lab bool,
    gratification real,
    skip bool,
    valid bool,
    constraint pk_conventions_student PRIMARY KEY (student),
    constraint fk_conventions_student FOREIGN KEY (student) REFERENCES students(email) on delete cascade on update cascade,
    constraint fk_conventions_tutor FOREIGN KEY (tutor) REFERENCES users(email) on update cascade
     -- no cascade delete for fk_conventions_tutor because we don't want to loose convention when we remove a duplicated tutor account
);

create table reports(
    student text,
    kind text,
    deadline timestamp with time zone,
    delivery timestamp with time zone,
    reviewed timestamp with time zone,
    grade integer,
    comment text,
    private boolean,
    cnt bytea,
    toGrade bool,                        
    constraint pk_reports_student PRIMARY KEY(student, kind),
    constraint fk_reports_student FOREIGN KEY(student) REFERENCES students(email) on delete cascade on update cascade
);

create table surveys(
    student text,
    kind text,
    deadline timestamp with time zone,
    delivery timestamp with time zone,
    cnt bytea,
    token text UNIQUE,
    constraint pk_surveys_student PRIMARY KEY(student, kind),
    constraint fk_surveys_student FOREIGN KEY(student) REFERENCES students(email) on delete cascade on update cascade
);

create table defenseSessions(
    date timestamp with time zone,
    room text,    
    constraint pk_defenseSessions PRIMARY KEY(date, room)
);

create table defenseJuries(
    date timestamp with time zone,
    room text,
    jury text,
    constraint fk_defenseJuries FOREIGN KEY(jury) REFERENCES users(email) on delete cascade on update cascade,
    constraint pk_defenseJuries PRIMARY KEY(date, room, jury),
    constraint fk_defenseJuries_session FOREIGN KEY(date, room) REFERENCES defenseSessions(date, room) on delete cascade    
);

create table defenses(
    date timeStamp with time zone,
    room text,
    student text,
    grade integer,
    private bool,
    local bool,
    constraint pk_defenses_student PRIMARY KEY(student),
    constraint fk_defenses_student FOREIGN KEY(student) REFERENCES students(email) on delete cascade on update cascade,
    constraint fk_defenses_session FOREIGN KEY(date, room) REFERENCES defenseSessions(date, room) on delete cascade on update cascade
);`
)
