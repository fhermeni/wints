var config;

var STUDENT_LEVEL = 0;
var TUTOR_LEVEL = 1;
var MAJOR_LEVEL = 2;
var HEAD_LEVEL = 3;
var ADMIN_LEVEL = 4;
var ROOT_LEVEL = 5;


String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.substring(1);
};

function showWait() {
	$("#cnt").html(waitingBlock);
}

function loadSuccess(data) {
	myself = data;
	$("#fullname").html(myself.Person.Lastname + ", " + myself.Person.Firstname);

	//my options
	for (i = 0; i <= level(myself.Role); i++) {
		$(".role-" + i).removeClass("hidden");
	}

	//homepage
	if (level(myself.Role) == STUDENT_LEVEL) {
		showStudent();
	} else if (level(myself.Role) >= MAJOR_LEVEL) {
		showWatchlist();
	} else {
		showTutored();
	}
}

function level(role) {
	if (role == "student") {
		return STUDENT_LEVEL;
	} else if (role == "tutor") {
		return TUTOR_LEVEL;
	} else if (role.indexOf("major") === 0) {
		return MAJOR_LEVEL;
	} else if (role == "head") {
		return HEAD_LEVEL;
	} else if (role == "admin") {
		return ADMIN_LEVEL;
	} else if (role == "root") {
		return ROOT_LEVEL;
	}
	return -1;
}

function showModal(next) {
	$("#modal").modal("show");
	$('#modal').find('[data-toggle="popover"]').popover();
	$('#modal').unbind('shown.bs.modal');
	if (next) {
		$('#modal').on('shown.bs.modal', function() {
			next();
		});
	}
	$("#modal").find(".date").datetimepicker();
	$('#modal').find('[data-toggle="confirmation"]').confirmation();
}

function tableCount() {
	count = $("tbody").find("tr").length;
	$(".count").html(count);
}

function ui() {

	$("#cnt").find(".tablesorter").tablesorter();
	$('#cnt').find('[data-toggle="popover"]').popover();
	$('#cnt').find('[data-toggle="confirmation"]').confirmation();

	$("#cnt").find(".shiftSelectable").shiftSelectable();
	$("#cnt").find("table").bind("sortEnd", function() {
		$("#cnt").find('.shiftSelectable').shiftSelectable();
	});

	$(".date").datetimepicker({
		format: "DD MMM YYYY"
	});

	tableCount();

	$(".globalSelect").change(function() {
		var ctx = $(this).data("context");
		$("#" + ctx).find("input:checkbox").prop("checked", this.checked);
	});
}

function showAlumni(student) {
	internship(student).done(function(i) {
		$("#modal").render("alumni-modal", i.Convention.Student, function() {
			var val = i.Convention.Student.Alumni.Position;

			if (val == "sabbatical") {
				$("#country").addClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "entrepreneurship") {
				$("#country").removeClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "study") {
				$("#country").removeClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
			} else if (val == "company") {
				$("#country").removeClass("hidden");
				$("#contract").removeClass("hidden");
				$("#company").removeClass("hidden");
		} else if (val == "looking") {
				$("#country").addClass("hidden");
				$("#contract").addClass("hidden");
				$("#company").addClass("hidden");
		}
		showModal();
		});
	});
}

function hideModal() {
	$('#modal').find('[data-toggle="popover"]').popover('destroy');
	$("#modal").modal("hide");
	$("#modal").html("");
}

function updateInternshipRow(em) {
	var partial = $("table").data("partial");
	var row = $("table").find("tr[data-email='" + em + "']");
	internship(em).done(function(u) {
		var cnt = Handlebars.partials[partial](u);
		row.replaceWith(cnt);
		$('table').trigger("update").trigger("updateCache");
	});
}

function showInternship(em, edit) {
	if (!edit || level(myself.Role) < ADMIN_LEVEL) {
		internship(em).done(function(i) {
			internshipModal(i, [], edit && level(myself.Role) >= ADMIN_LEVEL);
		}).fail(logFail);
	} else {
		$.when(internship(em), users()).done(function(i, uss) {
			i = i[0];
			uss = uss[0].filter(function(u) {
				return level(u.Role) != STUDENT_LEVEL && u.Person.Email != i.Convention.Tutor.Person.Email;
			});
			internshipModal(i, uss, edit);
		}).fail(logFail);
	}
}

function resetSurvey(btn, student, kind) {
	postResetSurvey(student, kind).done(function(data, status, xhr) {
		$(btn).attr("disabled","disabled");
		defaultSuccess(data, status, xhr);
	}).fail(notifyError);
}

function requestSurvey(btn, student, kind) {
	postRequestSurvey(student, kind).done(function(data, status, xhr) {
		defaultSuccess(data, status, xhr);
		$(btn).html(kind);
	}).fail(notifyError);
}

function internshipModal(i, uss, edit) {
	var dta = {
		I: i,
		Editable: edit,
		Teachers: uss
	}
	$("#modal").render("convention-detail", dta, showModal);
}