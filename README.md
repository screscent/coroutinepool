# coroutinepool
=============
go coroutine pool  

找到了一个[goroutine性能测试](http://www.mikespook.com/2012/01/goroutine%E6%80%A7%E8%83%BD%E6%B5%8B%E8%AF%95/)  
里面提到了goroutine可以并发开启上百万个goroutine，所以这个goroutine pool没有什么实际意义。  
当练习吧  

# License
===
author gonghh   
Copyright 2013 gonghh(screscent).  
Under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).  


## example
==============================


* create

```go
coroutinepool.Create(6)

//or

p:=coroutinepool.NewPool()
p.Create(6)
```
* run

```go
  for i := 0; i < 100; i++ {
		coroutinepool.Run(func() {
			fmt.Println("test", i)
		})
  }

//or
  for i := 0; i < 100; i++ {
		p.Run(func() {
			fmt.Println("test", i)
		})
  }
  
```
* exit

```go
coroutinepool.Exit()

//or
p.Exit()
```
