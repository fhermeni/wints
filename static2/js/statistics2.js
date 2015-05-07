var stats;
var allMajors;

var months = ["jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec"];

var pieCountry, pieSector, barGratification;

$( document ).ready(function () {	
waitingBlock = $("#cnt").clone().html();

    statistics(function(m) {
        stats = m;        
        majors(function(m) {
        	allMajors = m;
        	allMajors.unshift("ALL");
        	allMajors.push("n/a");        	
	    })
	    for (var i = 0; i < stats.length; i++) {
	    	if (stats[i].Major == "" || !stats[i].Major) {
	    		stats[i].Major = "n/a";
	    	}
	    	if (stats[i].Promotion.indexOf("si") >= 0) {
	    		stats[i].Promotion = "SI";
	    	} else {
	    		stats[i].Promotion = "master";
	    	}
	    	if (stats[i].Gratification > 10000) {
	    		stats[i].Gratification /= 100;
	    	}
	    	stats[i].Begin = new Date(stats[i].Begin)
	    	stats[i].End = new Date(stats[i].End)
	    }	    

	    Chart.defaults.global.responsive = true;	    
	    //basics()
	    sector()
	    country()
	    employers()
	    //durations()
	    gratification()
	    declared()
    })
});

String.prototype.endsWith = function(suffix) {
    return this.indexOf(suffix, this.length - suffix.length) !== -1;
};

