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

