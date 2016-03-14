var ints;

var ccountry = undefined,
	csector = undefined,
	calumni = undefined,
	cMidtermDelay = undefined,
	barGratification;

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
	});
}

String.prototype.endsWith = function(suffix) {
	return this.indexOf(suffix, this.length - suffix.length) !== -1;
};

function delays(kind, type) {
	$("li.delays").removeClass("active")
	if (kind) {
		delayKind = kind;
	} else {
		kind = delayKind;
	}

	if (type) {
		delayType = type;
	} else {
		type = delayType;
	}

	$("li.delays." + kind).addClass("active")
	$("li.delays." + type).addClass("active")

	var dates = []
	var now = new Date();
	var missing = 0;
	stats.forEach(function(s) {
		s.Reports.forEach(function(r) {
			if (r.Kind != kind) {
				return
			}
			var from = new Date(r.Deadline)
			if (type == "Delivery") {
				if (moment(from).isAfter(moment(now))) {
					//not passed						
					return
				}
				if (new Date(r.Delivery).getTime() > 0) {
					//delivered
					var d = nbDays(from, new Date(r.Delivery))
					if (d < 0) {
						d = 0;
					}
					if (d > 14) {
						d = 14
					}
					dates[d] = dates[d] ? dates[d] + 1 : 1
				} else {
					//not delivered
					missing++;
				}
			} else {
				var del = new Date(r.Delivery)
				var rev = new Date(r.Reviewed)

				if (del.getTime() > 0) {
					if (rev.getTime() <= 0) {
						if (r.Grade < 0) {
							//negative means not reviewed.
							//>0 means reviewed but no timestamp
							missing++;
						}
					} else {
						d = Math.round(nbDays(del, rev) / 7) //per week
						dates[d] = dates[d] ? dates[d] + 1 : 1
					}
				}
				/*else if (moment(from).isBefore(moment(now))) {
					//not delivered
					if (r.Grade >= 0 || rev.getTime() > 0) {
						//but reviewed
						d = Math.round(nbDays(del, rev) / 7) //per week
						console.log(d);
						dates[d] = dates[d] ? dates[d] + 1 : 1
					}
				}*/
			}
		})
	});
	$("#delays").closest(".hidden").removeClass('hidden')
	var keys = Object.keys(dates).map(function(x) {
		if (type == "Delivery") {
			return x + " d.";
		} else {
			return x + " w.";
		}
	})
	if (type == "Delivery") {
		keys[keys.length - 1] = "14 d. +";
	}
	keys.push("missing")
	var values = Object.keys(dates).map(function(x) {
		return dates[x];
	})
	values.push(missing)

	var late = $("#delays").html("").append("<canvas></canvas>").find("canvas").get(0).getContext("2d");
	var data = {
		labels: keys,
		datasets: [line(values)]
	}
	cMidtermDelay = new Chart(late).Bar(data, {
		showTooltip: false
	});
}

function nbDays(from, to) {
	return moment(to).dayOfYear() - moment(from).dayOfYear();
}

/*function declared() {
	var at = [];
	var ms = [];
	var i = 0;
	stats.forEach(function(s) {
		d = new Date(s.Creation);
		if (d.getUTCDate() > 15) {
			d.setUTCDate(15)
		} else {
			d.setUTCDate(1);
		}
		var m = d.getMonth();
		ms.push((d.getYear() - 100) + "-" + d.getUTCMonth() + "-" + d.getUTCDate());
	});
	ms.sort();
	var toShift = ms[0].endsWith("-15")
	if (toShift) {
		var x = ms[0].substring(0, ms[0].length - 3) + "-1"
		ms.unshift(x)
	}
	var labels = [];
	var count = [];
	var i = 0;
	ms.forEach(function(m) {
		if (labels.length == 0) {
			labels.push(m);
			count[i++] = 1;
		} else {
			if (labels[i - 1] == m) {
				count[i - 1]++;
			} else {
				labels[i] = m;
				count[i] = count[i - 1] + 1;
				i++;
			}
		}
	});
	//Padding if the first convention is established by the end of the month
	if (toShift) {
		count.unshift(0)
	}

	//percentage
	for (i = 0; i < count.length; i++) {
		count[i] = Math.round(count[i] / stats.length * 100)
	}
	var xx = [];
	labels.forEach(function(l) {
		var buf = l.split("-");
		xx.push(months[buf[1]] + " " + buf[0]);
	});

	var atLab = $("#conventions").get(0).getContext("2d");
	var data = {
		labels: ddply(xx),
		datasets: [line(count)]
	}
	var atc = new Chart(atLab).Line(data, {
		pointDotRadius: 2,
		tooltipTemplate: "<%= value %>",
		scaleOverride: true,
		scaleSteps: 6,
		scaleStepWidth: 20
	});
}*/

