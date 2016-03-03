$(document).ready(function() {

	$('[data-toggle="popover"]').popover();


	//Login part
	if (window.location.pathname == "/login") {
		var em = $.urlParam("email")
		if (em) {
			$("#loginEmail").val(em);
		}		

		$(function() {
			$("#login-diag").find("input").keypress(function(e) {			
				if ((e.which && e.which == 13) || (e.keyCode && e.keyCode == 13)) {					
					$('#login').click();
					return false;
				} else {
					return true;
				}
			});
		});
	} else if (window.location.pathname == "/home") {
		//home part
		waitingBlock = $("#cnt").clone().html();
		$( document ).ajaxError(function(e, xhr) {		
  			if (xhr.responseText.indexOf("expired") >0) {
  				window.location = "/#sessionExpired";
  			}
		});
		$.tablesorter.defaults.widgets = ["uitheme"]
		$.tablesorter.defaults.theme = 'bootstrap';
		$.tablesorter.defaults.headerTemplate = '{content} {icon}';

		getConfig().done(function(c) {
			config = c;
			cookie = getCookie("login")		
			if (cookie) { 			
				user(cookie).done(loadSuccess).fail(logFail);
				$(".release").html(config.Version)
			}
		})
	} else if (window.location.pathname == "/survey") {
		setLang(".fr",".en");
		loadSurvey();
	} else if (window.location.pathname == "/password") {
		var t = $.urlParam("token");
		if (!t) {
			$(".alert-danger").html("<strong>There is no token in the request</strong>. Initiate a reset request <a href='login'>here</a>.").removeClass("hidden");
			$(".btn").attr("disabled", "disabled");
		}	
	}

});