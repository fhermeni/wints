function showReport(email, kind) {
	getReport(email, kind).done(function(r) {
		showReportModal(r, email);
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


function updateReportDeadline(em, kind, e) {
	postReportDeadline(em, kind, e.date.toJSON()).done(function(data, status, xhr) {
		updateInternshipRow(em);
		defaultSuccess({}, status, xhr)
	}).fail(function(xhr) {
		input.val(e.oldDate.format(input.data("date-format")));
		notifyError(xhr);
	});
}

function showReportModal(r, em) {
	r.Email = em;
	console.log(r);
	$("#modal").render("report-modal", r, function() {
		$("#report-deadline").datetimepicker().on("dp.change", function(e) {
			updateReportDeadline(em, r.Kind, e);
		});
		showModal()
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