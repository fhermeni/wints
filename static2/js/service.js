function showService() {
	internships(function(interns) {
		var service = {};
		interns.forEach(function(i) {
			if (!service[i.Tutor.Email]) {
				service[i.Tutor.Email] = {
					P: i.Tutor,
					Internships: {}
				};
			}
			if (!service[i.Tutor.Email].Internships[i.Promotion]) {
				service[i.Tutor.Email].Internships[i.Promotion] = []
			}
			service[i.Tutor.Email].Internships[i.Promotion].push(i)
		});
		var html = Handlebars.getTemplate("service")(service);
		$("#cnt").html(html);

		$("#cnt").find("td .icheckbox_flat").icheck(shiftSelect());
		$('#cnt').find(".check_all").icheck(selectAllNone($("#cnt")));
	});
}