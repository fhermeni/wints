/**
 * Created by fhermeni on 03/07/2014.
 */



Handlebars.logger.level = 0;
$.handlebars({
	templatePath: '/assets/hbs',
	templateExtension: 'hbs',
	partialPath: '/assets/hbs',
	partials: [/*'person', 'users-user', 'placement-student', 'convention-student', 'convention-editor', 'student-dashboard-report', 'student-dashboard-company', 'student-dashboard-contacts', 'student-dashboard-alumni', 'tutored-student', 'watchlist-student'*/]
});

Handlebars.registerHelper('len', function(a) {
	if (a.constructor === Array) {
		return a.length;
	}
	return Object.keys(a).length;
});

Handlebars.registerHelper('dateFmt', function(d, fmt, loc, none, foo) {
	if (!d) {
		//None test for background compatibility
		if (typeof none === "string") {
			return none;
		}
		return "-";
	}
	if (typeof loc === "string") {
		return moment.tz(d,loc).format(fmt);
	}
	return moment(d).format(fmt);
});

var positions = {
	"looking" : "Looking for a job",
	"sabbatical" : "Sabattical leave",
	"company" : "Working in a company",
	"entrepreneurship" : "Entrepreneurship",
	"study" : "Pursuit of higher education"
};

Handlebars.registerHelper('alumniPosition', function(r) {
	return positions[r];
});


var roles = ["student", "tutor", "major", "head", "admin", "root"];

Handlebars.registerHelper('roleLevel', function(r) {
	return level(r);
});
Handlebars.registerHelper('optionRoles', function(r) {
	var res = "";
	roles.forEach(function(role) {
		if (role == "major") {
			config.Majors.forEach(function(m) {
				var full = role + "-" + m;
				var selected = full == r ? "selected" : "";
				res += ("<option value='" + full + "' " + selected + ">" + full + "</option>");
			});
		} else {
			var selected = role == r ? "selected" : "";
			res += "<option value='" + role + "' " + selected + ">" + role + "</option>";
		}

	});
	return new Handlebars.SafeString(res);
});

