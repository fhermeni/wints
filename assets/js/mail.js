//mail.js
function conventionMailing(ctx, t, w) {
	var to = [];
	var cc = [];
	$(ctx).find("input:checked").each(function(i, x) {
		var c = $(x);
		to.push(c.data(t));
		if (w) {
			cc.push(c.data(w));
		}
	});
	sendMail(to, cc);
}

function sendMail(to, cc) {
	if (to.length > 0) {
		window.location.href = "mailto:" + to.join(",") + (cc.length > 0 ? "?cc=" + cc.join(",") : "");
	}
}

function userMailing(ctx) {
	var to = [];
	$(ctx).find("input:checked").each(function(i, c) {
		var em = $(c).data("email");	
		to.push(em);
	});
	sendMail(to, []);
}