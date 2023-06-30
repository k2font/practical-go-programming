# Goのパッケージとモジュール
## メモ
- Goのパッケージの最低限のルール
  - 同一フォルダ内のファイルは同一のパッケージとなる
  - 実行ファイルのエントリーポイントはmainパッケージとなる
- 実行ファイルを含むモジュールの場合は、`/cmd` ディレクトリを作成し、そこにパッケージ名と同名のディレクトリを切る
  - 現にこのような構成を持つモジュールは多い
  - 例: https://github.com/Songmu/gocredits
    - 個人で作る規模のOSSとしては、パッケージ構成がめっちゃ参考になる。
- ちなみに、小規模のリポジトリだったら無理してモジュール分割しなくて良い
  - その場合は `main` packageにして、 `$ go run .`で全ファイルbuildすれば実行できる
    - 同じパッケージにファイルがあれば、 `import` も不要
  - Goはよりシンプルに書けるプログラミング言語なので、小規模ならばシンプルな構成にしよう。