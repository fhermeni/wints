/**
 * Created by fhermeni on 06/08/2014.
 */

$.urlParam = function(name) {
	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
	if (results == null) {
		return null;
	} else {
		return results[1] || 0;
	}
}

function flipForm(from, to) {
	$('[data-toggle="popover"]').popover('hide').closest(".form-group").removeClass("has-error");
	$("#" + from).slideToggle(400, function() {
		$("#" + to).slideToggle();
	});
}

function empty() {
	var args = Array.prototype.slice.call(arguments);
	return args.filter(function(id) {
		if ($(id).val() == 0) {
			reportError(id, "required")
			return true
		}
		return false
	}).length == args.length
}

function login() {
	if (empty("#loginEmail", "#loginPassword")) {
		return
	}
	signin($("#loginEmail").val(), $("#loginPassword").val())
		.done(function() {
			console.log("kk")
		}).fail(loginFail)

}

function loginFail(xhr) {
	if (xhr.status == 404) {
		reportError("#loginEmail", xhr.responseText)
	} else if (xhr.status == 403) {
		reportError("#loginPassword", xhr.responseText)
	}
}

function reportError(id, message) {
	$(id).data("content", message)
		.popover("show")
		.closest(".form-group").addClass("has-error");
}

function passwordLost() {
	if (empty("#lostEmail")) {
		return
	}
	resetPassword($("#lostEmail").val())
		.fail(passwordLostFail)
}

function passwordLostFail(xhr) {
	reportError("#lostEmail", xhr.responseText)
}

$(document).ready(function() {
	var em = decodeURIComponent($.urlParam("email"))
	if (em && em != "null") {
		$("#login").val(em);
	}
});