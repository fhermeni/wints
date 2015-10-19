function showProfileEditor() {
	$("#modal").render("profile-editor", myself, showModal)
}

function makeEditable(root) {
	$(root).find(".editable-role").each(function(i, e) {
		$(e).editable({
			source: editableRoles(),
			url: function(p) {
				return postUserRole($(e).data("user"), parseInt(p.value))
			}
		});
	});
}


function showPasswordEditor() {
	$("#modal").render("password-editor", {}, showModal)
}

function startResetPassword() {
	resetPassword(myself.Person.Email)
		.fail(resetFail).done(passwordLostOk);
}

function passwordLostOk(xhr) {
	$(".alert-success").removeClass("hidden");
	$("#modal").find(".btn").attr("disabled", "disabled");
}

function resetFail(xhr) {
	$("#modal").find(".alert-danger").removeClass("hidden");
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

function successUpdateProfile(u) {
	if (u.Person.Email == myself.Person.Email) {
		myself = u;
		$("#fullname").html(p.Lastname + ", " + p.Firstname);
	}
	hideModal();
}