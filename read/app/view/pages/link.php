<?php $_temp_header(); ?>
<div id="content" class="link-block m-l">
    <div class="markdown">
        <textarea class="md" style="display: none;">
<?php echo trim($article->content); ?>

        </textarea>
    </div>
    <div id="comment" data-id="<?php echo $article->id ?>"></div>
</div>
<script src="//at.alicdn.com/t/font_2434616_fos6fvq6jgf.js"></script>
<script src="//at.alicdn.com/t/font_1922445_kmgc1s3rvap.js"></script>
<?php $_temp_footer(); ?>
