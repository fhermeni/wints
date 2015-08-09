var interns;
var currentPage;
var waitingBlock;
var myself;
var currentConvention;
var shiftPressed = false;
var allMajors;


$(document).ready(function() {
	waitingBlock = $("#cnt").clone().html();

	//IE8 stuff
	if (typeof Array.prototype.forEach != 'function') {
		Array.prototype.forEach = function(callback) {
			for (var i = 0; i < this.length; i++) {
				callback.apply(this, [this[i], i, this]);
			}
		};
	}
	majors(function(m) {
		allMajors = m;
		user(getCookie("session"), function(u) {
			myself = u;
			$("#fullname").html(u.Firstname + " " + u.Lastname);
			showMyServices(u.Role);
			if (myself.Role == 0) {
				showDashboard();
			} else if (myself.Role >= 2) {
				showPage(undefined, "conventions");
			} else {
				showPage(undefined, "myStudents");
			}

		}, function() {
			window.location.href = "/login"
		});

	});

	$(document).keydown(function(e) {
		if (e.keyCode == 16) {
			shiftPressed = true;
		}
	});
	$(document).keyup(function(e) {
		if (e.keyCode == 16) {
			shiftPressed = false;
		}
	});
});

function showPage(li, id) {
	$("#cnt").html(waitingBlock);

	$("#menu-pages").find("li").removeClass("active");
	if (li) {
		$(li.parentNode).addClass("active");
	}
	currentPage = id;
	refresh();
}

function summary() {
	var sum = {
		Surveys: {},
		Reports: {}
	}
	interns.forEach(function(i) {
		i.Reports.forEach(function(r) {
			if (!sum.Reports[r.Kind]) {
				sum.Reports[r.Kind] = {
					Missing: 0,
					Pending: 0,
					Total: 0
				}
			}
			if (nbDaysLate(new Date(r.Deadline), new Date()) < 0 && new Date(r.Delivery).getTime() < 0) {
				sum.Reports[r.Kind].Missing++;
			}
			if (new Date(r.Delivery).getTime() > 0 && r.Grade < 0) {
				sum.Reports[r.Kind].Pending++;
			}
			sum.Reports[r.Kind].Total++
		});
		i.Surveys.forEach(function(s) {
			if (!sum.Surveys[s.Kind]) {
				sum.Surveys[s.Kind] = {
					Missing: 0,
					Total: 0
				}
			}
			sum.Surveys[s.Kind].Total++
				if (Object.keys(s.Answers) == 0) {
					sum.Surveys[s.Kind].Missing++
				}
		})
	});
	return sum
}

function editableMajors() {
	var arr = []
	allMajors.forEach(function(p, i) {
		arr.push({
			value: p,
			text: p
		})
	})
	return arr
}

function selectStudent(stu, ok) {
	var ic = $("input[data-email='" + stu + "']").closest(".icheckbox")
	if (ok) {
		ic.addClass("checked")
	} else {
		ic.removeClass("checked")
	}
}

function selectBy(nature, kind, f) {
	interns.forEach(function(i) {
		if (nature == "reports") {
			i.Reports.forEach(function(r) {
				if (r.Kind == kind) {
					if (f == "missing") {
						var late = nbDaysLate(new Date(r.Deadline), new Date()) < 0 && new Date(r.Delivery).getTime() < 0
						selectStudent(i.Student.Email, late)
					} else if (f == "pending") {
						var late = new Date(r.Delivery).getTime() > 0 && r.Grade < 0
						selectStudent(i.Student.Email, late)
					}
				}
				var delayWithDeadline = nbDaysLate(new Date(r.Deadline), new Date(r.Delivery))
				r.Late = nbDaysLate(new Date(r.Deadline), new Date(r.Delivery)) < 0
			})
		} else if (nature == "surveys") {
			i.Surveys.forEach(function(s) {
				if (s.Kind == kind) {
					if (f == "missing") {
						selectStudent(i.Student.Email, Object.keys(s.Answers) == 0)
					}
				}
			});
		}
	});
}

