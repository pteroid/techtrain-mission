/*
 * TechTrain MISSION Game API
 *
 * TechTrain MISSION ゲームAPI入門仕様  まずはこのAPI仕様に沿って機能を実装してみましょう。    この画面の各APIの「Try it out」->「Execute」を利用することで  ローカル環境で起動中のAPIにAPIリクエストをすることができます。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "techtrain-mission/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(router.Run(":8080"))
}