function declared() {
	var at = [];
	var ms = [];
	var i = 0;
	stats.forEach(function (s) {
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
		var x = ms[0].substring(0, ms[0].length-3) + "-1"
		ms.unshift(x)
	}		
	var labels = [];
	var count = [];	
	var i = 0;
	ms.forEach(function (m) {
		if (labels.length == 0) {
			labels.push(m);
			count[i++] = 1;
		} else {
			if (labels[i - 1] == m) {
				count[i - 1]++;
			} else {
				labels[i] = m; 
				count[i] = count[i-1] + 1;
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
	labels.forEach(function (l) {
		var buf = l.split("-");
		xx.push(months[buf[1]] + " " + buf[0]);
	});

	var atLab = $("#conventions").get(0).getContext("2d");			
	var data = {
		labels: ddply(xx),
		datasets : [line(count)]
	}	
	var atc = new Chart(atLab).Line(data, {tooltipTemplate : "<%= value %>", scaleOverride: true, scaleSteps: 6, scaleStepWidth: 20});
}

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

function durations() {
	var begin = new Date()
	begin.setMonth(1)
	begin.setDate(1)
	var end = new Date()
	end.setMonth(10);
	end.setDate(31);	
	console.log(begin)
	console.log(end)
	//Nb of weeks ?
	var step = 1000 * 60 * 60 * 24 * 7 * 2;
	var nbWeeks = (end - begin) / step;

	var labels = [];
	var cur = new Date(begin.getTime());
	while (cur.getTime() < end.getTime()) {
		labels.push(months[cur.getMonth()] + " " + (cur.getYear() - 100))
		cur = new Date(cur.getTime())
		cur.setDate(cur.getDate() + 15)
	};
	console.log(labels);
	var ins = [];
	var i = 0;
	var now = begin;	
	var lbls = [];
	for (i = 0; i <= nbWeeks; i++) {
		stats.forEach(function (s) {
			if (s.Begin.getTime() <= now.getTime() && s.End.getTime() >= now.getTime()) {
				if (!ins[i]) {
					ins[i] = 1;
				} else {
					ins[i]++;
				}						
			}
		});
	now = new Date(now.getTime() + step);
	lbls.push(months[now.getMonth()] + " " + (now.getYear() - 100))
	}		
	var labels = [];	
	var atLab = $("#periods").get(0).getContext("2d");	
	var data = {
		labels: ddply(lbls),
		datasets : [line(ins)]
	}		
	var atc = new Chart(atLab).Line(data, {pointDot: false, tooltipFontSize: 10, tooltipTemplate : "<%= value %>", scaleOverride: true, scaleSteps: 6, scaleStepWidth: 20});
}

function gratification(filter) {
	$("li.gratification").removeClass("active")
	if (!filter) {
		$("li.gratification.all").addClass("active")	
	} else {
		$("li.gratification." + filter).addClass("active")
	}
	var si = [], master = [];
	var lab = [], company = [];
	var fr = [], out = []
	var all = 0		
		stats.forEach(function (s) {
			all += s.Gratification
			if (s.Promotion == "SI") {
				si.push(s.Gratification)
			} else {
				master.push(s.Gratification)
			}	

			if (s.Lab) {
				lab.push(s.Gratification)
			} else {
				company.push(s.Gratification)
			}

			if (s.ForeignCountry) {
				out.push(s.Gratification)
			} else {
				fr.push(s.Gratification)
			}
		});	

	if (!filter) {		
		$("#how-much").html(Math.round(all / stats.length) + "€");
	} else {
		var a,b
			if (filter == "promotion") {
				a = Math.round(avg(si))
				b = Math.round(avg(master))
			} else if (filter == "sector") {
				a = Math.round(avg(lab))
				b = Math.round(avg(company))				
			} else if (filter == "country") {
				a = Math.round(avg(fr))
				b = Math.round(avg(out))								
			}
			$("#how-much").html(a + "€ vs. " + b + "€");
	}
}

function avg(values) {
	var sum = 0
	var nb = 0	
	values.forEach(function (v) {
		sum += v;
		nb++;
	})
	return sum / nb;
}

function employers() {
	var c2 = {}
	stats.forEach(function (s) {		
		if (s.Cpy.WWW) {
			c2[s.Cpy.WWW] = "<a href='" + s.Cpy.WWW + "'>" + s.Cpy.Name + "</a>";				
		} else {
			c2[s.Cpy.Name] = s.Cpy.Name			
		}		
	});		
	var c3 = []
	Object.keys(c2).forEach(function (k) {
		c3.push(c2[k]);
	});	
	$("#employers").html(c3.join("; "));
}

function country(filter) {
	$("li.country").removeClass("active")
	if (!filter) {
		$("li.country.all").addClass("active")	
	} else {
		$("li.country."+filter).addClass("active")	
	}
	
	qty = [0, 0];	
	stats.forEach(function (s) {
		if (filter == undefined || s.Promotion==filter) {
			qty[s.ForeignCountry ? 1 : 0] += 1;
		}
	});	
	var g = $("#canvas-country").get(0).getContext("2d");
	data = [
		{ value: qty[0], label: "France", color:"#F7464A", highlight: "#FF5A5E"},
		{ value: qty[1], label: "Foreign Country", color: "#FDB45C", highlight: "#FFC870"}
	];
	if (!pieCountry) {
		pieCountry = new Chart(g).Pie(data)		
	} else {		
		pieCountry.segments[0].value=qty[0]
		pieCountry.segments[1].value=qty[1]
		pieCountry.update()		
	}	
}

function sector(filter) {
	$("li.sector").removeClass("active")
	if (!filter) {
		$("li.sector.all").addClass("active")	
	} else {
		$("li.sector."+filter).addClass("active")	
	}

	qty = [0, 0];	
	stats.forEach(function (s) {
		if (filter == undefined || s.Promotion == filter) {
			qty[s.Lab ? 1 : 0] += 1;
		}
	});	
	var g = $("#canvas-sector").get(0).getContext("2d");
	var data = [
		{ value: qty[0], label: "Company", color:"#949FB1", highlight: "#A8B3C5"},
		{ value: qty[1], label: "Lab", color: "#4D5360", highlight: "#616774"}
	];
	if (!pieSector) {
		pieSector = new Chart(g).Pie(data)		
	} else {		
		pieSector.segments[0].value=qty[0]
		pieSector.segments[1].value=qty[1]
		pieSector.update()		
	}	

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
            data : dta
        };
}