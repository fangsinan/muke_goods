package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	G_ArmsType = 1
	F_ArmsType = 2
)

// 雇佣兵种
type arms struct {
	Name       string
	Type       int
	Price      int
	TotalBlood int
	Blood      int
	Reduce     map[string]int
}

type MyArms struct {
	// ID   int
	MyName string
	Arms   arms
}

// 几座森林
const forestLen int = 7

// 森林属性
type forest struct {
	ID           int
	monstersID   int
	monstersName string
}

var inputIsArms int
var inputArmsId int
var inputArmsName string
var inputArmsPrice int
var inputArmsIS int
var totalPrice = 1000

var armsData = map[int]arms{
	G_ArmsType: {Name: "弓箭兵", Type: G_ArmsType, Price: 100, Blood: 100, TotalBlood: 100,
		Reduce: map[string]int{"鹰妖": 20, "狼妖": 80},
	},
	F_ArmsType: {Name: "斧头兵", Type: F_ArmsType, Price: 120, Blood: 120, TotalBlood: 120,
		Reduce: map[string]int{"鹰妖": 80, "狼妖": 20},
	},
}

// 小妖
var monsters = map[int]string{1: "狼妖", 2: "鹰妖"}

// 获取初始信息
func InitMyArms() []MyArms {
	var initArms []MyArms
	for {
		fmt.Printf("当前灵石%d  \n", totalPrice)
		fmt.Printf("雇佣兵种 输入选项... 1：继续雇佣 2:完成雇佣  \n")
		fmt.Scan(&inputIsArms)
		if inputIsArms == 2 {
			if len(initArms) <= 0 {
				fmt.Printf("没有雇佣任何兵种  无法进入下一个环节 \n\n")
				continue
			}
			break
		}
		// 金额不够
		if totalPrice < 100 {
			fmt.Printf("金额不够，无法继续雇佣 \n")
			break
		}
		fmt.Printf("兵种信息如下  请输入雇佣编号：\n")
		for k, v := range armsData {
			fmt.Printf("编号:%d  %s(需%d灵石 %d血量 ", k, v.Name, v.Price, v.Blood)
			// 杀妖损耗信息
			for rk, rv := range v.Reduce {
				fmt.Printf("杀%s损耗血量:%d ", rk, rv)
			}
			fmt.Println(" )")

		}

		fmt.Scan(&inputArmsId)

		if arms, ok := armsData[inputArmsId]; ok {
			fmt.Printf("输入兵种名称: \n")
			fmt.Scan(&inputArmsName)
			// 原有雇佣兵 数量+1
			initArms = append(initArms, MyArms{
				Arms:   arms,
				MyName: inputArmsName,
			})
			if totalPrice-arms.Price < 0 {
				fmt.Printf("金额不足雇佣该兵种 滚 \n")
				continue
			}

			totalPrice -= arms.Price // 扣除金额
			fmt.Printf("已雇佣：%s,  这个兵叫 %s \n", arms.Name, inputArmsName)
		} else {
			fmt.Printf("啥也不是 滚 \n")
			continue
		}

		fmt.Printf("当前金额%d \n", totalPrice)
	}

	// 打印信息
	fmt.Printf("\n\n\n\n 完成雇佣  \n 当前金额：%d 我的雇佣信息：\n\n ", totalPrice)
	for _, v := range initArms {
		fmt.Printf("姓名:%s 兵种:%s 总血量：%d  ", v.MyName, v.Arms.Name, v.Arms.Blood)
		// 杀妖损耗信息
		for rk, rv := range v.Arms.Reduce {
			fmt.Printf("杀%s损耗血量：%d  ", rk, rv)
		}
		fmt.Println()
	}
	return initArms
}

// 初始化森林小妖
func InitMonsters() []forest {

	// 七座森林 每次出现一个妖
	var forests []forest
	var monstersName string

	randint := len(monsters)
	for i := 1; i <= forestLen; i++ {
		//将时间戳设置成种子数 生成随机妖
		rand.Seed(time.Now().UnixNano())
		//  包括 0 和 randint-1 之间的随机数 不可以是0  所以加 1
		monstersID := rand.Intn(randint) + 1
		if v, ok := monsters[monstersID]; ok {
			monstersName = v
		}
		forests = append(forests, forest{
			ID:           i,
			monstersID:   monstersID,
			monstersName: monstersName,
		})
		monstersName = ""
	}
	// 打印信息
	fmt.Printf("\n\n森林完成初始化信息：\n")
	for _, v := range forests {
		fmt.Printf("第%d个森林,怪是：%s \n", v.ID, v.monstersName)
	}

	return forests
}