function refresh() {
	//reset the div
	if (currentPage == "myStudents") {
		displayMyStudents();
	} else if (currentPage == "conventions") {
		displayMyConventions();
	} else if (currentPage == "assignments") {
		displayTutors();
	} else if (currentPage == "privileges") {
		showPrivileges();
	} else if (currentPage == "pending") {
		showPendingConventions();
	} else if (currentPage == "juries") {
		showJuryService();
	} else if (currentPage == "service") {
		showService();
	} else if (currentPage == "status") {
		showStatus();
	} else if (currentPage == "defenses-schedule") {
		showDefenses();
	} else if (currentPage == "defense-program") {
		showDefenseProgram()
	} else if (currentPage == "myDefenses") {
		showMyDefenses()
	} else if (currentPage == "alumni") {
		showAlumni()
	} else {
		console.log("Unsupported operation on '" + currentPage + "'");
	}
}

function showMyDefenses() {
	defenses(function(ss) {
		ss = ss.filter(function(s) {
			ok = false
			s.Juries.forEach(function(j) {
				if (j.Email == myself.Email) {
					ok = true
					return false;
				}
			})
			return ok
		});
		console.log(ss)
		var html = Handlebars.getTemplate("jury")(ss);
		var root = $("#cnt");
		root.html(html);
		root.find(":checkbox").icheck();
		root.find("select").selecter();
		root.find(".grade").each(function(i, e) {
			$(e).editable({
				pk: "1",
				url: function(p) {
					return postDefenseGrade($(e).data("email"), parseInt(p.value))
				},
				validate: isGrade
			});
		})
		root.find(".position").each(function(i, e) {
			$(e).editable({
				pk: "1",
				source: editablePositions(),
				url: function(p) {
					return setNextPosition($(e).data("email"), parseInt(p.value))
				}
			});
		})

	});
}

function showAlumni() {
	internships(function(ii) {
		var missings = 0
		ii.forEach(function(i) {
			if (i.Future.Position == 0) {
				missings++
			}
		})
		ii["Missings"] = missings
		var html = Handlebars.getTemplate("alumni")(ii)
		var root = $("#cnt");
		root.html(html).find("td .icheckbox_flat").icheck({
			callbacks: {
				ifChecked: shiftSelect
			}
		});
		$('#cnt').find(".check_all").icheck({
			callbacks: {
				ifChecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("checked")
				},
				ifUnchecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("unchecked")
				}
			}
		});

		root.find(".tablesorter").tablesorter({
			theme: 'bootstrap',
			widgets: ["uitheme"],
			headerTemplate: '{content} {icon}',
			headers: {
				0: {
					sorter: false
				}
			}
		});
	});
}

function showStatus() {
	getStudents(function(stus) {
		var placed = 0
		var known = [];
		var toPlace = stus.length;
		interns.filter(function(u) {
			known.push(u.Student)
		})
		conventions(function(cc) {
			cc.forEach(function(conv) {
				if (conv.Skip) {
					known.push(conv.Student)
				}
			})

			stus.forEach(function(s) {
				if (s.Internship != null && s.Internship.length > 0)  {
					known.forEach(function(k) {
						if (k.Email == s.Internship) {
							s.Internship = k;
							placed++;
						}
						return false;
					});
				} else if (s.Hidden) {
					toPlace--;
				}
			});
			var html = Handlebars.getTemplate("status")({
				Students: stus,
				Interns: known,
				Placed: placed,
				ToPlace: toPlace
			});
			var root = $("#cnt");
			root.html(html);

			$("#student-fullnames").on("click", function() {
				showFullnames(stus);
			});

			$('#cnt').find(":checkbox").icheck();
			var idx = 0;
			stus.forEach(function(stu) {
				s = $("#select-" + (idx++));
				var best = pickBestStudentMatching(stu, known)
				s.val(best);
			})

			//CSV upload      
			$('input[type=file]').filestyle({
				input: false,
				buttonText: "Upload student list",
				buttonName: "btn-success btn-sm",
				iconName: "glyphicon-cloud-upload",
				badge: false
			})
			$(':file').change(function() {
				var file = this.files[0];
				var name = file.name;
				var size = file.size;
				var type = file.type;
				if (type != "text/csv") {
					reportError("Must be a CSV file")
				} else if (size > 1000000) {
					reportError("The file cannot exceed 1MB")
				} else {
					var reader = new FileReader();
					reader.onload = function(e) {
						sendStudents(e.target.result, function() {
							refresh();
						});
					};
					reader.readAsText(file);
				}
			});
		})
	})
}

