window.onload = initPage;

function initPage() {
	$('#btnCalc').click(function() {
		var color_error = "Tomato";
		var color_OK = "MediumSeaGreen";
		var wait_mess = "Waiting for reply...";
		var norepl_mess = "No meaningful reply"
		var nothing_mess = "Nothing to show"
		var sum = $('#sum');
		var etime = $('#etime');
		var primes = $('#primes');
		var mess = $('#mess');
		var lst = $('#lst').prop('checked');
		var n = $('#n').val();
		var nr = $('#nr').val();
		
		sum.html(wait_mess);		
		etime.html(wait_mess);
		if (lst) {
			primes.html(wait_mess);
		} else {
			primes.html("");
		}

		$.ajax({
			url: '/sum',
			method: 'post',
			data: 
				'{"n": '+n+', "lst": '+lst+', "nr": '+nr+'}',
			dataType: 'json',
			contentType: 'application/json',
			success: function(data) {
						if (data.error) {
							mess.html(data.error);
							mess.css("color", color_error);
							sum.html(norepl_mess);
							etime.html(norepl_mess);
							primes.html(nothing_mess);
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
							primes.html(nothing_mess);
						}
						mess.html("OK");
						mess.css("color", color_OK);
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
						mess.css("color", color_error);
						sum.html(norepl_mess);
						etime.html(norepl_mess);
						primes.html(nothing_mess);
    		},
		});
	});
};