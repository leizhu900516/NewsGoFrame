//get news list
$(document).ready(function () {
    $.ajax({
        url:"/api/news",
        data:JSON.stringify({"page":"1","limit":"6"}),
        contentType: "application/json;charset=utf-8",
        type:"post",
        success:function (result) {
            $.each(result.data,function (i, j) {
                console.log(j)
                $("#newslist_left").append('<div class="layui-row grid-demo layui-col-space10 site-demo-flow">\n' +
                    '                            <div class="layui-col-md3">\n' +
                    '                                <img src="https://gw.alicdn.com/bao/uploaded/i2/162734861/TB2V5rsX_gc61BjSZFkXXcTkFXa_!!162734861.jpg_400x400q90.jpg?t=1518395023940">\n' +
                    '                            </div>\n' +
                    '                            <div class="layui-col-md9"  id="news_item_all_style">\n' +
                    '                                <div class="new_home_list">\n' +
                    '                                    <a href="/news/'+j.newid+'" class="news_title">'+j.title+'</a>\n' +
                    '                                    <br>\n' +
                    '                                    <a href="/news/'+j.newid+'" class="news_abstract">'+j.abstract+'</a>\n' +
                    '                                    <div class="bottom_refer_and_icon">\n' +
                    '                                        <div class="left">\n' +
                    '                                            <i class="layui-icon">&#xe6c6;</i><span>赞</span>\n' +
                    '                                        </div>\n' +
                    '                                        <div class="right">\n' +
                    '                                            <i class="layui-icon"> &#xe60e;</i><span>30分钟前</span>\n' +
                    '                                        </div>\n' +
                    '                                    </div>\n' +
                    '                                </div>\n' +
                    '                            </div>\n' +
                    '                        </div>')
            })
        }

    })
});
