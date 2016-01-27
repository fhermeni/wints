function showTutored() {
	showWait();
	internships().done(loadTutored);
}

function loadTutored(interns) {
	if (!interns) {
		interns = [];
	}
	managed = interns.filter(function (i) {
		return !i.Convention.Student.Skip;
	});
	//Filter people I am not following (due to the major/+ rights)
	var ii = managed.filter(function(i) {
		return i.Convention.Tutor.Person.Email == myself.Person.Email;
	});
	$("#cnt").render("tutored", {
		Internships: ii,
		Org: config
	}, function() {
		ui();
		$(".report").each(function(i, r) {
			var em = $(r).closest("tr").data("email");
			var k = $(r).data("report-kind");
			$(r).on("click", function() {
				showReport(em, k);
			});
		});
	});
}