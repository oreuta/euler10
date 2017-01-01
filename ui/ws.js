window.onload = initPage;

function initPage() {
	$('#btnCalc').click(function() {
		var wait_mess = "Waiting for reply...";
		var sum = $('#sum');
		var etime = $('#etime');
		var primes = $('#primes');
		var list = $('#list').prop('checked');
		var n = $('#n').val();
		
		sum.html(wait_mess);		
		etime.html(wait_mess);
		if (list) {
			primes.html(wait_mess);
		} else {
			primes.html("");
		}

		$.ajax({
			url: '/sum',
			method: 'post',
			data: 
				'{"n": '+n+', "list": '+list+'}',
			dataType: 'json',
			contentType: 'application/json',
			success: function(data) {
						if (data.error) {
							sum.html(data.error);
							etime.html("");
							primes.html("");
							return;							
						}
						sum.html(data.sum);
						etime.html(data.etime);
						if (data.primes) {
							var str = '';
							for(var p in data.primes) {
    							str += data.primes[p] + '\t';
							}
							primes.html(str.slice(0, -1));
						} else {
							primes.html("");
						}
			},
		});
	});
};