package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var i, n, duplikat, j int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for j = 0; j < n; j++{
		for i = j+1; i < n; i++{
			if A[j] == A[i]{
				duplikat+=1
				break
			}
		}
	}
	fmt.Print(duplikat)

}