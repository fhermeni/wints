var stats;
var allMajors;
$( document ).ready(function () {	
waitingBlock = $("#cnt").clone().html();

    statistics(function(m) {
        stats = m;        
        majors(function(m) {
        	allMajors = m;
        	allMajors.push("n/a");
        	local();
        	lab();
        	gratification();
	    })
    })
});

function max() {
	var m = arguments[0][0];
	arguments.forEach(function (arr) {
		arr.forEach(function (e) {
			if (e > m) {
				m = e;
			}
		})
	});
	return m;
}

function major_counter() {
	var cnt = [];
	for (var i = 0; i < allMajors.length; i++) {
		cnt.push(0);
	}
	return cnt;
}

function count(ok,ko, idx,o, key) {
	if (o[key]) {
		ok[idx]++;
	} else {
		ko[idx]++;
	}
}
function master(st) {
	return st.Promotion.indexOf("si") != 0;
}
function bar_width(dta) {
  if(dta.type === 'bar') {
    dta.element.attr({
      style: 'stroke-width: 30px'
    });
  }
}

function major_bars(arrays) {	
	return {labels : allMajors,series: arrays};	
}

function local() {
	var foreign_si = major_counter();
	var local_si = major_counter();	
	var foreign_ma = major_counter();
	var local_ma = major_counter();	
	
	stats.forEach(function (st) {		
		var ma = st.Major;
		if (ma=="" || !ma) {ma="n/a";}
		var idx = allMajors.indexOf(ma);
		if (!master(st)) {	
			count(foreign_si, local_si, idx, st, "ForeignCountry");
		} else {
			count(foreign_ma, local_ma, idx, st, "ForeignCountry");
		}
	});	

	var dta_si = major_bars([local_si, foreign_si]);
	var dta_ma = major_bars([local_ma, foreign_ma]);	
	var options = {
  		stackBars:true,
  		high: stats.length
  	};
	new Chartist.Bar('.local_si', dta_si, options).on('draw', bar_width);
    new Chartist.Bar('.local_ma', dta_ma, options).on('draw', bar_width);
}

function lab() {
	var foreign_si = major_counter();
	var local_si = major_counter();	
	var foreign_ma = major_counter();
	var local_ma = major_counter();	
	
	stats.forEach(function (st) {		
		var ma = st.Major;
		if (ma=="" || !ma) {ma="n/a";}
		var idx = allMajors.indexOf(ma);
		if (!master(st)) {	
			count(foreign_si, local_si, idx, st, "Lab");
		} else {
			count(foreign_ma, local_ma, idx, st, "Lab");
		}
	});	

	var dta_si = major_bars([local_si, foreign_si]);
	var dta_ma = major_bars([local_ma, foreign_ma]);	
	var options = {
  		stackBars:true,
  		high: stats.length
  	};
	new Chartist.Bar('.lab_si', dta_si, options).on('draw', bar_width);
    new Chartist.Bar('.lab_ma', dta_ma, options).on('draw', bar_width);
}

function gratification() {
	var si = major_counter();
	var ma = major_counter();	
	var nb_si = major_counter();
	var nb_ma = major_counter();
	stats.forEach(function (st) {		
		var ma = st.Major;
		if (ma=="" || !ma) {ma="n/a";}
		var idx = allMajors.indexOf(ma);
		if (!master(st)) {	
			si[idx]+= st.Gratification;
			nb_si[idx]++;			
		} else {
			ma[idx]+= st.Gratification;
			nb_ma[idx]++;			
		}
	});	

	//Average
	for (var i = 0; i < si.length; i++) {
		si[i] /= nb_si[i];
		ma[i] /= nb_ma[i];
	}
	var dta_si = major_bars([si, ma]);	
	var options = {seriesBarDistance: 10};
	new Chartist.Bar('.gratification_si', dta_si, options).on('draw', function (dta) {
  if(dta.type === 'bar') {
    dta.element.attr({
      style: 'stroke-width: 10px'
    });
  }
});
    //new Chartist.Bar('.gratification_ma', dta_ma, options).on('draw', bar_width);	
}