var mine;

function showStudent() {
	internship(myself.Person.Email).done(showStudentDashboard)
}

function showStudentDashboard(m) {
	mine = m;
	Object.keys(mine.Reports).forEach(function(kind) {
		var passed = moment(mine.Reports[kind].Deadline).isBefore(moment());
		var r = mine.Reports[kind];		
		mine.Reports[kind].Open = isReportUpdloadeable(r);
	});
	$("#cnt").render("student-dashboard", m, function() {
		syncAlumniEditor($("#position"))
		ui()
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

function updateCompany() {
	if (empty("lbl-name") || empty("lbl-title") || empty("lbl-www")) {
		return
	}
	var cpy = {
		Name: $("#lbl-name").val(),
		WWW: $("#lbl-www").val(),
		Title: $("#lbl-title").val()
	}
	postCompany(myself.Person.Email, cpy).done(refreshCompany)
}

function sendSupervisor() {
	if (empty("sup-fn") || empty("sup-ln") || empty("sup-tel") || empty("sup-email")) {
		return
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
	if (empty("#lbl-email")) {
		return
	}
	a = {
		Contact: $("#lbl-email").val(),
		Position: $("#position").val(),
		France: $('input[name=france]:checked').val() == "true",
		Permanent: $('input[name=permanent]:checked').val() == "true",
		SameCompany: $('input[name=sameCompany]:checked').val() == "true",
	}
	postAlumni(myself.Person.Email, a)
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
	reader.readAsDataURL(file);
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
	console.log(r);
	var passed = moment(r.Deadline).isBefore(moment());
	return !r.Reviewed && (!passed || !r.Delivery);
}

function refreshReports(r) {	
	hideModal();
	r.Open = isReportUpdloadeable(r);
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
/*


function showDashboard() {
	internship(myself.Email, function(i) {
		mine = i
		for (x in i.Reports) {
			var delivered = new Date(i.Reports[x].Delivery).getTime() > 0
			i.Reports[x].Uploaded = delivered
				//Locked= deadline passed || commented            
				//Locked: commented || (deadline passed && delivery)
			var passed = (new Date(i.Reports[x].Deadline).getTime() + (86400 * 1000)) < new Date().getTime()
			i.Reports[x].Locked = (passed && delivered) || i.Reports[x].Grade >= 0;
		}
		var html = Handlebars.getTemplate("internship")(i);
		var root = $("#cnt");
		root.html(html);
		$("#select-position").selecter();
		$("#major").selecter({
			callback: function(v) {
				setMajor(mine.Student.Email, v)
			}
		});
		$('input[type=file]').filestyle({
			input: false,
			buttonText: "",
			buttonName: "btn-success btn-sm",
			iconName: "glyphicon-cloud-upload",
			badge: false
		})
		$(':file').change(function() {
			var file = this.files[0];
			var name = file.name;
			var size = file.size;
			var type = file.type;
			if (type != "application/pdf") {
				reportError("The report must be in PDF format. Currently: " + type)
			} else if (size > 10000000) {
				reportError("The report cannot exceed 10MB")
			} else {
				var formData = new FormData();
				formData.append('report', file);
				var html = Handlebars.getTemplate("upload-progress")(mine);
				var root = $("#modal-hard");
				root.html(html).modal({
					backdrop: 'static',
					keyboard: false,
					show: true
				});
				setReportContent(mine.Student.Email, $(this).attr("data-kind"), formData, showProgress, function() {
					$("#modal-hard").modal("hide");
					reportSuccess("Report uploaded");
					showDashboard();
				}, function(o) {
					$("#modal-hard").modal("hide");
					reportError(o.responseText);
				})
			}
		});
	});
}

function showProgress(evt) {
	if (evt.lengthComputable) {
		var pct = evt.loaded / evt.total * 100;
		$("#progress-value").html(Math.round(pct) + "%")
		$("#progress-value").attr("aria-valuenow", pct)
		$("#progress-value").css("width", pct + "%");
	} else {
		// Unable to compute progress information since the total size is unknown
		console.log('unable to complete');
	}
}


function showReportComment(kind) {
	mine.Reports.forEach(function(r) {
		if (r.Kind == kind) {
			if (r.Comment.length > 0) {
				var html = Handlebars.getTemplate("raw");
				$("#modal").html(html).modal("show");
				$("#rawContent").html(r.Comment);

			}
		}
		return false;
	});
}

function sendSupervisor() {
	if (missing("lbl-fn") || missing("lbl-ln") || missing("lbl-email") || missing("lbl-tel")) {
		return
	}
	setSupervisor(myself.Email, $("#lbl-fn").val(), $("#lbl-ln").val(), $("#lbl-email").val(), $("#lbl-tel").val(), function() {
		showDashboard();
		$("#modal").modal('hide');
		reportSuccess("Operation succeeded");
	});
}

function sendAlumni() {
	if (missing("next-email") || missing("select-position")) {
		return
	}
	setAlumni(mine.Student.Email, parseInt($("#select-position").val()), $("#next-email").val(), undefined, function(jqr) {
		$("#select-position").val(mine.Future.Position);
		$("#next-email").val(mine.Future.Contact);
		reportError(jqr.responseText)
	})
}
*/