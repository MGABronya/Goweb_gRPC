# Goweb_gRPC

在使用前需要安装protobuf

进入protobuf并选择对应的版本进行下载。

设置编译目录

````
$ ./configure -- prefix=/usr/local/protobuf
````

用make编译命令安装，先运行make，在编译成功后在执行install命令

````
$ make
$ make install
````

配置环境变量，打开.bash_profile文件并编辑

````
$ cd ~
$ vim .bash_profile
````

在打开的bash_profile文件末尾添加

````
export PROTOBUF=/usr/local/protobuf
export PATH=$PROTOBUF/bin:$PATH
````

通过source命令使文件生效

````
$ source .bash_profile
````

后安装Go语言protobuf包即可