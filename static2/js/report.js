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
	postReportDeadline(em, kind, e.date).done(function(data, status, xhr) {
		updateInternshipRow(em);
		defaultSuccess({}, status, xhr)
	}).fail();
}

function showReportModal(r, em) {
	r.Email = em;
	$("#modal").render("report-modal", r, function() {
		showModal(function() {
			$(".date").datetimepicker().on("dp.change", function(e) {
				console.log("here " + em);
				updateReportDeadline(em, r.Kind, e);
			});
		});
	});
}