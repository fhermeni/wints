function showWatchlist() {
	showWait();
	internships().done(loadWatchlist);
}

function loadWatchlist(interns) {
	if (!interns) {
		interns = [];
	}
	managed = interns.filter(function (i) {
		return !i.Convention.Student.Skip;
	});
	$("#cnt").render("watchlist", {
		Internships: managed,
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
		editableDefenseGrade(".editable-defense-grade");
	});
}