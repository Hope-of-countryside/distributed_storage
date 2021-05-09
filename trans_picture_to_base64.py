# base64后文本也太多了，还是上传图片吧！
# coding=utf-8
import base64
import sys

f = open(sys.argv[1], 'rb')  # 二进制方式打开图文件
ls_f = base64.b64encode(f.read())  # 读取文件内容，转换为base64编码
f.close()
print(ls_f)
