//get news list
$(document).ready(function () {
    $.ajax({
        url:"/api/news",
        data:JSON.stringify({"page":"1","limit":"6"}),
        // data:{"page":1,"limit":6},
        contentType: "application/json;charset=utf-8",
        type:"post",
        success:function (result) {
            $.each(result.data,function (i, j) {
                console.log(j)
                $("#newslist_left").append('                        <div data-v-6ef8682e="">\n' +
                    '                            <div class="home-news-item" data-v-6ef8682e="">\n' +
                    '                                <a target="_blank" href="/articles/"'+j.newid+' class="home-news-item__cover">\n' +
                    '                                    <div class="wscn-lazyload lazy">\n' +
                    '                                        <img style="" src="'+j.show_url+'" class="">\n' +
                    '                                    </div>\n' +
                    '                                </a>\n' +
                    '                                <div class="home-news-item__main">\n' +
                    '                                    <a target="_blank" href="/articles/'+j.newid+'" class="home-news-item__main__title">'+j.title+'</a>\n' +
                    '                                    <a target="_blank" href="/articles/'+j.newid+'" class="home-news-item__main__summary">'+j.abstract+'</a>\n' +
                    '                                    <div class="home-news-item__main__meta">\n' +
                    '                                        <div class="home-news-item__main__meta__left">\n' +
                    '                                            <div class="left-item">\n' +
                    '                                                <a target="_blank" href="/authors/834515" class="display-name">新华社</a>\n' +
                    '                                            </div> <!----> <!---->\n' +
                    '                                            <div class="left-item">\n' +
                    '                                                <svg width="1em" height="1em" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" class="icon">\n' +
                    '                                                    <path d="M512 0C229.232 0 0 229.232 0 512s229.232 512 512 512 512-229.232 512-512S794.768 0 512 0zm0 960C264.976 960 64 759.024 64 512S264.976 64 512 64s448 200.976 448 448-200.976 448-448 448zm246.72-346.496L531.888 510.4V225.872c0-17.68-14.336-32-32-32s-32 14.32-32 32v305.12a31.944 31.944 0 0 0 18.768 29.12l245.6 111.632a31.577 31.577 0 0 0 13.216 2.88c12.16 0 23.776-6.976 29.136-18.752 7.312-16.096.208-35.056-15.888-42.368z"></path></svg>\n' +
                    '                                                <span class="text">30分钟前</span>\n' +
                    '                                            </div>\n' +
                    '                                        </div>\n' +
                    '                                        <div class="home-news-item__main__meta__right"><!----> <!---->\n' +
                    '                                        </div>\n' +
                    '                                    </div>\n' +
                    '                                </div>\n' +
                    '                            </div>\n' +
                    '                        </div>')
            })
        }

    })
});
