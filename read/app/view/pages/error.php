<?php $_temp_header(); ?>
<div id="content" class="article-block">
    <div class="markdown">
        <textarea class="md" style="display: none;">
        <?php if ($e->getCode() == '404') { ?>
# <?php echo $e->getCode(); ?>

`<?php echo $e->getMessage(); ?>`

<?php dump($e); ?>
        <?php } else { ?>
        <?php } ?>
        </textarea>
    </div>
</div>
<?php $_temp_footer(); ?>