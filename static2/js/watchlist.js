function showWatchlist() {
	internships().done(loadWatchlist);
}

function loadWatchlist(interns) {
	$("#cnt").render("watchlist", {
		Internships: interns,
		Org: config
	}, function() {
		ui();
		$(".report").each(function(i, r) {
			var em = $(r).closest("tr").data("email");
			var k = $(r).data("report-kind");
			$(r).on("click", function() {
				showReport(em, k, updateStudentWatchlist);
			});
		});
	});
}

function updateStudentWatchlist(em) {
	var row = $("#table-conventions").find("tr[data-email='" + em + "']");
	internship(em).done(function(u) {
		var cnt = Handlebars.partials['watchlist-student'](u);
		row.replaceWith(cnt);
		$('#table-users').trigger("update").trigger("updateCache");
		hideModal();
	});
}