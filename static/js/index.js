$.ajax({
    type: 'POST',
    url: "https://www.fembed.com/api/source/7y9w-yjjxvj",
    dataType: "jsonp",
    success: function (result) {
        console.log(result)
    }
});

    var $this = $(this);
    // $('#download-hidden').hide();
    // $this.addClass('is-loading');
    // $('#download-loading').removeClass('is-hidden');
    // timeOut = setTimeout(function(){countdown($this);}, 1000);
    $.post('https://www.fembed.com/api/source/7y9w-yjjxvj', function(res){
        if(res.success) {
            data = res.data;
            if(res.player.revenue) $.getScript(res.player.revenue);
            else revenue = false;
        }
    });

