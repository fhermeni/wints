var stats;
var allMajors;

var months = ["jan","feb","mar","apr","may","jun","jul","aug","sep","oct","nov","dec"];

var ccountry = undefined, csector = undefined, barGratification;

$( document ).ready(function () {	
waitingBlock = $("#cnt").clone().html();

    statistics(function(m) {
        stats = m;        
        majors(function(m) {
        	allMajors = m;        	
        	allMajors.push("n/a");        	

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
	    	Chart.defaults.global.pointDot = false;
	    	Chart.defaults.global.pointDotRadius = 2;	    
	    	//basics()
	    	sector()
	    	country()
	    	employers()
	    	durations()
	    	gratification()
	    	declared()
	    	$(":checkbox").iCheck()
	    })
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
	var atc = new Chart(atLab).Line(data, {pointDotRadius : 2,tooltipTemplate : "<%= value %>", scaleOverride: true, scaleSteps: 6, scaleStepWidth: 20});
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

	var labels = [];
	var cur = new Date(begin.getTime());
	var ins = [];	
	while (cur.getTime() < end.getTime()) {
		labels.push(months[cur.getMonth()])		
		cur.setMonth(cur.getMonth() + 1);		
		ins.push(0)
	};
	var now = begin;		
	stats.forEach(function (s) {
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
		datasets : [line(ins)]
	}		
	var atc = new Chart(atLab).Line(data, {pointDotRadius : 2, tooltipTemplate : "<%= value %>", scaleOverride: true, scaleSteps: 6, scaleStepWidth: 20});
}

function gratification(filter) {
	$("li.gratification").removeClass("active")
	if (!filter) {
		$("li.gratification.all").addClass("active")	
	} else {
		$("li.gratification." + filter).addClass("active")
	}
	
	var byMajor = [];
	var qty = [[], []];
	for (var i = 0; i < allMajors.length; i++) {
		byMajor.push([])
	}		
	var all = 0		
		stats.forEach(function (s) {			
			all += s.Gratification
			if (filter == "promotion") {
				qty[s.Promotion == "SI" ? 0 : 1].push(s.Gratification)
			} else if (filter == "sector") {
				qty[s.Lab ? 0 : 1].push(s.Gratification)
			} else if (filter == "country") {
				qty[s.ForeignCountry ? 0 : 1].push(s.Gratification)
			} else if (filter == "major") {								
				byMajor[allMajors.indexOf(s.Major)].push(s.Gratification);
			}
		});			
	if (!filter) {		
		$("#how-much").html(Math.round(all / stats.length) + "€");
	} else {
		if (filter == "major") {	
			var avgs = []
			for (var i = 0; i < byMajor.length; i++) {			
				avgs.push(Math.round(avg(byMajor[i])))
			}			
			c = $("<canvas>")
			data = {
    		labels: allMajors,
    		datasets: [
    			{ label: "Gratification", fillColor:"#F7464A", highlightFill: "#FF5A5E", data: avgs},    			
    		]
    		}       
			c = $("#how-much").html("").append(c).find("canvas")			
			ctx = c.get(0).getContext("2d");
			grat= new Chart(ctx).Bar(data);
		} else {			
			$("#how-much").html(Math.round(avg(qty[0])) + "€ vs. " + Math.round(avg(qty[1])) + "€");
		}		
	}
}

function avg(values) {
	var sum = 0
	var nb = 0	
	if (values.length == 0) {
		return nb
	}
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

function switchCountry(b) {
	c = $("#country").html("").append("<canvas>").first().get(0).getContent("2d")	
}

function toggleCountry(a) {
	 li  = $(a).closest("li");
	 if (li.hasClass("active")) {
	 	li.removeClass("active");	 	
	 } else {
	 	li.addClass("active");
	 }
	 //dump canvas
	 ccountry = undefined			 
	 $("#country").html("").append("<canvas>");
	 //the active filter	 
	 $("li.country.filter.active").find("a").click()
}	 

function country(filter) {	
	if (!filter) filter = "all"
	$("li.filter.country").removeClass("active")
	$("li.filter.country."+filter).addClass("active")		
	
	qty = [0, 0];	
	byMajor = [zeroes(allMajors.length),zeroes(allMajors.length)];	
	stats.forEach(function (s) {
		if (filter == "all" || s.Promotion==filter) {
			qty[s.ForeignCountry ? 1 : 0] += 1;
			byMajor[s.ForeignCountry ? 1 : 0][allMajors.indexOf(s.Major)] += 1;
		}
	});	
	
	if (!$("li.country.byMajor").hasClass("active")) {
		if (!ccountry) {
			var g = $("#country").find("canvas").get(0).getContext("2d");
			data = [
				{value: qty[0], label: "France", color:"#F7464A", highlight: "#FF5A5E"},
				{value: qty[1], label: "Foreign Country", color: "#FDB45C", highlight: "#FFC870"}
			];
			ccountry = new Chart(g).Pie(data)		
		} else {
			ccountry.segments[0].value=qty[0]
			ccountry.segments[1].value=qty[1]
			ccountry.update()		
		}
	} else {
		if (!ccountry) {
			data = {
    		labels: allMajors,
    		datasets: [
    			{ label: "France", fillColor:"#F7464A", highlightFill: "#FF5A5E", data: byMajor[0] },
    			{ label: "Foreign Country", fillColor: "#FFC870", highlightFill: "#FFC870", data: byMajor[1] }
    		]
    		}       
			var g = $("#country").find("canvas").get(0).getContext("2d");
			ccountry = new Chart(g).Bar(data);
		} else {
			for (var i = 0; i < allMajors.length; i++) {			
				ccountry.datasets[0].bars[i].value = byMajor[0][i]
				ccountry.datasets[1].bars[i].value = byMajor[1][i]				
			}
			ccountry.update();
		}
	}
}

function toggleSector(a) {
	 li  = $(a).closest("li");
	 if (li.hasClass("active")) {
	 	li.removeClass("active");	 	
	 } else {
	 	li.addClass("active");
	 }
	 //dump canvas
	 csector = undefined			 
	 $("#sector").html("").append("<canvas>");
	 //the active filter	 
	 $("li.sector.filter.active").find("a").click()
}

function sector(filter) {	
	if (!filter) filter = "all"
	$("li.sector.filter").removeClass("active")
	$("li.sector.filter."+filter).addClass("active")		

	qty = [0, 0];		
	byMajor = [zeroes(allMajors.length),zeroes(allMajors.length)];	
	stats.forEach(function (s) {
		if (filter == "all" || s.Promotion == filter) {
			qty[s.Lab ? 1 : 0] += 1;
			byMajor[s.Lab ? 1 : 0][allMajors.indexOf(s.Major)] += 1;
		}
	});	
	
	if (!$("li.sector.byMajor").hasClass("active")) {
		if (!csector) {
			var data = [
				{ value: qty[0], label: "Company", color:"#949FB1", highlight: "#A8B3C5"},
				{ value: qty[1], label: "Lab", color: "#4D5360", highlight: "#616774"}
			];			
			csector = new Chart($("#sector").find("canvas").get(0).getContext("2d")).Pie(data)					
		} else {
			csector.segments[0].value=qty[0]
			csector.segments[1].value=qty[1]
			csector.update()					
		}
	} else {
		if (!csector) {			
			data = {
    			labels: allMajors,
    			datasets: [
    				{ label: "Company", fillColor:"#949FB1", highlightFill: "#A8B3C5", data: byMajor[0] },
    				{ label: "Lab", fillColor: "#4D5360", highlightFill: "#616774", data: byMajor[1] }
    			]
    		}       		
			csector = new Chart($("#sector").find("canvas").get(0).getContext("2d")).Bar(data);		
		} else {			
			for (var i = 0; i < allMajors.length; i++) {			
				csector.datasets[0].bars[i].value = byMajor[0][i]
				csector.datasets[1].bars[i].value = byMajor[1][i]				
			}
			csector.update();			
		}
	}	
}

function zeroes(nb) {
	var x = []
	for (var i = 0; i < nb; i++) {
		x.push(0)
	}
	return x
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