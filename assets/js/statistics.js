var ints;

var calumni = undefined;

function loadStatistics() {
	//IE8 stuff
	if (typeof Array.prototype.forEach != 'function') {
		Array.prototype.forEach = function(callback) {
			for (var i = 0; i < this.length; i++) {
				callback.apply(this, [this[i], i, this]);
			}
		};
	}

	waitingBlock = $("#cnt").clone().html();

	internships().done(function (ii) {
		ints = ii;
		Chart.defaults.global.responsive = true;
		Chart.defaults.global.pointDot = false;
		Chart.defaults.global.pointDotRadius = 2;

		//Clean a bit the values
		for (var i = 0; i < ints.length; i++) {
				if (ints[i].Convention.Gratification > 10000) {
					ints[i].Convention.Gratification /= 100;
				}
				ints[i].Convention.Begin = moment(ints[i].Begin)
				ints[i].Convention.End = moment(ints[i].Begin)
		}

		employers(ints);
		gratifications();
		country();
		sector();
		grades('midterm');
		grades('final');
		surveys('midterm');
	});
}

String.prototype.endsWith = function(suffix) {
	return this.indexOf(suffix, this.length - suffix.length) !== -1;
};


function gratifications(filter) {
	$("li.gratification").removeClass("active")
	if (!filter) {
		$("li.gratification.all").addClass("active")
	} else {
		$("li.gratification." + filter).addClass("active")
	}

	var qty = [];
	var all = 0;
	ints.forEach(function(s) {
		g = s.Convention.Gratification;
		all += g;
		var idx = index(s, filter);
		if (!qty[idx]) {
			qty[idx] = [];
		}
		qty[idx].push(g);
	});
	for (var i = 0; i < qty.length; i++) {
		qty[i] = Math.round(avg(qty[i]));
	}
	if (!filter) {
		$("#how-much").html(Math.round(qty[0]) + "€");
	} else {
		if (filter == "country" || filter == "sector") {
			$("#how-much").html("<h1>" + Math.round(qty[0]) + "€ / " + Math.round(qty[1]) + "€</h1>");
		} else {
			c = $("<canvas>");
			data = dataset(filter == "promotion" ? config.Promotions : config.Majors, "gratifications", qty);
			c = $("#how-much").html(c).find("canvas");
			ctx = c.get(0).getContext("2d");
			new Chart(ctx).Bar(data);
		}
	}
}

function grades(kind, filter) {
	$("li.grades-"+kind).removeClass("active");
	if (!filter) {
		$("li.grades-"+kind+".all").addClass("active");
	} else {
		$("li.grades-"+kind+".filter").addClass("active");
	}

	var qty = [];
	var done = 0;
	ints.forEach(function(s) {
		s.Reports.forEach(function (r) {
			if (r.Kind != kind) {
				return;
			}
			if (r.Reviewed) {
				var idx = index(s, filter);
				if (!qty[idx]) {
					qty[idx] = [];
				}
				qty[idx].push(r.Grade);
				done++;
			}
		})
	});
	for (var i = 0; i < qty.length; i++) {
		qty[i] = avg(qty[i]);
		if (qty[i] != undefined) {
			qty[i] = qty[i].toFixed(1);
		}
	}
	if (done < ints.length / 2) {
		$("#grades-" + kind).html("<h4>Not enough data</h4>");
		return;
	}
	if (!filter) {
		$("#grades-" + kind).html("<h1>" + qty[0] + " / 20</h1>");
	} else {
		c = $("<canvas>");
		data = dataset(filter == "promotion" ? config.Promotions : config.Majors, "Grade", qty);
		c = $("#grades-" + kind).html(c).find("canvas");
		ctx = c.get(0).getContext("2d");
		new Chart(ctx).Bar(data);
	}
}

function surveys(kind) {
	$("li.surveys").removeClass("active")
	$("li.surveys-" + kind).addClass("active")

	var res={};
	ints.forEach(function(stat) {
		stat.Surveys.forEach(function(s) {
			if (!res[kind]) {
				res[kind] = [];
			}
			if (s.Kind != kind ||!s.Delivery) {
				return
			}

			var mark = s.Cnt["__MARK__"];
			if (mark === "false") {
				res[kind].push(0);
			} else if (mark === "true") {
				res[kind].push(100);
			} else {
				res[kind].push(parseInt(mark));
			}
		})
	});
	//Hide when there is less than 50% of answers
	var output = "<h4>not enough data</h4>";
	if (res[kind].length > ints.length / 2) {
		if (kind == "midterm") {
			output = "<h1>"+Math.round(avg(res[kind]))+"%</h1><h4>are satisfied by their intern</h4>";
		} else if(kind=="final") {
			output = "<h1>"+avg(res[kind]).toFixed(2) + " / 20</h1>";
		}
	}
	$("#surveys").html(output);
}


function employers(ints) {
	var c2 = {}
	ints.forEach(function(s) {
		www = s.Convention.Company.WWW;
		if (www) {
			c2[www] = "<a href='" + www + "'>" + s.Convention.Company.Name + "</a>";
		} else {
			c2[s.Convention.Company.Name] = s.Convention.Company.Name
		}
	});
	var c3 = []
	Object.keys(c2).forEach(function(k) {
		c3.push(c2[k]);
	});
	$("#employers").html(c3.join("; "));
}


