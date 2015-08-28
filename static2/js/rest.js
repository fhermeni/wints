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
			$(arguments[i]).closest(".form-group").removeClass("has-error")
		}
	}
	return count != 0;
}

function reportError(id, message) {
	var popover = $(id).data('bs.popover')
	popover.options.content = message
	$(id).popover("show")
		.closest(".form-group").addClass("has-error");
}

function getCookie(name) {
	var value = "; " + document.cookie;
	var parts = value.split("; " + name + "=");
	if (parts.length == 2) return parts.pop().split(";").shift();
}

function noCb(no) {
	if (no != undefined) {
		return no;
	}
	return function() {
		reportSuccess("Operation successful");
	};
}

function restError(no) {
	return function(jqr) {
		if (jqr.status == 408) {
			window.location.href = "/login?#sessionExpired"
		} else {

			if (no != undefined) {
				no(jqr)
			} else {
				$.notify(jqr.responseText, {
					className: "danger",
					globalPosition: "top center"
				})
			}
		}
	}
}

function missing(id) {
	var d = $("#" + id);
	if (d.val() == "") {
		d.notify("Required", {
			autoHide: true,
			autoHideDelay: 2000
		});
		return true;
	}
	return false;
}

function isGrade(v) {
	if (isNaN(v)) {
		return "number expected";
	}
	var x = parseInt(v)
	if (x < 0 || x > 20) {
		return "between 0 and 20, inclusive"
	}
	return undefined
}

function equals(f1, f2) {
	var v1 = $("#" + f1).val();
	if (v1.length < 8) {
		$("#" + f1).notify("Passwords must be 8 characters long minimum");
		return false;
	}
	var v2 = $("#" + f2).val();
	if (v1 != v2) {
		$("#" + f2).notify("Password do not match");
		return false;
	}
	return true;
}

var ROOT_API = "/api/v2/";

//Profile management

function post(URL, data) {
	return $.ajax({
		method: "POST",
		data: JSON.stringify(data),
		url: ROOT_API + URL,
	})
}

function put(URL, data, ok, no) {
	return $.ajax({
		method: "PUT",
		data: JSON.stringify(data),
		url: ROOT_API + URL,
	}).done(noCb(ok)).fail(restError(no));
}

function get(URL, ok, no) {
	return $.ajax({
		method: "GET",
		url: ROOT_API + URL,
	}).done(noCb(ok)).fail(restError(no));
}

function del(URL, ok, no) {
	return $.ajax({
		method: "DELETE",
		url: ROOT_API + URL,
	}).done(noCb(ok)).fail(restError(no));
}

function signin(login, password) {
	return post("signin", {
		Login: login,
		Password: password
	})
}

function resetPassword(email) {
	return post("resetPassword", email)
}

function newPassword(token, passwd) {
	return post("newPassword", {
		Token: token,
		Password: passwd
	})
}

function user(email, ok, no) {
	return get("/users/" + email, ok, no);
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

function newUser(fn, ln, tel, email, role, ok, no) {
	return post("/users/", {
		Firstname: fn,
		Lastname: ln,
		Tel: tel,
		Role: role,
		Email: email
	}, ok, no);
}

function rmUser(email, ok, no) {
	return del("/users/" + email, ok, no)
}

function setUserRole(email, r, ok, no) {
	return put("/users/" + email + "/role", r, ok, no);
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