function sendStudentHidden(em, flag) {
	hideStudent(em, flag, function() {
		reportSuccess("Status updated")
		refresh();
	})
}

function showFullnames(source) {
	var checked = $(".icheckbox.checked").find(":checkbox")
	var fns = [];
	checked.each(function(i, e) {
		var em = $(e).attr("data-email");
		source.forEach(function(src) {
			if (src.Email == em) {
				var fn = src.Firstname.charAt(0).toUpperCase() + src.Firstname.slice(1);
				var ln = src.Lastname.charAt(0).toUpperCase() + src.Lastname.slice(1);
				var n = fn + " " + ln;
				if (fns.indexOf(n) < 0) {
					fns.push(fn + " " + ln)
				}
			}
		})
	});
	if (fns.length > 0) {
		var html = Handlebars.getTemplate("raw");
		$("#modal").html(html);
		$("#rawContent").html(fns.join("\n"));
		$("#modal").modal('show');
	}
}

function sendStudentAlignement(stu, internship) {
	if (internship == "") {
		alignStudentWithInternship(stu, "", function() {
			reportSuccess("reset done");
			refresh();
		})
	} else {
		alignStudentWithInternship(stu, $("#" + internship).val(), function() {
			reportSuccess("alignment done");
			refresh();
		})
	}
}

function pickBestStudentMatching(student, students) {
	var best = "";
	students.forEach(function(s) {
		//email match        
		if (s.Email ==
			student.Email) {
			best = s.Email;
			return false;
		}
	});
	//Lastname match
	r = student.Lastname.toLowerCase();
	r = r.replace(new RegExp(/[àáâãäå]/g), "a");
	r = r.replace(new RegExp(/ç/g), "c");
	r = r.replace(new RegExp(/[èéêë]/g), "e");
	r = r.replace(new RegExp(/[ìíîï]/g), "i");
	r = r.replace(new RegExp(/[òóôõö]/g), "o");
	r = r.replace(new RegExp(/œ/g), "oe");
	r = r.replace(new RegExp(/[ùúûü]/g), "u");
	r = r.replace(new RegExp(/[ýÿ]/g), "y");
	students.forEach(function(s) {
		if (s.Lastname == r) {
			best = s.Email;
			return false;
		}
	});
	return best;
}


function displayMyStudents() {
	internships(function(data) {
		interns = data.filter(function(i) {
			return i.Tutor.Email == myself.Email
		});
		var html = Handlebars.getTemplate("tutoring")(interns);
		var root = $("#cnt");
		root.html(html);
		if (interns.length == 0) {
			return true
		}

		$("#cnt").find("td .icheckbox_flat").icheck({
			callbacks: {
				ifChecked: shiftSelect
			}
		});
		$('#cnt').find(".check_all").icheck({
			callbacks: {
				ifChecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("checked")
				},
				ifUnchecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("unchecked")
				}
			}
		});
		$("#table-conventions").tablesorter({
			theme: 'bootstrap',
			widgets: ["uitheme"],
			headerTemplate: '{content} {icon}',
			headers: {
				0: {
					sorter: false
				}
			}
		});
	});
}

