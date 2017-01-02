window.onload = initPage;

function initPage() {
	$('#btnCalc').click(function() {
		var wait_mess = "Waiting for reply...";
		var norepl_mess = "No meaningful reply"
		var sum = $('#sum');
		var etime = $('#etime');
		var primes = $('#primes');
		var mess = $('#mess');
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
						mess.html("OK");
						mess.css("color", "MediumSeaGreen");
			},
			error: function(jqXHR, exception) {
        				var msg = '';
        				if (jqXHR.status === 0) {
            				msg = 'Not connect.\n Verify Network.';
        				} else if (jqXHR.status == 404) {
            				msg = 'Requested page not found. [404]';
        				} else if (jqXHR.status == 500) {
            				msg = 'Internal Server Error [500].';
        				} else if (exception === 'parsererror') {
            				msg = 'Requested JSON parse failed.';
        				} else if (exception === 'timeout') {
            				msg = 'Time out error.';
        				} else if (exception === 'abort') {
            				msg = 'Ajax request aborted.';
        				} else {
            				msg = 'Uncaught Error.\n' + jqXHR.responseText;
        				}
        				mess.html(msg);
						mess.css("color", "Tomato");
						sum.html(norepl_mess);
						etime.html(norepl_mess);
    		},
		});
	});
};