# GoDaddy DDNS

## 一、 简介/Introduction

这是一个用 GO 语言编写的 DNS 记录更新程序，仅适用于 GoDaddy 域名提供商。程序运行后会将本地网络的出口 IP 更新到指定的 DNS 记录中。

## 二、 使用方法/Usage

1. 首先需要申请访问 GoDaddy API 所需要的 KEY 和 SECRET，申请地址：[https://developer.godaddy.com/keys](https://developer.godaddy.com/keys)

1. 编译或下载可执行文件 godaddyddns

1. 执行如下命令（以 Windows 环境为例）<br>`godaddyddns.exe --domain=example.com --key=your_key --secret=your_secret --host=your_sub_domain`

1. 如果当前 IP 地址与 DNS 记录中的不一致，则会看到提示<br> `Update dns successfully!` <br>说明更新 DNS 记录成功