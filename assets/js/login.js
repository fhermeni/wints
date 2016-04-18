/**
 * Created by fhermeni on 06/08/2014.
 */

function flipForm(from, to) {
	$('[data-toggle="popover"]').popover('hide').closest(".form-group").removeClass("has-error");
	var em = from == 'login-diag' ? $("#loginEmail").val() : $("#lostEmail").val();
	if (to == 'login-diag') {
		$("#loginEmail").val(em);
	} else {
		$("#lostEmail").val(em);
	}

	$("#" + from).slideToggle(400, function() {
		$("#" + to).slideToggle();
	});
}

function login() {
	cleanError("#loginEmail", "#loginPassword");
	if (invalidEmail("#loginEmail")) {
		return;
	}
	if (empty("#loginEmail", "#loginPassword")) {
		return;
	}
	console.log("here");
	signin($("#loginEmail").val().trim().toLowerCase(), $("#loginPassword").val())
		.done(loginSuccess).fail(loginFail);

}

function loginSuccess(session) {
	localStorage.setItem("token", session.Token);
	window.location.href = "/";
}

function loginFail(xhr) {
	if (xhr.status == 404) {
		reportError("#loginEmail", xhr.responseText);
	} else if (xhr.status == 401) {
		reportError("#loginPassword", xhr.responseText);
	}
}


function passwordLost() {
	if (invalidEmail("#lostEmail")) {
		return;
	}
	$("#reset-button").attr("disabled","disabled");
	resetPassword($("#lostEmail").val().trim().toLowerCase())
		.fail(passwordLostFail).done(passwordLostOk);
}

function passwordLostOk() {
	$(".alert-success").removeClass("hidden");
	$(".btn").attr("disabled", "disabled");
	cleanError("#lostEmail");
}

function passwordLostFail(xhr) {
	$("#reset-button").removeAttr("disabled");
	reportError("#lostEmail", xhr.responseText);
}