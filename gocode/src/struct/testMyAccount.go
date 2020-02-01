package main

import "fmt"

func main()  {

	input := ""
	flag := false
	moneyFlag := false

	money := 0.00 //每次收支的金额
	explain := "" //说明
	balance := 10000.00 //原始总金额

	detail := "收支\t账号金额\t收支金额\t说     明\n"

	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	//details := "收支\t账户金额\t收支金额\t说    明"

	for{
		fmt.Println("----------------家庭收支记账软件---------------")
		fmt.Println("1	收支明细")
		fmt.Println("2	登记收入")
		fmt.Println("3	登记支出")
		fmt.Println("4	退出软件")

		fmt.Printf("请选择（1-4）：")


		fmt.Scanln(&input)

		switch input {
			case "1":
				fmt.Printf("-----------------当前收支明细记录-----------------\n\n")

				if moneyFlag {
					fmt.Println(detail)
				} else {
					fmt.Println("当前没有收支明细... 来一笔吧!")
				}

			case "2":

				fmt.Printf("本次收入金额：")
				fmt.Scanln(&money)
				fmt.Printf("本次收入说明：")
				fmt.Scanln(&explain)

				balance += money

				detail += fmt.Sprintf("\n收入\t%v\t     %v\t         %v \n", balance, money, explain)

				moneyFlag = true
			case "3":
				fmt.Printf("本次支出金额：")
				fmt.Scanln(&money)

				//这里需要做一个必要的判断
				if money > balance {
					fmt.Println("余额的金额不足")
					break
				}

				fmt.Printf("本次支出说明：")
				fmt.Scanln(&explain)

				balance -= money

				detail += fmt.Sprintf("\n支出\t%v\t     %v\t                 %v \n", balance, money, explain)
				moneyFlag = true
			case "4":
				quit := ""
				fmt.Println("你确定要退出吗？ y/n")
				fmt.Scanln(&quit)

				for {
					if quit == "y" || quit == "n" {
						if quit == "y" {
							flag = true
						}
						break
					}else {
						fmt.Println("输入有误，请重新输入 y/n")
						fmt.Scanln(&quit)
					}
				}

			default:
				fmt.Println("请输入正确的选项..")
		}

		if flag {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")

}