function displayMyConventions() {
	internships(function(data) {
		interns = data;
		var html = Handlebars.getTemplate("watchlist")(interns);
		var root = $("#cnt");
		root.html(html);

		if (myself.Role >= 3) {
			var sum = summary()
			$(".summary").html(Handlebars.getTemplate("summary")(sum))
		}
		if (interns.length == 0) {
			return true
		}

		$("#cnt").find("td .icheckbox_flat").icheck({
			callbacks: {
				ifChecked: shiftSelect
			}
		});
		$('#cnt').find(".check_all").icheck({
			callbacks: {
				ifChecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("checked")
				},
				ifUnchecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("unchecked")
				}
			}
		});
		$("#table-conventions").tablesorter({
			theme: 'bootstrap',
			widgets: ["uitheme"],
			headerTemplate: '{content} {icon}',
			headers: {
				0: {
					sorter: false
				}
			}
		});
		if (myself.Role >= 3) {
			root.find(".grade").each(function(i, e) {
				$(e).editable({
					pk: "1",
					url: function(p) {
						return postDefenseGrade($(e).data("email"), parseInt(p.value))
					},
					validate: isGrade,
					success: function(r, v) {
						var p = $(e).parent()
						if (parseInt(v) < 10) {
							p.removeClass("bg-success").addClass("bg-danger")
						} else {
							p.removeClass("bg-danger").addClass("bg-success")
						}
					}
				});
			});

			users(function(uss) {
				uss = uss.filter(function(u) {
					return u.Role != 0; //get rid of students
				});
				var tutors = []
				uss.forEach(function(e) {
					tutors.push({
						value: e.Email,
						text: e.Firstname.capitalize() + " " + e.Lastname.capitalize()
					})
				});
				root.find(".tutor").each(function(i, e) {
					$(e).editable({
						pk: "1",
						source: tutors,
						value: getInternship($(e).data("email")).Tutor.Email,
						url: function(p) {
							return setTutor($(e).data("email"), p.value)
						}
					});
				})
			});
		}

		if (myself.Role >= 2) {
			//Major rights and more
			root.find(".major").each(function(i, e) {
				$(e).editable({
					pk: "1",
					source: editableMajors(),
					value: getInternship($(e).data("email")).Major,
					url: function(p) {
						return setMajor($(e).data("email"), p.value)
					}
				});
			});
		}
	});
}


function shiftSelect(e) {
	document.getSelection().removeAllRanges();
	if (shiftPressed) {
		//debugger

		var myTd = $(e).closest("td")
		var tr = myTd.parent()
		var col = tr.children().index(myTd)
		var p = tr.prev();
		while (p.length > 0) {
			var chk = $(p.children().get(col)).find(".icheckbox")
			if (chk.hasClass("checked")) {
				break;
			} else {
				chk.icheck("checked")
			}
			p = p.prev();
		}
	}
}

function showPrivileges() {
	sessions(function(ss) {

		users(function(data) {
			var teachers = [];
			var students = [];
			data.filter(function(u) {
				u.Last = ss[u.Email];
				if (u.Role == 0 && myself.Email != u.Email) {
					students.push(u)
				} else if (myself.Email != u.Email) {
					teachers.push(u);
				}
			})
			var html = Handlebars.getTemplate("privileges")({
				Students: students,
				Teachers: teachers
			});
			$("#cnt").html(html);
			$("#cnt").find("select").selecter();
			$('#cnt').find(":checkbox").icheck()
			$('[data-toggle="delteacher-confirmation"]').each(function(i, a) {
				var j = $(a);
				j.confirmation({
					onConfirm: function() {
						removeUser(j.attr("data-user"), j.parent().parent().parent())
					}
				});
			});
			$('[data-toggle="teacher-reinvite"]').each(function(i, a) {
				var j = $(a);
				j.confirmation({
					onConfirm: function() {
						reInvite(j.attr("data-user"))
					},
					btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'
				});
			});
			$('[data-toggle="student-reinvite"]').each(function(i, a) {
				var j = $(a);
				j.confirmation({
					onConfirm: function() {
						resetPassword(j.attr("data-user"))
					},
					btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'
				});
			});
			$('[data-toggle="delstudent-confirmation"]').each(function(i, a) {
				var j = $(a);
				j.confirmation({
					onConfirm: function() {
						removeUser(j.attr("data-user"), j.parent().parent())
					}
				});
			});
		});
	});
}

