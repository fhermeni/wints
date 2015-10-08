function showWatchlist() {
	internships().done(loadWatchlist);
}

function loadWatchlist(interns) {
	$("#cnt").render("watchlist", {
		Internships: interns,
		Org: config
	}, ui);
}