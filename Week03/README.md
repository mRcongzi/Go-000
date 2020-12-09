# Goroutine

#### goroutine泄露

1. 管住goroutine的生命周期

2. recover住goroutine的panic

   <!--野生的goroutine：指没有生命周期（即使用者不知道这个goroutine是什么时候结束的）或使用panic却没有做recover-->

#### 并发和并行

并发：两个队伍共用一台咖啡机

并行：两个队伍各自用一台咖啡机

![preview](https://pic2.zhimg.com/v2-674f0d37fca4fac1bd2df28a2b78e633_r.jpg?source=1940ef5c)