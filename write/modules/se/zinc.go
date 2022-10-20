package se

// import (
// 	_ "embed"
// 	"encoding/json"
// 	"strconv"
//
// 	"git.kicoe.com/blog/write/modules/setting"
// 	"github.com/davecgh/go-spew/spew"
// 	"github.com/go-resty/resty/v2"
// )
//
// var cli *resty.Client
// var url string
// var req *resty.Request
//
// //go:embed code-index.json
// var codeIndex string
//
// func Init() {
// 	cli = resty.New()
// 	url = "http://"+setting.SE.Host+":"+setting.SE.Port
// 	req = cli.R().
//       SetHeader("Content-Type", "application/json").
//       SetHeader("Authorization", "Basic YWRtaW46YWRtaW4xMjM=")
//
//     req.SetBody(codeIndex).Post(url+"/api/index")
// }
//
// func Insert(index string, id int64, data interface{}) error {
// 	str, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = req.SetBody(str).Put(url+"/api/"+index+"/_doc/"+strconv.FormatInt(id, 10))
// 	return err
// }
//
// func Update(index string, id int64, data interface{}) error {
// 	str, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = req.SetBody(str).Post(url+"/api/"+index+"/_update/"+strconv.FormatInt(id, 10))
// 	return err
// }
//
// func Delete(index string, id int64) error {
// 	_, err := req.Delete(url+"/api/"+index+"/_doc/"+strconv.FormatInt(id, 10))
// 	return err
// }
//
// func Search(index string, searchJson string) (re []map[string]interface{}, total int64, err error) {
// 	result, err := req.SetBody(searchJson).Post(url+"/api/"+index+"/_search")
// 	if err != nil {
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(result.String()), &m)
// 	if err != nil {
// 		return
// 	}
// 	spew.Dump(m)
//
// 	hits := m["hits"].(map[string]interface{})
// 	total = int64(hits["total"].(map[string]interface{})["value"].(float64))
// 	items := hits["hits"].([]interface{})
//
// 	re = make([]map[string]interface{}, 0)
// 	for _, v := range(items)  {
// 		re = append(re, v.(map[string]interface{}))
// 	}
// 	return
// }
