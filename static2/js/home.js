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

function logFail(xhr) {
	console.log(xhr.status + " " + xhr.responseText)
}