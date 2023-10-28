<?php $_temp_header(); ?>
<div id="curtain"></div>
<div id="content" class="article-block">
    <div class="article">
        <h1><a><?php echo $article->title ?></a></h1>
        <div class="mark">
<?php $math_flag = false ?>
            <svg class="icon i-date" aria-hidden="true">
                <use xlink:href="#icon-date"></use>
            </svg>
            <?php echo date('Y.m.d D', strtotime($article->created_time)) ?>

<?php if($article->getTags()){ ?>
            <svg class="icon i-tag" aria-hidden="true">
                <use xlink:href="#icon-tag"></use>
            </svg>
<?php } ?>
<?php foreach($article->getTags() as $tag){ ?>
           <span style="border-color:<?php echo $tag->color ?? 'none'; ?>;color:<?php echo $tag->color ?? 'none'; ?>">
<?php if ($math_flag != true) $math_flag = (strtolower($tag->name) === 'math'); ?>
               <a href="/?tag_id=<?php echo $tag->id ?>" style="background-color: <?php echo $tag->color ?>" class="tag"><?php echo $tag->name ?></a>
           </span>
<?php } ?>
        </div>
    </div>
    <div class="markdown">
        <textarea class="md" style="display: none;">
<?php echo trim($article->content); ?>

        </textarea>
    </div>
    <div class="agreement" id="agreement">powered by <a target="_blank" href="https://github.com/moonprism/kicoephp-src">kicoephp</a><br></div>
    <div id="comment" data-id="<?php echo $article->id ?>"></div>
</div>
<?php if ($math_flag) { ?>
    <script type="text/javascript" src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS-MML_HTMLorMML">
</script>
<?php } ?>
<?php $_temp_footer(); ?>
<link rel="stylesheet" type="text/css" href="https://cdn.staticfile.org/highlight.js/11.9.0/styles/tokyo-night-dark.min.css">
<script src="https://cdn.staticfile.org/highlight.js/11.9.0/highlight.min.js"></script>
<script>hljs.configure({
    ignoreUnescapedHTML: true
});hljs.highlightAll()</script>
