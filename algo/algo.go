package algo

func Calc(I_obl float64, temp float64) (I_ost float64) {
	I_ost = I_obl / (temp * 0.85)
	return
}
