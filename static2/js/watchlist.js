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
				showReport(em, k, updateInternshipRow);
			});
		});
	});
}