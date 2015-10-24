function showReport(email, kind, cb) {
	getReport(email, kind).done(function(r) {
		showReportModal(r, email, cb);
	});
}

function toggleReportConfidential(k, em, chk) {
	var b = $(chk).prop("checked");
	postReportConfidential(em, k, b).fail(function(xhr) {
		$(chk).prop("checked", !b);
		notifyError(xhr);
	});
}


function updateReportDeadline(em, kind, e, cb) {
	postReportDeadline(em, kind, e.date).done(function(data, status, xhr) {
		cb(em);
		defaultSuccess({}, status, xhr)
	}).fail();
}

function showReportModal(r, em, cb) {
	r.Email = em;
	$("#modal").render("report-modal", r, function() {
		showModal(function() {
			$(".date").datetimepicker().on("dp.change", function(e) {
				updateReportDeadline(em, r.Kind, e, cb);
			});
		});
	});
}