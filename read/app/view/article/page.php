<?php
$total_page = intval(ceil($count/number_format($limit, 1)));
?>
    <div class="page">
        <ul>
            <?php
            if ($total_page == 1) {
                // todo
            } else if ($page-1 < $total_page-$page) {
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="'.sprintf($url, $i).'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
                }
                if($page+3 < $total_page){
                    echo '<li><span>...</span></li>';
                    echo '<li><a href="'.sprintf($url, $total_page).'">'.$total_page.'</a></li>';
                }
            } else {
                if($page-3 > 1){
                    echo '<li><a href="'.sprintf($url, 1).'">1</a></li>';
                    echo '<li><span>...</span></li>';
                }
                $start_page = max($page-3, 1);
                $end_page = min($page+3, $total_page);
                for ($i = $start_page; $i <= $end_page; $i++) {
                    if ($i != $page) {
                        echo '<li><a href="'.sprintf($url, $i).'">'.$i.'</a></li>';
                    } else {
                        echo '<li><span>'.$i.'</span></li>';
                    }
            }
            }
            ?>

        </ul>
    </div>