// 砍怪逻辑
func Killmonster(monstersName string, myArms []MyArms) []MyArms {
	fmt.Printf("派个兵干活 叫啥名：  \n")
	for {
		fmt.Scan(&inputArmsName)
		// 校验我的

		// 校验兵在不在
		is_exis := 0
		for k, v := range myArms {
			flag := 0
			if v.MyName == inputArmsName {
				// 找到兵 赋值 1
				is_exis = 1
				for Rk, Rv := range v.Arms.Reduce {
					// 匹配 减伤
					if monstersName == Rk {
						// 正常打怪
						if v.Arms.Blood > Rv {
							myArms[k].Arms.Blood -= Rv

							flag = 1
						} else {
							// 血量不够
							flag = 2
						}
						break
					}
				}
				// 未能匹配到妖怪
				if flag == 0 {
					fmt.Printf("匹配错误 --\n")
					return myArms
				}
				// 正常完结  退出该函数
				if flag == 1 {
					if myArms[k].Arms.Blood == 0 {
						// 若血量为0 则也删除该函数
						myArms = append(myArms[:k], myArms[k+1:]...)
					}
					return myArms
				}
				if flag == 2 {
					// 兵死了  怪未打完  继续调兵战斗
					myArms = append(myArms[:k], myArms[k+1:]...)
					fmt.Printf("兵挂了 --\n")
					if len(myArms) <= 0 {
						fmt.Printf("兵力不够 错误：  \n")
						return myArms
					}
					fmt.Printf("  继续调兵\n\n")
					break
				}
				//  直接退出函数 返回 myArms
				return myArms
			}
		}
		if is_exis == 0 {
			fmt.Printf("没这个兵 重新输名\n\n")
			continue
		}
	}
}

// 补给
func supple(myArms []MyArms) []MyArms {
	fmt.Printf("剩余金额%d 剩余兵力 %v \n\n\n\n\n 是否用灵石给 战士补养 1是  2否\n", totalPrice, myArms)
	for {
		fmt.Scan(&inputArmsIS)
		if inputArmsIS == 1 {

			fmt.Printf("请输入补给战士名称... \n\n")
			fmt.Scan(&inputArmsName)
			is_exis := 0
			falg := 0
			for k, v := range myArms {
				if v.MyName == inputArmsName {
					is_exis = 1
					// 进行补给
					fmt.Printf("请输入补给灵石  1个灵石 一个血条... \n\n")
					fmt.Scan(&inputArmsPrice)
					//
					if totalPrice < inputArmsPrice {
						fmt.Printf("灵石不够...  是否继续 给战士补养  1是  2否 \n\n")
						falg = 1
						break
					}

					if (v.Arms.Blood + inputArmsPrice) > v.Arms.TotalBlood {
						fmt.Printf("补给超过该战士血条... \n\n")
						falg = 1
						break
					}

					// 进行补充
					myArms[k].Arms.Blood += inputArmsPrice
					// 总灵石减少
					totalPrice -= inputArmsPrice
					return myArms
				}
			}

			if is_exis == 0 {
				fmt.Printf("没有这个战士... 重新输入 \n\n")
				continue
			}
			if falg == 1 {
				continue
			}
		} else {
			fmt.Printf("继续... \n\n")
			return myArms
		}

	}

}
func main() {
	// 初始化我的雇佣
	myArms := InitMyArms()

	// 初始化 森林    一个森林一个妖
	forestData := InitMonsters()
	fmt.Printf("\n\n记住它  10秒后消失...")
	// 初始化完成  停止10s
	time.Sleep(time.Second * 10)
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n开始砍怪\n")
	// // 开始砍怪
	for _, v := range forestData {
		fmt.Printf("第%d个森林,怪是：%s \n", v.ID, v.monstersName)

		// 计算伤害 小怪id v.monstersID
		myArms = Killmonster(v.monstersName, myArms)

		// 砍完 补充能量
		if totalPrice <= 0 {
			fmt.Printf("灵石不够  你没资格充值.... 我继续了 \n\n")
		} else {
			myArms = supple(myArms)
		}
		fmt.Printf("第%d个森林过了,剩余金额%d 剩余兵力 %v \n\n\n\n\n \n", v.ID, totalPrice, myArms)
	}
	fmt.Printf("齐活！！")
}
