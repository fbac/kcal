package common

func CalculateMacro(kcal float32, leanmass float32) (float32, float32, float32){
        var fat, prot, ch float32

        fat = (kcal * 0.25)/9
        prot = leanmass * 2.2
        ch = (kcal - (fat * 9) - (prot * 4))/4

        return fat, prot, ch
}

func IsFloat(val float32) bool {
	return val == float32(int(val))
}