function showNewUser() {
	var buf = Handlebars.getTemplate("new-user")({});
	$("#modal").html(buf).modal('show');
	$("#modal").find("select").selecter();
}

function addUser() {
	if (missing("lbl-nu-fn") || missing("lbl-nu-ln") || missing("lbl-nu-email")) {
		return false;
	}
	newUser($("#lbl-nu-fn").val(), $("#lbl-nu-ln").val(),
		$("#lbl-nu-tel").val(), $("#lbl-nu-email").val(),
		parseInt($("#lbl-nu-priv").val()),
		function() {
			$("#modal").modal('hide');
			reportSuccess("Account created");
			refresh();
		},
		function(o) {
			if (o.status == 409) {
				$("#lbl-nu-email").notify(o.responseText);
			}
		});
}

function removeUser(email, div) {
	rmUser(email, function() {
		div.remove();
		reportSuccess("Account deleted")
	});
}

function pickBestMatching(tutor, users) {
	var res = undefined;
	users.forEach(function(t) {
		var known_ln = t.Lastname;
		if (tutor.Lastname.indexOf(known_ln) > -1 || known_ln.indexOf(tutor.Lastname) > -1) {
			res = t;
			return false;
		}
	});
	if (res == undefined) {
		var firstLetter = tutor.Lastname[0];
		users.forEach(function(t) {
			var knownFirstLetter = t.Lastname[0];
			if (firstLetter >= knownFirstLetter) {
				res = t;
			}
		});
	}
	if (!res) {
		res = users[0]
	}
	return res;
}

function showPendingConventions() {
	conventions(function(cc) {
		users(function(uss) {
			uss = uss.filter(function(u) {
				return u.Role != 0;
			});
			var ignored = []
			cc = cc.filter(function(c) {
				if (c.Skip) {
					ignored.push(c)
				}
				return !c.Skip
			});
			$("#pending-counter").html(" <span class='badge'>" + cc.length + (ignored.length > 0 ? "/" + ignored.length : "") + "</span>");
			if (cc.length == 0 && ignored.length == 0) {
				$("#pending-counter").html("");
			}
			currentConvention = cc[0]
			var html = Handlebars.getTemplate("pending")({
				c: currentConvention,
				users: uss,
				"Ignored": ignored
			});
			$("#cnt").html(html);
			if (cc.length > 0) {
				//Find the most appropriate predefined tutor            
				var best = pickBestMatching(currentConvention.Tutor, uss)
				$("#known-tutor-selector").val(best.Email)
				$("#cnt").find("select").selecter();

				if (best.Lastname == currentConvention.Tutor.Lastname) {
					d = $('#pick-theory');
					d.confirmation({
						onConfirm: pickTheory,
						placement: "right",
						title: "Sure ? It's a perfect match !",
						btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'
					});
					d.removeAttr("onclick");
					k = $("#btn-choose-known");
					k.attr("onclick", "pickKnown()");
					k.confirmation('destroy');
				} else {
					d = $('#pick-theory');
					d.attr("onclick", "pickTheory()");
					d.confirmation('destroy');

					k = $("#btn-choose-known");
					k.confirmation({
						onConfirm: pickKnown,
						placement: "left",
						title: "Sure ? They differ !",
						btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'
					});
					k.removeAttr("onclick");
				}
			}
		});
	});
}

