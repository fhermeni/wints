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
		});
	});
}

var dateTime = /(^\d+\/+\d+\/\d+)/; //2016/04/18 12:47:05

function showLog(f) {
	if (!f) {
		f = $("#logs").val();
	}
	getLog(f).success(function (l) {
		l = l.split('\n').map(function (x) {
			return x.replace(dateTime,"");
		}).join("\n");
		$("#log").html(l);
		$("html, body").animate({ scrollTop: $(document).height()-$(window).height()}, 500);
	});
}