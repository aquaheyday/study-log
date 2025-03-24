$(document).ready(function() {
    function eventList() {
        $(".firstBtn").on("click", function() {
            $.ajax({
                type: "post",
                url: "/main/first",
                data: {
                    "branch": $(this).data("branch"),
                    "country": $(this).data("country"),
                    "no": $(this).data("no")
                },
                success: function(result) {
                    if (result.success) {
                        location.reload();
                    } else {
                        alert(`${result.admin_no} 상담 창구에서 이미 호출하였습니다.`);
                    }
                }
            });
        });

        $(".secondBtn").on("click", function() {
            $.ajax({
                type: "post",
                url: "/main/second",
                data: {
                    "branch": $(this).data("branch"),
                    "country": $(this).data("country"),
                    "no": $(this).data("no")
                },
                success: function(result) {
                    location.reload();
                }
            });
        });

        $(".endBtn").on("click", function() {
            $.ajax({
                type: "post",
                url: "/main/end",
                data: {
                    "branch": $(this).data("branch"),
                    "country": $(this).data("country"),
                    "no": $(this).data("no")
                },
                success: function(result) {
                    location.reload();
                }
            });
        });
    }

    eventList();

    setInterval(function () {
        $.ajax({
            type: "get",
            url: "/main/list",
            success: function(result) {
                if (result.success) {
                    let html = '';

                    $.each(result.data, function(k, v) {
                        html += `<div><p>${k}</p>`;
                        $.each(v, function(i, e) {
                            let html2 = `<div><span>${e.NO}</span><span> ${e.STATUS === '01' ? '호출가능' : (e.STATUS === '02' ? '호출 중' : '상담 완료')} </span>`;
                            let button = '';
                            if (e.STATUS === '01') {
                                button = `<button type="button" class="firstBtn" ${ (result.count > 0) ? 'disabled' : '' } data-branch="${e.BRANCH}" data-country="${e.COUNTRY}" data-no="${e.NO}">호출 하기</button>`;
                            } else if (e.STATUS === '02') {
                                if (e.ADMIN_ID === '{{ Auth::user()->ID }}') {
                                    button = `<button type="button" class="secondBtn" data-branch="${e.BRANCH}" data-country="${e.COUNTRY}" data-no="${e.NO}">재 호출</button>`
                                        + `<button type="button" class="endBtn" data-branch="${e.BRANCH}" data-country="${e.COUNTRY}" data-no="${e.NO}">상담 완료</button>`;
                                } else {
                                    button = `<span>${e.ADMIN_NO} 상담</span><span>대기중</span>`;
                                }
                            } else {
                                button = `<span>${e.ADMIN_NO} 상담</span><span>${e.END_TIME}</span>`
                            }

                            html += html2 + button + '</div>';
                        });
                        html += '</div>';
                    });

                    $("#target").html(html);
                }
                eventList();
            }
        });
    }, 5000);
});
