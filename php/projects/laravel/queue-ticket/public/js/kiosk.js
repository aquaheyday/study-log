$(document).ready(function() {
    $(".callBtn").on('click', function() {
        $.ajax({
            type: "post",
            url: "/kiosk/call",
            data: {
                "country": $(this).data("country")
            },
            success: function(result) {
                if (result.success) {
                    const message = `${result.text} \n ${result.sub_text}: ${result.number}`;
                    showCustomAlert(message);
                }
            }
        });
    });

    setInterval(function () {
        $.ajax({
            type: "get",
            url: "/kiosk/waiting",
            success: function(result) {
                if (result.success) {
                    $.each(result.data, function (i, v) {
                        $(`#${i}_waiting_number`).html(v);
                    });
                }
            }
        });
    }, 10000);

    function showCustomAlert(message) {
        $("#alertMessage").text(message);
        $("#customAlert").removeClass("hidden");
    }

    $("#alertOkBtn").on("click", function() {
        $("#customAlert").addClass("hidden");
    });

});