function ddply(arr) {
	var res = [arr[0]];
	for (var i = 1; i < arr.length; i++) {
		if (arr[i] == arr[i - 1]) {
			res[i] = "";
		} else {
			res[i] = arr[i];
		}
	}
	return res
}

/*function durations() {
	var begin = new Date()
	begin.setMonth(1)
	begin.setDate(1)
	var end = new Date()
	end.setMonth(10);
	end.setDate(31);

	var labels = [];
	var cur = new Date(begin.getTime());
	var ins = [];
	while (cur.getTime() < end.getTime()) {
		labels.push(months[cur.getMonth()])
		cur.setMonth(cur.getMonth() + 1);
		ins.push(0)
	};
	var now = begin;
	stats.forEach(function(s) {
		var i = 0;
		var cur = new Date(begin.getTime());
		while (cur.getTime() < end.getTime()) {
			if (s.Begin.getTime() <= cur.getTime() && s.End.getTime() >= cur.getTime()) {
				ins[i]++;
			}
			cur.setMonth(cur.getMonth() + 1);
			i++;
		};
	});
	var atLab = $("#periods").get(0).getContext("2d");
	var data = {
		labels: ddply(labels),
		datasets: [line(ins)]
	}
	var atc = new Chart(atLab).Line(data, {
		pointDotRadius: 2,
		tooltipTemplate: "<%= value %>",
		scaleOverride: true,
		scaleSteps: 6,
		scaleStepWidth: 20
	});
}*/



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
		var idx = -1;		
		if (filter == "promotion") {
			idx = config.Promotions.indexOf(s.Convention.Student.Promotion);
		} else if (filter == "sector") {			
			idx  = s.Convention.Lab ? 0 : 1;
		} else if (filter == "country") {
			idx  = s.Convention.ForeignCountry ? 0 : 1;
		} else if (filter == "major") {
			idx = config.Majors.indexOf(s.Convention.Student.Major);
		}			
		if (!qty[idx]) {
			qty[idx] = [];
		}			
		qty[idx].push(g);
	});	
	if (!filter) {
		$("#how-much").html(Math.round(all / ints.length) + "€");
	} else {
		if (filter == "country" || filter == "sector") {
			$("#how-much").html("<h1>" + Math.round(avg(qty[0])) + "€ / " + Math.round(avg(qty[1])) + "€</h1>");
		} else {
			for (var i = 0; i < qty.length; i++) {
				qty[i] = Math.round(avg(qty[i]));
			}
			c = $("<canvas>");
			data = dataset(filter == "promotion" ? config.Promotions : config.Majors, "gratifications", qty);	
			c = $("#how-much").html(c).find("canvas");			
			ctx = c.get(0).getContext("2d");
			new Chart(ctx).Bar(data);
		} 
	}
}

