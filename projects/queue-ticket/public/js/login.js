$(document).ready(function() {
    document.querySelector("#loginBtn").addEventListener("click", function() {
        $.ajax({
            type: "post",
            url: "/login",
            data: {
                "id": $("input[name='id']").val(),
                "password": $("input[name='pw']").val()
            },
            success: function(result) {
                if (result.success) {
                    if (result.account_type === '01') {
                        location.replace('/main');
                    } else if (result.account_type === '02') {
                        location.replace('/kiosk');
                    } else if (result.account_type === '03') {
                        location.replace('/call');
                    }
                } else {
                    alert("아이디와 비밀번호가 올바르지 않습니다.\n다시 확인하고 입력해 주세요.");
                }
            }
        })
    });
});
