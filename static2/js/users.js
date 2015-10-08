$(document).on('change', '.btn-file :file', function() {
	var input = $(this),
		numFiles = input.get(0).files ? input.get(0).files.length : 1,
		label = input.val().replace(/\\/g, '/').replace(/.*\//, '');
	input.trigger('fileselect', [numFiles, label]);
});


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
	$('.btn-file :file').on('fileselect', csvLoaded);
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


function doneStudentImport(student) {
	stats.Success.push(student)
	var row = Handlebars.partials["users-user"](student.User);
	var config = $('#table-users')[0].config
	$.tablesorter.addRows(config, row, true, hideModal);
	usersUI()
	updateImportStatus()
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

function csvLoaded(event, numFiles, label) {
	$("#import-status").addClass("hidden")
	var files = event.target.files;
	for (var i = 0, f; f = files[i]; i++) {
		var reader = new FileReader();
		reader.onload = parseCSV
		reader.onerror = function() {
			console.log(arguments)
		}
		reader.readAsText(f)
	}
}

function parseCSV(e) {
	stats = {
		Nb: 0,
		Ignored: [],
		Errors: [],
		Success: []
	};
	var lines = e.target.result.split(/\r\n|\n|\r/);
	var first = true;
	lines.forEach(function(line) {
		if (first) {
			first = false;
			return true;
		}
		if (line.length == 0) {
			return true;
		}
		stats.Nb++;
		fields = line.split(';');
		var p = {
			User: {
				Person: {
					Lastname: fields[0].toLowerCase(),
					Firstname: fields[1].toLowerCase(),
					Email: fields[2].toLowerCase()
				}
			},
			Promotion: among(fields[3].toLowerCase(), config.Promotions),
			Major: among(fields[4].toLowerCase(), config.Majors)
		};
		var req = postStudent(p).done(doneStudentImport).fail(function(xhr) {
			failStudentImport(p, xhr)
		});
	});
	$("#csv-import").val("");
}

function updateImportStatus() {
	if (stats.Errors.length != 0) {
		$("#import-status").removeClass("hidden")
	}
}

function showImportError() {
	$("#modal").render("import-error", stats, showModal)
}

function failStudentImport(student, xhr) {
	if (xhr.status == 400) {
		stats.Errors.push({
			Student: student,
			Reason: xhr.responseText
		});
	} else if (xhr.status == 409) {
		stats.Ignored.push(student)
	}
	updateImportStatus()
}

function among(got, possibles) {
	var cleaned = got;
	possibles.forEach(function(p) {
		if (got.indexOf(" " + p) >= 0 || got.indexOf(p) == 0) {
			//a new word or starts with
			cleaned = p;
			return false;
		}
	});
	return cleaned;
}