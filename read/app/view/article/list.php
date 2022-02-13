<?php $_temp_header(); ?>
<div id="content" class="article-block">
<?php if($tag){ ?>
    <div style="border-color: <?php echo $tag->color ?>"><div class="tag-banner">当前标签<span class="t" style="background-color: <?php echo $tag->color; ?>"><?php echo $tag->name; ?></span>, 文章数<span class="c"><?php echo $count; ?><span></div></div>
<?php } ?>
    <div class="article-list">
<?php foreach($article_list as $article){ ?>
        <div class="article">
            <h1><a href="/article/<?php echo $article->id ?>"><?php echo stripslashes($article->title) ?></a></h1>
                <div class="mark">
                    <svg class="icon i-date" aria-hidden="true">
                        <use xlink:href="#icondate"></use>
                    </svg>
                    <?php echo date('Y.m.d D', strtotime($article->created_time)) ?>

<?php if($article->getTags()){ ?>
                    <svg class="icon i-tag" aria-hidden="true">
                        <use xlink:href="#icontag"></use>
                    </svg>
<?php } ?>
<?php foreach($article->getTags() as $t) { ?>
                    <span style="border-color:<?php echo $t->color; ?>;color:<?php echo $t->color; ?>">
                        <a href="/?tag_id=<?php echo $t->id ?>" style="<?php if ($tag && $t->id == $tag->id) echo 'background-color: inherit;color: inherit;border-color: inherit'; else echo 'background-color: '.$t->color; ?>" class="tag"><?php echo $t->name ?></a>
                    </span>
<?php } ?>
                </div>
                <div class="summary markdown">
                    <script type='text/html' style="display:none" class="md"><?php echo $article->summary?></script>
                </div>
            </div>
<?php } ?>
    </div>
<?php include(dirname(__DIR__).'/article/page.php') ?>
</div>
<script src="//at.alicdn.com/t/font_1747613_gq90x707awu.js"></script>
<?php $_temp_footer(); ?>
<link rel="stylesheet" type="text/css" href="https://cdn.staticfile.org/highlight.js/10.0.0/styles/monokai.min.css">
