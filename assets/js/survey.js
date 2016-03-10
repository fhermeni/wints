var lang;

function setLang(to, from) {
	lang = to;
	$(to).removeClass("hidden");
	$(from).addClass("hidden");	
}

function submit() {
	var answers = {};
	var missing = false;
	var mark = undefined;
	$(".form-group").each(function (idx, f) {		
		var d =$(f);
		var id = d.attr("id");				

		//Grab the value
		var val = undefined;		
		switch (d.data("type")) {
			case "yesno":
				val = $("[name='" + id + "']:checked").val();
				break;
			case "comment":
			case "rank":
			case "number":
				val = $("[name='" + id + "']").val();
				break;
		}

		//the mark value
		if (d.data("ismark")) {			
			answers["__MARK__"] = val;
		}
		//Control and storage
		if (d.data("required") && !val) {			
			empty("#" + id);
			missing = true;
		} else {			
			cleanError("#" + id);
			if (val) {
				answers[id] = val; 				
			}			
		}
	});
	if (missing) {
		$("#missingRequired").removeClass("hidden");
		return;
	} else {
		$("#missingRequired").addClass("hidden");
	}
	postSurvey($.urlParam("token"), answers).done(commitPost);
}

//Notify the survey has been uploaded
//no double upload, no edit + notification message
function commitPost(s) {
	$(".delivery").html(s.Delivery);
	if ($.urlParam("token")) {	
		$(".commit").removeClass("hidden");
	}
	readOnly();
}

function readOnly() {	
	$("#submit").addClass("hidden");
	$("textarea").attr('readonly','readonly');
	$("input[type=number]").attr('readonly','readonly');
	$("input:radio").attr('disabled','disabled');		
	$("select").attr('disabled','disabled');
}
function loadSurvey(ok) {
	var kind = $.urlParam("kind");
	var token = $.urlParam("token");
	if (token) {		
		surveyFromToken(token).done(function (s) {
			$("#submit").removeClass("hidden");					
			fill(s.Student, s.Tutor, s.Survey);			
		});
	} else {		
		var email = $.urlParam("student");		
		internship(email).done(function(i) {			
			stu = i.Convention.Student.User.Person;			
			tut = i.Convention.Tutor.Person;
			survey = i.Surveys.filter(function(s) {
				return s.Kind == kind;
			})[0]			
			fill(stu, tut, survey);
		});		
	}
}

function fill(stu, tut, survey) {	
	$(".student").html(fn(stu));
	$(".tutor").html(fn(tut));		
	if (survey.Delivery) {				
		commitPost(survey);
		Object.keys(survey.Cnt).forEach(function (k) {
			if (k == "__MARK__") {
				return true;
			}
			var v = survey.Cnt[k];			
			var d = $("#" + k);				
			switch (d.data("type")) {
				case "yesno":
					$("[name='" + k + "'][value='" + v + "']").attr("checked","checked");					
					break;
				case "comment":
				case "number":
				case "rank":
					$("[name='" + k + "']").val(v);
					break;
			}
		});
		$("#submit").addClass("hidden");		
	}
	if (!$.urlParam("token")) {		
		readOnly();		
	}
}

function fn(p) {
	return "<a class='fn' href='mailto:" + p.Email + "'>" + p.Lastname + ", " + p.Firstname + "</a>"
}