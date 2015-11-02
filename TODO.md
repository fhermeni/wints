# TODO

## Features

- Normalise relative deadline for reports: the next day at 2am wrt. the theoretical value
- automate surveys management:
	1. pre-made mails are send directly at a given deadline (mail include deadline)
	2. automatic reminder
	3. email notification to the academic tutor 
- iCal sync for each user:	
	- student get their deadlines
	- tutors get students deadline & review deadlines
	- no file. Sync with the server to stay aligned with the current status (new deadlines, report submitted, etc.)
- rewrite status page
- ease the daemon installation
	1. normalise parameters

## Development
- clean a bit the code
	1. Reduce/compose structures better
	2. Reduce the number of queries ?
	3. nicer rest middleware. Explicit context
- tests and comments
- simplify DB with less fake values and more optional fields hidden in JSONs

## Mailing
- when I act on behalf of tutor, cc tutor
- cc. emitter as well


- Convocation officielle par mail
- prévenir tuteurs
- pecho jury parmis les tuteurs
- organisation de visio
- saisie des notes



placement:
- bad names & validation ?
my student page


student page


reports
	- grade, penalties

cron

alumni (new)


defense


# login page
- remove error when logged in

# mailing
always have tutor in cc. when doing sth. about the student




## By page

### Student
   - notify company update OK
   - notify tutor when supervisor status changed OK
   	-> mail ?
   - notify future update OK

   - robust upload report
   - grade + penalty computation

   - Grading
   	<D -
   	>D 
   		R grade - lates
   	    !R
   	    	uploaded -> ?
   	    	!uploaded 
### Login
	RAS

### Reset password
	RAS

### Users
	RAS

### Service
	RAS

### Conventions

	support refresh, cache

	skip student (might have or not an internship)
		-> will be ignored by the mailing lists if he has an internship OK
		-> ignored in the placed % if it does not have an internship OK

	change tutor OK
		-> mail student/old tutor/new tutor OK
	change promo OK
	change major OK

## Tutored/Watchlist
	
	report workflow
	grade report
	late penalties

	report deadline: confirmation button to reduce spam


Roles: OK
Student
Tutor (teacher in fact)
MAJOR_x (see his major)
HEAD (see everything)
ADMIN (+ add student, teachers, validate convention, plan defenses)

### Caching

1 cache per user/request
POST invalidate all the caches
GET populate local cache


## Notes
rendu/pas rendu
en retard/ a l'heure
noté/pas noté

Sur la case= note net
onhover = note brut - penalty