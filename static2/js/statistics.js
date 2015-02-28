var stats;
var allMajors;
$( document ).ready(function () {	
waitingBlock = $("#cnt").clone().html();

    statistics(function(m) {
        stats = m;        
        majors(function(m) {
        	allMajors = m;

        	who_where();
	    })
    })
});


/*
per major/promotion
 -> lab/or not
 -> local or not
 */
function who_where() {
	var masters = [];
	var si = [];	
	var majors = allMajors;
	majors.push("n/a");	
	majors.forEach(function (m) {
		si.push(0);
		masters.push(0);
	});
	stats.forEach(function (st) {
		var p = st.Promotion;
		var ma = st.Major;
		if (ma=="" || !ma) {
			ma="n/a";
		}
		if (p.indexOf("si") == 0) {
			si[majors.indexOf(ma)]++;
		} else {
			masters[majors.indexOf(ma)]++;
		}
	});	
	var dta = {
		labels: allMajors,
		series: [
			si,masters
		]
	}
	console.log(dta);
	var options = {
  		stackBars:true,
  		axisX: {
    		offset: 60
  		},
  		width: 600,
  		height: 400
  };
	new Chartist.Bar('.who-where', dta, options)
	.on('draw', function(data) {
  if(data.type === 'bar') {
    data.element.attr({
      style: 'stroke-width: 30px'
    });
  }
});

	/*.on('draw', function(data) {
  		if(data.type === 'bar') {
    		data.element.attr({style: 'stroke-width: 30px'});
  		}
	})*/;
}