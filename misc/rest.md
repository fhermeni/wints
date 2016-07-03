
//user API 7 to 5
ed.get("/users/", users)
ed.post("/users/", newUser)
ed.post("/users/:u/email", setEmail)
ed.post("/users/:u/person", setUserPerson)
ed.post("/users/:u/role", setUserRole)
ed.del("/users/:u", delUser)
ed.get("/users/:u", user)

GET /users
POST /users

GET /users/:u
DEL /users/:u
PUT /users/:u
	(email, person, role)

//students 7 to 4
ed.post("/students/:s/major", setMajor)
ed.post("/students/:s/promotion", setPromotion)
ed.post("/students/:s/male", setMale)
ed.post("/students/:s/alumni", setAlumni)
ed.post("/students/:s/skip", setStudentSkippable)
ed.get("/students/", students)
ed.post("/students/", newStudent)

GET /students
POST /students
GET /students/:s
PUT /students/:s
	-> major, promotion, male, alumni, skip

//internships
GET /internships
GET /internships/:u
POST /internships
	-> update begin, end, company, tutor, supervisor