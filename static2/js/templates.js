/**
 * Created by fhermeni on 03/07/2014.
 */

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
	//     return /* 0-23 hour, given meridiem token and hour 1-12 */
	// },
	meridiem: function(hours, minutes, isLower) {
		return hours < 12 ? 'PD' : 'MD';
	},
	week: {
		dow: 1, // Monday is the first day of the week.
		doy: 4 // The week that contains Jan 4th is the first week of the year.
	}
});

Handlebars.getTemplate = function(name) {
	if (Handlebars.templates === undefined || Handlebars.templates[name] === undefined) {
		$.ajax({
			url: '/static/tpls/' + name + '.handlebars',
			success: function(data) {
				if (Handlebars.templates === undefined) {
					Handlebars.templates = {};
				}
				Handlebars.templates[name] = Handlebars.compile(data);
			},
			async: false
		});
	}
	return Handlebars.templates[name];
};

Handlebars.registerHelper('fullname', function(p) {
	return p.Firstname + " " + p.Lastname;
});

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

Handlebars.registerHelper('majorOptions', function(m) {
	var b = "";
	if (!m) {
		b += "<option selected>?</option>";
	}
	allMajors.forEach(function(o) {
		var selected = m == o ? " selected " : "";
		b += "<option value='" + o + "' " + selected + " >" + o + "</option>";
	});
	return new Handlebars.SafeString(b);
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

Handlebars.registerHelper('reportStatus', function(r) {
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

Handlebars.registerHelper('surveyAnnotation', function(s) {
	if (Object.keys(s.Answers).length == 0) {
		return -10;
	}
	if (s.Kind == "midterm" && s.Answers[19] == "false") {
		return -1;
	}

	if (s.Kind == "final" && s.Answers["q17"] < 10) {
		return -1;
	}

	return 0;
});


Handlebars.registerHelper('surveyGrade', function(s) {
	if (Object.keys(s.Answers).length == 0) {
		return "-"
	}
	if (s.Kind == "midterm") {
		if (s.Answers[19] == "true") {
			return new Handlebars.SafeString("&#10003;");
		}
		return new Handlebars.SafeString("x");
	} else {
		return s.Answers["q17"];
	}
	return "-";
});

Handlebars.registerHelper('surveyStatus', function(s) {
	if (Object.keys(s.Answers).length == 0) {
		return ""
	}
	if (s.Kind == "midterm") {
		if (s.Answers[19] == "true") {
			return "bg-success";
		}
		return "bg-danger";
	}
	if (s.Kind == "final") {
		if (s.Answers["q17"] >= 10) {
			return "bg-success";
		}
		return "bg-danger";
	}
	return "";
});


function nbDayLates(d1, d2) {
	var m1 = Math.floor(new Date(d1) / 8.64e7);
	var m2 = Math.floor(new Date(d2) / 8.64e7);
	return m1 - m2;
}

function fmtNbDayLate(d) {
	return new Handlebars.SafeString("<i class='glyphicon glyphicon-time'></i><small> " + d + " d.</small>");
}

Handlebars.registerHelper('reportGrade', function(r) {
	var passed = (new Date(r.Deadline).getTime() + 86400 * 1000) < new Date().getTime()
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