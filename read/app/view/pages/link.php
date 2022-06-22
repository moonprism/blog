<?php $_temp_header(); ?>
<div id="content" class="link-block m-l">
    <div class="markdown">
        <textarea class="md" style="display: none;">
<?php echo trim($article->content); ?>

        </textarea>
    </div>
<div class="p5-st" id="p5-st">
	<div style="text-align:right"></div>
    <div class="p5-form">
        <div class="p5-close"><canvas id="p5-close" width="64" height="62"></canvas></div>
		<div class="p5-f-i">
			<label for="p5-f-name" style="margin-left:5px">NAM<span>E</span></label>
			<input id="p5-f-name" name="p5-f-name" autocomplete="off" placeholder="kicoe's blog" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-intro">INT<span>RO</span></label>
			<input id="p5-f-intro" name="p5-f-intro" autocomplete="off" placeholder="joker." type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-email">EMAI<span>L</span></label>
			<input id="p5-f-email" name="p5-f-email" autocomplete="off" placeholder="ankicoe@gmail.com" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-icon" style="margin-left:3px">IC<span>ON</span></label>
			<input id="p5-f-icon" name="p5-f-icon" autocomplete="off" placeholder="https://gravatar/avatar/*.jpg" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-url" style="margin-left:10px">UR<span>L</span></label>
			<input id="p5-f-url" name="p5-f-url" autocomplete="off" placeholder="https://kicoe.com" type="text">
		</div>
		<div class="p5-f-i">
			<label for="p5-f-emoji">EMO<span>JI</span></label>
			<input id="p5-f-emoji" name="p5-f-emoji" autocomplete="off" placeholder="⭐️" type="text">
		</div>
        <div class="markdown" id="p5-link"><blockquote><p><a target="_blank" rel="noopener" href="https://kicoe.com"><img src="https://gravatar.cat.net/avatar/7c6d3737a25a9ec47b5439ec123bd1df">kicoe's Blog<code>⭐️<i>joker.</i></code></a></p></blockquote></div>
        <div style="height:300px"><canvas id="p5-talk" width="860" height="250"></canvas><div contenteditable="true" id="p5-talk-msg"></div></div>
        <div class="p5-post"><canvas id="p5-post" width="250" height="80"></div>
        <div id="p5-post-message" class="markdown"></div>
	</div>
</div>
    <div class="p5-load" id="p5-load"></div>
    <div id="comment" data-id="<?php echo $article->id ?>"></div>
</div>
<script src="//at.alicdn.com/t/font_2434616_fos6fvq6jgf.js"></script>
<script src="//at.alicdn.com/t/font_1922445_kmgc1s3rvap.js"></script>
<?php $_temp_footer(); ?>
