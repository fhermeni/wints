//students already having their defense planned
var placed = [];
//available teachers
var teachers = [];

function showDefensePlanner() {
		showWait();
		$.when(defenses(), users()).done(showDefenses).fail(notifyError);
}

function groupById(sessions) {
	var groups = {};
	sessions.map(function (s) {
		var b64 = window.btoa(s.Id);
		if (!groups[b64]) {
			groups[b64] = {
				Id: s.Id,
				B64: b64,
				Sessions: []
			};
		}
		groups[b64].Sessions.push(s);
	});
	return groups;
}

function showDefenses(defs, us) {
	//split students and potential juries
	teachers = us[0].filter(function(u) { return level(u.Role) != STUDENT_LEVEL});
	teachers.sort(function(a, b) { a.Person.Lastname.localeCompare(b.Person.Lastname);});
	if (defs[0] == null) {
		defs[0] = [];
	}
	//catch the placed students
	defs[0].forEach(function(def) {
		def.Defenses.forEach(function (d) {
			placed.push(d.Student.User.Person.Email);
		});
	});
	//Print the defenses
	$("#cnt").render("defense-planner-2", groupById(defs[0]), function() {
		ui();
		$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
		userMailing(".defense-groups");
	});
}

function newDefenseSlot(elmt, id, room) {
	$(elmt).closest(".panel").addClass("active");
	unPlaced  = [];
	internships().done(function (ints) {
		ints.forEach(function (i) {
			if (!placed.includes(i.Convention.Student.User.Person.Email)) {
				unPlaced.push(i.Convention.Student);
			}
		});
		unPlaced.sort(function (a,b) {return a.Major.localeCompare(b.Major)});
		var def = {
			Room: room,
			Id: id,
			SessionId: id, //to please edition
			Students: unPlaced,
			Local: true,
			Public: true,
			Time: new Date(),
		};
		$("#modal").render("defense-editor",def,
			function () {
				$("#defense-time").datetimepicker({stepping: 30, timeZone: 'Europe/Paris'});
				showModal()
			}
		);
	});
}

function showNewSession(id) {
	var dta = {};
	if (id) {
		dta.Id = id;
	}
	$("#modal").render("defense-session-creator", dta, function() {
		showModal(function () {
			$("#modal").find("#date").datetimepicker({stepping: 30, timeZone: 'Europe/Paris'});
		});
	});
}

function addDefenseSession(id) {
	if (!id) {
		if (empty("#date","#room")) {
			return;
		}
		var period = $("input[name='period']:checked").val();
		var date = moment($("#date").data("DateTimePicker").date());
		id = date.format("DD MMM") + " " + period;
	}
	var dta = {
		Id: id,
		Room: $("#room").val(),
	}
	postDefense(dta).done(function (def) {
		var b64 = window.btoa(def.Id);
		//existst, append
		if($("#" +b64).length) {
			$("#" + b64).find(".row").append(Handlebars.partials["defense-session-3"](def));
		} else {
			//create the group, render it
			$(".defense-groups").append(Handlebars.partials["session-group"](groupById([def])[b64]));

		}
		$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
		hideModal();
	}).fail(notifyError);
}

function rmDefenseSession(room, id, td) {
	delDefenseSession(room, id).done(function () {
		td.closest(".panel").remove();
		var b64 = window.btoa(id);
		if(!$("#" +b64).find(".panel").length) {
			$("#" +b64).remove();
		}
	}).fail(notifyError);
}


function setStudentDefense(room, id) {
	if (empty("#defense-time")) {
		return;
	}
	var em = $("#student-selecter").val();
	var public = !$("#private").prop("checked");
	var local = !$("#remote").prop("checked");
	var time = moment($('#defense-time').data("DateTimePicker").date());
	var date = moment(id.substring(0, id.length - 3),"DD MMM");
	date.set('year',new Date().getFullYear());
	date.set('hour',time.get('hour'));
	date.set('minute',time.get('minute'));
	date.set('second',0);
	date.set('millisecond',0);
	var def = {
		Public: public,
		Local: local,
		Time: date,
	}
	postStudentDefense(room, id, em, public, local, date).done(function () {
		getDefenseSession(room, id).done(function (s) {
			$(".active").replaceWith(Handlebars.partials['defense-session-3'](s))
			$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
			$(".active").removeClass("active");
			hideModal();
			placed.push(em);
		});
	}).fail(notifyError);
}

function updateStudentDefense(em, room, id) {
	if (empty("#defense-time")) {
		return;
	}
	var public = !$("#private").prop("checked");
	var local = !$("#remote").prop("checked");
	var time = moment($('#defense-time').data("DateTimePicker").date());
	var date = moment(id.substring(0, id.length - 3),"DD MMM");
	date.set('year',new Date().getFullYear());
	date.set('hour',time.get('hour'));
	date.set('minute',time.get('minute'));
	date.set('second',0);
	date.set('millisecond',0);
	var def = {
		Public: public,
		Local: local,
		Time: date,
	}
	putStudentDefense(em, public, local, date).done(function (def) {
		getDefenseSession(room, id).done(function (s) {
			$(".active").replaceWith(Handlebars.partials['defense-session-3'](s))
			$(".active").removeClass("active");
			$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
			hideModal();
		});
	}).fail(notifyError);
}

function editStudentDefense(elmt,stu, room, id) {
	$(elmt).closest(".panel").addClass("active");
	getDefense(stu).done(function (def) {
		$("#modal").render("defense-editor",def, function () {
			$("#defense-time").datetimepicker({stepping: 30, timeZone: 'Europe/Paris'})
			$("#modal").find(".btn-primary").attr("onclick", "updateStudentDefense('"+ stu + "','" + room + "','" + id + "')").html("Update");
			showModal(function() {
				$("#defense-time").data("DateTimePicker").date(moment(def.Time));
			})
		});
	});
}


function delStudentDefense(stu, room, id) {
	rmStudentDefense(stu).done(function () {
		getDefenseSession(room, id).done(function (def) {
			s = $(".active");
  			s.replaceWith(Handlebars.partials['defense-session-3'](def));
  			$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
  			$(".active").removeClass("active");
  			hideModal();
		});
	});
}

function addDefenseJury(room, id, elmt) {
	var em = $(elmt).prev().val();
	newDefenseJury(room, id, em).done(function(def) {
		session = $(elmt).closest(".panel");
  		session.replaceWith(Handlebars.partials['defense-session-3'](def));
  		$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
	}).fail(notifyError);
}

function delDefenseJury(elmt, room, id, em) {
	rmDefenseJury(room, id, em).done(function (def) {
		session = $(elmt).closest(".panel");
  		session.replaceWith(Handlebars.partials['defense-session-3'](def));
  		$(".jury-selecter").html(Handlebars.helpers["optionUsers"](teachers).string);
	}).fail(notifyError);
}