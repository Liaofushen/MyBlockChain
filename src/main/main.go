package main

import "../core"

func main() {
	blockChain := core.BlockChain{}
	blockChain.Write("Hello")
	blockChain.Write("World")
	blockChain.Read()
}
