var config;

String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.substring(1)
}


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
});

function loadSuccess(data) {
	myself = data;
	$("#fullname").html(myself.Person.Lastname + ", " + myself.Person.Firstname);

	//my options
	for (i = 0; i <= myself.Role; i++) {
		$(".role-" + i).removeClass("hidden");
	}

	//homepage	
	if (myself.Role == 1) { //student
		showStudent();
	} else if (myself.Role >= 4) { //admin +
		showWatchlist();
	} else {
		//tutor || major
		showTutored();
	}
}

function showModal(next) {
	$("#modal").modal("show");
	$('#modal').find('[data-toggle="popover"]').popover()
	if (next) {
		$('#modal').on('shown.bs.modal', function(e) {
			next()
		});
	} else {
		$('#modal').unbind('shown.bs.modal');
	}

	$("#modal").find(".date").datetimepicker();

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
				return postUserRole($(e).data("user"), parseInt(p.value))
			}
		});
	});

	$(".date").datetimepicker({
		format: "DD MMM YYYY"
	});
	/*$("#cnt").find(".editable-promotion").each(function(i, e) {
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
	});*/

	$(".globalSelect").change(function() {
		var ctx = $(this).data("context");
		$("#" + ctx).find("input:checkbox").prop("checked", this.checked);
	});
}

function hideModal() {
	$('#modal').find('[data-toggle="popover"]').popover('destroy');
	$("#modal").modal("hide");
}

function showInternship(em) {
	internship(em).done(internshipModal).fail(logFail)
}

function internshipModal(i) {
	$("#modal").render("convention-detail", i, showModal);
}