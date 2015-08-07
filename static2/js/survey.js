var survey;

function fn(p) {
	return "<a href='mailto:" + p.Email + "'>" + p.Firstname + " " + p.Lastname + "</a>"
}

function setLang(l) {
	var lib = lang();
	if (lib['en'].length != lib['fr'].length) {
		console.log("Warning language library differ: en=" + lib['en'].length + "; fr=" + lib['fr'].length)
	}
	$(".tr").each(function(i, e) {
		$(e).html(lib[l][i]);
	});
	//Basic words
	$('#missing-fields').html(l == "fr" ? "Certains champs n'ont pas été remplis" : "missing required fields");
	$('#committed').html(l == "fr" ? "Votre avis a été remonté à son tuteur. Merci pour votre implication" : "Your feedbacks have been sent to the tutor. Thanks for your involvement.");
	$('.y').html(l == "fr" ? "oui" : "yes")
	$('.n').html(l == "fr" ? "non" : "no")
	$('.submit').html(l == "fr" ? "Soumettre" : "Submit")
	$('.c').html(l == "fr" ? "Commentaires" : "Comments")
	$('.stu').html(l == "fr" ? "Stagiaire" : "Trainee")
	$('.tut').html(l == "fr" ? "Tuteur" : "Tutor")
	$('.ex').html(l == "fr" ? "Excellent" : "Excellent")
	$('.go').html(l == "fr" ? "Bon" : "Good")
	$('.sat').html(l == "fr" ? "Satisfaisant" : "Satisfactory")
	$('.unsat').html(l == "fr" ? "Insatisfaisant" : "Non satisfactory")
	$('.na').html(l == "fr" ? "Non applicable" : "Non applicable")
}

function choices() {
	$(".choice").each(function(idx, e) {
		var buf = "<select name='" + e.id + "' class='col-md-3'>";
		buf += "<option value=''> - </option>";
		buf += "<option value='4' class='ex'></option>";
		buf += "<option value='3' class='go'></option>";
		buf += "<option value='2' class='sat'></option>";
		buf += "<option value='1' class='unsat'></option>";
		buf += "<option value='0' class='na'></option>";
		buf += "</select>";
		$(e).after(buf);
	});
}

function showAnswers(answers) {
	Object.keys(answers).forEach(function(k) {
		var v = answers[k];
		if (v) {
			if (v == "true" || v == "false") {
				var q = $(":radio[name='" + k + "'][value='" + v + "']").iCheck("check")
				if (q.length > 0) {
					q.iCheck("check")
				} else {
					$("[name='" + k + "']").val(v);
				}
			} else {
				$("[name='" + k + "']").val(v);
			}
		}
	})
}

function readOnly() {
	$("#submit").remove()
	$("input").prop('disabled', true)
	$("textarea").prop('disabled', true)
	$("select").selecter("disable")
}

function checked(v) {
	var tag = $("[name='" + v + "']");
	var nn = tag[0].nodeName
	var b
	if (nn == "SELECT" || (nn == "INPUT" && tag.length == 1)) {
		b = tag.val().length > 0;
	} else if (nn == "INPUT" && tag.length == 2) {
		b = $("[name='" + v + "']:checked").val() != undefined;
	} else {
		console.log("unsupported: " + nn);
		return false;
	}
	if (!b) {
		$("#" + v).css('color', 'red');
	} else {
		$("#" + v).css('color', '');
	}
	return b;
}

function yesno() {
	$(".yesno").each(function(idx, e) {
		var buf = " <input type='radio' name='" + e.id + "' value='true'/><span class='y'></span>"
		buf += " <input type='radio' name='" + e.id + "' value='false'/><span class='n'></span>"
		$(e).after(buf);
	});
}

function textarea() {
	$(".textarea").each(function(idx, e) {
		var buf = "<br/><textarea cols='80' name='" + e.id + "'></textarea><br/>";
		$(e).after(buf);
	});
}

function number() {
	$(".number").each(function(idx, e) {
		var buf = " <input type='number' name='" + e.id + "'/>";
		$(e).after(buf);
	});
}

function loadSurvey(ok) {
	var token = $.urlParam("token")
	if (token) {
		longSurvey(token, function(s) {
			var d = new Date(s.Timestamp)
			survey = s;
			ok()
		}, function(jqr) {
			$("#errorMessage").html(jqr.responseText)
			$("#modal").modal('show')
		})
	} else {
		var email = $.urlParam("student")
		var kind = $.urlParam("kind")
		internship(email, function(i) {
			var url = window.location.protocol + "//" + window.location.host + "/surveys/" + kind + "?token="
			survey = i.Surveys.filter(function(s) {
				return s.Kind == kind
			})[0]
			survey.Student = i.Student
			survey.Tutor = i.Tutor
			ok()
		})
	}
}

function fullfill() {
	$("#student").html(fn(survey.Student))
	$("#tutor").html(fn(survey.Tutor))
	var d = new Date(survey.Timestamp)
	if (d.getTime() > 0) {
		if ($.urlParam("token")) {
			$(".alert-success").removeClass("hidden");
		}
		showAnswers(survey.Answers);
		readOnly();
	}
}

String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.substring(1)
}

function tplEvaluationMail() {
	var email = $.urlParam("student")
	var kind = $.urlParam("kind")
	internship(email, function(i) {
		var txt = Handlebars.getTemplate("eval-" + kind)({
			I: i,
			URL: getToken(i, kind)
		});
		var to = encodeURIComponent(i.Sup.Email)
		var s = encodeURIComponent(i.Student.Firstname.capitalize() + " " + i.Student.Lastname.capitalize() + " - Evaluation");
		window.location.href = "mailto:" + to + "?subject=" + s + "&body=" + encodeURIComponent(txt);
	});
}

function expand() {
	yesno();
	textarea();
	choices();
	number();
	setLang("fr")
}

function ui() {
	$(":radio").iCheck();
	$("select").selecter();
	var d = new Date(survey.Timestamp)
	if (d.getTime() > 0) {
		readOnly();
	}
	var token = $.urlParam("token")
	if (!token && d.getTime() < 0) {
		var url = window.location.protocol + "//" + window.location.host + "/surveys/" + survey.Kind + "?token=" + survey.Token
			//supervisor version
		d = $(".alert-warning");
		d.find(".token").html(url)
		d.removeClass("hidden")
		readOnly();
		$(".content").addClass("hidden")
	}
}

function submit() {
	//required fields
	var ok = true;
	$('.r').each(function(idx, e) {
		if (!checked(e.id)) {
			ok = false;
		}
	});
	if (!ok) {
		$(".alert-danger").show()
		return
	} else {
		$(".alert-danger").hide()
	}
	//Extract the content and format it
	//input fields are binary, textarea contains string
	cnt = {}
	$("input:checked").each(function(idx, i) {
		cnt[i.name] = i.value;
	});
	$("textarea").each(function(idx, i) {
		cnt[i.name] = i.value;
	});
	$("select").each(function(idx, i) {
		cnt[i.name] = i.value;
	});
	$("input[type=number]").each(function(idx, i) {
		cnt[i.name] = i.value;
	});
	setSurveyAnswers($.urlParam("token"), cnt,
		function() {
			$(".alert-success").removeClass("hidden")
			readOnly();
			window.location.href = "#"
		});

}