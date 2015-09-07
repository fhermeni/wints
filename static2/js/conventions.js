function full(u) {
	return u.Person.Lastname + "," + u.Person.Firstname + " (" + u.Person.Email + ")";
}

var regexps = [
	[new RegExp(/[àáâãäå]/g), "a"],
	[new RegExp(/ç/g), "c"],
	[new RegExp(/[èéêë]/g), "e"],
	[new RegExp(/[ìíîï]/g), "i"],
	[new RegExp(/[òóôõö]/g), "o"],
	[new RegExp(/œ/g), "oe"],
	[new RegExp(/[ùúûü]/g), "u"],
	[new RegExp(/[ýÿ]/g), "y"],
	[new RegExp(/\-/g), ""],
	[new RegExp(/\s/g), ""]
];

function showConventions() {
	if (myself.Role >= 4) {
		$.when(conventions(), users()).done(reduce).fail(logFail)
	} else {
		$.when(conventions(), users()).done(loadConventions).fail(logFail)
	}


}

function merge(u, sames) {
	if (sames.length > 0) {
		console.log(full(u) + " gathers " + sames.length + " accounts");
		sames.forEach(function(s) {
			replaceUserWith(s.Person.Email, u.Person.Email).done(doneMerge).fail(logFail);
		});
	}
}

function doneMerge() {
	console.log("merged");
}

function reduce(convs, us) {
	console.log("Reducing user accounts");
	var convs = convs[0];
	var knownTeachers = us[0].filter(function(u) {
		return u.Role > 1;
	});
	var unknownTeachers = us[0].filter(function(u) {
		return u.Role == 0;
	});
	students = us[0].filter(function(u) {
		return u.Role == 1;
	});

	//We trust known teachers so try to identify homonyms in the unknown teachers
	knownTeachers.forEach(function(known) {
		var sames = homonyms(known, unknownTeachers);
		merge(known, sames);
		//Maintain the remaining pool
		unknownTeachers = unknownTeachers.filter(function(el) {
			return sames.indexOf(el) < 0;
		});
	});
	//Now we reduce the pool of unknown teachers
	while (unknownTeachers.length != 0) {
		u = unknownTeachers[0];
		var sames = homonyms(u, unknownTeachers);
		merge(u, sames);
		sames.push(u); //to converge
		unknownTeachers = unknownTeachers.filter(function(el) {
			return sames.indexOf(el) < 0;
		});
	}

	//Now we do the display with fresh data
	$.when(conventions(), users()).done(loadConventions).fail(logFail);
}

function loadConventions(conventions, users) {
	var conventions = conventions[0];
	$("#cnt").render("conventions-header", conventions, ui);
}

function clean(u) {
	var r = u;
	regexps.forEach(function(rx) {
		r = r.replace(rx[0], rx[1]);
	});
	return r;
}

//return users in the pool having the same firstname and lastname
function homonyms(user, pool) {
	var res = [];
	var fn = clean(user.Person.Firstname);
	var ln = clean(user.Person.Lastname);
	pool.forEach(function(u) {
		if (user.Person.Email != u.Person.Email && clean(u.Person.Firstname) == fn && clean(u.Person.Lastname) == ln) {
			res.push(u);
		}
	});
	return res;
}