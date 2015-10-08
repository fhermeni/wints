var allStudents, allConventions, allTeachers;

function showConventionValidator() {
	$.when(students(), internships(), conventions(), users()).done(loadConventionValidator).fail(logFail)
}

function studentSort(a, b) {
	return a.User.Person.Lastname.toLowerCase().localeCompare(b.User.Person.Lastname.toLowerCase());
}

function userSort(a, b) {
	return a.Person.Lastname.toLowerCase().localeCompare(b.Person.Lastname.toLowerCase());
}


function convSort(a, b) {
	return a.Student.User.Person.Lastname.toLowerCase().localeCompare(b.Student.User.Person.Lastname.toLowerCase());
}

function loadConventionValidator(students, internships, convs, us) {
	allStudents = students[0];
	us = us[0];
	allStudents.sort(studentSort);
	internships = internships[0];
	allConventions = convs[0];
	allConventions.sort(convSort);

	allTeachers = us.filter(function(u) { 
		return u.Role > 1
	});
	allTeachers.sort(userSort);
	internships.forEach(function(i) {
		allStudents.forEach(function(s, idx, arr) {
			if (s.User.Person.Email == i.Convention.Student.User.Person.Email) {
				allStudents[idx].Placed = true;
				allStudents[idx].I = i;
			}
		});
	});
	$("#cnt").render("placement-header", {
		Students: allStudents,
		Ints: internships.length
	}, ui);
}

function conventionValidator(em) {
	data = {
		Conventions: allConventions,
		C: allConventions[0],
		Teachers: allTeachers
	};

	allStudents.forEach(function(s) {
		if (s.User.Person.Email == em) {
			data.S = s;
			return false;
		}
	});

	var u = propose(data.S.User, allConventions.map(function(c) {
		return c.Student.User;
	}));
	if (u) {
		data.C = allConventions.filter(function(c) {
			return c.Student.User.Person.Email == u.Person.Email;
		})[0];
	}
	var t = propose(data.C.Tutor, allTeachers);

	$("#modal").render("convention-validator", data, function() {
		$("#convention-selecter").val(data.C.Student.User.Person.Email);
		if (t) {
			$("#tutor-selecter").val(t.Person.Email);
		}
		showModal(function() {
			checkConventionAlignment();
			checkTutorAlignment();
		});
	});
}

function showInternshipEditor(i) {
	stus = allStudents.map(function(s) {
		return s.User;
	});
	$("#internship-editor").render("convention-editor", {
		I: i,
		Teachers: allTeachers,
		Students: stus
	}, function() {
		propose(i.Convention.Student.User.Person, stus, $("#student-group"));
		propose(i.Convention.Tutor.Person, allTeachers, $("#tutor-group"));
	});
}

function showConvention(cnt) {
	checkConventionAlignment();
	var em = $(cnt).val();
	var dta = {
		Teachers: allTeachers
	};
	allConventions.forEach(function(c) {
		if (c.Student.User.Person.Email == em) {
			dta.C = c;
			$("#convention-detail").html(Handlebars.partials['convention-editor'](dta));
			cleanError("#convention-selecter", "#tutor-selecter");
			$('#convention-detail').find('[data-toggle="confirmation"]').confirmation('destroy');
			propose(c.Tutor, allTeachers, $("#tutor-group"));
			return false;
		}
	});

}

function updateStudentSkipable(stu, btn) {
	allStudents.forEach(function(s, idx) {
		if (s.User.Person.Email == stu) {
			s.Skip = !s.Skip;
			allStudents[idx] = s;
			postStudentSkippable(stu, s.Skip).done(function() {
				var row = $("#table-placement").find("tr[data-email='" + stu + "']");
				if (s.Skip) {
					$(btn).removeClass("glyphicon-eye-open text-success").addClass("glyphicon-eye-close text-warning");
					row.addClass("skip");
				} else {
					$(btn).removeClass("glyphicon-eye-close text-warning").addClass("glyphicon-eye-open text-success");
					row.removeClass("skip");
				}

			}).fail(logFail);
			return false;
		}
	});
}

function validateConvention(stu) {
	if (checkConventionAlignment() && checkTutorAlignment()) {
		prepareValidation();
	} else {
		$('#modal').find('[data-toggle="confirmation"]').confirmation('show');
	}
}

