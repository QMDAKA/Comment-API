# postprime-interviews

## desgin-docs

[以下のフォルダーを確認お願いします](https://github.com/QMDAKA/comment-mock/tree/master/design-doc)
流れのイメージがうまく表示しなければ、mermail-extensionで見てください。また、pdfもあるので、確認お願いします。

## coding-challenge

### 実行手順

```
go install 
go run main.go
```
短い時間なので、docker-composeを追加していません。余裕があれば、追加します。

### API endpoint
- 前提：先にuserとpostのデータを準備しました。[post_id = 1, user_uuid = 3oxsad] などを利用お願いします
  - Header : [Authorization: 3oxsad] 
- comment作成： endpoint POST "posts/1/comments"; body: "{"content": ""}"
- comment編集： endpoint PATCH "comments/:id"; body: "{"content": ""}"
- comment削除： endpoint DELETE "comments/:id"; 

### TODO
余裕があれば、追加したいやつ
- golang-migrate導入
- JWT
- error handling詳しく導入
- recovery
