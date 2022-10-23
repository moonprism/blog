function drawLine(ctx, dots, color = "#fff") {
    ctx.beginPath()
    ctx.strokeStyle = color
    ctx.lineWidth = 1
    let start = dots.shift()
    ctx.moveTo(start[0], start[1])
    for (let i=0; i<dots.length; i++) {
        ctx.lineTo(dots[i][0], dots[i][1])
    }
    dots.unshift(start)
    ctx.lineTo(start[0], start[1])
    ctx.stroke()
    ctx.shadowOffsetX = 0
    ctx.shadowOffsetY = 0
    ctx.shadowBlur = 0
}

function draw2d(ctx, dots, color = "#fff", after) {
    ctx.beginPath()
    ctx.fillStyle = color
    let start = dots.shift()
    ctx.moveTo(start[0], start[1])
    for (let i=0; i<dots.length; i++) {
        ctx.lineTo(dots[i][0], dots[i][1])
    }
    dots.unshift(start)
    if (after) {
        after()
    }
    ctx.fill()
    ctx.shadowOffsetX = 0
    ctx.shadowOffsetY = 0
    ctx.shadowBlur = 0
}

var butCanvas = document.getElementById('p5');
if (butCanvas) {
    // button
    var butCtx = butCanvas.getContext('2d');
    function drawBut(type) {
        draw2d(butCtx, [[0,0], [0, 1000], [1000, 1000], [1000, 0]])

        draw2d(butCtx, [
            [6, 14],
            [13, 54],
            [253, 51],
            [257, 9],
        ], '#000')

        drawButContent()
    }

    function drawButContent(shadow = false) {

        draw2d(butCtx, [
            [10, 18],
            [16, 51],
            [250, 47],
            [253, 12],
        ], '#f00')

        butCtx.font = "bold 18px 'Roboto Mono',Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,'Noto Sans SC',Microsoft YaHei,SimSun,sans-serif";
        butCtx.fillStyle="#fff";

        if (shadow) {
            butCtx.shadowOffsetX = 1
            butCtx.shadowOffsetY = 1
            butCtx.shadowBlur = 2
            butCtx.shadowColor = "#000"
            butCtx.strokeStyle = '#000'
            butCtx.strokeText("申 请 友 链", 78, 40)
        }

        butCtx.fillText("申 请 友 链", 78, 40)

        butCtx.shadowOffsetX = 0
        butCtx.shadowOffsetY = 0
        butCtx.shadowBlur = 0
    }

    drawBut()

    butCanvas.onmouseover = (e) => {
        draw2d(butCtx, [
            [3, 13],
            [10, 56],
            [255, 52],
            [262, 7],
        ], '#000')

        draw2d(butCtx, [
            [10, 16],
            [15, 60],
            [261, 55],
            [267, 12],
        ], '#000')

        drawButContent(true);
    }

    butCanvas.onmouseout = (e) => {
        drawBut()
    }

    butCanvas.onclick = (e) => {
        window.scrollTo(0, 0)
        $("p5-st").style.display = "block"
        inputs[0].focus()
    }

    // todo 前端的双向绑定实现，有时间研究下
    let inputs = $('p5-st').getElementsByTagName('input');
    let preview = () => {
        let name = $('p5-f-name').value;
        let intro = $('p5-f-intro').value
        let email = $('p5-f-email').value
        let icon = $('p5-f-icon').value
        let url = $('p5-f-url').value
        let emoji = $('p5-f-emoji').value
        if (icon) {
            $('p5-link').innerHTML = '<blockquote><p><a target="_blank" rel="noopener" href="'+url+'"><img onerror="" src="'+icon+'">'+name+'<code>'+emoji+'<i>'+intro+'</i></code></a></p></blockquote>'
        } else {
            $('p5-link').innerHTML = '<blockquote><p><a target="_blank" rel="noopener" href="'+url+'"><img onerror="" src="https://gravatar.cat.net/avatar/123">'+name+'<code>'+emoji+'<i>'+intro+'</i></code></a></p></blockquote>'
        }
    }
    [].slice.call(inputs).forEach(i => {
        if (i.id == 'p5-f-icon') {
            i.onchange = preview
        } else {
            i.oninput = preview
        }
        i.onfocus = () => {
            i.parentNode.style.backgroundColor = '#000'
            i.parentNode.style.color = '#fff'
        }
        i.onblur = () => {
            i.parentNode.style.backgroundColor = '#fff'
            i.parentNode.style.color = '#000'
        }
    })

    // close
    var closeCanvas = document.getElementById('p5-close');
    var closeCtx = closeCanvas.getContext('2d');

    function drawClose(hover = false) {
        draw2d(closeCtx, [[0,0], [0, 1000], [1000, 1000], [1000, 0]], '#fff')
        draw2d(closeCtx, [
            [2, 31],
            [42, 58],
            [56, 49],
            [11, 13],
        ], '#f00')
        draw2d(closeCtx, [
            [8, 56],
            [19, 61],
            [63, 19],
            [52, 2],
        ], '#f00')
        let cc = '#fff'
        if (hover == true) {
            cc = '#000'
        }
        draw2d(closeCtx, [
            [6, 31],
            [42, 54],
            [50, 49],
            [12, 20],
        ], cc)
        draw2d(closeCtx, [
            [13, 55],
            [19, 58],
            [59, 19],
            [51, 10],
        ], cc)
    }

    drawClose(true);

    $('p5-close').onmouseover = (e) => {
        // drawClose(true)
    }
    $('p5-close').onmouseout = (e) => {
        // drawClose()
    }
    
    $('p5-close').onclick = (e) => {
        $("p5-st").style.display = "none"
    }
    
    var talkCanvas = document.getElementById('p5-talk');
    var talkCtx = talkCanvas.getContext('2d');
    draw2d(talkCtx, [[0,0], [0, 1000], [1000, 1000], [1000, 0]], '#000')

    draw2d(talkCtx, [
        [6, 17],
        [14, 44],

        [36, 44],
        [39, 47],
        [29, 48],
        [72, 71],
        [53, 49],
        [63, 48],
        [61, 44],

        [110, 44],
        [120, 6],
    ], '#f00');

    function talkTitle(world, x, y, style = 1) {
        talkCtx.rotate(-4 * Math.PI/180)

        if (style == 1) {
            talkCtx.fillStyle="#fff";
            talkCtx.shadowColor = "#000"
            talkCtx.strokeStyle = '#000'
        } else {
            draw2d(talkCtx, [
                [x-1, y-19],
                [x-1, y+5],
                [x+21, y+5],
                [x+21, y-19],
            ], '#fff');
            talkCtx.fillStyle="#000";
            talkCtx.shadowColor = "#fff"
            talkCtx.strokeStyle = '#fff'
        }

        talkCtx.font = "bold 18px 'Roboto Mono',Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,'Noto Sans SC',Microsoft YaHei,SimSun,sans-serif";
        talkCtx.shadowOffsetX = 1
        talkCtx.shadowOffsetY = 1
        talkCtx.shadowBlur = 2

        talkCtx.strokeText(world, x, y)
        talkCtx.fillText(world, x, y)

        talkCtx.rotate(4 * Math.PI/180)
        talkCtx.shadowOffsetX = 0
        talkCtx.shadowOffsetY = 0
        talkCtx.shadowBlur = 0
    }

    talkTitle('悄', 25, 38);
    talkTitle('悄', 48, 38);
    talkTitle('话', 72, 38, 2);

    // draw2d(talkCtx, [
    //     [30, 115],
    //     [30, 126],
    //     [100, 201],
    //     [115, 181],
    //     [171, 233],
    //     [196, 207],
    // 
    //     [770, 248], // 右下转折点
    // 
    //     [741, 0],
    //     [142, 67],
    //     [134, 124],
    //     [91, 95],
    //     [81, 126],
    //     [81, 126],
    // ], '#000')
    // 
    // draw2d(talkCtx, [
    //     [41, 119],
    //     [101, 185],
    //     [119, 170],
    //     [171, 220],
    //     [191, 201],
    // 
    //     [771, 249], // 右下转折点
    // 
    //     [731, 13],
    //     [145, 74],
    //     [138, 134],
    //     [95, 101],
    //     [89, 136],
    // ], '#fff')
    // 
    // draw2d(talkCtx, [
    //     [67, 136],
    //     [101, 170],
    //     [119, 154],
    //     [168, 203],
    //     [186, 187],
    // 
    //     [764, 248], // 右下转折点
    // 
    //     [729, 24],
    //     [152, 83],
    //     [142, 143],
    //     [100, 118],
    //     [92, 151],
    //     [67, 136],
    // ], '#000')
    
    // 五角星
    function p5Star(x, y, r, color = '#fff') {
        let r1 = r/2
        let r2 = r1/2.5
        let x1,x2,y1,y2
        postCtx.beginPath();
        for (var i = 0; i < 5; i++) {
            x1 = r1 * Math.cos((54 + i*72)/180*Math.PI);
            y1 = r1 * Math.sin((54 + i*72)/180*Math.PI);
            x2 = r2 * Math.cos((18 + i*72)/180*Math.PI);
            y2 = r2 * Math.sin((18 + i*72)/180*Math.PI);
            postCtx.lineTo(x+x2, y+y2);
            postCtx.lineTo(x+x1, y+y1);
        }
        postCtx.fillStyle = color
        postCtx.closePath();
        if (color == '#fff') {
            postCtx.strokeStyle = '#000'
            postCtx.stroke();
        }
        postCtx.fill()
    }

    var postCanvas = document.getElementById('p5-post');
    var postCtx = postCanvas.getContext('2d');
    // draw2d(postCtx, [[0,0], [0, 1000], [1000, 1000], [1000, 0]], '#000')

    draw2d(postCtx, [
        [17, 20],
        [17, 63],
        [222, 70],
        [231, 10],
    ], '#000');
    draw2d(postCtx, [
        [19, 22],
        [19, 60],
        [219, 67],
        [229, 12],
    ], '#f00');
    postCtx.font = "bold 25px 'Roboto Mono',Helvetica Neue,Helvetica,PingFang SC,Hiragino Sans GB,'Noto Sans SC',Microsoft YaHei,SimSun,sans-serif";
    postCtx.fillStyle="#fff";
    postCtx.shadowOffsetX = 2
    postCtx.shadowOffsetY = 2
    postCtx.shadowBlur = 2
    postCtx.shadowColor = "#000"
    postCtx.strokeStyle = '#000'
    postCtx.rotate(-2 * Math.PI/180)
    postCtx.strokeText("发", 80, 54)
    postCtx.fillText("发", 80, 54)
    postCtx.strokeText("布", 110, 54)
    postCtx.fillText("布", 110, 54)
    postCtx.rotate(2 * Math.PI/180)
    postCtx.shadowOffsetX = 0
    postCtx.shadowOffsetY = 0
    postCtx.shadowBlur = 0

    p5Star(156, 41, 24)
    p5Star(156, 41, 12, '#f00')

    postCanvas.onmouseover = (e) => {
        p5Star(156, 41, 24)
        p5Star(156, 41, 13, '#000')
    }

    postCanvas.onmouseout = (e) => {
        p5Star(156, 41, 24)
        p5Star(156, 41, 12, '#f00')
    }

    // 真不该弄这个样式，太麻烦了随便写吧
    var postAct = false
    postCanvas.onclick = (e) => {
        document.body.style.overflowY = 'hidden'
        let name = $('p5-f-name').value;
        let intro = $('p5-f-intro').value
        let email = $('p5-f-email').value
        let icon = $('p5-f-icon').value
        let url = $('p5-f-url').value
        let emoji = $('p5-f-emoji').value

        $('p5-load').style.width = '100%'
        if (postAct) {
            return
        }
        postAct = true
        f.fetch('/api/link', {
            method: "POST",
            data: {
                link: url,
                email: email,
                name: name,
                icon: emoji,
                avatar: icon,
                message: intro,
                data: $('p5-talk-msg').innerText
            },
            success: function (response) {
                setTimeout(() => {
                    $('p5-post-message').innerHTML = '<code>{code:'+response.code+', message:"'+response.message+'"}</code>';
                    $('p5-load').style.left = '100%'
                    $('p5-load').style.width = '0%'
                    if (response.code == 200) {
                        $('p5-st').style.display = 'none'
                    }
                }, 1000)
                document.body.style.overflowY = 'auto'
                setTimeout(() => {
                    $('p5-load').style.left = '0'
                    postAct = false
                }, 2000)
            },
        });
    }
}