function country(filter) {
	//Toggles coordination
	$("li.country").removeClass("active")
	if (!filter) {
		$("li.country.all").addClass("active")
	} else {
		$("li.country." + filter).addClass("active")
	}

	var qty = [];
	var total = [];

	ints.forEach(function(s) {
		var idx = index(s, filter);
		if (!qty[idx]) {
			qty[idx] = 0;
			total[idx] = 0;
		}
		if (s.Convention.ForeignCountry) {
			qty[idx]++;
		}
		total[idx]++;
	});


	for (var i = 0; i < qty.length; i++) {
		qty[i] = pct(qty[i],total[i]);
	}

	if (!filter) {
		$("#country").html("<h1>" + qty[0] + "%</h1>");
	} else {
		c = $("<canvas>");
		data = dataset(filter == "promotion" ? config.Promotions : config.Majors, "foreign country", qty);
		c = $("#country").html(c).find("canvas");
		ctx = c.get(0).getContext("2d");
		new Chart(ctx).Bar(data);
	}
}

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

var colors = ["#4D4D4D", "#5DA5DA", '#FAA43A', '#60BD68', '#F17CB0', '#B2912F', '#B276B2', '#DECF3F', '#F15854', "#DDDDDD"];


function showAlumniStatistics(filter) {
	if (!filter) filter = "all"
	$("li.alumni.filter").removeClass("active")
	$("li.alumni.filter." + filter).addClass("active")

	qty = zeroes(possiblePositions.length);
	stats.forEach(function(s) {
		if (s.Reports.length > 0 && (filter == "all" || s.Promotion == filter)) {
			qty[s.Future.Position] += 1;
		}
	});
	var legend = ""
	for (i = 0; i < colors.length; i++) {
		legend += "<li><i class='box' style='background-color: " + colors[i] + ";'></i> " + possiblePositions[i] + "</li>";
	}
	$("#alumni-legend").html(legend);
	if (!calumni) {
		var data = [];
		for (i = 0; i < possiblePositions.length; i++) {
			d = {
				value: qty[i],
				color: colors[i],
				label: possiblePositions[i]
			}
			data.push(d);
		}
		calumni = new Chart($("#alumni").find("canvas").get(0).getContext("2d")).Pie(data, {
			segmentStrokeWidth: 1,
			maintainAspectRatio: true
		})
	} else {
		for (i = 0; i < colors.length; i++) {
			calumni.segments[i].value = qty[i]
		}
		calumni.update()
	}
}

function sector(filter) {
	$("li.sector").removeClass("active")
	if (!filter) {
		$("li.sector.all").addClass("active")
	} else {
		$("li.sector." + filter).addClass("active")
	}

	var qty = [];
	var total = [];

	ints.forEach(function(s) {
		var idx = index(s, filter);
		if (!qty[idx]) {
			qty[idx] = 0;
			total[idx] = 0;
		}
		if (!s.Convention.Lab) {
			qty[idx]++;
		}
		total[idx]++;
	});

	for (var i = 0; i < qty.length; i++) {
		qty[i] = pct(qty[i],total[i]);
	}

	if (!filter) {
		$("#sector").html("<h1>" + qty[0] + "%</h1>");
	} else {
		c = $("<canvas>");
		data = dataset(filter == "promotion" ? config.Promotions : config.Majors, "in company", qty);
		c = $("#sector").html(c).find("canvas");
		ctx = c.get(0).getContext("2d");
		new Chart(ctx).Bar(data);
	}
}


/*
* Helpers
*/

/**
* returns an identifier for the category depending on the filter.
* 0 when there is no filter
*/
function index(i, filter) {
		if (filter == "promotion") {
			return config.Promotions.indexOf(i.Convention.Student.Promotion);
		} else if (filter == "sector") {
			return i.Convention.Lab ? 0 : 1;
		} else if (filter == "country") {
			return i.Convention.ForeignCountry ? 0 : 1;
		} else if (filter == "major") {
			return config.Majors.indexOf(i.Convention.Student.Major);
		}
		return 0
}

function avg(values) {
	var sum = 0
	var nb = 0
	if (!values || values.length == 0) {
		return undefined;
	}
	values.forEach(function(v) {
		sum += v;
		nb++;
	})
	return sum / nb;
}

function pct(a, b) {
	return a === undefined ? a : Math.round(a / b * 100);
}

function zeroes(nb) {
	var x = []
	for (var i = 0; i < nb; i++) {
		x.push(0)
	}
	return x
}

function dataset(labels, title, d) {
	var i = d.length;
	while (i--) {
		if (isNaN(d[i])) {
			d.splice(i, 1);
			labels.splice(i, 1);
		}
	}
	return { labels: labels,
			datasets: [{
					label: title,
					fillColor: "rgba(220,220,220,0.5)",
            		strokeColor: "rgba(220,220,220,0.8)",
            		highlightFill: "rgba(220,220,220,0.75)",
            		highlightStroke: "rgba(220,220,220,1)",
					data: d,
				}, ]
	};
}