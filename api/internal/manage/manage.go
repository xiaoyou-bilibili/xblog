package manage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"xBlog/internal/db"
	"xBlog/pkg/database"
	"xBlog/tools"
)

// Init 菜单栏初始化检测
func Init() {
	//获取启动的参数
	if len(os.Args) > 1 && os.Args[1] == "blog" {
		//mongo数据库连接池初始化
		database.DbInit()
		fmt.Println("===============博客系统管理面板[v1.0]=============")
		fmt.Println("      1:已废弃    2:重置管理员密码")
		fmt.Println("      3:修改博客端口      4:取消")
		fmt.Printf("==================================================\n请输入命令编号:")
		var input int
		var inputS string
		if _, err := fmt.Scan(&input); err == nil {
			switch input {
			case 1:
				fmt.Printf("该接口已废弃~")
				break
			case 2:
				fmt.Printf("请输入新的管理员密码，按回车即可重置:")
				if _, err := fmt.Scan(&inputS); err == nil {
					if database.NewDb(db.CollUser).SetFilter(bson.M{"user_id": 1}).Set(bson.M{"password": tools.Encrypt(inputS), "status": 1, "login_fail": 0}).UpdateOne() == nil {
						fmt.Println("你的密码已重置,请重新登录博客系统")
					} else {
						fmt.Println("密码重置失败")
					}
				}
				break
			case 3:
				fmt.Printf("输入新的端口号:")
				if _, err := fmt.Scan(&input); err == nil {
					if tools.SetConfig("site", "apiPort", tools.Int2Str(input)) {
						fmt.Println("修改端口成功，请重新启动容器")
					} else {
						fmt.Println("修改端口失败！")
					}
				}
				break
			}
		} else {
			fmt.Println("命令有误,退出管理系统")
		}
		os.Exit(1)
	} else if len(os.Args) > 1 && os.Args[1] == "export" {
		articleExport()
	}
}

// 文章导出
func articleExport() {
	fmt.Println("文章导出中。。。")

	// 首先我们删除所有文件夹避免重复
	dir, err := ioutil.ReadDir("post")
	if err == nil {
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{"post", d.Name()}...))
		}
	}

	articles := new([]db.Article)
	// 创建对应的文件夹
	os.Mkdir("post", os.ModePerm)
	os.Mkdir("post/post", os.ModePerm)
	os.Mkdir("post/doc", os.ModePerm)
	os.Mkdir("post/diary", os.ModePerm)
	os.Mkdir("post/images", os.ModePerm)
	// 查找所有的文件
	if err := database.NewDb(db.CollArticle).FindMore(articles); err == nil {
		fmt.Println("已获取到", len(*articles), "篇文章")
		for _, v := range *articles {
			// 文章标题和日记
			path := "post/post/"
			title := "post"
			switch v.PostType {
			case "post":
				title = v.Title
				path = "post/post/"
			case "doc":
				title = v.Title
				path = "post/doc/"
			case "diary":
				// 如果是日记，那么标题就是对应的时间
				title = tools.Time2String(v.PostTime, false)
				path = "post/diary/"
			}

			// 文章内容
			content := v.Md
			if content == "" {
				content = v.Content
			}

			// 替换内容
			content = imageReplace(content)

			err := ioutil.WriteFile(path+title+".md", []byte(content), 755)
			if err != nil {
				fmt.Println("导出", title, "失败！")
			} else {
				fmt.Println("导出", title, "成功！")
			}
		}
	}
}

// 替换掉文件里面的图片
func imageReplace(content string) string {
	//	首先提取出所有的markdown图片
	imgs := tools.FindAllMatch(content, "\\!\\[]\\((.*?)\\)", false)
	// 然后提取出普通的图片标签
	imgs2 := tools.FindAllMatch(content, "<img.*?src=\"(.*?)\"", false)
	imgs = append(imgs, imgs2...)
	// 遍历图片下载
	for _, img := range imgs {
		filename := downloadImage(img)
		if filename != "" {
			fmt.Println("下载", img, "成功!")
			// 替换所有的图片
			content = strings.Replace(content, img, "../images/"+filename, -1)
		} else {
			fmt.Println("下载", img, "失败!")
		}
	}
	return content
}

// 下载文件
func downloadImage(url string) string {
	imgPath := "post/images"
	// 获取文件夹图片数
	info, err := ioutil.ReadDir(imgPath)
	if err != nil {
		fmt.Println("获取文件夹失败", err)
		return ""
	}
	// 获取文件夹后缀
	suffix := ""
	if strings.LastIndex(url, ".") > 0 {
		suffix = url[strings.LastIndex(url, "."):]
	}

	filename := strconv.Itoa(len(info)+1) + suffix

	out, err := os.Create(imgPath + "/" + filename)
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
