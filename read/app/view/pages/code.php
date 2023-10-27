<?php $_temp_header(); ?>
<div id="content" class="m-c">
    <div <?php if($next_page == -1){ echo 'style="display:none"'; } ?> class="search-input">
        <input id="searchInput" autocomplete="off" autofocus>
    </div>
    <div id="load"></div>
    <div id="searchResult" class="search-result">
<?php foreach($code_list as $code) { ?>
        <div class="search-item">
            <blockquote>
                <svg class="icon icon-code" aria-hidden="true">
                    <use xlink:href="#icon-<?php echo $code->lang ?>"></use>
                </svg>
                <?php echo $code->description ?> <span>.<?php echo $code->lang ?></span>
            </blockquote>
                <?php if($code->tags) { ?><div class="tags"><?php echo $code->tags ?></div><?php } ?>

                <pre class="c-<?php echo $code->lang ?>"><code class="<?php echo $code->lang ?>"><?php echo htmlspecialchars($code->content) ?></code><div class="m-cc"></div></pre>
        </div>
<?php } ?>
        <?php if($next_page>0) { ?><div class="next"><a href="/code?page=<?php echo $next_page;?>">下一页</a></div><?php } ?>

    </div>
</div>

<?php $_temp_footer(); ?>

<link rel="stylesheet" type="text/css" href="https://cdn.staticfile.org/highlight.js/11.9.0/styles/tokyo-night-dark.min.css">
<script src="//cdn.staticfile.org/highlight.js/11.9.0/highlight.min.js"></script>
<script>hljs.configure({
    ignoreUnescapedHTML: true
})</script>
