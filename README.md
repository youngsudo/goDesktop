# 同步传 local

一个「局域网PC与手机互传文件，且不想借助微信/QQ等骚扰软件」的软件

必须要有谷歌浏览器

编译   go build -ldflags -H=windowsgui -o "local.exe"

` go run . `会有问题,因为会找不到  ` config.ini ` 文件



![img](https://github.com/young-sudo/goDesktop/blob/main/frontend/src/images/synk.png alt="同步传" style="zoom: 50%;" />

![img](https://github.com/young-sudo/goDesktop/blob/main/.%5Cfrontend%5Clocal.png /)

`local_test.exe`可以直接在windows上运行

`localsetup.exe`是Windows安装程序

这是某位大佬的教学项目,并不是我自己写的
