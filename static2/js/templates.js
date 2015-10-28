/**
 * Created by fhermeni on 03/07/2014.
 */

/*
moment.locale('fr', {
	months: "janvier_février_mars_avril_mai_juin_juillet_août_septembre_octobre_novembre_décembre".split("_"),
	monthsShort: "janv._févr._mars_avr._mai_juin_juil._août_sept._oct._nov._déc.".split("_"),
	weekdays: "dimanche_lundi_mardi_mercredi_jeudi_vendredi_samedi".split("_"),
	weekdaysShort: "dim._lun._mar._mer._jeu._ven._sam.".split("_"),
	weekdaysMin: "Di_Lu_Ma_Me_Je_Ve_Sa".split("_"),
	longDateFormat: {
		LT: "HH:mm",
		LTS: "HH:mm:ss",
		L: "DD/MM/YYYY",
		LL: "D MMMM YYYY",
		LLL: "D MMMM YYYY LT",
		LLLL: "dddd D MMMM YYYY LT"
	},
	calendar: {
		sameDay: "[Aujourd'hui à] LT",
		nextDay: '[Demain à] LT',
		nextWeek: 'dddd [à] LT',
		lastDay: '[Hier à] LT',
		lastWeek: 'dddd [dernier à] LT',
		sameElse: 'L'
	},
	relativeTime: {
		future: "dans %s",
		past: "il y a %s",
		s: "quelques secondes",
		m: "une minute",
		mm: "%d minutes",
		h: "une heure",
		hh: "%d heures",
		d: "un jour",
		dd: "%d jours",
		M: "un mois",
		MM: "%d mois",
		y: "une année",
		yy: "%d années"
	},
	ordinalParse: /\d{1,2}(er|ème)/,
	ordinal: function(number) {
		return number + (number === 1 ? 'er' : 'ème');
	},
	meridiemParse: /PD|MD/,
	isPM: function(input) {
		return input.charAt(0) === 'M';
	},
	// in case the meridiem units are not separated around 12, then implement
	// this function (look at locale/id.js for an example)
	// meridiemHour : function (hour, meridiem) {
	//     return  0-23 hour, given meridiem token and hour 1-12 
	// },
	meridiem: function(hours, minutes, isLower) {
		return hours < 12 ? 'PD' : 'MD';
	},
	week: {
		dow: 1, // Monday is the first day of the week.
		doy: 4 // The week that contains Jan 4th is the first week of the year.
	}
});*/

$.handlebars({
	templatePath: '/static/hbs/',
	templateExtension: 'hbs',
	partialPath: '/static/hbs',
	partials: ['person', 'users-user', 'placement-student', 'convention-student', 'convention-editor', 'student-dashboard-report', 'student-dashboard-company', 'student-dashboard-contacts', 'student-dashboard-alumni', 'tutored-student', 'watchlist-student']
});

Handlebars.registerHelper('len', function(a) {
	if (a.constructor === Array) {
		return a.length;
	}
	return Object.keys(a).length;
});

Handlebars.registerHelper('dateFmt', function(d, fmt) {
	if (!d) {
		return "-";
	}
	return moment(d).format(fmt);
});

var roles = ["none", "student", "tutor", "major", "admin", "root"];

Handlebars.registerHelper('role', function(r) {
	return roles[r];
});

