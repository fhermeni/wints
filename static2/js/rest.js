/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries

$.urlParam = function(name) {
	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
	if (results == null) {
		return null;
	} else {
		return results[1] || 0;
	}
}

function empty() {
	var count = 0;
	for (var i = 0; i < arguments.length; i++) {
		if ($(arguments[i]).val() == 0) {
			reportError(arguments[i], "required");
			count++;
		} else {
			$(arguments[i]).closest(".form-group").removeClass("has-error");
			$(arguments[i]).popover("hide");
		}
	}
	return count != 0;
}

function reportError(id, message) {
	var popover = $(id).data('bs.popover');
	if (popover) {
		popover.options.content = message;
	} else {
		$(id).popover({
			template: '<div class="popover popover-error" role="tooltip"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>',
			content: message
		});
	}
	$(id).popover("show").closest(".form-group").addClass("has-error");
}

function defaultSuccess(data, status, xhr) {
	$.notify({
		message: xhr.responseText ? xhr.responseText : status
	}, {
		type: "success",
		delay: 1000,
		placement: {
			from: "bottom",
			align: "right"
		}
	});
}

function cleanError() {
	for (var i = 0; i < arguments.length; i++) {
		$(arguments[i]).closest(".form-group").removeClass("has-error");
		$(arguments[i]).popover("destroy");
	}
}

function getCookie(name) {
	var value = "; " + document.cookie;
	var parts = value.split("; " + name + "=");
	if (parts.length == 2) return parts.pop().split(";").shift();
}

function equals(f1, f2) {
	if ($(f1).val() != $(f2).val()) {
		reportError(f2, "Password do not match");
		return false;
	}
	return true;
}

var ROOT_API = "/api/v2";

//Profile management

function post(URL, data) {
	return $.ajax({
		method: "POST",
		data: JSON.stringify(data),
		url: ROOT_API + URL,
	});
}

function get(URL) {
	return $.ajax({
		method: "GET",
		url: ROOT_API + URL,
	}).fail(logFail)
}

function del(URL, ok, no) {
	return $.ajax({
		method: "DELETE",
		url: ROOT_API + URL,
	}).fail(logFail);
}

function signin(login, password) {
	return post("/signin", {
		Login: login,
		Password: password
	});
}

function delSession() {
	return del("/users/" + myself.Person.Email + "/session");
}

function resetPassword(email) {
	return post("/resetPassword", email);
}

function newPassword(token, passwd) {
	return post("/newPassword", {
		Token: token,
		Password: passwd
	});
}

function sendPassword(current, now) {
	return post("/users/" + myself.Person.Email + "/password", {
		Current: current,
		Now: now
	});
}

function user(email) {
	return get("/users/" + email);
}

function sendProfile(p) {
	return post("/users/" + p.Email + "/person", p);
}

function internships() {
	return get("/internships/");
}

function internship(em) {
	return get("/internships/" + em);
}

function getConfig() {
	return get("/config/");
}

function users() {
	return get("/users/");
}

function postNewUser(u) {
	return post("/users/", u);
}

function delUser(email) {
	return del("/users/" + email);
}

function postUserRole(email, r) {
	return post("/users/" + email + "/role", r);
}

function postStudentMajor(email, m) {
	return post("/students/" + email + "/major", m);
}

function postStudentPromotion(email, p) {
	return post("/students/" + email + "/promotion", p);
}

function students() {
	return get("/students/");
}

function conventions() {
	return get("/conventions/");
}

function pendingConventions() {
	return get("/conventions/");
}

function postStudent(p) {
	return post("/students/", p);
}

function postStudentSkippable(stu, skip) {
	return post("/students/" + stu + "/skip", skip);
}

function postUserEmail(old, now) {
	return post("/users/" + old + "/email", now);
}

function internship(email) {
	return get("/internships/" + email);
}

function convention(email) {
	return get("/conventions/" + email);
}

function newInternship(c) {
	return post("/internships/", c);
}

function resetPassword(email) {
	return post("/resetPassword", email)
}

function logFail(xhr) {
	if (xhr.status == 403) {
		var e = $("#modal").render("error", xhr.responseText, showModal)
	}
	console.log(xhr.status + " " + xhr.responseText)
}

function notifyError(xhr) {
	$.notify({
		message: xhr.responseText ? xhr.responseText : status
	}, {
		type: "danger",
		delay: 1000,
		placement: {
			from: "bottom",
			align: "right"
		}
	});
}

function postCompany(em, cpy) {
	return post("/internships/" + em + "/company", cpy)
}

function postTitle(em, title) {
	return post("/internships/" + em + "/title", title)
}

function postSupervisor(em, sup) {
	return post("/internships/" + em + "/supervisor", sup)
}

function postAlumni(em, a) {
	return post("/students/" + em + "/alumni", a)
}

function getReport(stu, kind) {
	return get("/reports/" + stu + "/" + kind)
}

function postReportConfidential(stu, kind, b) {
	return post("/reports/" + stu + "/" + kind + "/private", b)
}

function postReportDeadline(stu, kind, date) {
	return post("/reports/" + stu + "/" + kind + "/deadline", date);
}

/*function setReportContent(email, kind, d, progressFunc, ok, no) {
	return $.ajax({
		method: "POST",
		url: ROOT_API + "/internships/" + email + "/reports/" + kind + "/content",
		data: d,
		processData: false,
		contentType: false,
		xhr: function() { // custom xhr
			myXhr = $.ajaxSettings.xhr();
			if (myXhr.upload) { // check if upload property exists
				myXhr.upload.addEventListener('progress', progressFunc, false); // for handling the progress of the upload
			}
			return myXhr;
		}
	}).done(noCb(ok)).fail(restError(no));
}*/

