var allStudents, allConventions, allTeachers;

var feederWarning = false;
function showConventionValidator() {
	showWait();
	if (level(myself.Role) >= ADMIN_LEVEL) {
		$.when(students(), internships(), conventions(), users()).done(loadConventionValidator).fail(function(xhr) {
			if (xhr.responseText.indexOf("feeder")) {
				//conventions() failed, so, an error message and we retry without this call
				//we don't notify error because it is done by the default fail				
				//feederWarning = xhr.responseText
				$.when(students(), internships()).done(loadConventionValidator).fail(notifyError);
			}
		});
	} else {
		$.when(students(), internships()).done(loadConventionValidator).fail(notifyError);
	}
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
	errors = {
		Warnings : []
	};
	allStudents = students[0];
	if (!allStudents) {
		allStudents = [];
	}	
	allTeachers = [];
	allConventions = [];
	var managed = 0;
	allStudents.forEach(function(s, idx) {
		allStudents[idx].Warn = !s.Skip;
		if (!s.Skip) {
			managed++;
		}
	});

	allStudents.sort(studentSort);
	internships = internships[0];
	if (!internships) {
		internships = [];
	}
	if (convs) {
		allConventions = convs[0].Conventions;
		errors = convs[0].Errors;
		if (!allConventions) {
			allConventions = [];
		}
		allConventions.sort(convSort);	

		//Remove the conventions that are already validated
		var checked = [];
		internships.forEach(function (i) {
			checked.push(i.Convention.Student.User.Person.Email);
		});
		allConventions = allConventions.filter(function (c) {					
			return !(checked.indexOf(c.Student.User.Person.Email) >= 0);
		});		
	}

	if (us) {
		us = us[0];		
		if (!us) {
			us = [];
		}
		allTeachers = us.filter(function(u) { 
			return level(u.Role) != STUDENT_LEVEL;
		});		
	}	
	allTeachers.sort(userSort);	
	internships.forEach(function(i) {
		allStudents.forEach(function(s, idx, arr) {
			if (s.User.Person.Email == i.Convention.Student.User.Person.Email) {
				allStudents[idx].Placed = true;
				allStudents[idx].I = i;
				allStudents[idx].Warn = !i && !allStudents[idx].Skip;
			}
		});
	});

	$("#cnt").render("placement-header", {
		Students: allStudents,
		Managed: managed,
		Ints: internships.length,	
		Errors: errors,
	}, ui);
}

function conventionValidator(em) {
	if (level(myself.Role) < ADMIN_LEVEL) {
		notifyError("Only an admin can validate a convention");
		return
	}


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
			cleanError("#tutor-selecter");
			$('#modal').find('[data-toggle="confirmation"]').confirmation();
			propose(c.Tutor, allTeachers, $("#tutor-group"));
			return false;
		}
	});

}

function updateStudentSkipable(stu, btn) {
	allStudents.forEach(function(s, idx) {
		if (s.User.Person.Email == stu) {
			postStudentSkippable(stu, !s.Skip).done(function(dta, status, xhr) {
				s.Skip = !s.Skip;
				var row = $("#table-placement").find("tr[data-email='" + stu + "']");
				s.I = undefined;
				
				internship(stu).done(function(i) {
					s = i.Convention.Student;
					s.I = i;
					s.Warn = false;
					allStudents[idx] = s;
					var cnt = Handlebars.partials['placement-student'](s);
					row.replaceWith(cnt);
					$('.tablesorter').trigger("update").trigger("updateCache");
					var v = parseInt($('#managed_cnt').html());
					$('#managed_cnt').html(s.Skip ? v - 1 : v + 1);					
				}).fail(function(xhr) {
					s.Warn = !s.Skip
					allStudents[idx] = s;
					var cnt = Handlebars.partials['placement-student'](s);
					row.replaceWith(cnt);
					$('.tablesorter').trigger("update").trigger("updateCache");
				});
				defaultSuccess({}, "OK");
			}).fail(notifyError);
			return false;
		}
	});
}

function validateConvention(stu) {
	if (checkConventionAlignment() && checkTutorAlignment()) {
		prepareValidation();
	}
	/*else {
		$('#modal').find('[data-toggle="confirmation"]').confirmation('show');
	}*/
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
		reportError("#convention-selecter", "does not match")
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

	var e = exists(tutor, allTeachers);
	if (em == "_new_") {
		if (e) {
			reportError("#tutor-selecter", e.Lastname + ", " + e.Firstname + " ?");
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
	//Catch the line, remove it by the new one
	var em = $("legend").data("email");
	var row = $("#table-placement").find("tr[data-email='" + em + "']");
	var dta = i.Convention.Student;
	dta.I = i;
	dta.Warn = !i && !dta.Skip;
	var cnt = Handlebars.partials['placement-student'](dta);
	row.replaceWith(cnt);
	$('.tablesorter').trigger("update").trigger("updateCache");
	hideModal();
	var v = $('#placed_cnt').html();
	$('#placed_cnt').html(parseInt(v) + 1);
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

function updateInternship(xhr, em) {
	internship(em).done(function(i) {
		var row = $("#table-placement").find("tr[data-email='" + em + "']");
		var dta = i.Convention.Student;
		dta.I = i;
		dta.Warn = !i && !dta.Skip
		var cnt = Handlebars.partials['placement-student'](dta);
		row.replaceWith(cnt);
		$('.tablesorter').trigger("update").trigger("updateCache");
	});
}

function updateMajor(em, old, sel) {
	var v = $(sel).val();
	postStudentMajor(em, v).done(
		function(dta, status, xhr) {
			updateInternship(xhr, em)
			defaultSuccess({}, "OK");
		}
	).fail(function(xhr) {
		notifyError(xhr)
		$(sel).val(old);
	});
}

function updatePromotion(em, old, sel) {
	var v = $(sel).val();
	postStudentPromotion(em, v).done(
		function(dta, status, xhr) {
			updateInternship(xhr, em);
			defaultSuccess({}, "OK");
		}
	).fail(function(xhr) {
		notifyError(xhr)
		$(sel).val(old);
	});
}

function switchTutor(stu, old) {
	var now = $("#tutor-selecter").val();
	if (now == old) {
		return;
	}
	postNewTutor(stu, now).done(function(u, status, xhr) {
		defaultSuccess({}, "OK");
		$("#tutor-group").find("p").html(Handlebars.partials['person'](u.Person));
		updateInternship(xhr, stu);
	})
}