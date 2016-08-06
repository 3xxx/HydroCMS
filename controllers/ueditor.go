package controllers

import (
	// "bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	// "quick/models"
	"encoding/base64"
	"io/ioutil"
	"quick/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UeditorController struct {
	beego.Controller
}

type UploadimageUE struct {
	url      string
	title    string
	original string
	state    string
	// "url": fmt.Sprintf("/static/upload/%s", filename),
	// "title": "demo.jpg",
	// "original": header.Filename,
	// "state": "SUCCESS"
}

// func (c *UeditorController) ControllerUE(w http.ResponseWriter, r *http.Request) {
// 	action := r.URL.Query()["action"][0]
// 	beego.Info(action)
// 	fmt.Println(r.Method, action)
// 	if r.Method == "GET" {
// 		if action == "config" {
// 			Configs(w, r)
// 		}
// 	} else if r.Method == "POST" {
// 		if action == "uploadimage" {
// 			UploadImage(w, r)
// 		}
// 	}
// }

func (c *UeditorController) ControllerUE() {
	op := c.Input().Get("action")
	key := c.Input().Get("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
	switch op {
	case "config": //这里还是要优化成conf/config.json

		// $CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		// $action = $_GET['action'];

		// switch ($action) {
		//     case 'config':
		//         $result =  json_encode($CONFIG);

		// var configJson []byte // 当客户端请求/controller?action=config 返回的json内容
		file, err := os.Open("conf/config.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		// fi, err := os.Open("d:/config.json")
		// if err != nil {
		// 	panic(err)
		// }
		// defer fi.Close()
		fd, err := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/") //参考php的$CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		//将php中的正则移植到go中，需要将/ \/\*[\s\S]+?\*\/  /去掉前后的/，然后将\改成2个\\
		//参考//去除所有尖括号内的HTML代码，并换成换行符
		// re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		// src = re.ReplaceAllString(src, "\n")
		//当把<和>换成/*和*\时，斜杠/和*之间加双斜杠\\才行。
		src = re.ReplaceAllString(src, "")
		tt := []byte(src)

		// buf := bytes.NewBuffer(nil)
		// buf.ReadFrom(file)
		// configJson = buf.Bytes()
		// beego.Info(configJson)
		// w.Write(configJson)
		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码
		c.Data["json"] = r
		c.ServeJSON()

		//下面这段是测试用的
		// b := []byte(`{
		//             "imageActionName": "uploadimage",
		//             "imageFieldName": "upfile",
		//             "imageMaxSize": 2048000,
		//             "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
		//             "imageCompressEnable": true,
		//             "imageCompressBorder": 1600,
		//             "imageInsertAlign": "none",
		//             "imageUrlPrefix": "",
		//             "imagePathFormat": "/static/upload/{yyyy}{mm}{dd}/{time}{rand:6}"
		//       }`)
		// var r interface{}
		// json.Unmarshal(b, &r)
		// c.Data["json"] = r
		// c.ServeJSON()
	case "uploadimage", "uploadfile", "uploadvideo":
		// file, header, err := c.GetFile("upfile") // r.FormFile("upfile")
		// if err != nil {
		// 	panic(err)
		// }
		// defer file.Close()
		// filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		switch key {
		case "diary": //添加文章
			// title := c.Input().Get("title")
			// tnumber := c.Input().Get("tnumber")
			// content := c.Input().Get("content")
			// category := c.Input().Get("category")
			categoryid := c.Input().Get("categoryid")
			// beego.Info(categoryid)
			//保存上传的图片
			_, h, err := c.GetFile("upfile")
			if err != nil {
				beego.Error(err)
			}

			// category1, _, err := models.GetCategory(categoryid)
			// if err != nil {
			// 	beego.Error(err)
			// 	return
			// }
			url, diskdirectory, err := models.GetCategoryUrl(categoryid)
			if err != nil {
				beego.Error(err)
			}
			var filesize int64
			fileSuffix := path.Ext(h.Filename)
			// random_name
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			// if err != nil {
			// 	beego.Error(err)
			// }
			path1 := diskdirectory + newname    //h.Filename
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			filesize, _ = FileSize(path1)
			filesize = filesize / 1000.0
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": url + newname, "title": newname, "original": h.Filename}
			c.ServeJSON()
		case "wiki": //添加wiki
			// title := c.Input().Get("title")
			// tnumber := c.Input().Get("tnumber")
			// content := c.Input().Get("content")
			// category := c.Input().Get("category")
			// categoryid := c.Input().Get("categoryid")
			//保存上传的图片
			_, h, err := c.GetFile("upfile")
			if err != nil {
				beego.Error(err)
			}

			// category1, err := models.GetCategory(categoryid)
			// if err != nil {
			// 	beego.Error(err)
			// 	return
			// }
			var filesize int64
			fileSuffix := path.Ext(h.Filename)
			// random_name
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			// if err != nil {
			// 	beego.Error(err)
			// }
			year, month, _ := time.Now().Date()

			err = os.MkdirAll(".\\attachment\\wiki\\"+strconv.Itoa(year)+month.String()+"\\", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
			path1 := ".\\attachment\\wiki\\" + strconv.Itoa(year) + month.String() + "\\" + newname //h.Filename
			Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			filesize, _ = FileSize(path1)
			filesize = filesize / 1000.0
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": Url + newname, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()
		default: //添加封面、简介
			number := c.Input().Get("number")
			name := c.Input().Get("name")
			err := os.MkdirAll(".\\attachment\\"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
			// err = os.MkdirAll(path.Join("static", "upload"), 0775)
			// if err != nil {
			// 	panic(err)
			// }
			// outFile, err := os.Create(path.Join("static", "upload", filename))
			// if err != nil {
			// 	panic(err)
			// }

			// diskdirectory := ".\\attachment\\" + number + name + "\\"
			// url := "/attachment/" + number + name + "/"

			//保存上传的图片
			//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
			_, h, err := c.GetFile("upfile")
			// beego.Info(h)
			if err != nil {
				beego.Error(err)
			}
			// var attachment string
			// var path string
			var filesize int64
			// var route string
			//保存附件
			// attachment = h.Filename
			// beego.Info(attachment)
			path1 := ".\\attachment\\" + number + name + "\\" + h.Filename
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			//如果扩展名为jpg
			// if strings.ToLower(path.Ext(h.Filename)) == ".jpg" {
			// }
			//如果包含jpg，则进行压缩——压缩导致UEditor里显示尺寸过大。
			// if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
			// 	// 随机名称
			// 	// to := path + random_name() + ".jpg"
			// 	origin := path1 //path + file.Name()
			// 	fmt.Println("正在处理" + origin + ">>>" + origin)
			// 	cmd_resize(origin, 2048, 0, origin)
			// 	//defer os.Remove(origin)//删除原文件
			// }
			filesize, _ = FileSize(path1)
			filesize = filesize / 1000.0
			// route = "/attachment/" + number + name + "/" + h.Filename
			// outFile, err := os.Create(path.Join(".\\attachment\\"+number+name+"\\", filename))
			// if err != nil {
			// 	beego.Error(err)
			// }
			// defer outFile.Close()
			// io.Copy(outFile, file)
			// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + filename, "title": filename, "original": filename}
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/attachment/" + number + name + "/" + h.Filename, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()
			// 		{
			//     "state": "SUCCESS",
			//     "url": "upload/demo.jpg",
			//     "title": "demo.jpg",
			//     "original": "demo.jpg"
			// }
		}
	case "uploadscrawl":
		number := c.Input().Get("number")

		name := c.Input().Get("name")
		err := os.MkdirAll(".\\attachment\\"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		path1 := ".\\attachment\\" + number + name + "\\"
		//保存上传的图片
		//upfile为base64格式文件，转成图片保存
		ww := c.Input().Get("upfile")
		ddd, _ := base64.StdEncoding.DecodeString(ww)           //成图片文件并把文件写入到buffer
		newname := strconv.FormatInt(time.Now().Unix(), 10)     // + "_" + filename
		err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err != nil {
			beego.Error(err)
		}
		var filesize int64
		filesize, _ = FileSize(path1)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{
			"state":    "SUCCESS",
			"url":      "/attachment/" + number + name + "/" + newname + ".jpg",
			"title":    newname + ".jpg",
			"original": newname + ".jpg",
		}
		c.ServeJSON()
	case "listimage":
		type List struct {
			Url string `json:"url"`
			// Source string
			// State  string
		}

		type Listimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
			Start int    `json:"start"`
			Total int    `json:"total"`
			// Name        string
			// Age         int
			// Slices      []string //slice
			// Mapstring   map[string]string
			// StructArray []List            //结构体的切片型
			// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
			//	Desks  List
		}

		// var m map[string]string = make(map[string]string)
		// m["Go"] = "No.1"
		// m["Java"] = "No.2"
		// m["C"] = "No.3"
		// fmt.Println(m)

		list := []List{
			{"/static/upload/1.jpg"},
			{"/static/upload/2.jpg"},
			// {"upload/1.jpg", "http://a.com/1.jpg", "SUCCESS"},
			// {"upload/2.jpg", "http://b.com/2.jpg", "SUCCESS"},
		}
		// var mm map[string][]List = make(map[string][]List)
		// mm["Go"] = list
		// mm["Java"] = list
		// fmt.Println(mm)

		listimage := Listimage{"SUCCESS", list, 1, 21}
		// beego.Info(listimage){SUCCESS [{/static/upload/1.jpg} {/static/upload/2.jpg}] 1 21}
		// fmt.Println(listimage)
		// b, _ := json.Marshal(listimage)
		// mystruct := { ... }
		// c.Data["jsonp"] = listimage
		// beego.Info(string(b)){"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}
		// c.ServeJSONP()
		c.Data["json"] = listimage
		c.ServeJSON()
		// c.Data["json"] = map[string]interface{}{"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}

		// 需要支持callback参数,返回jsonp格式
		// {
		//     "state": "SUCCESS",
		//     "list": [{
		//         "url": "upload/1.jpg"
		//     }, {
		//         "url": "upload/2.jpg"
		//     }, ],
		//     "start": 20,
		//     "total": 100
		// }
	case "catchimage":

		type List struct {
			Url    string `json:"url"`
			Source string `json:"source"`
			State  string `json:"state"`
		}

		type Catchimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
			// Start int
			// Total int
			// Name        string
			// Age         int
			// Slices      []string //slice
			// Mapstring   map[string]string
			// StructArray []List            //结构体的切片型
			// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
			//	Desks  List
		}

		// var m map[string]string = make(map[string]string)
		// m["Go"] = "No.1"
		// m["Java"] = "No.2"
		// m["C"] = "No.3"
		// fmt.Println(m)

		list := []List{
			{"/static/upload/1.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			{"/static/upload/2.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			// {"upload/1.jpg", "http://a.com/1.jpg", "SUCCESS"},
			// {"upload/2.jpg", "http://b.com/2.jpg", "SUCCESS"},
		}
		// var mm map[string][]List = make(map[string][]List)
		// mm["Go"] = list
		// mm["Java"] = list
		// fmt.Println(mm)

		catchimage := Catchimage{"SUCCESS", list}
		// beego.Info(catchimage){SUCCESS [{/static/upload/1.jpg} {/static/upload/2.jpg}] 1 21}
		// fmt.Println(catchimage)
		// b, _ := json.Marshal(catchimage)
		// mystruct := { ... }
		// c.Data["jsonp"] = catchimage
		// beego.Info(string(b)){"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}
		// c.ServeJSONP()
		c.Data["json"] = catchimage
		c.ServeJSON()

		file, header, err := c.GetFile("source") // r.FormFile("upfile")
		beego.Info(header.Filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			panic(err)
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		io.Copy(outFile, file)
		// 需要支持callback参数,返回jsonp格式
		// list项的state属性和最外面的state格式一致
		// {
		//     "state": "SUCCESS",
		//     "list": [{
		//         "url": "upload/1.jpg",
		//         "source": "http://b.com/2.jpg",
		//         "state": "SUCCESS"
		//     }, {
		//         "url": "upload/2.jpg",
		//         "source": "http://b.com/2.jpg",
		//         "state": "SUCCESS"
		//     }, ]
		// }

		// f := &UploadimageUE{
		// 	state:    "SUCCESS",
		// 	url:      fmt.Sprintf("/static/upload/%s", filename),
		// 	title:    "demo.jpg",
		// 	original: header.Filename,
		// }
		// c.Data["json"] = f
		// c.ServeJSON()
		// 	reply := &Comment{
		// Tid:     tidNum,
		// Name:    nickname,
		// Content: content,

		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(string(b))
		// w.Write(b)
		// 	c.Data["json"] = map[string]string{
		// 	"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
		// 	"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
		// 	"original": h.Filename,                                     //原始文件名
		// 	"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
		// }
		// var s interface{}
		// json.Unmarshal(f, &s)
		// s, err := json.Marshal(f)
		// if err == nil {
		// c.Ctx.WriteString(string(f))

		// default:
	}
	// c.Write(configJson)
}

// func Configs(w http.ResponseWriter, r *http.Request) {
// 	w.Write(configJson)
// }

// var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

// func init() {
// 	file, err := os.Open("conf/config.json")
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	defer file.Close()
// 	buf := bytes.NewBuffer(nil)
// 	buf.ReadFrom(file)

// 	configJson = buf.Bytes()
// }

// func (c *UeditorController) UploadImage() {
// 	name := "111"    //c.Input().Get("name")
// 	number := "222"  //c.Input().Get("number")
// 	content := "333" //c.Input().Get("test-editormd-html-code")
// 	path := "c"      //c.Input().Get("tempString")

// 	diskdirectory := ".\\attachment\\" + "test" + "\\"
// 	url := "/attachment/" + "test" + "/"
// 	//保存上传的图片
// 	//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
// 	_, h, err := c.GetFile("upfile") //editormd-image-file
// 	beego.Info(h)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// var attachment string
// 	// var path string
// 	var filesize int64
// 	var route string
// 	if h != nil {
// 		//保存附件
// 		path1 := ".\\attachment\\" + "test" + "\\" + h.Filename
// 		err = c.SaveToFile("upfile", path1) //editormd-image-file  .Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			beego.Error(err)
// 		}

// 		if strings.Contains(strings.ToLower(h.Filename), ".jpg") { //ToLower转成小写
// 			// 随机名称
// 			// to := path + random_name() + ".jpg"
// 			origin := path1 //path + file.Name()
// 			fmt.Println("正在处理" + origin + ">>>" + origin)
// 			cmd_resize(origin, 2048, 0, origin)
// 			//				defer os.Remove(origin)//删除原文件
// 		}
// 		filesize, _ = FileSize(path1)
// 		filesize = filesize / 1000.0
// 		route = "/attachment/" + "test" + "/" + h.Filename
// 	} else {
// 		img := CreateRandomAvatar([]byte(number + name))
// 		fi, _ := os.Create("./attachment/" + "test" + "/u1.png")
// 		png.Encode(fi, img)
// 		fi.Close()
// 		route = "/attachment/" + "test" + "/u1.png"
// 	}

// 	uname := "4"

// 	//存入数据库
// 	_, err = models.AddCategory(name, number, content, path, route, uname, diskdirectory, url)
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		// f := Uploadimage{
// 		// 	url:     route,
// 		// 	success: 1,
// 		// 	message: "ok",
// 		// }
// 		// beego.Info(f)2016/01/17 01:40:03 [category.go:549] [I] {/attachment/test/u1.png ok 1}

// 		// c.Data["json"] = map[string]interface{}{"success": 1, "message": "111", "url": route}

// 		// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": route}
// 		// c.Data["json"] = f
// 		c.Data["json"] = map[string]string{
// 			"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
// 			"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
// 			"original": h.Filename,                                     //原始文件名
// 			"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
// 		}
// 		c.ServeJSON()
// 		// beego.Info(c.Data["json"])
// 		// 2016/01/17 01:42:00 [category.go:554] [I] map[success:1 message:111 url:/attachm
// 		// ent/test/u1.png]
// 		// 		{
// 		//     "state": "SUCCESS",
// 		//     "url": "upload/demo.jpg",
// 		//     "title": "demo.jpg",
// 		//     "original": "demo.jpg"
// 		//      }
// 	}

// 	// c.Data["Uname"] = ck.Value
// 	// id1 := strconv.FormatInt(id, 10)
// 	// c.Redirect("/category?op=view&id="+id1, 301)
// 	return //???
// }

func UploadImage(w http.ResponseWriter, r *http.Request) { //这个没用
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, file)
	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w.Write(b)
}

func (c *UeditorController) UploadImage() { //对应这个路由 beego.Router("/controller", &controllers.UeditorController{}, "post:UploadImage")

	file, header, err := c.GetFile("upfile") // r.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, file)
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/" + filename, "title": "111", "original": "demo.jpg"}
	c.ServeJSON()
	// "state": "SUCCESS",
	// "url": "upload/demo.jpg",
	// "title": "demo.jpg",
	// "original": "demo.jpg"
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
	// w.Write(b)
	// 	c.Data["json"] = map[string]string{
	// 	"url":      fmt.Sprintf("/attachment/test/%s", h.Filename), //保存后的文件路径
	// 	"title":    "",                                             //文件描述，对图片来说在前端会添加到title属性上
	// 	"original": h.Filename,                                     //原始文件名
	// 	"state":    "SUCCESS",                                      //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	// }
	// c.Data["json"] = b
	// c.ServeJSON()
}
