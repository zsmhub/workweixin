package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/ChimeraCoder/gojson"
	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

// 生成企微api代码
type Field struct {
	Name       string
	IsRequired bool
	Desc       string
}

type Api struct {
	IsGet      bool
	DocURL     string
	Name       string
	StructName string
	Method     string
	MethodCaml string
	URL        string
	ReqJson    string
	ReqCode    string
	ReqFields  []Field
	RespJson   string
	RespCode   string
	RespFields []Field
}

var ignoreAccessTokenFields = []string{"access_token", "suite_access_token"}

var docVar = flag.String("doc", "", "[必填]企微文档地址")
var prefixVar = flag.String("prefix", "", "[选填]生成的文件名前缀")

func main() {
	flag.Parse()

	var docURL, savePath, filePrefix string

	if docVar != nil {
		docURL = *docVar
	}
	if docURL == "" {
		fmt.Println("请输入参数doc(企微文档地址):")
		_, _ = fmt.Scanf("%s", &docURL)
	}
	if docURL == "" {
		die("必传参数 doc=?")
	}

	if prefixVar != nil {
		filePrefix = *prefixVar
	}

	// get the fresh documentation!
	var doc *goquery.Document
	resp, err := http.Get(docURL)
	if err != nil {
		die("http get of errcode documentation failed: %+v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		die("non-200 app: %+v\n", resp)
	}

	tmp, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		die("parse document failed: %+v\n", err)
	}

	doc = tmp

	titleHtml, err := doc.Find("title").Html()
	if err != nil {
		die("failed to get html: %+v\n", err)
	}
	titleHtml = titleHtml[:strings.Index(titleHtml, " ")]
	savePath = fmt.Sprintf("./apis/%s.go", titleHtml)
	if filePrefix != "" {
		savePath = fmt.Sprintf("./apis/%s-%s.go", filePrefix, titleHtml)
	}
	fmt.Printf("开始抓取和生成API代码，文档地址:%s，代码保存路径:%s\n", docURL, savePath)

	rawHtml, err := doc.Find(".ep-layout-cnt").Html()
	if err != nil {
		die("failed to get html: %+v\n", err)
	}

	apis := make([]*Api, 0)
	rawHtml = regexp.MustCompile(`<h2 id="([0-9a-zA-Z\-%]+)" data-sign="\w+" data-lines="\w+">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtml = regexp.MustCompile(`<h2 data-sign="\w+" data-lines="\w+" id="([0-9a-zA-Z\-%]+)">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtmlSections := strings.Split(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtmlType := "h2" // 一个页面多个接口
	if len(rawHtmlSections) == 1 {
		rawHtmlType = "h1" // 一个页面一个接口
	}
	for index, rawHtmlSection := range rawHtmlSections {
		fmt.Printf("\n\n开始处理第%d个接口\n", index)
		api := &Api{
			DocURL: docURL,
		}

		if rawHtmlType == "h1" {
			api.Name = titleHtml
		} else {
			rawHtmlSection = `<h2 class="rawHtmlSection">` + rawHtmlSection
			apiNameRegexp := regexp.MustCompile(`<a class="anchor" href=".+"></a>(.+?)</h2>`)
			api.Name = pickSubMatchString(apiNameRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		}

		if api.Name == "" {
			fmt.Println("接口标题为空，跳过处理")
			continue
		}
		fmt.Printf("%s\n", api.Name)

		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式：</strong>`, `<strong>请求方式:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式:</strong> `, `<strong>请求方式:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求示例：</strong>`, `<strong>请求示例：</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>返回结果 ：</strong>`, `<strong>返回结果:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `POST（<strong>HTTPS</strong>）`, `POST(<strong>HTTPS</strong>)`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `GET（<strong>HTTPS</strong>）`, `GET(<strong>HTTPS</strong>)`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `：`, `:`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `: </strong>`, `:</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `:</strong>`, `</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `（<strong>HTTPS</strong>）`, `(<strong>HTTPS</strong>)`) // 兼容中文括号
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `POST(<strong>HTTPS</strong>)`, `POST<strong>HTTPS</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `GET(<strong>HTTPS</strong>)`, `GET<strong>HTTPS</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式</strong>:<strong>POST</strong>(<strong>HTTPS</strong>)`, `<strong>请求方式</strong>POST<strong>HTTPS</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式</strong>:<strong>GET</strong>(<strong>HTTPS</strong>)`, `<strong>请求方式</strong>GET<strong>HTTPS</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式</strong>:POST<strong>HTTPS</strong>`, `<strong>请求方式</strong>POST<strong>HTTPS</strong>`)
		rawHtmlSection = strings.ReplaceAll(rawHtmlSection, `<strong>请求方式</strong>:GET<strong>HTTPS</strong>`, `<strong>请求方式</strong>GET<strong>HTTPS</strong>`)
		// 过滤掉不是接口的节点
		if !strings.Contains(rawHtmlSection, `<strong>请求方式</strong>`) {
			fmt.Println("没有请求方式，跳过处理")
			continue
		}

		buf := bytes.NewBufferString(rawHtmlSection)
		subDoc, err := goquery.NewDocumentFromReader(buf)
		if err != nil {
			fmt.Printf("goquery.NewDocumentFromReader failed: %+v\n", err)
			continue
		}

		apiMethodRegexp := regexp.MustCompile(`<strong>请求方式</strong>(\w+)<strong>HTTPS</strong>`)
		api.Method = pickSubMatchString(apiMethodRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		if strings.ToUpper(api.Method) != "POST" {
			api.IsGet = true
			api.Method = "GET"
		}

		apiURLRegexp := regexp.MustCompile(`<strong>请求地址</strong>(.+?)</p>`)
		api.URL = pickSubMatchString(apiURLRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)

		cutIndex := strings.Index(api.URL, "<br/>")
		if cutIndex > 0 {
			api.URL = api.URL[:cutIndex]
		}

		api.URL = strings.Trim(api.URL, " ")
		api.URL = strings.ReplaceAll(api.URL, "amp;", "")

		if strings.Contains(rawHtmlSection, `<strong>请求示例</strong>`) {
			api.ReqJson = subDoc.Find("pre > code").Eq(0).Text()
		} else if !api.IsGet {
			api.ReqJson = subDoc.Find("pre > code").Eq(0).Text()
		}

		if api.ReqJson == "" && api.IsGet {
			getUrl, _ := url.Parse(api.URL)
			if getUrl != nil {
				getParams := getUrl.Query()
				var jsonParams []string
				for k, v := range getParams {
					if len(v) > 0 && !inSlice(k, ignoreAccessTokenFields) { // 忽略access_token请求字段
						jsonParams = append(jsonParams, fmt.Sprintf(`"%s":""`, k))
					}
				}
				if len(jsonParams) > 0 {
					api.ReqJson = "{" + strings.Join(jsonParams, ",") + "}"
				}
			} else {
				fmt.Println(api.URL)
				continue
			}
		}

		if strings.Contains(rawHtmlSection, `<strong>返回结果</strong>`) {
			if subDoc.Find("pre > code").Eq(1).Text() != "" {
				api.RespJson = subDoc.Find("pre > code").Eq(1).Text()
			} else if subDoc.Find("pre > code").Eq(0).Text() != "" {
				api.RespJson = subDoc.Find("pre > code").Eq(0).Text()
			}
		}

		api.MethodCaml = strcase.ToCamel(strings.ToLower(api.Method))
		api.URL = strings.ReplaceAll(api.URL, "https://qyapi.weixin.qq.com", "")

		// 循环两个table
		tableTmp, err := goquery.NewDocumentFromReader(bytes.NewBufferString(rawHtmlSection))
		if err != nil {
			die("parse document failed: %+v\n", err)
		}
		// 传入参数和响应参数处理
		var isContainDataType bool // 新文档的参数说明多了一列：数据类型
		tableTmp.Find(".cherry-table-container .cherry-table").Find("thead").Find("tr").Find("th").Each(func(i int, s *goquery.Selection) {
			filed, _ := s.Html()
			if filed == "类型" {
				isContainDataType = true
			}
		})
		tableTmp.Find(".cherry-table-container .cherry-table").Find("tbody").Each(func(i int, s *goquery.Selection) {
			s.Find("tr").Each(func(i int, ss *goquery.Selection) {
				var matches []string
				ss.Find("td").Each(func(i int, sss *goquery.Selection) {
					content, _ := sss.Html()
					matches = append(matches, content)
				})

				// 传入参数，每行有3或4个字段，利用字段数量来提取
				var (
					reqColumnNum  = 4 // 说明的字段数量
					reqNameColumn = 0 // 参数列
					// reqTypeColumn     = 1 // 类型列
					reqRequiredColumn = 2 // 是否必须列
					reqDesColumn      = 3 // 说明列
				)
				if !isContainDataType {
					reqColumnNum = 3
					reqRequiredColumn = 1
					reqDesColumn = 2
				}
				if len(matches) == reqColumnNum {
					// 忽略access_token请求字段
					if inSlice(pickSubMatchString(matches, reqColumnNum, reqNameColumn), ignoreAccessTokenFields) {
						return
					}

					var isRequired bool
					if pickSubMatchString(matches, reqColumnNum, reqRequiredColumn) == "是" {
						isRequired = true
					}
					api.ReqFields = append(api.ReqFields, Field{
						Name:       pickSubMatchString(matches, reqColumnNum, reqNameColumn),
						Desc:       pickSubMatchString(matches, reqColumnNum, reqDesColumn),
						IsRequired: isRequired,
					})
				}

				// 响应参数，每行有2或3个字段，利用字段数量来提取
				var (
					respColumnNum  = 3 // 字段数量
					respNameColumn = 0 // 参数列
					// respTypeColumn     = 1 // 类型列
					respDesColumn = 2 // 说明列
				)
				if !isContainDataType {
					respColumnNum = 2
					respDesColumn = 1
				}
				if len(matches) == respColumnNum {
					api.RespFields = append(api.RespFields, Field{
						Name: pickSubMatchString(matches, respColumnNum, respNameColumn),
						Desc: pickSubMatchString(matches, respColumnNum, respDesColumn),
					})
				}
			})
		})
		apis = append(apis, api)
		fmt.Printf("第%d个h2完成\n", index)
	}

	if len(apis) <= 0 {
		die("没有获取到api\n")
	}

	result, err := generateCode(apis)
	if err != nil {
		die("generateCode failed: %+v\n", err)
	}

	err = ioutil.WriteFile(savePath, result, os.ModePerm)
	if err != nil {
		die("ioutil.WriteFile failed: %+v\n", err)
	}
	err = exec.Command("gofmt", "-w", savePath).Run()
	if err != nil {
		die("exec gofmt failed: %+v\n", err)
	}
	fmt.Printf("保存文件成功:%s\n", savePath)
}

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
	// unreachable
}

func pickSubMatchString(matches []string, mustTotal int, pickIndex int) string {
	if len(matches) == mustTotal {
		return matches[pickIndex]
	}
	return ""
}

func generateStruct(rawJson string, structName string, subStruct bool, fields []Field) (code string, err error) {
	if rawJson == "" {
		code = fmt.Sprintf("type %s struct{}", structName)
		return
	}
	var rawCode []byte
	rawJson = strings.ReplaceAll(rawJson, "\",\n}", "\"\n}")
	rawJson = strings.ReplaceAll(rawJson, "　", " ")
	rawJson = strings.ReplaceAll(rawJson, "，", ",")
	// 修改腾讯文档的json格式错误
	rawJson = regexp.MustCompile(`([\d"]+)\n([ ]+")`).ReplaceAllString(rawJson, "$1,\n$2")
	rawCode, err = gojson.Generate(strings.NewReader(rawJson),
		gojson.ParseJson,
		structName,
		"apis",
		[]string{"json"}, subStruct, true,
	)
	if err != nil {
		fmt.Println(rawJson)
		err = errors.Wrap(err, "gojson.Generate failed")
		return
	}

	code = string(rawCode)
	structFieldRegexp := regexp.MustCompile("(\\w+?)([ ]+?)([\\w\\[\\]]+?)([ ]+?)`json:\"(.+?)\"`")
	allMatches := structFieldRegexp.FindAllStringSubmatch(code, -1)
	code = strings.Replace(code, "package apis\n\n", "", 1)
	replacedFields := make(map[string]Field)
	for _, matches := range allMatches {
		for _, field := range fields {
			if field.Name == "errcode" || field.Name == "errmsg" {
				continue
			}
			requiredTips := ""
			jsonFlag := ""
			if strings.HasPrefix(structName, "Req") {
				jsonFlag = ",omitempty"
			}
			if field.IsRequired {
				requiredTips = "，必填"
				jsonFlag = ""
			}
			// 支持文档里面字段名的点号语法
			if strings.Contains(field.Name, ".") {
				tmps := strings.Split(field.Name, ".")
				field.Name = tmps[len(tmps)-1]
			}
			if len(matches) == 6 && matches[5] == field.Name {
				code = strings.Replace(code,
					matches[0],
					fmt.Sprintf("// %s %s%s\n\t%s", matches[1], field.Desc, requiredTips, matches[0]),
					1,
				)
				code = strings.Replace(code,
					fmt.Sprintf("`json:\"%s\"`", field.Name),
					fmt.Sprintf("`json:\"%s%s\"`", field.Name, jsonFlag),
					1,
				)

				replacedFields[field.Name] = field
			} else {
				if _, ok := replacedFields[field.Name]; ok {
					continue
				}
				code = strings.Replace(code,
					fmt.Sprintf("} `json:\"%s\"`", field.Name),
					fmt.Sprintf("} `json:\"%s%s\"` // %s%s", field.Name, jsonFlag, field.Desc, requiredTips),
					1,
				)
				replacedFields[field.Name] = field
			}
		}
	}

	code = strings.ReplaceAll(code, "int64", "int")
	commonRespRegexp := regexp.MustCompile("Errcode.+\\n.+?Errmsg.+")
	code = commonRespRegexp.ReplaceAllString(code, "CommonResp")
	return
}

func generateCode(apis []*Api) (result []byte, err error) {
	tpl, err := template.ParseFiles("./api_generate/api.tmpl")
	if err != nil {
		err = errors.Wrap(err, "template.ParseFiles failed")
		return
	}
	for _, api := range apis {
		api.URL = strings.ReplaceAll(api.URL, "&amp;", "?")
		api.URL = strings.Split(api.URL, "?")[0]
		segs := strings.Split(api.URL, "/")
		api.StructName = strcase.ToCamel(segs[len(segs)-1]) + strcase.ToCamel(segs[len(segs)-2])
		api.ReqCode, err = generateStruct(api.ReqJson, "Req"+api.StructName, false, api.ReqFields)
		if err != nil {
			fmt.Printf("generate reqStruct failed: %+v\n", err)
			continue
		}

		api.RespCode, err = generateStruct(api.RespJson, "Resp"+api.StructName, false, api.RespFields)
		if err != nil {
			fmt.Printf("generate RespStruct failed: %+v\n", err)
			continue
		}
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, apis)
	if err != nil {
		err = errors.Wrap(err, "tpl.Execute failed")
		return
	}

	result = buf.Bytes()

	return
}

func inSlice(find string, slices []string) bool {
	var isExist bool
	for _, v := range slices {
		if v == find {
			isExist = true
			break
		}
	}
	return isExist
}
