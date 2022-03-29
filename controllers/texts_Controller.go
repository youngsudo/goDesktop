package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//  实现接口1: 上传文件
/*
	思路:
	1,获取go执行文件(.exe文件)所在目录
	2,在该目录创建 uploads 目录
	3,将文本保存为一个文件
	4,返回该文件的下载路径
*/
func TextsController(c *gin.Context) {
	var json struct {
		Raw string
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		// os.Getwd会输出实际的工作目录
		// os.Executable会输出一个临时文件的路径，毕竟os.Executable就是要返回当前运行的程序路径，
		// 所以会返回一个go run生成的临时文件路径
		exe, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		// filepath.Dir()函数用于返回指定路径中除最后一个元素以外的所有元素
		/*
			Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；
			如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。
			Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。
			结果是经过简化的，所有的空字符串元素会被忽略。
		*/
		dir := filepath.Dir(exe)
		if err != nil {
			log.Fatal(err)
		}
		// uuid是谷歌开发的生成16字节UUID的模块
		filename := uuid.New().String()
		uploads := filepath.Join(dir, "uploads")
		// os.Mkdir	 创建目录
		// 初次创建dir时成功，再次创建dir时，如果path已经是一个目录，Mkdir会报错
		// os.MkdirAll  创建多级目录,如果path已经是一个目录，MkdirAll什么也不做，并返回nil
		// 必须分成两步：先创建文件夹、再修改权限
		err = os.MkdirAll(uploads, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fullpath := path.Join("uploads", filename+".txt")
		err = ioutil.WriteFile(filepath.Join(dir, fullpath), []byte(json.Raw), 0644)
		if err != nil {
			log.Fatal(err)
		}

		/*
			exe:  "c:\\Users\\young\\Desktop\\lorcademo\\demo3\\__debug_bin.exe"
			dir: "c:\\Users\\young\\Desktop\\lorcademo\\demo3"
			filename: "79a89ddf-5025-4c75-9716-6405e01b37c2"
			uploads: "c:\\Users\\young\\Desktop\\lorcademo\\demo3\\uploads"
			fullpath: "uploads/79a89ddf-5025-4c75-9716-6405e01b37c2.txt"
		*/
		c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
	}

	/*
		    获取当前目录
			os.Getwd()

			创建文件
			f1, _ := os.Create("./1.txt")
			defer f1.Close()

			以读写方式打开文件，如果不存在则创建文件，等同于上面os.Create
			f4, _ := os.OpenFile("./4.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			defer f4.Close()

			用os.path.join()连接两个文件名地址的时候，就比如
			os.path.join("D:\","test.txt")      \\结果是D:\test.txt

			删除指定目录下所有文件
			os.Remove("abc/d/e/f")

			删除指定目录
			os.RemoveAll("abc")

			重命名文件
			os.Rename("./2.txt", "./2_new.txt")
	*/

}
