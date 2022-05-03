package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"xBlog/tools"
)

func TestExport(t *testing.T) {
	row := `
<p>[success]昨天晚上突然什么也不想做，所以索性就直接躺在床上看番去了，这一看没想到直接肝到两点多，直接一口气把这部番给看完了，感想也很多，所以今天就特地花点时间把这部番剧的感想写一下吧。[/success]</p>
<p>说实话，很久没这么肝的看番了，但是怎么说呢，好的番就是那种可以让你忘记时间，让你心中没有任何杂念，让你完全沉浸于其中。这部番给我的感觉就是如此，我这次看的时候还特意把自己感觉很有感触或者很有哲理的话还截了屏，后面一看截了300多张，同时我也看了一下豆瓣的评论，评分居然高达9.3分，感觉这部番真的不错，所以这次的内容可能会非常多，特别是图片，请大家做好准备。</p>
<p>这部番还是百合番（话说最近的百合番也太多了吧！），讲的就是四个女孩去南极的故事，这个算是一部有点励志的番吧，风格很像《前进吧登山少女》，女主玉木真理（小决）是一个害怕失败的人，番剧一开头就说她想逃一次课，然后去东京旅游，但是却还是中途放弃了，后来在她等车的时候捡到了一个装有100万日元的信封，然后在寻找失主的时候认识了小渕泽报濑，一个想去南极的女生，她辛苦攒的100万就是为了去南极。于是小决就打算和报濑一起去南极，但是她也很犹豫，因为害怕失败，而且她的朋友也不看好这次计划。</p>
<p><img src="https://img.xiaoyou66.com/images/2019/05/26/QRzQ.jpg" /></p>`
	//	首先提取出所有的markdown图片
	imgs := tools.FindAllMatch(row, "\\!\\[]\\((.*?)\\)", false)
	// 然后提取出普通的图片标签
	imgs2 := tools.FindAllMatch(row, "<img.*?src=\"(.*?)\"", false)
	imgs = append(imgs, imgs2...)
	// 遍历图片下载
	for _, img := range imgs {
		filename := downloadImage(img)
		if filename != "" {
			fmt.Println("下载", img, "成功！")
			row = strings.Replace(row, img, "../images/"+filename, 1)
		} else {
			fmt.Println("下载", img, "失败")
		}
	}
	ioutil.WriteFile("a.md", []byte(row), 755)
}

// 下载文件
func downloadImage(url string) string {
	path := "images"
	// 获取文件夹图片数
	info, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("获取文件夹失败", err)
		return ""
	}
	// 获取文件夹后缀
	suffix := url[strings.LastIndex(url, "."):]
	filename := strconv.Itoa(len(info)+1) + suffix

	out, err := os.Create(path + "/" + filename)
	if err != nil {
		fmt.Println("创建文件失败", err)
		return ""
	}
	resp, err := http.Get(url)
	if err != nil {
		_ = out.Close()
		fmt.Println("发送请求失败", err)
		return ""
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if _, err = io.Copy(out, resp.Body); err != nil {
		_ = out.Close()
		fmt.Println("拷贝文件失败", err)
		return ""
	}
	_ = out.Close()

	return filename
}
