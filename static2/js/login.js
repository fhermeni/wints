/**
 * Created by fhermeni on 06/08/2014.
 */

function flipForm(from, to) {
	$('[data-toggle="popover"]').popover('hide').closest(".form-group").removeClass("has-error");
	$("#" + from).slideToggle(400, function() {
		$("#" + to).slideToggle();
	});
}

function login() {
	if (empty("#loginEmail", "#loginPassword")) {
		return
	}
	signin($("#loginEmail").val(), $("#loginPassword").val())
		.done(loginSuccess).fail(loginFail)

}

function loginSuccess(session) {
	localStorage.setItem("token", session.Token);
	window.location.href = "/"
	console.log("success, relocation")
}

function loginFail(xhr) {
	if (xhr.status == 404) {
		reportError("#loginEmail", xhr.responseText)
	} else if (xhr.status == 401) {
		reportError("#loginPassword", xhr.responseText)
	}
}


function passwordLost() {
	if (empty("#lostEmail")) {
		return
	}
	resetPassword($("#lostEmail").val())
		.fail(passwordLostFail).done(passwordLostOk);
}

function passwordLostOk(xhr) {
	console.log("clean");
	cleanError("#lostEmail")
}

function passwordLostFail(xhr) {
	reportError("#lostEmail", xhr.responseText)
}

$(document).ready(function() {
	var em = $.urlParam("login")
	if (em) {
		$("#loginEmail").val(em);
	}
	$('[data-toggle="popover"]').popover()
});