package 0_3_types

//
//// PHPのenumに相当にするGoの型定義
//// PHPではenum Status: string{ case Todo = 'Todo'}
//
//type Status string
//
//const (
//	StatusTodo       Status = "Todo"
//	StatusInProgress Status = "In Progress"
//	StatusCompleted  Status = "Completed"
//)
//
//func (status Status) IsValid() bool  {
//	switch status {
//	case StatusTodo, StatusInProgress, StatusCompleted:
//
//		return true
//	}
//
//	return false
//}


structはオブジェクトの形をしたデータの集合
typeはそのstructで定義をしたオブジェクトの集合に名前付けをする