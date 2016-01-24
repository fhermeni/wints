//log management

function showLogs() {
	logs().success(function(list) {
		$("#cnt").render("logs", list, function() {
			showLog();
		})
	});
}

function showLog() {
	var f = $("#logs").val();
	getLog(f).success(function (l) {
		$("#log").html(l);
	})
}