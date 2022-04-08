# 同步传 local

一个「局域网PC与手机互传文件，且不想借助微信/QQ等骚扰软件」的软件

必须要有谷歌浏览器
首先到frontend目录下创建一个 dist目录,然后再当前路径下打开终端,使用` npm install` 或 ` yarn `下载node_module,然后再使用 ` npm run build`编译文件

编译   go build -ldflags -H=windowsgui -o "local.exe"

` go run . `会有问题,因为会找不到  ` config.ini ` 文件



<img src=".\frontend\src\images\synk.png" alt="同步传" style="zoom: 50%;" />

![img](.\frontend\local.png)

`local_test.exe`可以直接在windows上运行

`localsetup.exe`是Windows安装程序

这是某位大佬的教学项目,并不是我自己写的