function grades(kind, filter) {
	$("li.grades-" + kind).removeClass("active")
	if (!filter) {
		$("li.grades-" + kind + ".all").addClass("active")
	} else {
		$("li.grades-" + kind + "." + filter).addClass("active")
	}

	var byMajor = [];
	var qty = [
		[],
		[]
	];
	for (var i = 0; i < allMajors.length; i++) {
		byMajor.push([])
	}
	var all = 0
	var nb = 0
	stats.forEach(function(s) {
		s.Reports.forEach(function(r) {
			if (r.Kind == kind && r.Grade > 0 && new Date(r.Delivery).getTime() > 0) {
				g = r.Grade
				nb++
				all += g
				if (filter == "promotion") {
					qty[s.Promotion == "SI" ? 0 : 1].push(g)
				} else if (filter == "major") {
					byMajor[allMajors.indexOf(s.Major)].push(g);
				}
			}
		})
	});
	//Hide when no data
	if (nb) {
		$("#grades-" + kind).closest(".hidden").removeClass('hidden')
	}

	if (!filter) {
		num = all / nb
		$("#grades-" + kind).html(num.toFixed(2) + " / 20");
	} else {
		if (filter == "major") {
			var avgs = []
			for (var i = 0; i < byMajor.length; i++) {
				avgs.push(avg(byMajor[i]).toFixed(2))
			}
			c = $("<canvas>")
			data = {
				labels: allMajors,
				datasets: [{
					label: "Gratification",
					fillColor: "#F7464A",
					highlightFill: "#FF5A5E",
					data: avgs
				}, ]
			}
			c = $("#grades-" + kind).html("").append(c).find("canvas")
			ctx = c.get(0).getContext("2d");
			grat = new Chart(ctx).Bar(data);
		} else {
			avgs = [avg(qty[0]).toFixed(2), avg(qty[1]).toFixed(2)]
			c = $("<canvas>")
			data = {
				labels: ["SI", "master"],
				datasets: [{
					label: "Gratification",
					fillColor: "#F7464A",
					highlightFill: "#FF5A5E",
					data: avgs
				}, ]
			}
			c = $("#grades-" + kind).html("").append(c).find("canvas")
			ctx = c.get(0).getContext("2d");
			grat = new Chart(ctx).Bar(data);
		}
	}
}

function surveys(kind) {
	$("li.surveys").removeClass("active")
	$("li.surveys-" + kind).addClass("active")

	var all = 0
	var nb = 0
	var grades = 0;
	stats.forEach(function(stat) {
		stat.Surveys.forEach(function(s) {
			if (s.Kind != kind) {
				return
			}
			var q = s.Answers
			if (q && Object.keys(q).length > 0) {
				nb++
				if (kind == "midterm" && q[19] == "true") {
					all++
				} else if (kind == "final") {
					grades += parseInt(q["q17"])
				}
			}
		})
	});
	//Hide when no data
	if (nb) {
		var num = Math.round(all / nb * 100)
		if (kind == "midterm") {
			$("#surveys").html(num + "%");
		} else if (kind == "final") {
			var avg = grades / nb;
			$("#surveys").html(avg.toFixed(2) + " / 20");
		}
		$("#surveys").closest(".hidden ").removeClass('hidden')

	}
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
	$("li.country").removeClass("active")
	if (!filter) {
		$("li.country.all").addClass("active")
	} else {
		$("li.country." + filter).addClass("active")
	}

	var qty = [];
	var total = [];
	
	ints.forEach(function(s) {
		var idx = 0;		
		if (filter == "promotion") {
			idx = config.Promotions.indexOf(s.Convention.Student.Promotion);
		} else if (filter == "major") {
			idx = config.Majors.indexOf(s.Convention.Student.Major);
		}		
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
		qty[i] = Math.round(qty[i] / total[i] * 100);
		if (!qty[i]) {
			qty[i] = 0;
		}
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


function showAlumni(filter) {
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
		var idx = 0;		
		if (filter == "promotion") {
			idx = config.Promotions.indexOf(s.Convention.Student.Promotion);
		} else if (filter == "major") {
			idx = config.Majors.indexOf(s.Convention.Student.Major);
		}		
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
		qty[i] = Math.round(qty[i] / total[i] * 100);
		if (!qty[i]) {
			qty[i] = 0;
		}
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

function avg(values) {
	var sum = 0
	var nb = 0
	if (!values || values.length == 0) {
		return nb
	}
	values.forEach(function(v) {
		sum += v;
		nb++;
	})
	return sum / nb;
}

function zeroes(nb) {
	var x = []
	for (var i = 0; i < nb; i++) {
		x.push(0)
	}
	return x
}

function dataset(labels, title, d) {
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

function line(dta) {
	return {
		label: "Pick up date",
		fillColor: "rgba(220,220,220,0.2)",
		strokeColor: "rgba(220,220,220,1)",
		pointColor: "rgba(220,220,220,1)",
		pointStrokeColor: "#fff",
		pointHighlightFill: "#fff",
		pointHighlightStroke: "rgba(220,220,220,1)",
		data: dta
	};
}