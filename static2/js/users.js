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