package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/miku/zek"
	"gitlab.ydjdev.com/ydjai/workweixin/callbacks"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Api struct {
	DocUrl     string
	FileName   string
	StructName string
	XmlStr     string
	StructStr  string
	MsgType    string
	EventType  string
	ChangeType string
	Msg        callbacks.CallbackMessage
}

var docVar = flag.String("doc", "", "微信文档地址")

func main() {
	flag.Parse()

	var docURL, savePath string

	if docVar != nil {
		docURL = *docVar
	}

	if docURL == "" {
		fmt.Println("请输入参数doc(企业微信开发文档地址):")
		_, _ = fmt.Scanf("%s", &docURL)
	}
	if docURL == "" {
		die("必传参数 doc=?")
	}

	tag := strings.Index(docURL, "#")
	if tag > 0 {
		docURL = docURL[:tag]
	}

	tpl, err := template.ParseFiles("./callback_generate/temp.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles failed:", err)
		return
	}

	var doc *goquery.Document
	resp, err := http.Get(docURL)
	if err != nil {
		die("http get of errcode documentation failed: %+v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		die("non-200 app: %s\n", resp)
	}

	tmp, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		die("parse document failed: %+v\n", err)
	}

	doc = tmp

	rawHtml, err := doc.Find("#js_doc_apiShow_cnt").Html()
	if err != nil {
		die("failed to get html: %+v\n", err)
	}

	fmt.Println("开始抓取和生成API代码，文档地址:", docURL)

	rawHtml = regexp.MustCompile(`<h2 id="([0-9a-zA-Z\-_]+)">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtml = regexp.MustCompile(`<h1 id="\w+" class="\w+">`).ReplaceAllString(rawHtml, `<h1 class="rawHtmlSection">`)
	rawHtmlSections := strings.Split(rawHtml, `<h2 class="rawHtmlSection">`)
	rawHtmlType := "h2" // 一个页面多个接口
	if len(rawHtmlSections) == 1 {
		rawHtml = regexp.MustCompile(`<h3 id="([0-9a-zA-Z\-_]+)">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
		rawHtml = regexp.MustCompile(`<h4 id="([0-9a-zA-Z\-_]+)">`).ReplaceAllString(rawHtml, `<h2 class="rawHtmlSection">`)
		rawHtmlSections = strings.Split(rawHtml, `<h2 class="rawHtmlSection">`)

		if len(rawHtmlSections) == 1 {
			rawHtmlType = "h1" // 一个页面一个接口
			rawHtmlSections = strings.Split(rawHtml, `<h1 class="rawHtmlSection">`)
		}
	}

	for index, rawHtmlSection := range rawHtmlSections {
		fmt.Printf("\n\n开始处理第%d个接口\n", index)
		api := &Api{
			DocUrl: docURL,
		}

		if rawHtmlType == "h1" {
			rawHtmlSection = `<h1 class="rawHtmlSection">` + rawHtmlSection
			apiNameRegexp := regexp.MustCompile(`<h1 class="rawHtmlSection">(.+?)</h1>`)
			api.FileName = pickSubMatchString(apiNameRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		} else {
			rawHtmlSection = `<h2 class="rawHtmlSection">` + rawHtmlSection
			apiNameRegexp := regexp.MustCompile(`<a name="(.+?)" class="reference-link"></a>`)
			api.FileName = pickSubMatchString(apiNameRegexp.FindStringSubmatch(rawHtmlSection), 2, 1)
		}

		if api.FileName == "" {
			fmt.Println("接口标题为空，跳过处理")
			continue
		}
		fmt.Println(api.FileName)

		api.DocUrl = docURL + "#" + api.FileName
		tmp, err := goquery.NewDocumentFromReader(bytes.NewBufferString(rawHtmlSection))
		if err != nil {
			fmt.Println("html解析失败:", err)
			continue
		}

		api.XmlStr, err = tmp.Find("pre").Find("code").Html()
		if err != nil {
			fmt.Println("xml解析失败:", err)
			continue
		}

		api.XmlStr = strings.ReplaceAll(api.XmlStr, "&lt;", "<")
		api.XmlStr = strings.ReplaceAll(api.XmlStr, "&gt;", ">")

		api.Msg, _ = callbacks.CallbackMessage{}.ParseMessageFromXml([]byte(api.XmlStr))

		api.StructName = api.Msg.GetStructName()

		api.MsgType = string(api.Msg.MsgType)
		api.EventType = string(api.Msg.EventType)
		api.ChangeType = string(api.Msg.ChangeType)

		api.StructStr, err = gen(api.XmlStr)
		if err != nil || api.StructStr == "" {
			fmt.Println("struct解析失败:", err)
			fmt.Println(tmp.Html())
			continue
		}

		api.StructStr = strings.ReplaceAll(api.StructStr, "type XML struct", "type "+api.StructName+" struct")

		buf := bytes.NewBufferString("")
		err = tpl.Execute(buf, api)
		if err != nil {
			fmt.Println("写入模板失败:", err)
			continue
		}

		savePath = "./callbacks/" + api.MsgType + api.FileName + ".go"
		err = ioutil.WriteFile(savePath, bytes.Replace(buf.Bytes(), []byte("&#34;"), []byte("\""), -1), os.ModePerm)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			continue
		}
		fmt.Println("保存文件成功:", savePath)
	}
}

func pickSubMatchString(matches []string, mustTotal int, pickIndex int) string {
	if len(matches) == mustTotal {
		return matches[pickIndex]
	}
	return ""
}

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func gen(s string) (string, error) {
	node := new(zek.Node)

	raw := bytes.NewBufferString(s)

	if _, err := node.ReadFrom(raw); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	sw := zek.NewStructWriter(&buf)

	if err := sw.WriteNode(node); err != nil {
		return "", err

	}
	return buf.String(), nil
}
