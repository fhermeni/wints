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
		if ($(arguments[i]).val() == "") {
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
	var msg = status;
	if (xhr && xhr.responseText) {
		msg = xhr.responseText
	}
	$.notify({
		message: msg
	}, {
		type: "success",
		delay: 1,
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
	}).fail(notifyError)
}

function del(URL, ok, no) {
	return $.ajax({
		method: "DELETE",
		url: ROOT_API + URL,
	}).fail(notifyError);
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
	return get("/config");
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
	return post("/resetPassword", email);
}

function surveyFromToken(token) {
	return get("/surveys/" + token);
}

function postSurvey(token, answers) {
	return post("/surveys/" + token, answers);
}

function postResetSurvey(student, kind) {
	return del("/surveys/" + student + "/" + kind)
}

function postRequestSurvey(student, kind) {
	return post("/surveys/" + student + "/" + kind)
}

function logFail(xhr) {	
	if (xhr.status == 403 || xhr.status == 401) {
		$("#modal").render("error", xhr.responseText, showModal)
	}
	console.log(xhr.status + " " + xhr.responseText)
}

function notifyError(xhr) {
	var msg = xhr
	if (typeof xhr !== "string") {
		msg = xhr.responseText ? xhr.responseText : xhr.status;
	}
	$.notify({
		message: msg
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

function getLog(kind) {
	return get("/logs/" + kind)
}
function logs() {
	return get("/logs/")
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

function postNewTutor(stu, now) {
	return post("/internships/" + stu + "/tutor", now);
}

function postReview(e, k, c, g) {
	buf = {
		Comment: c
	};
	if (g) {
		buf.Grade = g;
	}
	return post("/reports/" + e + "/" + k + "/grade", buf);
}