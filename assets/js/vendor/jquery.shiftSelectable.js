$.fn.shiftSelectable = function() {
	var lastChecked,
		$boxes = this;
	$boxes.off('click');		
	$boxes.click(function(evt) {
		if (!lastChecked) {
			lastChecked = this;
			return;
		}
		if (evt.shiftKey) {
			var start = $boxes.index(this),
				end = $boxes.index(lastChecked);
			$boxes.slice(Math.min(start, end), Math.max(start, end) + 1)
				.prop('checked', lastChecked.checked)
				.trigger('change');
		}

		lastChecked = this;
	});
};