Handlebars.registerHelper('optionRoles', function(r) {
	var res = "";
	for (i = 2; i < roles.length; i++) {
		if (r == i) {
			res += "<option value='" + i + "' selected>" + roles[i] + "</option>";
		} else {
			res += "<option value='" + i + "'>" + roles[i] + "</option>";
		}
	}
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

Handlebars.registerHelper('optionUsers', function(users, u) {
	var b = "";
	users.forEach(function(o) {
		b += "<option value='" + o.Person.Email + "'" + (o.Person.Email == u.Person.Email ? "selected" : "") + ">" + o.Person.Lastname + ", " + o.Person.Firstname + "</option>";
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
	})
	return res;
}

function editableMajors() {
	var res = [];
	config.Majors.forEach(function(p) {
		res.push({
			"value": p,
			"text": p
		});
	})
	return res;
}

Handlebars.registerHelper('roleSelect', function(m) {
	var b = "";
	for (i = 2; i < roles.length; i++) {
		b += "<option value='" + i + "'>" + roles[i] + "</option>";
	}
	return new Handlebars.SafeString(b);
});

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
	if (moment(d).isBefore(moment()))
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('ifRole', function(d, opts) {
	if (myself.Role >= d)
		return opts.fn(this);
	else
		return opts.inverse(this);
});

Handlebars.registerHelper('fullname', function(p, pretty) {
	if (pretty) {
		return p.Lastname.capitalize() + ", " + p.Firstname.capitalize();
	}
	return p.Lastname + ", " + p.Firstname;
});

Handlebars.registerHelper('survey', function(s, stu) {
	var value = -10;
	var grade = "-";
	var bg = "";
	if (s.Delivery) {
		if (s.Kind == "midterm") {
			if (Object.keys(s.Cnt).length > 0) {
				if (s.Cnt[19] == "false") {
					value = -1;
					grade = "x";
					bg = "bg-danger";
				} else {
					value = 1;
					grade = "&#10003;";
					bg = "bg-success";
				}
			}
		} else { //final
			if (Object.keys(s.Cnt).length > 0) {
				value = s.Cnt[q17];
				grade = s.Cnt[q17];
				bg = grade >= 10 ? "bg-success" : "bg-danger";
			}
		}
	} else if (moment(s.Deadline).isBefore(new Date())) {
		bg = "bg-warning";
		value = -20;
	}
	var buf = '<td class="click ' + bg + ' text-center" data-text="' + value + '" onclick="showSurvey(\'' + stu + '\', \'' + s.Kind + '\')">' + grade + '</td>';
	return new Handlebars.SafeString(buf);
});

Handlebars.registerHelper('report', function(r, em, cb) {
	var value = -10;
	var grade = "-";
	var bg = "";
	/*if (s.Delivery) {
		if (s.Kind == "midterm") {
			if (Object.keys(s.Cnt).length > 0) {
				if (s.Cnt[19] == "false") {
					value = -1;
					grade = "x";
					bg = "bg-danger";
				} else {
					value = 1;
					grade = "&#10003;";
					bg = "bg-success";
				}
			}
		} else { //final
			if (Object.keys(s.Cnt).length > 0) {
				value = s.Cnt[q17];
				grade = s.Cnt[q17];
				bg = grade >= 10 ? "bg-success" : "bg-danger";
			}
		}
	}*/
	var buf = '<td onclick="showReport(\'' + em + '\', \'' + r.Kind + '\')" class="click ' + bg + 'text-center" data-text="' + value + '">' + grade + '</td>';
	return new Handlebars.SafeString(buf);
});


/*Handlebars.registerHelper('reportStatus', function(r) {

	var passed = (new Date(Date.parse(r.Deadline)).getTime() + 86400 * 1000) < new Date().getTime()
	var style = "bg-link";

	//Deadline passed, nothing
	if (passed && r.Grade == -2) {
		style = "bg-warning";
	} else if (r.Grade == -1) {
		//waiting for beging reviewed
		style = "bg-info";
	} else if (r.ToGrade && r.Grade >= 0 && r.Grade < 10) {
		style = "bg-danger";
	} else if ((!r.ToGrade && r.Grade >= 0) || r.Grade >= 10) {
		style = "bg-success";
	}

	return style;
	return "";
});

Handlebars.registerHelper('gradeAnnotation', function(r) {
	var passed = (new Date(Date.parse(r.Deadline)).getTime() + 86400 * 1000) < new Date().getTime()
	var d = 0
	if (passed && r.Grade == -2) {
		var d = nbDayLates(r.Deadline, new Date()) //moment(r.Deadline).dayOfYear() - moment(new Date).dayOfYear()
		return -2000 + d;
	} else if (passed & r.Grade == -1) {
		var d = nbDayLates(r.Delivery, new Date()) //moment(r.Delivery).dayOfYear() - moment(new Date).dayOfYear()
		return -1000 + d
	} else if (r.Grade >= 0) {
		return r.Grade
	}
	return -9999
});

Handlebars.registerHelper('reportGrade', function(r) {
	return "";
	/*var passed = (new Date(r.Deadline).getTime() + 86400 * 1000) < new Date().getTime()
	if (!r.ToGrade) {
		if (passed && r.Grade == -2) {
			return fmtNbDayLate(nbDayLates(new Date(), r.Deadline))
		} else if (passed && r.Grade == -1) {
			return fmtNbDayLate(nbDayLates(new Date(), r.Delivery))
		} else {
			return new Handlebars.SafeString("<i title='no grade needed'>n/a</i>");
		}
	}
	if (r.Grade == -2) {
		if (passed) {
			return fmtNbDayLate(nbDayLates(new Date(), r.Deadline))
		} else {
			return "-";
		}
	} else if (r.Grade == -1) {
		return fmtNbDayLate(nbDayLates(new Date(), r.Delivery))
	}
	return r.Grade;
});
Handlebars.registerHelper('gradeInput', function(r) {
	if (!r.Gradeable || r.Grade <= 0) {
		return "";
	}
	return r.Grade;
});

Handlebars.registerHelper('studentGrade', function(r) {
	var passed = (new Date(r.Deadline).getTime() + 86400 * 1000) < new Date().getTime()
	if (!r.Uploaded) {
		if (r.Grade == 0) {
			return r.Grade
		} else if (passed) {
			return "deadline passed. Hurry!";
		} else {
			return "";
		}
	} else {
		if (r.ToGrade) {
			if (r.Grade == -1) { //waiting for review
				return "?";
			} else {
				return r.Grade;
			}
		} else {
			if (r.Grade == -1) { //waiting for review
				return "(waiting for the review)";
			} else {
				return "n/a";
			}
		}
	}
});

Handlebars.registerHelper('URIemails', function(students) {
	var l = [];
	students.forEach(function(s) {
		if (s && s.P) {
			l.push(s.P.Email);
		}
	});
	return encodeURI(l.join(","));
});


Handlebars.registerHelper('student', function(g) {
	if (!g || g.length == 0) {
		return "break";
	}

	var buf = g.P.Firstname + " " + g.P.Lastname;
	if (defenses.private[g.P.Email]) {
		buf += " <span class='glyphicon glyphicon-eye-close'></span>";
	}
	if (defenses.visio[g.P.Email]) {
		buf += " <span class='glyphicon glyphicon-facetime-video'></span>";
	}
	var c = getConvention(g.P.Email);
	if (!c.SupReport.IsIn) {
		buf += " <span class='glyphicon glyphicon-file alert-danger'></span>";
	}
	return new Handlebars.SafeString(buf);
});

Handlebars.registerHelper("offset", function(from, offset, tz) {
	if (tz === undefined) {
		tz = true;
	}
	var d = moment(from)
	if (tz) {
		d = d.tz('Europe/Paris');
	}
	d = d.add(offset * 30, "minutes")
	return d.format("HH:mm") + "-" + d.add(30, "minutes").format("HH:mm");
})

Handlebars.registerHelper("offset_local", function(from, offset, tz) {
	var d
	if (typeof from == "string") {
		console.log("string");
		d = moment(new Date(from));
	} else {
		d = moment(from);
	}
	d = d.add(offset * 30, "minutes")
	return d.format("dddd D MMMM HH:mm") + "-" + d.add(30, "m").format("HH:mm");
})

Handlebars.registerHelper('longDate', function(d) {
	var m = moment(d);
	m.locale("fr").tz('Europe/Paris')
	return m.format("dddd D MMMM");
});

Handlebars.registerHelper('dateFmt', function(d, fmt) {
	return moment(d).format(fmt)
});
*/


/*
Handlebars.registerHelper('prettyFullname', function(p) {
	var fn = p.Firstname
	var ln = p.Lastname
	return fn.charAt(0).toUpperCase() + fn.substring(1) + " " + ln.charAt(0).toUpperCase() + ln.substring(1)
});

Handlebars.registerHelper('shortFullname', function(p) {
	var fn = p.Firstname + " " + p.Lastname;
	if (fn.length > 20) {
		fn = p.Firstname[0] + ". " + p.Lastname;
	}
	return fn;
});

Handlebars.registerHelper('abbrvFullname', function(p) {
	return p.Firstname[0].toUpperCase() + ". " + p.Lastname;
});

Handlebars.registerHelper('shortProm', function(p) {
	if (p.indexOf("master") == 0) {
		p = "ma. " + p.substring(p.lastIndexOf(" "));
	}
	return p;
});

Handlebars.registerHelper('len', function(a) {
	return a.length
});

Handlebars.registerHelper('company', function(c) {
	if (c.WWW && c.WWW != "") {
		return new Handlebars.SafeString("<a target='_blank' href='" + c.WWW + "'>" + c.Name + "</a>");
	}
	return c.Name;
});

Handlebars.registerHelper('shortCompany', function(c) {
	n = c.Company;
	if (n.length > 30) {
		n = n.substring(0, 27) + "...";
	}
	if (c.CompanyWWW != "") {
		return new Handlebars.SafeString("<a target='_blank' href='" + c.CompanyWWW + "'>" + n + "</a>");
	}
	return n;
});

Handlebars.registerHelper('rawFullname', function(p) {
	var fn = p.Firstname + " " + p.Lastname;
	for (i = fn.length; i < 40; i++) {
		fn = fn + " ";
	}
	return fn;
});

Handlebars.registerHelper('raw', function(p) {
	var fn = p;
	for (i = p.length; i < 60; i++) {
		fn = fn + " ";
	}
	return fn;
});


Handlebars.registerHelper('userSelecter', function(users) {
	var b = "";
	users.forEach(function(o) {
		b += "<option value='" + o.Email + "'>" + o.Firstname + " " + o.Lastname + " (" + o.Tel + ") </option>";
	});
	return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('jurySelecter', function(users) {
	var b = "";
	users.forEach(function(o) {
		b += "<option value='" + o.Email + "'>" + o.Firstname + " " + o.Lastname + " </option>";
	});
	return new Handlebars.SafeString(b);
});


Handlebars.registerHelper('studentSelecter', function(users) {
	var b = "";
	Object.keys(users).forEach(function(em) {
		var i = users[em];
		b += "<option value='" + i.Student.Email + "'>" + i.Student.Firstname + " " + i.Student.Lastname + " (" + i.Major + ") </option>";
	})
	return new Handlebars.SafeString(b);
});


var possiblePositions = [
	"not available",
	"sabbatical leave",
	"looking for a job",
	"pursuit of higher education",
	"fixed term contract in the internship company",
	"fixed term contract in another company",
	"permanent contract in the internship company",
	"permanent contract in another company",
	"entrepreneurship",
	"repeat the year"
];

function editablePossiblePositions() {
	var arr = []
	possiblePositions.forEach(function(p, i) {
		arr.push({
			value: i,
			text: p
		})
	})
	return arr
}

Handlebars.registerHelper('nextPosition', function(pos) {
	return new Handlebars.SafeString("<i>" + possiblePositions[pos] + "</i>");
});

Handlebars.registerHelper('nextContact', function(c) {
	if (!c || c.length == 0) {
		return new Handlebars.SafeString("<i>not available</i>");
	}
	return c;
});

Handlebars.registerHelper('positionSelector', function(pos) {
	var b = "";
	possiblePositions.forEach(function(o, i) {
		b += "<option value='" + i + "' " + (i == pos ? " selected " : "") + ">" + o + "</option>";
	});
	return new Handlebars.SafeString(b);
});

Handlebars.registerHelper('roleOptions', function(m) {
	var b = "";
	var i = 1;
	["tutor", "major", "admin", "root"].forEach(function(k) {
		var selected = i == m ? " selected " : "";
		b += "<option value='" + (i++) + "' " + selected + " >" + k + "</option>";
	});
	return new Handlebars.SafeString(b);
});



Handlebars.registerHelper('majors', function(emails) {
	var majors = {};
	emails.forEach(function(e) {
		if (e && e.Major) {
			majors[e.Major] = true;
		}
	});
	return Object.keys(majors).join(", ");
});

function twoD(d) {
	return d <= 9 ? "0" + d : d;
}



Handlebars.registerHelper('shortSlotEntry', function(s) {
	if (e) {
		var c = getConvention(e);
		var fn = c.Stu.P.Firstname + " " + c.Stu.P.Lastname;
		if (fn.length > 30) {
			fn = fn.substring(0, 27) + "...";
		}
		return fn + " (" + s.Major + ")";
	}
	return "Break";
});

Handlebars.registerHelper('shortKind', function(r) {
	return r.Kind.toLowerCase();
	//return r.Kind.substring(0, 3)
});


Handlebars.registerHelper('defenseStatus', function(g) {
	if (!g.Defenses) {
		return "bg-link";
	}
	var passed = new Date(g.Date).getTime() < new Date().getTime()
	var style = "bg-link";
	if (g.Defenses[0].Grade >= 0 && g.Defenses[0].Grade < 10) {
		return "bg-danger";
	} else if (g.Defenses[0].Grade >= 10) {
		return "bg-success";
	} else if (passed) {
		return "bg-info"
	}
	return "bg-link"
});

Handlebars.registerHelper('defenseGrade3', function(g) {
	if (g >= 0) {
		return g;
	}
	return "?"
});

Handlebars.registerHelper('defenseGrade', function(g) {
	if (!g.Defenses) {
		return "-";
	}
	d = moment(g.Date).add(30 * g.Defenses[0].Offset)
	grade = g.Defenses[0].Grade
	if (grade >= 0) {
		return grade
	}
	var passed = d.toDate().getTime() < new Date().getTime()
	if (grade == -1) {
		return passed ? "?" : "-"
	}
	return "-"
});

Handlebars.registerHelper('defenseGrade2', function(from, offset, grade) {
	d = moment(from).add(30 * offset)
	if (grade > 0) {
		return grade
	}
	var passed = d.toDate().getTime() < new Date().getTime()
	if (grade == -1) {
		return passed ? "?" : "-"
	}
	return "-"
});



function nbDayLates(d1, d2) {
	var m1 = Math.floor(new Date(d1) / 8.64e7);
	var m2 = Math.floor(new Date(d2) / 8.64e7);
	return m1 - m2;
}

function fmtNbDayLate(d) {
	return new Handlebars.SafeString("<i class='glyphicon glyphicon-time'></i><small> " + d + " d.</small>");
}
*/