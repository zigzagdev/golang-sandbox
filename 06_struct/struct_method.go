package 0_6_struct

package main

// ----------------------------------------------------------------
// PHPとの比較
// ----------------------------------------------------------------

// PHP：クラスの中にメソッドを書く
// class CreateTodoUseCase {
//     private TodoRepository $repo;
//     public function __construct(TodoRepository $repo) {
//         $this->repo = $repo;
//     }
//     public function execute(TodoRequest $request): void {}
// }

// Go：structの外でメソッドを定義する
// ① structでフィールドを定義
type CreateTodoUseCase struct {
	repo TodoRepository  // interfaceに依存（DIP）
}

// ② New関数がPHPのコンストラクタに相当
func NewCreateTodoUseCase(repo TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{repo: repo}
}

// ③ ポインタレシーバでメソッドを定義
// (u *CreateTodoUseCase) がPHPの $this に相当
func (u *CreateTodoUseCase) Execute(request TodoRequest) error {
	// ビジネスロジック
	return nil
}

// ----------------------------------------------------------------
// メモ
// ----------------------------------------------------------------

// PHPと違い、Goはstructの外でメソッドを定義する
// New関数 → PHPのコンストラクタに相当
// (u *CreateTodoUseCase) → PHPの $this に相当
// *（ポインタレシーバ）→ 同じメモリを参照するため必要