function sendSkipConvention(email, s) {
	skipConvention(email, s, function() {
		if (s) {
			reportSuccess("Convention ignored")
		} else {
			reportSuccess("The convention is no longer ignored")
		}
		refresh()
	})
}

function sendDeleteConvention(email) {
	deleteConvention(email, function() {
		reportSuccess("Convention deleted")
		refresh()
	})
}

function pickTheory() {
	if (missing("th-tutor-fn") || missing("th-tutor-ln") || missing("th-tutor-email") || missing("th-tutor-tel")) {
		return false;
	}
	var fn = $("#th-tutor-fn").val();
	var ln = $("#th-tutor-ln").val();
	var email = $("#th-tutor-email").val();
	var tel = $("#th-tutor-tel").val();
	newUser(fn, ln, tel, email, -1, function() {
		currentConvention.Tutor = {
				Firstname: fn,
				Lastname: ln,
				Tel: tel,
				Email: email,
				Role: 1
			} //tutor     
		reportSuccess("Tutor account created")
		newInternship(currentConvention, function() {
			reportSuccess("Internship added");
			refresh()
		})
	}, function(o) {
		if (o.status == 409) {
			reportSuccess("Tutor account already exists")
			newInternship(currentConvention, function() {
				reportSuccess("Internship added");
				refresh()
			})
		} else {
			reportError(o.responseText);
		}
	})
}

function gradeReport(toGrade, email, kind) {
	if ((toGrade && missing("grade")) || missing("comment")) {
		return
	}
	var grade = $("#grade").val();
	var comment = $("#comment").val();
	if (comment.length < 3) {
		$("#comment").notify("required", {
			className: "danger"
		});
		return
	}
	setReportGrade(email, kind, parseInt(grade), comment, function() {
		reportSuccess("Operation successfull")
		$("#modal").modal('hide');
		refresh();
	})
}

function pickKnown() {
	user($("#known-tutor-selector").val(), function(us) {
		currentConvention.Tutor = us
		newInternship(currentConvention, function() {
			reportSuccess("Internship added");
			refresh()
		})
	});
}

function serviceMailing() {
	var to = [];
	$(".icheckbox.checked").find(":checkbox").each(function(i, c) {
		var em = $(c).attr("data-email")
		to.push(em);
	});
	if (to.length > 0) {
		window.location.href = "mailto:" + to.join(",");
	}
}

function mailing(t, w) {
	var to = [];
	var cc = [];
	$(".icheckbox.checked").find(":checkbox").each(function(i, c) {
		var em = $(c).attr("data-email")
		var i = getInternship(em);
		if (i) {
			if (t == "students") {
				to.push(i.Student.Email)
			} else if (t == "tutors") {
				to.push(i.Tutor.Email)
			} else if (t == "supervisors") {
				to.push(i.Sup.Email)
			}

			if (w == "students") {
				cc.push(i.Student.Email)
			} else if (w == "tutors") {
				cc.push(i.Tutor.Email)
			} else if (w == "supervisors") {
				cc.push(i.Sup.Email)
			}
		}
	});
	if (to.length > 0) {
		window.location.href = "mailto:" + to.join(",") + (cc.length > 0 ? "?cc=" + cc.join(",") : "");
	}
}

function getInternship(email) {
	var i = undefined;
	interns.forEach(function(e) {
		if (e.Student.Email == email) {
			i = e;
			return false;
		}
	})
	return i;
}

