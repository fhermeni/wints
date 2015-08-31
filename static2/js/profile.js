function showProfileEditor() {
	$("#modal").render("profileEditor", myself, showModal)
}

function showPasswordEditor() {
	$("#modal").render("passwordEditor", {}, showModal)
}

function updatePassword() {
	if (empty("#password-current", "#password-new", "#password-confirm") || !equals("#password-new", "#password-confirm")) {
		return
	}
	sendPassword($("#password-current").val(), $("#password-new").val())
		.done(hideModal)
		.fail(failUpdatePassword)
}

function failUpdatePassword(xhr) {
	if (xhr.status == 401) {
		reportError("#password-current", xhr.responseText)
	} else if (xhr.status == 400) {
		reportError("#password-new", xhr.responseText)
	}
}

function updateProfile() {
	if (empty("#profile-firstname", "#profile-lastname")) {
		return
	}
	p = {
		Email: myself.Person.Email,
		Firstname: $("#profile-firstname").val(),
		Lastname: $("#profile-lastname").val(),
		Tel: $("#profile-tel").val()
	}
	sendProfile(p)
		.done(successUpdateProfile)
}

function logout() {
	delSession()
		.done(function() {
			window.location.href = "/"
		})
		.fail(logFail);
}

function successUpdateProfile(p) {
	myself.Person = p
	$("#fullname").html(p.Firstname + " " + p.Lastname);
	hideModal()
}