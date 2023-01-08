# 使用現有的 container 變更後建立新的 image
[lab source](https://github.com/sixeyed/diamol/tree/master/ch03/lab)

```shell
# 執行 container 並進入互動介面
docker run -it --name ch03lab diamol/ch03-lab

# 更改 container 內的資訊
echo Shou >> ch03.txt

# 對變更後的 container 做成 image 並命名
# docker commit 7ad4b web-ping:v6
docker commit [container-id] [image:tag]
```
