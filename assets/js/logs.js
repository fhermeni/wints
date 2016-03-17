//log management

function showLogs() {
	logs().success(function(list) {
		$("#cnt").render("logs", list, function() {
			$("a[href='#top']").click(function() {
  				$("html, body").animate({ scrollTop: 0 }, 500);
  				return false;
			});
			$("a[href='#bottom']").click(function() {				
				showLog();
  				return false;
			});
	
			events = list.filter(function (i) {
				return i.indexOf("event") >= 0;
			});			
			toShow = events[events.length-1];			
			$("#logs").val(toShow);
			showLog(toShow);
		})
	});
}

var re = /\[\d+\/\d+\/\d+ /;

function showLog(f) {
	if (!f) {
		f = $("#logs").val();
	}	
	getLog(f).success(function (l) {			
		l = l.split('\n').map(function (x) {
			return x.replace(re,'[');
		}).join("\n");	
		$("#log").html(l);
		$("html, body").animate({ scrollTop: $(document).height()-$(window).height()}, 500);		
	});
}