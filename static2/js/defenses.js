var active = undefined;
var teachers = []
var students = []
Array.prototype.remove = function(val) {
	var i = this.indexOf(val);
	return i > -1 ? this.splice(i, 1) : [];
};

var allDefenses = {}
var defenseSessions = []


function hash(s) {
	return s.Date.format("DDhha") + s.Room.replace("+", "plus").replace("-", "minus")
}

function getSession(sid) {
	var s = undefined;
	defenseSessions.forEach(function(ss) {
		if (hash(ss) == sid) {
			s = ss;
			return false;
		}
	});
	return s;
}

function sortByMajor(a, b) {
	if (a.Major < b.Major) return -1;
	if (a.Major > b.Major) return 1;
	return 0;
}

function showDefenses() {
	users(function(data) {
		teachers = [];
		data.filter(function(u) {
			if (u.Role != 0) {
				teachers.push(u);
			}
		});
		students = interns.slice()
		students.sort(sortByMajor);
		defenses(function(defs) {
			var html = Handlebars.getTemplate("defense-editor")();
			var root = $("#cnt");
			root.html(html);

			//prepare the modal
			var html = Handlebars.getTemplate("defenseSessionEditor")({
				Room: "TBA"
			});
			$('.date').datetimepicker();
			root.find(".date").datepicker({
				format: 'd M yyyy',
				autoclose: true,
				minViewMode: 0,
				weekStart: 1
			}).on("changeDate", function(e) {
				setDefenseDay(id, e.date)
			})
			var root = $("#modal");
			root.html(html)
			load(defs)
			students.sort(sortByMajor);
			$('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)
			$("#cnt").find("ul.students").sortable({
				connectWith: "sortable"
			}).bind('sortupdate', updateSessionOrder);
			if (active) {
				$('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)
			}
		})
	});
}

function updateSessionOrder(e, li) {
	var sid = $(li.item).closest(".session").attr("id")
	s = getSession(sid)
	s.Students = []

	$(li.item).closest("ul").find("li").each(function(i, input) {
			em = $(input).find("input").data("email")
			if (em) {
				allDefenses[em].Offset = i
			}
			s.Students.push(em);
		})
		//Get the emails
}

function addDefenseSession() {
	var rooms = $("#room").val().split(" ");
	var d = $("#date").data("DateTimePicker").date()
	var sessions = []
	rooms.forEach(function(r) {
		var s = {
			Pause: -1,
			Room: r,
			Date: d,
			Students: [],
			Juries: []
		};
		if ($("#" + hash(s)).length > 0) {
			$("#room").closest(".form-group").addClass("has-error");
		} else {
			sessions.push(s)
		}
	});
	if (rooms.length == sessions.length) {
		$("#room").closest(".form-group").removeClass("has-error");
		$("#modal").modal('hide');
		sessions.forEach(function(s) {
			drawSession(s);
			defenseSessions.push(s);
			activeSession(hash(s));
		})
	}
}

function rmSession() {
	s = getSession(active);
	activeSession(active)
	s.Students.forEach(function(stu) {
		if (stu) {
			students.push(getInternship(stu));
			delete allDefenses[stu]
		}
	});
	$("#" + hash(s)).remove()
	defenseSessions.remove(s)
	students.sort(sortByMajor);
	$('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)
}


function drawSession(s) {
	var html = "<div id='" + hash(s) + "' class='col-md-3 session'>" +
		"<div class='defense-panel panel'>" +
		"<div class='panel-heading' onclick='activeSession(\"" + hash(s) + "\")'>" +
		"<div class='panel-title'><i class='glyphicon glyphicon-calendar'> </i> <span class='date'>" + s.Date.format("D MMM - HH:mm") + "</span> &nbsp;" +
		"<span class='where'><i class='glyphicon glyphicon-map-marker'> </i> <span class='room'>" + s.Room + "</span></span></div>" +
		"</div>" +
		"<div class='panel-body'>" +
		"<label>Agenda</label> " +
		"<ul class='fn students sortable list-unstyled'></ul>" +
		"<label>Jury</label> " +
		"<ul class='fn juries list-unstyled'></ul>" +
		"</div>" +
		"</div>";
	var rId = s.Date.month() + "" + s.Date.date()
	var period = s.Date.hour() < 12 ? "am" : "pm";

	var day = $("#" + rId);
	if (day.length == 0) {
		day = $("#days").append("<div id='" + rId + "' class='day'><div class='row am'></div><div class='row pm'></div></div>").children().last();
	}
	var row = day.find("." + period)
	row.append(html)
}

function saveDefenseSession() {
	var newRoom = $("#room").val();
	var newDate = $("#date").data("DateTimePicker").date();
	var newPeriod = newDate.hour() < 12 ? "am" : "pm";
	var newDay = newDate.date();

	var oldSession = getSession(active)
	var newSession = {
		Pause: -1,
		Room: newRoom,
		Date: newDate,
		Juries: oldSession.Juries.slice(),
		Students: oldSession.Students.slice()
	};
	if (hash(oldSession) != hash(newSession) && $("#" + hash(newSession)).length > 0) {
		//new ID but it already exists
		$("#room").closest(".form-group").addClass("has-error");
		console.log("no way")
		return
	}
	$("#room").closest(".form-group").removeClass("has-error");
	$("#modal").modal('hide');

	var oldPeriod = oldSession.Date.hour() < 12 ? "am" : "pm";

	if (oldPeriod != newPeriod || newDay != oldSession.Date.date()) {
		rmSession(oldSession)
		drawSession(newSession)
	} else {
		console.log("new stuff")
			//Just the room change,         
		var d = $("#" + active);
		d.find(".room").html(newRoom)
		d.find(".date").html(newDate.format("D MMM - HH:mm"))
		d.attr("id", hash(newSession))
		d.attr("onclick", "activeSession(\"" + hash(newSession) + "\")")
	}

}

function showNewSession() {
	var root = $("#modal");
	$("#room").val("TBA")
	root.find("button").attr("onclick", "addDefenseSession()")
	d = new Date()
	if (defenseSessions.length == 0) {
		d.setMinutes(0)
	} else {
		d = defenseSessions[defenseSessions.length - 1].Date
	}
	root.find('#date').datetimepicker({
		inline: true,
		sideBySide: true,
		format: "dd/mm/yyyy HH:mm"
	}).data("DateTimePicker").date(d);
	root.modal('show');
}

function showEditSession() {
	s = getSession(active)
	var root = $("#modal");
	$("#room").val(s.Room)
	root.find("button").attr("onclick", "saveDefenseSession()")
	root.find('#date').datetimepicker({
		inline: true,
		sideBySide: true,
		format: "dd/mm/yyyy HH:mm"
	}).data("DateTimePicker").date(s.Date)
	root.modal('show');
}

function activeSession(sid) {

	if (active) {
		$("#" + active).find(".panel").removeClass("panel-success")
	}
	if (active != sid) {
		active = sid
		$("#" + sid).find(".panel").addClass("panel-success");
		$("#cnt").find(".activable").removeClass("disabled")
		$('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)
	} else {
		$("#cnt").find(".activable").addClass("disabled")
		$('#jury-selecter').html("")
		active = undefined
	}
}

function removeByEmail(arr, em) {
	return arr.filter(function(u) {
		return u.Email != em;
	});
}

function addStudent(d) {
	if (!active) {
		return
	}
	var em = $("#student-selecter").val()
	var i = getInternship(em);

	var def = {
		Remote: false,
		Private: false,
		Offset: $("#" + active).find("ul.students li").length
	}
	allDefenses[em] = def;
	drawStudent(em)
	$("#cnt").find("ul.students").sortable({
		connectWith: "sortable"
	}).bind('sortupdate', updateSessionOrder);
	students.remove(i)
	$('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)
	$('#student-selecter').selecter("destroy");
	$('#student-selecter').selecter("update");
	s = getSession(active)
	s.Students.push(em)
}

function drawStudent(em) {
	var i = getInternship(em);
	var def = allDefenses[em];
	var glyphRemote = "glyphicon-picture"
	var glyphPrivate = "glyphicon-eye-open"
	if (def.Remote) {
		glyphRemote = "glyphicon-facetime-video"
	}
	if (def.Private) {
		glyphPrivate = "glyphicon-eye-close text-danger"
	}
	var html = "<li>" +
		"<input type='checkbox' data-email='" + em + "'>" +
		" <a class='fn' onclick='showInternship(\"" + em + "\")'>" + Handlebars.helpers.abbrvFullname(i.Student) + " (" + i.Major + ")</a>" +
		" <i class='glyphicon " + glyphRemote + "' onclick='toggleRemote(this, \"" + em + "\")'></i> <i class='glyphicon " + glyphPrivate + "' onclick='togglePrivate(this,\"" + em + "\")'></i>" +
		" &nbsp; <i class='glyphicon glyphicon-remove-circle pull-right' onclick='removeStudent(this,\"" + em + "\")'></i>" +
		"</li>";
	$("#" + active).find("ul.students").append(html).find(":checkbox").last().icheck()
}


function addPause() {
	s = getSession(active)
	if (s) {
		s.Pause = s.Students.length
		s.Students.push(undefined)
		drawPause()
	}
}

function drawPause() {
	if (active) {
		s = getSession(active)
		var html = "<li><i>pause</i> <i onclick='rmPause(this)' class='glyphicon glyphicon-remove-circle pull-right'></i></li>";
		//$("#"+active).find("ul.students").append(html)    
		$("#" + active).find("ul.students li:nth-child(" + s.Pause + ")").after(html)
		$("#" + active).find("ul.students").sortable().bind('sortupdate', updateSessionOrder);
	}
}

function rmPause(p) {
	$(p).closest("li").remove()
}

function toggleRemote(i, em) {
	allDefenses[em].Remote = !allDefenses[em].Remote;
	var j = $(i)
	if (j.hasClass("glyphicon-picture")) {
		j.removeClass("glyphicon-picture").addClass("glyphicon-facetime-video")
	} else {
		j.addClass("glyphicon-picture").removeClass("glyphicon-facetime-video")
	}
}

function togglePrivate(i, em) {
	allDefenses[em].Private = !allDefenses[em].Private;
	var j = $(i)
	if (j.hasClass("glyphicon-eye-open")) {
		j.removeClass("glyphicon-eye-open").addClass("glyphicon-eye-close").addClass("text-danger")
	} else {
		j.addClass("glyphicon-eye-open").removeClass("glyphicon-eye-close").removeClass("text-danger")
	}
}

function removeStudent(b, em) {
	$(b).closest("li").remove();
	s = getSession(active);
	delete allDefenses[em]
	s.Students.remove(em);
	i = getInternship(em)
	students.push(i);
	students.sort(sortByMajor);
	$('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)
}


function overlap(d1, d2) {
	var day1 = d1.date();
	var day2 = d2.date();
	var p1 = d1.hour() < 12 ? "am" : "pm";
	var p2 = d2.hour() < 12 ? "am" : "pm";
	return day1 == day2 && p1 == p2;
}

function addJury(d) {
	if (!active) {
		return
	}
	var em = $("#jury-selecter").val()
	drawJury(em)
	s = getSession(active)
	s.Juries.push(em);
	$('#jury-selecter').find('option:selected', this).remove();
}

function drawJury(em) {
	teachers.forEach(function(t) {
		if (t.Email == em) {
			var html = "<li>" +
				"<input type='checkbox' data-toggle='checkbox' class='icheckbox_flat check_all' data-email='" + em + "'/>" +
				" <a href='mailto:" + em + "'>" + Handlebars.helpers.abbrvFullname(t) + "</a>" +
				" <i class='glyphicon glyphicon-remove-circle pull-right' onclick='removeJury(this,\"" + em + "\")'></i>" +
				"</li>";
			$("#" + active).find("ul.juries").append(html).find(":checkbox").icheck();

			return false;
		}
	});
}

function removeJury(b, em) {
	s = getSession(active);
	s.Juries.remove(em);
	$(b).closest("li").remove();
	$('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)
}

function availableTeachers(sid) {
	var s = getSession(sid);
	var ok = teachers.slice();
	defenseSessions.forEach(function(ss) {
		if (overlap(ss.Date, s.Date)) {
			ss.Juries.forEach(function(t) {
				ok = removeByEmail(ok, t)
			})
		}
	});
	return ok
}

function save() {
	ss = [];
	defenseSessions.forEach(function(session) {
		s = {
			Room: session.Room,
			Date: session.Date.toDate(),
			Defenses: [],
			Juries: []
		}
		session.Juries.forEach(function(j) {
			s.Juries.push({
					Email: j
				}) //Only the email is usefull
		})
		session.Students.forEach(function(stu) {
			if (stu) {
				d = allDefenses[stu]
				d["Student"] = {
					Email: stu
				}
				s.Defenses.push(d)
			}
		})
		ss.push(s)
	})
	postDefenses(ss)
}

function load(defs) {
	active = undefined
	allDefenses = {}
	defenseSessions = []
	defs.forEach(function(s) {
		s.Date = moment(s.Date)
		defenseSessions.push(s);
		drawSession(s);
		activeSession(hash(s));
		juries = []
		s.Juries.forEach(function(j) {
			juries.push(j.Email)
			drawJury(j.Email)
		});
		s.Juries = juries
		i = 0 //rank
		s.Students = []
		s.Defenses.forEach(function(def) {
			if (def.Offset != s.Students.length) {
				addPause()
			}
			s.Students.push(def.Student.Email);
			students.remove(getInternship(def.Student.Email))
			allDefenses[def.Student.Email] = def;
			drawStudent(def.Student.Email)
		});
		delete s.Defenses
	});
}