function checkConventionAlignment() {
	var em = $("#convention-selecter").val(); //the one in the convention
	var th = $(".modal-body").find("legend").data("email");
	var student;
	var th_student;

	allConventions.forEach(function(c) {
		if (c.Student.User.Person.Email == em) {
			student = c.Student;
			return false;
		}
	});
	allStudents.forEach(function(s) {
		if (s.User.Person.Email == th) {
			th_student = s;
			return false;
		}
	});
	//Check if the students match
	if (!matching(student.User, th_student.User)) {
		reportError("#convention-selecter", "The student does not match")
		return false;
	}
	cleanError("#convention-selecter");
	return true;
}

function checkTutorAlignment() {
	var stu = $("#convention-selecter").val(); //the one in the convention
	var em = $("#tutor-selecter").val(); //the one in the DB	
	var tutor; //the one in the convention
	allConventions.forEach(function(c) {
		if (c.Student.User.Person.Email == stu) {
			tutor = c.Tutor;
			return false;
		}
	});

	e = exists(tutor, allTeachers);
	if (em == "_new_") {
		if (e) {
			reportError("#tutor-selecter", "A registered teacher exists");
			return false;
		} else {
			cleanError("#tutor-selecter");
		}
	} else {
		if ((e && e.Email != em) || !e) {
			reportError("#tutor-selecter", "Teachers don't match");
			return false;
		} else {
			cleanError("#tutor-selecter");
		}
	}
	return true;
}

function prepareValidation() {
	cleanError("#convention-selecter", "#tutor-selecter");

	//si tuteur new, on crée le compte
	var tut = $("#tutor-selecter").val();
	var stu = $("#convention-selecter").val();
	var c = allConventions.filter(function(c) {
		return c.Student.User.Person.Email == stu
	})[0];

	//We update the student profile
	var knownEmail = $("#modal").find("legend").data("email");
	p = c.Student.User.Person;
	var newEmail = p.Email;
	p.Email = newEmail;
	postUserEmail(knownEmail, newEmail).done(function() {
		sendProfile(p).done(function() {
			commitValidation(c, tut);
		}).fail(function(xhr) {
			reportError("#convention-selecter", xhr.responseText);
		});
	}).fail(function(xhr) {
		reportError("#convention-selecter", xhr.responseText);
	});
}

function commitValidation(c, tut) {
	if (tut == "_new_") {
		tut.Role = 2;
		postNewUser(c.Tutor).done(function() {
			allTeachers.push(c.Tutor);
			allTeachers.sort(userSort);
			newInternship(c).done(updateConventionTable).fail(function(xhr) {
				reportError("#convention-selecter", xhr.responseText);
			});
		}).fail(function(xhr) {
			reportError("#tutor-selecter", xhr.responseText);
		});
	} else {
		c.Tutor.Person.Email = tut;
		newInternship(c).done(updateConventionTable).fail(function(xhr) {
			reportError("#convention-selecter", xhr.responseText);
		});
	}
}

function updateConventionTable(i) {
	hideModal();
	//Catch the line, remove it by the new one
	var em = $("legend").data("email");
	var row = $("#table-placement").find("tr[data-email='" + em + "']");
	var dta = i.Convention.Student;
	dta.I = i;
	var cnt = Handlebars.partials['placement-student'](dta);
	row.replaceWith(cnt);
}

function matching(u1, u2) {
	if (u1.Person.Email == u2.Person.Email) {
		return true;
	}
	if (u1.Person.Lastname.indexOf(u2.Person.Lastname) > -1 || u2.Person.Lastname.indexOf(u1.Person.Lastname) > -1) {
		return true;
	}
	return false;
}

function exists(u, users) {
	var res = undefined;
	users.forEach(function(t) {
		var known_ln = t.Person.Lastname;
		if (u.Person.Email == t.Person.Email) {
			res = t.Person;
			return false;
		}
		if (u.Person.Lastname.indexOf(known_ln) > -1 || known_ln.indexOf(u.Person.Lastname) > -1) {
			res = t.Person;
			return false;
		}
	});
	return res;
}

function propose(u, us) {
	var res = undefined;
	us.forEach(function(t) {
		var known_ln = t.Person.Lastname;
		if (u.Person.Email == t.Person.Email) {
			res = t;
			return false;
		} else if (u.Person.Lastname.indexOf(known_ln) > -1 || known_ln.indexOf(u.Person.Lastname) > -1) {
			res = t;
			return false;
		}
	});
	return res;
}