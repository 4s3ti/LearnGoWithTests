package main

import "fmt"

const (
  french = "French"
  spanish = "Spanish"
  portuguese = "Portuguese"  

  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  frenchHelloPrefix = "Bonjour, "
  portugueseHelloPrefix = "Olá, "
)  

func Hello(name string, language string) string {
  if name == "" { 
    name = "World"
  }
  
  return greetingPrefix(language) + name

}


func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
      prefix = frenchHelloPrefix
  case spanish:
      prefix = spanishHelloPrefix
  case portuguese:
      prefix = portugueseHelloPrefix
  default:
      prefix = englishHelloPrefix
  }
    
  return
}

func main() {
	fmt.Println(Hello("world", ""))
}