function postReport(email, kind, dta, progress) {
	return $.ajax({
		method: "POST",
		url: ROOT_API + "/reports/" + email + "/" + kind + "/content",
		data: dta,
		processData: false,
		contentType: "application/pdf",
		xhr: function() { // custom xhr
			myXhr = $.ajaxSettings.xhr();
			if (myXhr.upload) { // check if upload property exists
				myXhr.upload.addEventListener('progress', progress, false); // for handling the progress of the upload
			}
			return myXhr;
		}
	});

}

/*function user(email, ok, no) {
	return get("/users/" + email, ok, no);
}

function users(ok, no) {
	return get("/users/", ok, no);
}

function sessions(ok, no) {
	return get("/sessions/", ok, no);
}



function setUser(fn, ln, tel, ok, no) {
	var email = getCookie("session")
	return put("/users/" + email + "/profile", {
		Firstname: fn,
		Lastname: ln,
		Tel: tel
	}, ok, no);
}

function setPassword(old, n, ok, no) {
	var email = getCookie("session")
	return put("/users/" + email + "/password", {
		Old: old,
		New: n
	}, ok, no);
}

function resetPassword(email, ok, no) {
	return del("/users/" + email + "/password", ok, no)
}

function reInvite(email, ok, no) {
	return del("/users/" + email + "/password?invite=true", ok, no)
}


function internships(ok, no) {
	return get("/internships/", ok, no);
}

function internship(email, ok, no) {
	return get("/internships/" + email, ok, no);
}

function newInternship(c, ok, no) {
	return post("/internships/", c, ok, no);
}

function conventions(ok, no) {
	return get("/conventions/", ok, no);
}

function skipConvention(email, skip, ok, no) {
	return post("/conventions/" + email + "/skip", skip, ok, no)
}

function deleteConvention(email, ok, no) {
	return del("/conventions/" + email, ok, no)
}

function reportHeader(email, kind, ok, no) {
	return get("/internships/" + email + "/reports/" + kind, ok, no);
}

function setReportDeadline(email, kind, d, ok, no) {
	return post("/internships/" + email + "/reports/" + kind + "/deadline", d, ok, no)
}

function setReportContent(email, kind, d, progressFunc, ok, no) {
	return $.ajax({
		method: "POST",
		url: ROOT_API + "/internships/" + email + "/reports/" + kind + "/content",
		data: d,
		processData: false,
		contentType: false,
		xhr: function() { // custom xhr
			myXhr = $.ajaxSettings.xhr();
			if (myXhr.upload) { // check if upload property exists
				myXhr.upload.addEventListener('progress', progressFunc, false); // for handling the progress of the upload
			}
			return myXhr;
		}
	}).done(noCb(ok)).fail(restError(no));

}

function setTutor(e, t, ok, no) {
	return post("/internships/" + e + "/tutor", t, ok, no)
}

function setMajor(e, m, ok, no) {
	return post("/internships/" + e + "/major", m, ok, no)
}

function setCompany(e, n, w, ok, no) {
	return post("/internships/" + e + "/company", {
		Name: n,
		WWW: w
	}, ok, no)
}

function setSupervisor(e, fn, ln, email, tel, ok, no) {
	return post("/internships/" + e + "/supervisor", {
		Firstname: fn,
		Lastname: ln,
		Email: email,
		Tel: tel
	}, ok, no)
}

function setTitle(e, t, ok, no) {
	return post("/internships/" + e + "/title", t, ok, no)
}

function setReportPrivate(e, k, b, ok, no) {
	return post("/internships/" + e + "/reports/" + k + "/private", b, ok, no)
}

function setReportGrade(e, k, g, c, ok, no) {
	return post("/internships/" + e + "/reports/" + k + "/grade", {
		Grade: g,
		Comment: c
	}, ok, no)
}

function majors(ok, no) {
	return get("/majors/", ok, no);
}

function longSurvey(token, ok, no) {
	return get("/surveys/" + token, ok, no)
}

function setSurveyAnswers(token, cnt, ok, no) {
	return post("/surveys/" + token, cnt, ok, no)
}

function alumni(student, ok, no) {
	return get("/internships/" + student + "/alumni", ok, no)
}

function setAlumni(student, pos, email, ok, no) {
	return post("/internships/" + student + "/alumni", {
		Contact: email,
		Position: pos
	}, ok, no)
}

function setNextPosition(student, pos, ok, no) {
	return post("/internships/" + student + "/alumni/position", pos, ok, no)
}

function statistics(ok, no) {
	return get("/statistics/", ok, no)
}

function getStudents(ok, no) {
	return get("/students/", ok, no)
}

function alignStudentWithInternship(stu, internship, ok, no) {
	return post("/students/" + stu + "/internship", internship, ok, no)
}

function hideStudent(stu, flag, ok, no) {
	return post("/students/" + stu + "/hidden", flag, ok, no)
}

function sendStudents(cnt, ok, no) {
	return post("/students/", cnt, ok, no)
}

function defenses(ok, no) {
	return get("/defenses/", ok, no)
}

function defenseProgram(ok, no) {
	return get("/program/", ok, no)
}

function postDefenses(cnt, ok, no) {
	return post("/defenses/", cnt, ok, no)
}

function postDefenseGrade(stu, g, ok, no) {
	return post("/internships/" + stu + "/defense/grade", g, ok, no)
}

function defense(student, ok, no) {
	return get("/internships/" + student + "/defense", ok, no)
}

function requestSurvey(student, kind, ok, no) {
	return get("/internships/" + student + "/surveys/" + kind + "/request", ok, no)
}*/