package main

/*
	见https://colobu.com/2017/08/04/talk-about-getting-goroutine-id-again/
	1 通过堆栈信息解析出ID
	2 通过汇编获取runtime·getg()的调用结果
	3 修改运行时代码，export一个可以外部调用的getID()

	1比较慢，2因为是hack的方式(Go Team并不想暴露goroutine id)，针对不同Go版本需要特殊的hack手段
	3需要定制Go运行时，不通用
	提供了1的方法，在2方法不起作用时作为备选
 */
