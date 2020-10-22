<?php

return [
    '/'                   => 'article/index',
    'article/page'        => 'Article@fetchList',
    'article/tag'         => 'Article@tag',
    'article/id'          => 'Article@detail',
    // post
    'comment/up'          => 'CommentController@up',
    // json api
    'comment/list'        => 'CommentController@ls',
    // fake
    'json/comment'        => 'CommentController@cat',

    // page
    'page/link'           => 'Article@linkDetail',
    'page/about'          => 'Article@aboutDetail',
    'page/code'           => 'Code@index',

    'code/preview'        => 'Code@preview',

    // search
    'code/search'         => 'Code@search',
    // api
    'api/articles'         => 'Api@articleList',
    'api/article'         => 'Api@article'
];
