var config;

$(document).ready(function() {
	waitingBlock = $("#cnt").clone().html();


	$.tablesorter.defaults.widgets = ["uitheme"]
	$.tablesorter.defaults.theme = 'bootstrap';
	$.tablesorter.defaults.headerTemplate = '{content} {icon}';

	getConfig().done(function(c) {
		config = c;
	})
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
	$('#cnt').find('[data-toggle="popover"]').popover();
	$('#cnt').find('[data-toggle="confirmation"]').confirmation();

	$("#cnt").find(".shiftSelectable").shiftSelectable();
	$("#cnt").find("table").bind("sortEnd", function() {
		$("#cnt").find('.shiftSelectable').shiftSelectable();
	});

	$("#cnt").find(".editable-role").each(function(i, e) {
		$(e).editable({
			source: editableRoles(),
			url: function(p) {
				return postUserRole($(e).data("user"), parseInt(p.value));
			}
		});
	});

	$("#cnt").find(".editable-promotion").each(function(i, e) {
		$(e).editable({
			source: editablePromotions(),
			url: function(p) {
				return postStudentPromotion($(e).data("email"), p.value);
			}
		});
	});

	$("#cnt").find(".editable-major").each(function(i, e) {
		$(e).editable({
			source: editableMajors(),
			url: function(m) {
				return postStudentMajor($(e).data("email"), m.value);
			}
		});
	});
}

function hideModal() {
	$('#modal').find('[data-toggle="popover"]').popover('destroy');
	$("#modal").modal("hide");
}

function showWatchlist() {
	internships().done(loadWatchlist);
}

function loadWatchlist(interns) {
	$("#cnt").render("watchlist", {
		Internships: interns,
		Org: config
	});
}

function logFail(xhr) {
	console.log(xhr.status + " " + xhr.responseText)
}