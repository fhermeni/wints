$(document).ready(function() {
	waitingBlock = $("#cnt").clone().html();

	$.tablesorter.defaults.widgets = ["uitheme"]
	$.tablesorter.defaults.theme = 'bootstrap';
	$.tablesorter.defaults.headerTemplate = '{content} {icon}';

	user(getCookie("login"), loadSuccess, function() {
		window.location.href = "/login"
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
	/*showMyServices(myself.Role);
	if (myself.Role == 0) {
		showDashboard();
	} else if (myself.Role >= 2) {
		displayMyStudents()
	} else {
		displayMyConventions()
	}*/

}

function showProfileEditor() {
	$("#modal").render("profileEditor", myself).modal("show")
}

function showPasswordEditor() {
	$("#modal").render("passwordEditor", myself).modal("show")
}