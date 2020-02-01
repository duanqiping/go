package utils

import "fmt"

type familyAccount struct {
	input string //用户的输入
	flag bool	//退出程序标志
	moneyFlag bool //是否有收支标志

	money float64 //每次收支的金额
	explain string  //收支说明
	balance float64  //原始总金额

	detail string  //收支表头
}

func NewFamilyAccount() *familyAccount  {
	return &familyAccount{
		input:"",
		flag:false,
		moneyFlag:false,
		money:0.00,
		explain:"",
		balance:10000.00,
		detail:"",
	}
}

//主菜单
func (f *familyAccount) MainMenu()  {
	for{
		fmt.Println("----------------家庭收支记账软件---------------")
		fmt.Println("1	收支明细")
		fmt.Println("2	登记收入")
		fmt.Println("3	登记支出")
		fmt.Println("4	退出软件")

		fmt.Printf("请选择（1-4）：")


		fmt.Scanln(&f.input)

		switch f.input {
		case "1":
			f.moneyDetail()

		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()

		default:
			fmt.Println("请输入正确的选项..")
		}

		if f.flag {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}

//明细
func (f *familyAccount) moneyDetail()  {
	fmt.Printf("-----------------当前收支明细记录-----------------\n\n")

	if f.moneyFlag {
		fmt.Println(f.detail)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

//收入
func (f *familyAccount) income()  {
	fmt.Printf("本次收入金额：")
	fmt.Scanln(&f.money)
	fmt.Printf("本次收入说明：")
	fmt.Scanln(&f.explain)

	f.balance += f.money

	f.detail += fmt.Sprintf("\n收入\t%v\t     %v\t         %v \n", f.balance, f.money, f.explain)

	f.moneyFlag = true
}

//支出
func (f *familyAccount) pay()  {
	fmt.Printf("本次支出金额：")
	fmt.Scanln(&f.money)

	//这里需要做一个必要的判断
	if f.money > f.balance {
		fmt.Println("余额的金额不足")
		return
	}

	fmt.Printf("本次支出说明：")
	fmt.Scanln(&f.explain)

	f.balance -= f.money

	f.detail += fmt.Sprintf("\n支出\t%v\t     %v\t                 %v \n", f.balance, f.money, f.explain)
	f.moneyFlag = true
}

//退出
func (f *familyAccount) exit() {
	quit := ""
	fmt.Println("你确定要退出吗？ y/n")
	fmt.Scanln(&quit)

	for {
		if quit == "y" || quit == "n" {
			if quit == "y" {
				f.flag = true
			}
			break
		}else {
			fmt.Println("输入有误，请重新输入 y/n")
			fmt.Scanln(&quit)
		}
	}
}