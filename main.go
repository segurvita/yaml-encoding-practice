package main

import "fmt"
import "gopkg.in/yaml.v2"

func main() {
	// 入力YAMLを定義
	yamlInput := []byte(`
swagger: '2.0'
info:
  description: これはアパートに関するAPIです。
  version: 0.0.1
  title: アパートAPI
paths:
  '/rooms/{room-id}':
    get:
      summary: 部屋情報API
      description: 指定されたroom-idの情報を返します
      parameters:
        - name: room-id
          in: path
          description: 取得したい部屋のID
          required: true
          type: integer
          format: int64
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              id:
                type: integer
                format: int64
                example: 404
              comment:
                type: string
                example: 404号室です。どこにも存在しない部屋かも。
  `)

	// YAMLをオブジェクトに変換
	var objInput interface{}
	err := yaml.Unmarshal(yamlInput, &objInput)
	if err != nil {
		fmt.Println("エラー: ", err)
	}

	// オブジェクトをYAMLに変換
	yamlOutput, err := yaml.Marshal(&objInput)
	if err != nil {
		fmt.Println("エラー: ", err)
	}

	// 標準出力
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println("# 入力YAML:")
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println(string(yamlInput), "\n")
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println("# オブジェクト:")
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println(objInput, "\n")
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println("# 出力YAML:")
	fmt.Println("# ------------------------------------------------------------")
	fmt.Println(string(yamlOutput), "\n")
}