function showService() {
	internships(function(interns) {
		var service = {};
		interns.forEach(function(i) {
			if (!service[i.Tutor.Email]) {
				service[i.Tutor.Email] = {
					P: i.Tutor,
					Internships: {}
				};
			}
			if (!service[i.Tutor.Email].Internships[i.Promotion]) {
				service[i.Tutor.Email].Internships[i.Promotion] = []
			}
			service[i.Tutor.Email].Internships[i.Promotion].push(i)
		});
		var html = Handlebars.getTemplate("service")(service);
		$("#cnt").html(html);

		$("#cnt").find("td .icheckbox_flat").icheck({
			callbacks: {
				ifChecked: shiftSelect
			}
		});
		$('#cnt').find(".check_all").icheck({
			callbacks: {
				ifChecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("checked")
				},
				ifUnchecked: function(e) {
					$("#cnt").find("td .icheckbox_flat").icheck("unchecked")
				}
			}
		});
	});
}

function alumniToCSV() {
	internships(function(interns) {
		var csv = "firstname;lastname;oldEmail;newEmail;position;txtPosition\n";
		interns.forEach(function(i) {
			csv += i.Student.Firstname + ";" + i.Student.Lastname + ";" + i.Student.Email + ";" + i.Future.Contact + ";" + i.Future.Position + ";" + possiblePositions[i.Future.Position] + "\n";
		});
		var html = Handlebars.getTemplate("raw")();
		$("#modal").html(html).modal('show');
		$("#rawContent").html(csv);
	});
}

function showRawService() {
	internships(function(interns) {
		var service = {};
		interns.forEach(function(i) {
			if (!service[i.Tutor.Email]) {
				service[i.Tutor.Email] = {
					P: i.Tutor,
					Internships: {}
				};
			}
			if (!service[i.Tutor.Email].Internships[i.Promotion]) {
				service[i.Tutor.Email].Internships[i.Promotion] = []
			}
			service[i.Tutor.Email].Internships[i.Promotion].push(i)
		});
		var html = Handlebars.getTemplate("raw-service")(service);
		$("#modal").html(html).modal('show');
	});
}

function showRawFullname(kind) {
	var checked = $(".icheckbox.checked").find(":checkbox")
	var fns = [];
	checked.each(function(i, e) {
		var em = $(e).attr("data-email");
		if (em) {
			var i = getInternship(em);
			var p = i.Student;
			if (kind == "supervisors") {
				p = i.Sup

			} else if (kind == "tutors") {
				p = i.Tutor
			}
			var fn = p.Firstname.charAt(0).toUpperCase() + p.Firstname.slice(1);
			var ln = p.Lastname.charAt(0).toUpperCase() + p.Lastname.slice(1);
			var n = fn + " " + ln;
			if (fns.indexOf(n) < 0) {
				fns.push(fn + " " + ln)
			}
		}
	});
	if (fns.length > 0) {
		var html = Handlebars.getTemplate("raw");
		$("#modal").html(html);
		$("#rawContent").html(fns.join("\n"));
		$("#modal").modal('show');
	}
}

function align(txt, to) {
	var b = txt;
	for (i = b.length; i < to; i++) {
		b += " ";
	}
	return b;
}

function selectText(elm) {
	// for Internet Explorer
	if (document.body.createTextRange) {
		var range = document.body.createTextRange();
		range.moveToElementText(elm);
		range.select();
	} else if (window.getSelection) {
		// other browsers
		var selection = window.getSelection();
		var range = document.createRange();
		range.selectNodeContents(elm);
		selection.removeAllRanges();
		selection.addRange(range);
	}
}

function showInternship(s) {
	internship(s, function(i) {
		users(function(uss) {
			uss = uss.filter(function(u) {
				return u.Role != 0; //get rid of students
			});
			i.Defense.Privileges = false;
			if (i.Defense.Defenses) {
				i.Defense.Mine = i.Defense.Defenses[0]
				i.Defense.Privileges = i.Defense.Defenses[0].Remote || i.Defense.Defenses[0].Private
			}

			buf = Handlebars.getTemplate("student")({
				I: i,
				SurveyAdmin: myself.Email == i.Tutor.Email || myself.Role >= 3,
				URL: window.location.protocol + "//" + window.location.host + "/surveys/"
			})
			$("#modal").html(buf).modal('show');
			var c = $("#modal").find("select.select-tutor");
			c.val(i.Tutor.Email)
			$("#modal").find("select.select-major").selecter({
				callback: function(v) {
					sendMajor(i.Student.Email, v)
				}
			});
			$("#modal").find("select.select-tutor").selecter({
				callback: function(v) {
					sendTutor(i.Student.Email, v)
				}
			});
		});
	});
}

