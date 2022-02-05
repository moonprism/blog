<?php

namespace app\controller;

use app\model\response\ViewResponse;
use kicoe\core\Request;
use kicoe\core\Config;
use app\model\Code;
use Protodata\CodeClient;
use Protodata\CodeDetail;
use Protodata\SearchRequest;

class CodeController
{
    /**
     * @route get /code/search
     * @param Config $config
     * @param string $text
     * @return array
     */
    public function search(Request $request, Config $config)
    {
        $text = $request->query('text');
        $client = new CodeClient($config->get('grpc.host').':'.$config->get('grpc.port'), [
            'credentials' => \Grpc\ChannelCredentials::createInsecure(),
        ]);
        $request = new SearchRequest();
        $request->setText($text);
        list($response, $status) = $client->Search($request)->wait();
        $res = [];
        if ($status->code === 0) {
            /** @var CodeDetail $data */
            foreach ($response->getData() as $data) {
                $res[] = [
                    'id' => $data->getId(),
                    'lang' => $data->getLang(),
                    'content' => $data->getContent(),
                    'description' => $data->getTitle(),
                    'tags' => $data->getTags(),
                ];
            }
        }
        return $res;
    }

    /**
     * @route get /page/code
     * @route get /page/code/{page}
     * @param ViewResponse $response
     * @param int $page
     * @return ViewResponse
     */
    public function index(ViewResponse $response, int $page = 1)
    {
        $limit = 10;
        $code_list = Code::where('deleted_at is null')
            ->orderBy('updated_time', 'desc')
            ->limit(($page-1)*$limit, $limit)
            ->get();
        $next_page = count($code_list) == 10 ? $page+1 : 0;
        return $response->view('pages/code', compact('code_list', 'next_page'));
    }
}
