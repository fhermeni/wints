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
				showReport(em, k, updateStudentTutored);
			});
		});
	});
}

function updateStudentTutored(em) {
	var row = $("#table-tutoring").find("tr[data-email='" + em + "']");
	internship(em).done(function(u) {
		var cnt = Handlebars.partials['tutored-student'](u);
		row.replaceWith(cnt);
		$('#table-users').trigger("update").trigger("updateCache");
		hideModal();
	});
}