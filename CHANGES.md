## Version 1.0
- authentication OK
- listings OK
- mailto: OK
- update profile OK
- tutor alignment OK
- major alignment OK
- permissions: -(students, tutor), major, admin, root   OK
  major: see all conventions + manipulate major OK
  admin: pending conventions OK
  root: users permission OK
- new user OK
- delete user OK
- init mode OK
- rest API in "rest/v1/"  OK
- confirmation for sensible data
   -> user deletion OK
   -> alignment with exact matching OK
- change password OK
- shift click support for mailing OK
- tutor can adapt the report deadline OK
- hide static pages, stay inside "/" OK
- defenses preparation support OK
- checkbox to mail supervisor/tutor in the first two tabs
   - watchlist OK
- upload/download reports OK
- grade reports
- watchlist: mode general, mode grades OK, no need in practice

- verify email format when aligning and changing
- sort by date correctly
- logging
- error reporting
- encoding when pulling conventions
- les fiches d'évaluations finales devraient être accessibles au jury naturellement
- UI reports, encoding
-

defenses:
 - coarse grain view for the student schedule
 - detailed view for
   - jury, room, date
   - confidential defenses
   - mail checkbox + button
   - download all fiche d'évaluation
 - upload fiches d'évaluation final depuis watchlist

defenses(_id_, cnt)
  - new,
  - save,
  - delete,
  - save as


midtermReport(_id_, cnt)

finalReport(_id, cnt)



## Version 2.0 ##

- tutor view
- student view

## Version 3.0 ##
- midterm evaluation

- submit grade
- submit comment
- notify student by mail

- company survey && student|advisor notification

- another page with 2 tables: one with received grades, one with delayed, one with no hurry ?
