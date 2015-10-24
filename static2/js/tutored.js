function showTutored() {
	internships().done(loadTutored);
}

function loadTutored(interns) {
	//Filter people I am not following (due to the major/+ rights)
	var ii = interns.filter(function(i) {
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