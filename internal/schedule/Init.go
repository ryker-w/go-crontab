package schedule

// Sch 全局调度器
var Sch Schedule


func New() {
	Sch = &schedule{
	}
	Sch.Init()
}
