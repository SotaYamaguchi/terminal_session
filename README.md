# terminal_session

`terminal_session` は、現在のターミナルセッションを `.terminal_session/logs` ディレクトリにタイムスタンプ付きで記録する Go プログラムです。  
セッションの開始と終了時刻がログファイルに記録され、後から確認することができます。

## 特徴

- **セッションの記録**：`script` コマンドを利用して、ターミナルで実行したコマンドがログファイルに記録されます。
- **開始時刻と終了時刻**：セッションの開始と終了時刻がログファイルに追記され、セッションのタイミングを確認できます。
- **タイムスタンプ付きのログファイル名**：ログファイル名に実行時刻が含まれており、履歴管理がしやすくなっています。

## ログファイルの保存形式

- ログファイルは、ホームディレクトリ以下の `.terminal_session/logs` に保存されます。
- ファイル名は `terminal_session_YYYYMMDDhhmm.log` 形式です（例: `terminal_session_202411132216.log`）。

## インストール手順

1. **ソースコードをクローンまたはコピー**

   このリポジトリをクローンするか、ソースコードをダウンロードします。

   ```bash
   git clone https://github.com/username/terminal_session.git
   cd terminal_session 
    ```

2.	Goを使用してコンパイル
      go build を使って、terminal_session コマンドを $HOME/.terminal_session/bin にインストールします。
      ```bash
      mkdir -p $HOME/.terminal_session/bin
      go build -o $HOME/.terminal_session/bin/terminal_session terminal_session.go
      ```

3. PATHに追加
$HOME/.terminal_session/bin が PATH に含まれているか確認します。含まれていない場合は、以下を .zshrc（または .bashrc）に追加し、シェルを再読み込みします。
    ```bash
    echo 'export PATH="$HOME/.terminal_session/bin:$PATH"' >> ~/.zshrc
    source ~/.zshrc
    ```

## 使用方法

1.	ターミナルで terminal_session を実行
ターミナルで以下のコマンドを入力すると、セッションの記録が開始されます。
      ```bash
      terminal_session
      ```

2.	セッションの終了
セッションを終了するには、Ctrl+D を押すか、exit コマンドを入力します。これにより、終了時刻がログファイルに記録されます。
 
3. ログファイルの確認
ログファイルは .terminal_session/logs フォルダ内にタイムスタンプ付きファイル名で保存されます。

## 注意事項

	•	script コマンドがシステムにインストールされている必要があります（macOSやLinuxでは通常標準でインストールされています）。
	•	.terminal_session/logs ディレクトリが存在しない場合、自動的に作成されます。