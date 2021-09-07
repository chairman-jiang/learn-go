package main

func frontendInterView() {
	/*
		* 1. template下为什么有且只有一个dom元素 (3.x 支持了Fragment 可以多节点)
		* 2. data为什么要求是函数(3.x中要求必须是函数, js原生的Object将不再支持)
		* 3. 如何在子组件中访问父组件的实例 (3.x移除了 $children)
		* 4. 动态给vue的data添加一个新的属性时会发生什么？怎样解决？
		* 5. 如果现在让你从vue/react/angular三个中选择一个，你会选哪个？说说你的理由
		* 6. 怎样让一个自定义组件支持v-model -> 依赖哪个属性 model:{ prop: 'value', event: 'input' } (3.x中 统一了v-model和.sync; .sync也是一个语法糖(@update:val @customEvent="val= $event"), 写法为v-model:title="curValue")
		* 7. 能够实现一个类似$message或者$confirm的组件吗, 大致思路 extend
		* 8. vueX的应用场景, 之前的项目都拿他做过什么, 为什么不推荐在mutations中使用异步
		* 9. 在v-for 时,(v-for 和 v-if谁的优先级高) 为什么不推荐 内部item使用 v-if 有什么优化方法吗(3.x中 v-if的优先级比v-for高)
		* 10.provide inject是做什么的 initProvide 将 _provided 赋值  (resolveInject (通过while循环), 不停找父组件有没有_provided一个属性并且找到provideKey, 如果找到就使用defineReactive响应式定义在自己身上)
		*    渲染和更新都走patch方法 遇到节点会先存入到VNodeQueue 父节点create -> created -> render -> 子组件 create -> created -> mounted -> 父组件mounted(更新,销毁也是同理)
		* 11. 逻辑抽离 mixin用过吗(mergeOptions, mergeHook), 有什么缺点, (3.x中混入只合并第一层数据) nextTick原理(eventLoop)
		* 12. vue3 有了解吗 都更新了什么 CompositionAPI, BlockTree, CacheHandlers, Fragments, Suspense, Teleport (defineProperty 和 proxy)
		* 13. css说几个居中方法
		* 14. 平时怎么做适配 rem原理 flex grid vh vw @media
		* 15. 登录流程 刷新页面各个组件都依赖当前用户信息 你会怎么存储用户信息
		* 16. find, findIndex, some, every, reduce
		* 17. js数据类型
		* 18. Promise 解决了什么问题
		* 19. ajax、fetch、axios 这三都有什么区别
		* react
		*   1.生命周期
		*   2.key是什么
		*   3.ref是什么
		npm list -g depth 0
	* */

	/**
	script标签 defer,async 属性是做什么的?
	通过js动态添加的script标签是同步还是异步的?
		答案: 是async,如果要设置成同步需要在head标签当中添加
		<link rel="preload" href="style.css" as="style">
		<link rel="preload" href="main.js" as="script">
		preload预加载,告诉浏览器先去缓存这个文件,否则阻塞主线程
	后端返回blob或者arraybuffer格式数据,怎么导出文件
	virtual dom 是什么
	反转字符串
	伪数组,类数组,怎么转化真实数组
	object == object 和 object === object 结果一样吗
	async await 是什么
	显示一个处理图片加载失败的方法
	回流和重排怎么优化？
	js为什么需要放在body(更好的回答其实是浏览器的渲染引擎和js解析引擎的冲突，当然回答js是单线程执行也没问题,如何优化)？
	操作DOM为什么是昂贵的？
	js如何执行(even Loop/宏任务、微任务，事件队列，promise, async/await)？
	js的作用域？
	闭包的理解(防抖和节流)？
	(通过一些题进行考察)，基础类型以及如何判断类型？
	事件机制以及如何实现一个事件队列？
	oop编程和原型链？
	最优的继承方式，es6 super的作用（进阶），this指向问题和new的过程(bind函数&&new函数手写)？
	js深拷贝？（JSON方法实现拷贝有什么问题？）
	*/

	/**
	vue3变化
		key 在v-if, v-else-if v-else 中不必要在声明key, 内部会自动声明
		ref 在v-for和嵌套v-for声明ref 都需要与当前实例绑定到$refs去, 降低效率
			<div v-for="item in list" :ref="setItemRef"></div>
			setup() {
				const refList = [];
				const setItemRef = el => {
					refList.push(el)
				}
				onBeforeUpdate(() => {
					refList = [];
				})
				onUpdate(() => {
					console.log(refList);
				})
			}
		$attr 2.x是不包含class和style的 3.x都将囊括在内
		移除$children 使用refs找到子组件
		移除$on, $off, $once 不在支持兄弟组件通信时使用的eventHub(const hub = new Vue(); export default hub) 改为第三个库 mitt 和 tiny-emitter
		移除过滤器 | 因为他打破了在插值中只有js表达式这种规定 增加了学习和使用成本
		函数式组件 functional 属性 在template上 移除
		dom 等级
		事件流,冒泡和穿透是什么
		array.splice 和 slice 区别
		promise 的三个状态
		for in:
				for ... in 循环返回的值都是数据结构的 键值名。
				遍历对象返回的对象的key值,遍历数组返回的数组的下标(key)。
				for ... in 循环不仅可以遍历数字键名,还会遍历原型上的值和手动添加的其他键
				对象当中属性分为 常规属性和 排序属性, 数字属性(排序属性)被最先打印出来了，并且按照数字⼤⼩的顺序打印的, 设置的字符串属性依然是按照之前的设置顺序打印
		for of:
				获取值, (for in获取键)
				一个数据结构只要部署了 Symbol.iterator 属性, 就被视为具有 iterator接口, 就可以使用 for of循环。
				提供了遍历所有数据结构的统一接口
		null == undefined
		匿名函数自调用
	*/

	/*
	为什么 0.1 + 0.2 不等于 0.3 ?
		js中数字均基于IEEE754标准的双精度64位浮点数
		64位
			第一位叫 sign bit(符号): 用来表示正负数
			下来的11位叫 exponent(指数): 用来表示次方
			在下来的52位叫 mantissa(尾数): 用来表示精确度
		也就是说一个数字的范围只能在 -(2^53 -1) 至 2^53 -1 之间
		由于浮点数用二进制数表达式是无穷的, 且最多53位, 必须截断, 进而产生误差.
	浮点型转换二进制公式 (乘 2, 取整)
	举例: 0.1 ->
	0.2=> 0, 0.4=>0, 0.8=>0, 1.6=>1, 1.2=>1, 0.2=>0 ... 无穷, 但是尾数位有限


	进制转换
	2 -> 8
	将二进制的额没三位为一组,从低位开始 转成对应的8进制数即可
	举例
	11010101 ->
	2 -> 16



	8 -> 2
	规则将8进制数的每一位转成对应的一个3位的二进制数即可
	举例
	0237 -> 10 011 111 (10011111)
	16 -> 2
	规则将16进制数的每一位转成对应的一个4位的二进制数即可
	0x237 -> 0010 0011 0111 (1000110111)
	*/
}
