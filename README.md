# gocntool
A tool for gocn/news conflict check.

## Installation

Go get with:

```
go get github.com/songjiayang/gocntool
```

## Usage

There is a file named `gocn.md` with content:

```
## GoCN每日新闻(2019-03-22)

1. 从 dep 迁移到 mod http://elliot.land/post/migrating-projects-from-dep-to-go-modules
2. Go http.Flusher 在实际项目中的应用 http://www.songjiayang.com/posts/go-http-zhong-flusher-zai-shi-ji-kai-fa-zhong-de-ying-yong
3. 结合 docker 做 Go 项目集成测试 https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
4.从 nginx 热更新聊一聊 Go 中的热更新（下） https://zhuanlan.zhihu.com/p/59196185
5. 使用排队论做容量规划 https://hackernoon.com/why-capacity-planning-needs-queueing-theory-without-the-hard-math-342a851e215c

* 第五届GopherChina大会门票 https://www.bagevent.com/event/gocn5th

编辑: lwhile
订阅新闻: http://tinyletter.com/gocn
GoCN归档: https://gocn.vip/question/3213

```

Run command `gocntool -file ./gocn.md` then will get output like:

```
2019/03/22 15:24:35 start check: http://elliot.land/post/migrating-projects-from-dep-to-go-modules
2019/03/22 15:24:36 after check: http://elliot.land/post/migrating-projects-from-dep-to-go-modules
2019/03/22 15:24:37 start check: http://www.songjiayang.com/posts/go-http-zhong-flusher-zai-shi-ji-kai-fa-zhong-de-ying-yong
2019/03/22 15:24:38 after check: http://www.songjiayang.com/posts/go-http-zhong-flusher-zai-shi-ji-kai-fa-zhong-de-ying-yong
2019/03/22 15:24:39 start check: https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
2019/03/22 15:24:39 after check: https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
2019/03/22 15:24:40 start check: https://zhuanlan.zhihu.com/p/59196185
2019/03/22 15:24:41 after check: https://zhuanlan.zhihu.com/p/59196185
2019/03/22 15:24:42 start check: https://hackernoon.com/why-capacity-planning-needs-queueing-theory-without-the-hard-math-342a851e215c
2019/03/22 15:24:42 after check: https://hackernoon.com/why-capacity-planning-needs-queueing-theory-without-the-hard-math-342a851e215c

There are 5 conflicts:
http://elliot.land/post/migrating-projects-from-dep-to-go-modules
http://www.songjiayang.com/posts/go-http-zhong-flusher-zai-shi-ji-kai-fa-zhong-de-ying-yong
https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
https://zhuanlan.zhihu.com/p/59196185
https://hackernoon.com/why-capacity-planning-needs-queueing-theory-without-the-hard-math-342a851e215c
```

If everything is ok, you will get output like:

```
Congratulations！check passed.
```

If you have any requestion let me known please, now just try it.
