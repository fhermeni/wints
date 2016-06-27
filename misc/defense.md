# Backend

defense_session(#id, room, date)
For each student:
 datetime, room, private/public, remote/local

For each jury member:
 datetime, location


slots
unique(#(date, am|pm), room)


 #defense
 #student, (datetime, location), private/public, remote/local

 #juries
 teacher, location,



 # Frontend:
 visu permettant de voir a fond
Par etudiant:
-> publique/privée
-> local/remote

Par groupe:
-> jury
-> salle
-> date

DB:

defense_sessions(#id, room, date)
defense_juries(#id, email)

