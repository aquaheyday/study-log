$(document).ready(function() {
    setInterval(function () {
        $.ajax({
            type: "get",
            url: "/call/list",
            success: function(result) {
                if (result.success) {
                    $.each(result.data, function (i, v) {
                        let html = '';
                        $.each(v, function (i2, v2) {
                            html += `<div><p>${v2.number}</p><p>${v2.counsel}</p></div>`
                        });

                        $(`.${i}_list`).html(html);
                    });
                }
            }
        });
    }, 5000);
});
