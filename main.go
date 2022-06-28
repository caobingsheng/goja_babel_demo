package main

import (
	"fmt"

	babel "github.com/jvatic/goja-babel"
)

func main() {
	babel.Init(4)
	res, err := babel.TransformString(`
	const patientName = '李四';
	function sleep(t) {
		return new Promise((resolve,reject) => {
			const id = setTimeout(()=>{
				resolve()
				clearTimeout(id);
			},t);
		});
	}
	(async (n) => {
		console.log(n);
		await sleep(10000)
		console.log(n);
	})(patientName);
	`, map[string]interface{}{
		"plugins": []string{
			"transform-async-to-generator",
			"transform-exponentiation-operator",
		}})
	if err != nil {
		fmt.Println("转换出错", err)
		return
	}
	fmt.Println(res)
}
