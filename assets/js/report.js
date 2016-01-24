function showReport(email, kind) {
	internship(email).done(function(i) {
		showReportModal(i, kind);
	});
}

function toggleReportConfidential(k, em, chk) {
	var b = $(chk).prop("checked");
	postReportConfidential(em, k, b).fail(function(xhr) {
		$(chk).prop("checked", !b);
		notifyError(xhr);
	}).done(function(data, status, xhr) {
		defaultSuccess({}, status, xhr);
	});
}


function updateReportDeadline(em, kind) {
	var now = $('#report-deadline').data("DateTimePicker").date();
	postReportDeadline(em, kind, now.toJSON()).done(function(data, status, xhr) {
		updateInternshipRow(em);
		defaultSuccess({}, status, xhr)
	}).fail(function(xhr) {
		input.val($("#report-deadline").data("old-date"));
		notifyError(xhr);
	});
}

function showReportModal(i, kind) {
	var r;
	i.Reports.forEach(function (rr) {
		if (rr.Kind == kind) {
			r = rr;
			return false;
		}
	})		
	r.Email = i.Convention.Student.User.Person.Email;
	r.Tutor = i.Convention.Tutor.Person.Email;
	$("#modal").render("report-modal", r, function() {
		$("#report-deadline").datetimepicker()/*.on("dp.change", function(e) {
			updateReportDeadline(r.Email, r.Kind, e);
		});*/
		showModal()
	});
}

function reviewDelay(r) {
	return Math.floor(moment.duration(moment(r.Delivery).diff(moment(r.Deadline))).asDays());
}
function netGrade(r) {
	var duration = moment.duration(moment(r.Delivery).diff(moment(r.Deadline)));
	var days = Math.floor(duration.asDays());
	return Math.max(0, r.Grade - config.LatePenalty * days);
}

function review(student, kind, toGrade) {
	if (empty("#comment")|| (toGrade && empty("#grade"))) {
		return;
	}	

	var comment = $("#comment").val();
	var g = $("#grade").val();	
	postReview(student, kind, comment, parseInt(g)).done(function(dta, status, xhr) {
		updateInternshipRow(student);
		defaultSuccess({}, status, xhr)
		hideModal();
	});
}

function penalty(deadline, delivery) {
	var dead = moment(deadline);
	var del = delivery ? moment(delivery) : moment();
	var diff = delivery.unix() - deadline.unix();
	if (diff < 0) {
		return 0;
	}

	return Math.floor(diff / 60 / 60 / 24) * config.LatePenalty;
}