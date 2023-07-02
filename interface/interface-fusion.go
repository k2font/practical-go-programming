package main

type HumanResource interface {
	getPersonnel()
	getSalary()
	getFacility()
}

type Engineering interface {
	getSourceCode()
	getTestCode()
	getInfrastructure()
	getFacility() // HumanResourceと同じメソッドを持っていても合成する上では問題ない
}

// インターフェースの合成
// メソッドは各インターフェースごとに細かく分けて最後に合成するのがおすすめ
type Genneralist interface {
	HumanResource
	Engineering
}

func FusionTest() {
	var a Genneralist
	a.getPersonnel()
	a.getTestCode()
}
