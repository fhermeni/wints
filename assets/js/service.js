function showService() {
	showWait();
	$.when(internships(), defenses())
		.done(computeService);

}

function computeService(ints, ss) {
	ints = ints[0];
	ss = ss[0];

	var service = {};
	ints.forEach(function(i) {
		var em = i.Convention.Tutor.Person.Email;
		var p = i.Convention.Student.Promotion;
		if (!service[em]) {
			service[em] = {
				U: i.Convention.Tutor,
				Ints: {},
				Defs: 0,
				TotalInts: 0,
			};
		}
		if (!service[em].Ints[p]) {
			service[em].Ints[p] = []
		}
		service[em].Ints[p].push(i)
		service[em].TotalInts++
	});

	ss.forEach(function (s) {
		s.Juries.forEach(function(j) {
			var em = j.Person.Email;
			service[em].Defs++;
		});
	})
	$("#cnt").render("service", service, ui);
}