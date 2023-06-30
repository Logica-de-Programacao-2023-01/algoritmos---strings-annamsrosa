package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var letras string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&s)

	letras = "a,A,b,B,c,C,d,D,e,E,f,F,g,G,h,H,i,I,j,J,k,K,l,L,m,M,n,N,o,O,p,P,q,Q,r,R,s,S,t,T,u,U,v,V,w,W,x,X,y,Y,z,Z,á,Á,à,À,ã,Ã,â,Â,ä,Ä,é,É,è,È,ê,Ê,ë,Ë,í,Í,ì,Ì,î,Î,ï,Ï,ó,Ó,ò,Ò,õ,Õ,ô,Ô,ö,Ö,ú,Ú,ù,Ù,û,Û,ü,Ü,ç,Ç,ñ,Ñ,.,?,!,@,#,$,%,&,*,-,_,+,=,/,|,\\,\",',:,;,<,>,(,),[,],{,},´,`"

	if strings.ContainsAny(s, letras) {
		fmt.Print("A string possui pelo menos uma letra ")
	} else {
		fmt.Print("A string só possui números ")
	}
}
