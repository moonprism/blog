<?php

return [
    '/'                   => 'article/index',
    'article/page'        => 'Article@fetchList',
    'article/id'          => 'Article@detail',
    // post
    'comment/up'          => 'Comment@up',
    // json api
    'comment/list'        => 'Comment@ls',
    // fake
    'json/comment'        => 'Comment@cat',

    // page
    'page/link'           => 'Article@linkDetail',
    'page/about'          => 'Article@aboutDetail',
    'page/code'           => 'Code@index',

    // search
    'code/search'         => 'Code@search'
];