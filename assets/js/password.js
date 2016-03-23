function equals(f1, f2) {
	var v1 = $(f1).val();
	var v2 = $(f2).val();
	if (v1 != v2) {
		reportError(f2, "Passwords do not match");
		return false;
	}
	return true;
}

function setPassword() {
	$(".alert-danger").addClass("hidden");
	if (empty("#passwd1", "#passwd2")) {
		return;
	}
	if (!equals("#passwd1", "#passwd2")) {
		return;
	}
	newPassword($.urlParam("token"), $("#passwd1").val())
		.fail(failSetPassword)
		.done(doneSetPassword);
}

function doneSetPassword(em) {
	$(".btn").attr("disabled", "disabled");
	window.location = "/login?email=" + em;
}

function failSetPassword(xhr) {
	cleanError("#passwd1", "#passwd2");
	$(".alert-danger").html(xhr.responseText).removeClass("hidden");
}