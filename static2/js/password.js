function equals(f1, f2) {
	var v1 = $(f1).val();
	var v2 = $(f2).val();
	if (v1 != v2) {
		reportError(f2, "Passwords do not match")
		return false;
	}
	return true;
}

function setPassword() {
	if (empty("#passwd1", "#passwd2") || !equals("#passwd1", "#passwd2")) {
		return
	}
	newPassword($.urlParam("token"), $("#passwd1").val())
		.fail(failSetPassword)
		.done(doneSetPassword)
}

function doneSetPassword(xhr) {
	window.location.href = "login.html?login=" + xhr.responseText;
}

function failSetPassword(xhr) {
	reportError("#passwd1", xhr.responseText)
}
$(document).ready(function() {
	$('[data-toggle="popover"]').popover()
});