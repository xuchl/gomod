package gomod
import "fmt"

func Say(s string) string{
    return fmt.Sprintf("hello,%s",s)
}

func nosay(s string) string{
    return fmt.Sprintf("say,%s",s)
}
