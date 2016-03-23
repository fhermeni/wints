function showAlumni() {
	students().done(loadAlumni).fail(logFail);
}

function loadAlumni(students) {
	$("#cnt").render("alumni", students, ui);
}