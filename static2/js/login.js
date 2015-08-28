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
	$('[data-toggle="popover"]').popover('hide')
	$("#" + from).slideToggle(400, function() {
		$("#" + to).slideToggle();
	});
}

function empty(id) {
	var q = $(id)
	if (q.val() == 0) {
		q.closest(".form-group").addClass("has-error")
		return true
	}
	return false
}

function login() {
	/*	$(function() {
			$('[data-toggle="popover"]').popover()
		})*/
	if (empty("#loginEmail") ||  empty("#loginPassword")) {
		return
	}
	signin($("#loginEmail").val(), $("#loginPassword").val())
		.done(function() {
			console.log("kk")
		})
		.fail(function(xhr) {
			if (xhr.status == 404) {
				$("#loginEmail").closest(".form-group").addClass("has-error")
				$("#loginEmail").data("content", xhr.responseText).popover("show")

			} else if (xhr.status == 403) {
				$("#loginPassword").closest(".form-group").addClass("has-error")
				$("#loginPassword").data("content", xhr.responseText).popover("show")
			}
		})

}

function passwordLost() {
	if (empty("#lostEmail")) {
		return
	}
	var email = $("#lostEmail").val();
	resetPassword(email)
		.fail(function(xhr) {
			$("#lostEmail").closest(".form-group").addClass("has-error")
			$("#lostEmail").data("content", xhr.responseText).popover("show")
		})
}

$(document).ready(function() {
	var em = decodeURIComponent($.urlParam("email"))
	if (em && em != "null") {
		$("#login").val(em);
	}
});