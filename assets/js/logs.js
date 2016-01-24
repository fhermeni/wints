//log management

function showLogs() {
	logs().success(function(list) {
		$("#cnt").render("logs", list, function() {
			$("a[href='#top']").click(function() {
  				$("html, body").animate({ scrollTop: 0 }, 500);
  				return false;
			});
			$("a[href='#bottom']").click(function() {
  				$("html, body").animate({ scrollTop: $(document).height()-$(window).height()}, 500);
  				return false;
			});
	
			$('.pull-down').each(function() {
    			$(this).css('margin-top', $(this).parent().height()-$(this).height())
			});
			events = list.filter(function (i) {
				return i.indexOf("event") == 0;
			});
			toShow = events[0];
			events.forEach(function (v,i) {
				if (v > toShow) {					
					toShow = v;
				} 
			});
			$("#logs").val(toShow);
			showLog(toShow);
		})
	});
}

function showLog(f) {
	if (!f) {
		f = $("#logs").val();
	}	
	getLog(f).success(function (l) {
		$("#log").html(l);
	})
}