Handlebars.registerHelper('optionMajors', function(m) {
	var b = "";
	if (!m) {
		b += "<option selected>?</option>";
	}
	config.Majors.forEach(function(o) {
		var selected = m == o ? " selected " : "";
		b += "<option value='" + o + "' " + selected + " >" + o + "</option>";
	});
	return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('optionPromotions', function(p) {
	var b = "";
	if (!p) {
		b += "<option selected>?</option>";
	}
	config.Promotions.forEach(function(o) {
		var selected = p == o ? " selected " : "";
		b += "<option value='" + o + "' " + selected + " >" + o + "</option>";
	});
	return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('daysSince', function(d1, d2) {
	if (!d2) {
		d2 = moment();
	} else {
		d2 = moment(d2);
	}
	var duration = moment.duration(moment(d1).diff(d2));
	return -Math.floor(duration.asDays());
});


Handlebars.registerHelper('optionUsers', function(users, u) {
	var b = "";
	users.forEach(function(o) {
		b += "<option value='" + o.Person.Email + "'" + ((u != undefined && o.Person.Email == u.Person.Email) ? "selected" : "") + ">" + o.Person.Lastname + ", " + o.Person.Firstname + "</option>";
	});
	return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('optionStudents', function(students, s) {
	var b = "";
	students.forEach(function(o) {
		b += "<option value='" + o.User.Person.Email + "'>" + o.User.Person.Lastname + ", " + o.User.Person.Firstname + "(" + o.Major + "/" + o.Promotion + ")</option>";
	});
	return new Handlebars.SafeString(b);
});

function editableRoles() {
	var res = [];
	for (i = 2; i < roles.length; i++) {
		res.push({
			"value": i,
			"text": roles[i]
		});
	}
	return res;
}

function editablePromotions() {
	var res = [];
	config.Promotions.forEach(function(p) {
		res.push({
			"value": p,
			"text": p
		});
	});
	return res;
}

function editableMajors() {
	var res = [];
	config.Majors.forEach(function(p) {
		res.push({
			"value": p,
			"text": p
		});
	});
	return res;
}

Handlebars.registerHelper('student', function(r, opts) {
	if (r <= 1)
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifEq', function(n, val, opts) {
	if (n == val)
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifLate', function(d, opts) {
	if (!d || moment(d).isBefore(moment()))
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifAfter', function(d1, d2, opts) {
	if ((!d1 && moment().isAfter(d2)) || moment(d1).isAfter(d2)) {
		return opts.fn(this);
	} else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifRole', function(d, opts) {
	if (level(myself.Role) >= d)
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifManage', function(r, opts) {
	if (myself.Person.Email == r.Tutor || level(myself.Role) >= ADMIN_LEVEL) {
		return opts.fn(this);
	} else {
		return opts.inverse(this);
	}
});

Handlebars.registerHelper('fullname', function(p, pretty) {
	if (pretty) {
		return p.Lastname.capitalize() + ", " + p.Firstname.capitalize();
	}
	return p.Lastname + ", " + p.Firstname;
});

Handlebars.registerHelper('grade', function(r) {
	var buf;
	if (!r.Reviewed) {
		buf = "-";
	} else if (!r.ToGrade) {
		buf= "&#10003;";
	} else {
		var duration = moment.duration(moment(r.Delivery).diff(moment(r.Deadline)));
		var days = Math.floor(duration.asDays());
		if (days <= 0) {
			buf = ""+r.Grade;
		} else {
			buf = "<span title='Tutor: " + r.Grade + "; Late penalty: -" + (config.LatePenalty * days) + "'>" + netGrade(r) + "</span>";
		}
	}
	return new Handlebars.SafeString(buf);
});

Handlebars.registerHelper('defenseGrade', function(d) {
	var buf = "";
	if (!d.Time) {
		return "-"; //no defense planned
	}
	if (d.Grade == -1) {
		return "?"
	}
	return d.Grade;
});

/*
	if (d.Grade == -1) {
		return "<td data-text='-1'>?</span>"; //not passed
	} else {
		if (d.Grade < 10) {
			return "<td data-text='" + d.Grade + "' class='bg-danger'>" + d.Grade + "</span>"; //baad
		}
		return "<td data-text='" + d.Grade + "' class='bg-danger'>" + d.Grade + "</span>"; //baad
	}
*/

Handlebars.registerHelper('survey', function(s, stu) {
	var value = -10;
	var grade = "-";
	var bg = "";
	if (s.Delivery) {
		var mark = s.Cnt["__MARK__"];
		if (mark == "false") {
			bg = "bg-danger";
			value = -1;
			grade = "x";
		} else if (mark == "true") {
			value = 1;
			grade = "&#10003;";
		} else {
			value = mark;
			grade = mark;
			if (mark < 10) {
				bg = "bg-danger";
			}
		}
		var buf = '<td class="' + bg + ' text-center" data-text="' + value + '">';
		buf += "<a target='_blank' href='/survey?kind=" + s.Kind + "&student=" + stu  +"'>";
		buf += grade;
		buf += '</a></td>';
		return new Handlebars.SafeString(buf);
	}

	var delay;
	if (moment(s.Deadline).isBefore(new Date())) {
			bg = "bg-danger";
			delay = Math.floor(moment.duration(moment().diff(moment(s.Deadline))).asDays());
			value = -1000 + delay;
			grade = "<i class='glyphicon glyphicon-time'></i> " + delay + " d.";
	} else if (moment(s.Invitation).isBefore(new Date())) {
			bg = "bg-warning";
			delay = Math.floor(moment.duration(moment().diff(moment(s.Invitation))).asDays());
			value = -100 + delay;
			grade = "<i class='glyphicon glyphicon-time'></i> " + delay + " d.";
	}

	var buf = '<td class="' + bg + ' text-center" data-text="' + value + '">';
	buf += grade;
	buf += '</td>';
	return new Handlebars.SafeString(buf);
});

Handlebars.registerHelper('report', function(r, em, cb) {
	var bg = "";
	var value = "-999999", grade;
	var cnt = '-';
	if (r.Delivery && moment(r.Deadline).isBefore(moment()) && !r.Reviewed) {
		//uploaded, deadline passed, not review
		bg = "info";
		var delay = reviewDelay(r);
		value = -100 + delay;
		if (delay === 0) {
			delay = "< 1";
		}
		cnt = "<i class='glyphicon glyphicon-time'></i> " + delay + " d.";
	}

	if (!r.Delivery && moment(r.Deadline).isBefore(new Date())) {
		bg = "danger";
		var delay = deadlineDelay(r);
		value = -50 + delay;
		if (delay === 0) {
			delay = "< 1";
		}
		cnt = "<i class='glyphicon glyphicon-time'></i> " + delay + " d.";
	}
	if (r.Reviewed) {
			grade = netGrade(r);
			cnt = Handlebars.helpers['grade'](r).string;
		if (r.ToGrade) {
			value = grade;
		} else {
			value = -1;
		}
	}
	if (r.Reviewed && r.ToGrade && grade < 10) {
		bg = "danger";
	}

	var buf = '<td data-text=' + value + ' onclick="showReport(\'' + em + '\', \'' + r.Kind + '\')" class="click ' + bg + ' text-center">';

	buf = buf + cnt + '</td>';
	return new Handlebars.SafeString(buf);
});