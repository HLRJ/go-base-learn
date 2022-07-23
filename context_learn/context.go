package context_learn

import (
	"context"
	"fmt"
)

//context 基本用法
//context使用场景 withCancel()

func NewCtx() {
	bg := context.Background()
	fmt.Println(bg)

	todo := context.TODO()
	fmt.Println(todo)
}

//CtxWithValue ctx中存储值
//突出两点
//    1.ctx 不可变
//	  2.ctx WithValue  后会基于ParentContext 生成 childContext

func CtxWithValue() {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")
	ctxF := context.WithValue(ctxC, "f", "F")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)
}

//AccessParentContextValue  访问父Context的值
//程序中 Context 之间 的关系  树结构
//          ctxB
//		  /
//  ctxA
//		 \
//        ctxC
//          \
//          ctxF
//突出3点
//1.child context 可以访问 parent context
//2.分支路径上的context之间不能互相访问
//3.parent context 不能访问 child context 中的值
func AccessParentContextValue() {
	ctxA := context.Background()
	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")
	ctxF := context.WithValue(ctxC, "f", "F")

	//防止报错，调用一下ctxB
	fmt.Println(ctxB)
	//当前context 访问 self context内的值
	fmt.Println(ctxF.Value("f"))
	//1.child context 可以访问 parent context
	fmt.Println(ctxF.Value("c"))
	//2.分支路径上的context之间不能互相访问
	fmt.Println(ctxF.Value("b"))
	//3.parent context 不能访问 child context 中的值
	fmt.Println(ctxA.Value("b"))
}

//WithCancelContext  演示with
