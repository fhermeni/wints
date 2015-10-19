function showProfileEditor(em) {
	if (!em) {
		$("#modal").render("profile-editor", myself, showModal)
	} else {
		user(em).done(function(ctx) {
			$("#modal").render("profile-editor", ctx, showModal);
		});
	}
}

/*
function showPasswordEditor() {
	$("#modal").render("password-editor", {}, showModal)
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
		reportError("#password-current", xhr.responseText);
	} else if (xhr.status == 400) {
		reportError("#password-new", xhr.responseText);
	}
}
*/
function updateProfile(em) {
	if (empty("#profile-firstname", "#profile-lastname")) {
		return
	}
	user(em).done(function(u) {
		p = u.Person;
		p.Firstname = $("#profile-firstname").val();
		p.Lastname = $("#profile-lastname").val();
		p.Tel = $("#profile-tel").val();
		sendProfile(p).done(successUpdateProfile);
	});
}

function logout() {
	delSession()
		.done(function() {
			window.location.href = "/"
		})
		.fail(logFail);
}

function successUpdateProfile(p) {
	if (p.Email == myself.Person.Email) {
		myself.Person = p
		$("#fullname").html(p.Firstname + " " + p.Lastname);
	}
	hideModal()
}