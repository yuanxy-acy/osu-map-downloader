# osu-map-downloader

------
## 使用方法

### 设置配置文件

    可以直接运行一次程序，会自动生成配置文件模板到程序的同路径下

#### 获取osu_session
    需要先登陆osu官网

按下ctrl+shift+I打开开发者工具

选择网络

点击类型为document的项目

转到标头栏

下翻找到Cookie对应的值

复制osu_session=后面的值

将<your-osu-session-cookie-value>替换为你复制的值

#### 设置下载文件存放路径

将你希望的文件存放路径替换掉配置文件模板中的<your-path>

### 运行程序

    目前仅支持下载所有ranked图，不支持筛选

#### 遍历mapset

下载前需要先拉取mapset信息

用命令行打开存放程序的文件目录

运行指令：

````[shell]
./downloader.exe -m find
````

#### 下载mapset

遍历完成后就可以进行下载了

````[shell]
./downloader.exe -m download
````

--------