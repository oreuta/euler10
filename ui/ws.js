window.onload = initPage;

function initPage() {
	$('#btnCalc').click(function() {
		var wait_mess = "Waiting for reply...";
		var sum = $('#sum');
		var etime = $('#etime');
		var n = $('#n').val();
		
		sum.html(wait_mess);		
		etime.html(wait_mess);

		$.ajax({
			url: '/sum',
			method: 'post',
			data: 
				'{"n": '+n+', "list": false}',
			dataType: 'json',
			contentType: 'application/json',
			success: function(data) {
						sum.html(data.sum);
						etime.html(data.etime);
			},
		});
	});
};