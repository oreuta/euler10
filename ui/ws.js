window.onload = initPage;

function initPage() {
	$('#btnCalc').click(function() {
		var sum = $('#sum');
		sum.html("Waiting for reply...");
		var n = $('#n').val();
		$.ajax({
			url: 'http://localhost:5050/sum',
			method: 'post',
			data: 
				'{"n": '+n+', "list": false}',
			dataType: 'json',
			contentType: 'application/json',
			success: function(data) {
						sum.html(data.sum);
			},
		});
	});
};