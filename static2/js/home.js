$(document).ready(function() {
	waitingBlock = $("#cnt").clone().html();


	$.tablesorter.defaults.widgets = ["uitheme"]
	$.tablesorter.defaults.theme = 'bootstrap';
	$.tablesorter.defaults.headerTemplate = '{content} {icon}';

	user(getCookie("login")).done(loadSuccess).fail(function() {
		window.location.href = "/login";
	});

	$(document).keydown(function(e) {
		if (e.keyCode == 16) {
			shiftPressed = true;
		}
	});
	$(document).keyup(function(e) {
		if (e.keyCode == 16) {
			shiftPressed = false;
		}
	});
});

function loadSuccess(data) {
	myself = data;
	$("#fullname").html(myself.Person.Firstname + " " + myself.Person.Lastname);

	//my options
	for (i = 0; i <= myself.Role; i++) {
		$(".role-" + i).removeClass("hidden");
	}
}

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

function showModal() {
	$("#modal").modal("show")
	$('#modal').find('[data-toggle="popover"]').popover()
}

function ui() {

	$("#cnt").find(".tablesorter").tablesorter();
	$('#cnt').find('[data-toggle="popover"]').popover()
	$('#cnt').find('[data-toggle="confirmation"]').confirmation()

	$("#cnt").find(".shiftSelectable").shiftSelectable();
	$("#cnt").find("table").bind("sortEnd", function() {
		$("#cnt").find('.shiftSelectable').shiftSelectable();
	})
}

function hideModal() {
	$('#modal').find('[data-toggle="popover"]').popover('destroy')
	$("#modal").modal("hide")
}

function showWatchlist() {
	$.when(internships(), config()).then(loadWatchlist);
}

function loadWatchlist(interns, organization) {
	$("#cnt").render("watchlist", {
		Internships: interns[0],
		Org: organization[0]
	})
}

function showUsers() {
	users().done(loadUsers)
}

function showNewUser() {
	$("#modal").render("new-user", {}, showModal)
}

function newUser() {
	if (empty("#new-firstname", "#new-lastname", "#new-email")) {
		return
	}
	u = {
		Person: {
			Firstname: $("#new-firstname").val().toLowerCase(),
			Lastname: $("#new-lastname").val().toLowerCase(),
			Email: $("#new-email").val().toLowerCase(),
			Tel: $("#new-tel").val().toLowerCase(),
		},
		Role: parseInt($("#new-role").val())
	}
	postNewUser(u)
		.done(successNewUser)
		.fail(failNewUser)
}


function usersUI() {
	$("#cnt").find(".editable-role").each(function(i, e) {
		$(e).editable({
			source: editableRoles(),
			url: function(p) {
				return postUserRole($(e).data("user"), parseInt(p.value));
			}
		});
	});
	ui();
}

function loadUsers(users) {
	$("#cnt").render("users-header", users, usersUI);
}


function successNewUser(p) {
	var row = Handlebars.partials["users-user"](p);
	var config = $('#table-users')[0].config
	$.tablesorter.addRows(config, row, true, hideModal);
	usersUI()
}

function failNewUser(xhr) {
	if (xhr.status == 409 || xhr.status == 400) { //user exists or invalid email
		reportError("#new-email", xhr.responseText)
	}
	console.log(xhr.status + " " + xhr.responseText);
}

function rmUser(em) {
	delUser(em).done(function() {
		successDelUser(em)
	}).fail(logFail)
}

function successDelUser(em) {
	var config = $('#table-users')[0].config
	$("#table-users").find('[data-user="' + em + '"]').closest('tr').remove()
	$.tablesorter.update(config, false)
}

function logFail(xhr) {
	console.log(xhr.status + " " + xhr.responseText)
}