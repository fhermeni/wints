//show the defenses I am the jury of
function showJuries() {
		showWait();
		defenses().done(loadDefenses);
}

function loadDefenses(defs) {
	defs = defs.filter(function (def) {
		var mine = false;
		def.Juries.forEach(function (j) {
			if (j.Person.Email == myself.Person.Email) {
				mine = true;
			}
		});
		return mine;
	});
	$("#cnt").render("my-defenses",defs,function() {
		ui();
		editableDefenseGrade(".editable-grade");
	});
}

function editableDefenseGrade(ctx) {
	$(ctx).each(function (idx, e) {
			var em = $(e).data("student");
			$(e).editable({
			url: function(params) {
				return postDefenseGrade(em, params.value)}
			});
	});
}