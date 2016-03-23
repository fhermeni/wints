function showProfileEditor() {
	$("#modal").render("profile-editor", myself, showModal);
}

function makeEditable(root) {
	$(root).find(".editable-role").each(function(i, e) {
		$(e).editable({
			source: editableRoles(),
			url: function(p) {
				return postUserRole($(e).data("user"), parseInt(p.value));
			}
		});
	});
}


function showPasswordEditor() {
	$("#modal").render("password-editor", {}, showModal);
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

function updateProfile(em) {
	if (empty("#profile-firstname", "#profile-lastname")) {
		return;
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
			window.location.href = "/";
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