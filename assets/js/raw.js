function showRawFullname(ctx, kind) {
	var k = kind ? kind + "-fn" : "email";
	var fns = [];
	$(ctx).find("input:checked").each(function(i, x) {
		var c = $(x);
		fns.push(c.data(k).capitalize());
	});
	if (fns.length > 0) {
		$("#modal").render("raw", {
			Title: fns.length + " person(s)",
			Cnt: fns.join("\n")
		}, showModal);
	}
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