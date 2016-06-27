function loadDefenseProgram() {
	program().done(function (dd) {
		moment.locale('fr');
		$("#cnt").render("defense-program", idsToDay(dd), function() {
			ui();
			moment.locale('en');
		});
	});
}

function idsToDay(dd) {
	dd.forEach(function (d, idx) {
		if (d.Defenses.length > 0) {
			dd[idx].Day = d.Defenses[0].Time;
		}
	});
	return dd;
}