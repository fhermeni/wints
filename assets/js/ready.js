$(document).ready(function() {

moment.locale('fr', {
	months: "janvier_février_mars_avril_mai_juin_juillet_août_septembre_octobre_novembre_décembre".split("_"),
	monthsShort: "janv._févr._mars_avr._mai_juin_juil._août_sept._oct._nov._déc.".split("_"),
	weekdays: "dimanche_lundi_mardi_mercredi_jeudi_vendredi_samedi".split("_"),
	weekdaysShort: "dim._lun._mar._mer._jeu._ven._sam.".split("_"),
	weekdaysMin: "Di_Lu_Ma_Me_Je_Ve_Sa".split("_"),
	longDateFormat: {
		LT: "HH:mm",
		LTS: "HH:mm:ss",
		L: "DD/MM/YYYY",
		LL: "D MMMM YYYY",
		LLL: "D MMMM YYYY LT",
		LLLL: "dddd D MMMM YYYY LT"
	},
	calendar: {
		sameDay: "[Aujourd'hui à] LT",
		nextDay: '[Demain à] LT',
		nextWeek: 'dddd [à] LT',
		lastDay: '[Hier à] LT',
		lastWeek: 'dddd [dernier à] LT',
		sameElse: 'L'
	},
	relativeTime: {
		future: "dans %s",
		past: "il y a %s",
		s: "quelques secondes",
		m: "une minute",
		mm: "%d minutes",
		h: "une heure",
		hh: "%d heures",
		d: "un jour",
		dd: "%d jours",
		M: "un mois",
		MM: "%d mois",
		y: "une année",
		yy: "%d années"
	},
	ordinalParse: /\d{1,2}(er|ème)/,
	ordinal: function(number) {
		return number + (number === 1 ? 'er' : 'ème');
	},
	meridiemParse: /PD|MD/,
	isPM: function(input) {
		return input.charAt(0) === 'M';
	},
	// in case the meridiem units are not separated around 12, then implement
	// this function (look at locale/id.js for an example)
	// meridiemHour : function (hour, meridiem) {
	//     return  0-23 hour, given meridiem token and hour 1-12
	// },
	meridiem: function(hours, minutes, isLower) {
		return hours < 12 ? 'PD' : 'MD';
	},
	week: {
		dow: 1, // Monday is the first day of the week.
		doy: 4 // The week that contains Jan 4th is the first week of the year.
	}
});
moment.locale('en');


	$('[data-toggle="popover"]').popover();


	getConfig().done(function(c) {
		config = c;
		$(".release").html(config.Version);
	});

	//Login part
	if (window.location.pathname == "/login") {
		cookie = getCookie("login");
		if (cookie) {
			//Already identified, go to the dashboard
			window.location="/home";
		}
		var em = $.urlParam("email");
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
		$.tablesorter.defaults.widgets = ["uitheme"];
		$.tablesorter.defaults.theme = 'bootstrap';
		$.tablesorter.defaults.headerTemplate = '{content} {icon}';

		cookie = getCookie("login");
		if (cookie) {
			user(cookie).done(loadSuccess).fail(function() {
				window.location="/login#sessionExpired";
			});
		} else {
			window.location="/login";
		}
	} else if (window.location.pathname == "/survey") {
		setLang(".fr",".en");
		loadSurvey();
	} else if (window.location.pathname == "/password") {
		var t = $.urlParam("token");
		if (!t) {
			$(".alert-danger").html("<strong>There is no token in the request</strong>. Initiate a reset request <a href='login'>here</a>.").removeClass("hidden");
			$(".btn").attr("disabled", "disabled");
		}
	} else if (window.location.pathname == "/statistics") {
		loadStatistics();
	} else if (window.location.pathname == "/defense-program") {
		loadDefenseProgram();
	}

});