function showReport(email, kind) {
	reportHeader(email, kind, function(r) {
		r.Passed = new Date(r.Deadline).getTime() < new Date()
		r.Late = nbDaysLate(new Date(r.Deadline), new Date(r.Delivery)) < 0
		r.In = r.Grade != -2
		r.Reviewable = r.Grade != -2 || r.Passed
		r.Gradeable = r.ToGrade && (r.Grade != -2 || r.Passed)
		r.Email = email
		var i = getInternship(email);
		r.Mine = i.Tutor.Email == myself.Email;
		r.Reviewed = r.Comment.length > 0 || r.Grade >= 0
		buf = Handlebars.getTemplate("reportEditor")(r)
		$("#modal").html(buf).modal('show');
		$('#modal').find(':checkbox').icheck()
			.on('ifChecked', function() {
				setReportPrivate(email, kind, true)
			})
			.on('ifUnchecked', function() {
				setReportPrivate(email, kind, false)
			})
		$(".date").datepicker({
			format: 'd M yyyy',
			autoclose: true,
			minViewMode: 0,
			weekStart: 1
		}).on("changeDate", function(e) {
			setReportDeadline(email, kind, e.date)
		})
	});
}

function nbDaysLate(deadline, now) {
	if (!now) {
		now = new Date()
	}
	var d1 = moment(deadline)
	var d2 = moment(now)
	return d1 - d2;
}

function sendMajor(e, m) {
	setMajor(e, m, function() {
		reportSuccess("Major updated")
		refresh()
	})
}

function sendTutor(e, s) {
	setTutor(e, s, function() {
		reportSuccess("Tutor updated")
		refresh();
	})
}

function showDefense(stu) {
	i = getInternship(stu)
	def = i.Defense
	def.Student = i.Student
	def.Gradeable = myself.Role >= 3
		//Correct date
	if (def.Defenses) {
		def.Date = moment(def.Date).add(def.Defenses[0].Offset * 30, "minutes")
		def.Remote = def.Defenses[0].Remote
		def.Private = def.Defenses[0].Private
		def.Grade = def.Defenses[0].Grade
	}
	def.Juries.forEach(function(j) {
		if (j.Email == myself.Email) {
			def.Gradeable = true
			return false
		}
	})
	buf = Handlebars.getTemplate("defense-modal")(def)
	$("#modal").html(buf).modal('show');
	$('#modal').find(":checkbox").icheck();
	$('#modal').find("select").selecter();
}

function gradeDefense(stu) {
	var g = $("#modal").find("#grade").val()
	postDefenseGrade(stu, parseInt(g), function() {
		$("#modal").find("#grade").val(g).removeClass("has-error")
		$("#modal").modal('hide')
		refresh();
	})
}
String.prototype.capitalize = function() {
	return this.charAt(0).toUpperCase() + this.substring(1)
}

function tplEvaluationMail(kind, student, url) {
	var i = getInternship(student);
	var txt = Handlebars.getTemplate("eval-" + kind)({
		I: i,
		URL: url
	});
	var to = encodeURIComponent(i.Sup.Email)
	var s = encodeURIComponent(i.Student.Firstname.capitalize() + " " + i.Student.Lastname.capitalize() + " - Evaluation");
	window.location.href = "mailto:" + to + "?subject=" + s + "&body=" + encodeURIComponent(txt);
}

function requestSurveys(kind) {
	$(".icheckbox.checked").find(":checkbox").each(function(i, c) {
		var em = $(c).attr("data-email")
		requestSurvey(em, kind)
	});
}