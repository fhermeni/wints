$(document).on('change', '.btn-file :file', function() {
	var input = $(this),
		numFiles = input.get(0).files ? input.get(0).files.length : 1,
		label = input.val().replace(/\\/g, '/').replace(/.*\//, '');
	input.trigger('fileselect', [numFiles, label]);
});


//Ignored: XX
//Errors: {foo: bar}
var stats;

function showPlacementStatus() {
	$.when(students(), conventions()).done(loadPlacementStatus).fail(logFail)
}


function loadPlacementStatus(students, conventions) {
	students = students[0];
	conventions = conventions[0];
	placed = {};
	students.forEach(function(s) {
		placed[s] = false;
	});
	conventions.forEach(function(c) {
		placed[c.Student.User.Person.Email] = true;
	});
	for (i = 0; i < students.length; i++) {
		students[i].Placed = placed[students[i].User.Person.Email];
	}
	$("#cnt").render("placement-header", students, placementUI);
}

function placementUI() {
	$('.btn-file :file').on('fileselect', csvLoaded);
	ui();
}

function csvLoaded(event, numFiles, label) {
	$("#import-status").addClass("hidden")
	var files = event.target.files; // FileList object
	for (var i = 0, f; f = files[i]; i++) {
		var reader = new FileReader();
		reader.onload = parseCSV
		reader.onerror = function() {
			console.log(arguments)
		}
		reader.readAsText(f)
	}
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

function updateStudentSkipable(stu, skip) {
	postStudentSkippable(stu, skip).done(refreshStudentTable).fail(logFail)
}

function refreshStudentTable() {
	var config = $('#table-placement')[0].config
	$.tablesorter.update(config, false)
}

function showImportError() {
	$("#modal").render("import-error", stats, showModal)
}

function doneStudentImport(student) {
	stats.Success.push(student)
	var row = Handlebars.partials['placement-student'](student);
	var config = $('#table-placement')[0].config
	$.tablesorter.addRows(config, row, true, hideModal);
	usersUI()
	updateImportStatus()
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