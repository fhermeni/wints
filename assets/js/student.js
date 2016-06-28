var mine;

function showStudent() {
	showWait();
	internship(myself.Person.Email).done(showStudentDashboard)
}

function showStudentDashboard(m) {
	mine = m;
	var data = {
		Config: config,
		Internship: m
	};
	Object.keys(mine.Reports).forEach(function(kind) {
		var passed = moment(mine.Reports[kind].Deadline).isBefore(moment());
		var r = mine.Reports[kind];
		mine.Reports[kind].Open = isReportUpdloadeable(r);
		mine.Reports[kind].Email = myself.Person.Email;
	});
	$("#cnt").render("student-dashboard", data, function() {
		syncAlumniEditor($("#position"))
		ui()
		showMyDefense();
	});
}

function showCompanyEditor() {
	$("#modal").render("company-editor", mine, showModal)
}

function showSupervisorEditor() {
	$("#modal").render("student-dashboard-supervisor-editor", mine.Convention.Supervisor, showModal);
}

function showAlumniEditor() {
	$("#modal").render("student-dashboard-alumni-editor", mine.Convention.Student.Alumni, showModal);
}

function showMyDefense() {
	getDefense(myself.Person.Email).done(function (def) {
		getDefenseSession(def.Room, def.SessionId).done(function(ss) {
			ss.Defense = def;
			$("#dashboard-defense").render("student-dashboard-defense",ss);
		});
	});
}
function updateCompany() {
	if (empty("#lbl-name") || empty("#lbl-title") || empty("#lbl-www")) {
		return;
	}
	var cpy = {
		Name: $("#lbl-name").val(),
		WWW: $("#lbl-www").val(),
		Title: $("#lbl-title").val()
	}
	postCompany(myself.Person.Email, cpy).done(refreshCompany).fail(notifyError);
}

function sendSupervisor() {
	if (empty("#sup-fn") || empty("#sup-ln") || empty("#sup-tel") || empty("#sup-email")) {
		return;
	}

	if (invalidEmail("#sup-email")) {
		return;
	}
	var sup = {
		Firstname: $("#sup-fn").val(),
		Lastname: $("#sup-ln").val(),
		Email: $("#sup-email").val(),
		Tel: $("#sup-tel").val()
	}
	postSupervisor(myself.Person.Email, sup).done(refreshContacts)
}

function refreshContacts(p) {
	mine.Convention.Supervisor = p;
	var buf = Handlebars.partials["student-dashboard-contacts"](mine);
	$("#student-dashboard-contacts").html(buf);
	hideModal();
}

function refreshCompany(cpy) {
	mine.Convention.Company = cpy;
	var buf = Handlebars.partials["student-dashboard-company"](mine);
	$("#dashboard-company").html(buf);
	hideModal();
}

function sendAlumni() {
	if (invalidEmail("#lbl-email")) {
		return;
	}
	a = {
		Contact: $("#lbl-email").val(),
		Position: $("#position").val(),
		France: $('input[name=france]:checked').val() == "true",
		Permanent: $('input[name=permanent]:checked').val() == "true",
		SameCompany: $('input[name=sameCompany]:checked').val() == "true",
	}
	postAlumni(myself.Person.Email, a).done(function(){defaultSuccess({}, "Done");}).fail(notifyError);
}


function syncAlumniEditor(sel) {
	var val = $(sel).val();
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
}

function loadReport(input, kind) {
	var file = input.files[0];
	var name = file.name;
	var size = file.size;
	var type = file.type;
	if (type != "application/pdf") {
		reportError("The report must be in PDF format. Currently: " + type)
		return
	}
	if (size > 10000000) {
		reportError("The report cannot exceed 10MB")
		return
	}
	var reader = new FileReader();
	reader.onload = function(e) {
		return reportLoaded(kind, reader.result);
	};
	reader.readAsArrayBuffer(file);
}

function reportLoaded(kind, cnt) {
	$("#modal").render("progress", {}, function() {
		showModal(function() {
			postReport(myself.Person.Email, kind, cnt, progress).done(refreshReports);
		});
	});

}

//can upload if not reviewed and either the deadline passed or not passed but not uploaded
function isReportUpdloadeable(r) {
	var passed = moment(r.Deadline).isBefore(moment());
	return !r.Reviewed && (!passed || !r.Delivery);
}

function isReportReviewable(r) {
	var passed = moment(r.Deadline).isBefore(moment());
	return passed;
}

function refreshReports(r) {
	hideModal();
	r.Open = isReportUpdloadeable(r);
	r.Email = myself.Person.Email;
	var buf = Handlebars.partials["student-dashboard-report"](r);
	$("#report-" + r.Kind).replaceWith(buf);
}


function progress(ev) {
	var val = (ev.loaded / ev.total) * 100;
	$("#progress-value")
		.attr('aria-valuenow', Math.round(val))
		.width(val + "%")
		.html(val + "%");
}

function showReportComment(kind) {
	mine.Reports.forEach(function(r) {
		if (r.Kind == kind) {
			$("#modal").render("raw", {
				Title: "Comments for report '" + kind + "'",
				Cnt: r.Comment
			}, showModal);
		}
		return false;
	});
}