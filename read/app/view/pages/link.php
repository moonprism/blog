<?php $_temp_header(); ?>
<div id="content" class="link-block m-l">
    <div class="markdown">
        <textarea class="md" style="display: none;" id="mb">
<?php echo trim($article->content); ?>

        </textarea>
    </div>
    <div id="comment" data-id="<?php echo $article->id ?>"></div>
</div>

<!-- Take Your Heart -->
<div class="p5-st" id="p5-st">
	<div style="text-align:right"></div>
    <div class="p5-form">
        <div class="p5-title"><canvas id="p5-title" width="200" height="60"></div>
        <div class="p5-close"><canvas id="p5-close" width="64" height="62"></canvas></div>
		<div class="p5-f-i">
			<label for="p5-f-name" style="margin-left:5px">NAM<span>E</span></label>
			<input id="p5-f-name" name="p5-f-name" autocomplete="off" placeholder="kicoe's blog" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-email">EMAI<span>L</span></label>
			<input id="p5-f-email" name="p5-f-email" autocomplete="off" placeholder="ankicoe@gmail.com" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-url" style="margin-left:15px">UR<span>L</span></label>
			<input id="p5-f-url" name="p5-f-url" autocomplete="off" placeholder="https://kicoe.com" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-intro">INT<span>RO</span></label>
			<input id="p5-f-intro" name="p5-f-intro" autocomplete="off" placeholder="静默空白." type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-icon" style="margin-left:3px">AVA<span>TAR</span></label>
			<input id="p5-f-icon" name="p5-f-icon" autocomplete="off" placeholder="https://cravatar.cn/avatar/7c6d3737a25a9ec47b5439ec123bd1df" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-emoji">EMO<span>JI</span></label>
			<input id="p5-f-emoji" name="p5-f-emoji" autocomplete="off" placeholder="🍥" type="text">
		</div>
        <div class="m-l"><div class="markdown" id="p5-link"><blockquote><p><a target="_blank" rel="noopener" href="https://kicoe.com"><img src="https://q1.qlogo.cn/g?b=qq&nk=1370808234&s=640">kicoe's Blog<code>🍥<i>静默空白.</i></code></a></p></blockquote></div></div>
        <div style="height:300px"><canvas id="p5-talk" width="814" height="250"></canvas><div contenteditable="true" id="p5-talk-msg"></div></div>
        <div class="p5-post"><canvas id="p5-post" width="250" height="80"></div>
        <div id="p5-post-message" class="markdown"></div>
	</div>
</div>
    <div class="p5-load" id="p5-load"></div>

<script>
// 简单写下 随机排序的处理
const url = window.location.href
var seed
if (url.indexOf("?") != -1) {
    seed = parseInt(url.split('?')[1])
} else {
    seed = Math.round(Math.random() * 8999999) + 1000000
}
var linkText = document.getElementById('mb').innerHTML
linkText = linkText.replaceAll(/\n(@@@[\s\S]*?\n@@@\+?)/g, (match, capture) => {
    let ll = match.split('\n\n')
    ll.shift();ll.pop()
    ll.map((v, i) => {
        let r = seed / v.split('').map(v => v.charCodeAt()).reduce((c, sum) => sum + c)
        r = Math.floor(r) % ll.length;
        [ll[i], ll[r]] = [ll[r], ll[i]]
    })
    let text = ll.join("\n\n")
    if (match.charAt(match.length-1) === '+') {
        return text + '\n\n<div class="seed"> 🌱 random seed: <span>'+seed+'<span></div>'
    }
    return text
})
document.getElementById('mb').innerHTML = linkText
window.history.pushState({}, 0, "?"+seed)
</script>

<?php $_temp_footer(); ?>
