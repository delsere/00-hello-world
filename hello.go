// Ogni programma Go è composto di packages. 
// I programmi cominciano l'esecuzione nel package main. 
package main

// Questo programma sta utilizzando i packages tramite percorsi d'importazione "fmt", "time" e "math/rand". 
// Puoi anche scrivere multiple importazioni (Ma è buona norma usare importazioni fattorizzate.), come: 
// import "fmt"
// import "math"
// import ...
import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
	"runtime"
)

// L'istruzione "var" dichiara una lista di variabili; cosí come per le liste di argomenti di una funzione, il tipo è alla fine. 
// Un'istruzione var può essere a livello di package o.. (1) 
var c, python, java bool
// Una dichiarazione var può includere inizializzazioni, uno per variabile. 
var i, j int = 1, 2

// L'esempio mostra variabili di differenti tipi, e anche che la dichiarazione di variabili può essere "fattorizzata" in blocchi, cosí come per le istruzioni 
// di import. I tipi int, uint, and uintptr sono grandi normalmente 32 bit su sistemi da 32-bit e 64 bit su sistemi da 64-bit. 
// Quando hai bisogno di un valore intero dovresti usare int a meno che tu non abbia particolari necessità di usare tipi di interi di dimensioni diverse o unsigned. 
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Le costanti sono dichiarate come variabili, ma con la parola chiave "const"; possono essere carattere, stringa, booleano, o valori numerici. 
// Non possono essere dichiarate con la sintassi :=.
const Pi = 3.14

// Costanti numeriche sono valori ad alta precisione. Una costante non tipizzata prende il tipo richiesto dal suo contesto. 
// Prova anche a stampare a video needInt(Big). 
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	
	fmt.Println("fmt.Println stampa una riga")
	
	fmt.Println("Data e ora:", time.Now())
	
	// Nota: l'ambiente in cui questi programmi sono eseguiti è deterministico, quindi rand.Intn ritornerá sempre lo stesso numero. 
	// (Per vedere numeri diversi, inizializza il generatore di numeri; guarda rand.Seed.) 
	fmt.Println("La funzione rand.Intn(x int) restituisce un numero intero da 0 a x.", "Nell'esempio seguente x=100", "E rand.Intn(100) = ", rand.Intn(100))

	fmt.Printf("La radice quadrata di 7 è %g\n", math.Sqrt(7))
	
	// Dopo aver importato un package ci si può riferire ai nomi che esso esporta (Inizia sempre con Maiuscola). 
	fmt.Printf("Il P Greco è %g\n", math.Pi)
	fmt.Println("Funzione add(42, 13) = ", add(42, 13))
	
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	fmt.Println(split(17))
	
	// (1) ... o di funzione. 
	var i int
	fmt.Println(i, c, python, java)
	
	// Nella dichiarazione successiva il tipo può essere omesso; la variabile prenderá il tipo dell'inizializzazione. 
	// var c, python, java = true, false, "no!"
	c, python, java := true, false, "no!"
	
	fmt.Println(i, j, c, python, java)
	
	// All'interno di una funzione, l'istruzione breve di assegnamento := può essere usata al posto dell'istruzione var con tipo implicito. 
    // Al di fuori di una funzione, ogni istruzione comincia con una parola chiave (var, func, e cosí via) e quindi il costrutto := non è disponibile. 
	k := 3

	fmt.Println(i, k, c, python, java)
	
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	
	zeroValues()
	
	conversioneTipiSempreEsplicita()
	
	interferenzaDiTipo()
	
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	
	sommaNumeriDa0A10()
	
	forSenzaIstruzioniPreEPost()
	
	fmt.Println(sqrt(2), sqrt(-4))
	
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	
	whenIsFriday()
	
	buongiornoBuonasera()
	
	funzioniDifferite()
	
	deferSovrapposti()
	
	fmt.Println("Mario")
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func zeroValues() {
	// Variabili dichiarate senza un valore iniziale esplicito vengono inizializzate con il loro valore zero. 
	// Il valore zero è: 0 per tipi numerici, false per il tipo booleano, e "" (stringa vuota) per le stringhe.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}


//Una funzione può avere zero o piú argomenti. In questo esempio, add ha due parametri di tipo int e restituisce un int
// Quando due o piú variabili di una funzione, posizionate consecutivamente, condividono il tipo, si può omettere il tipo su tutte le variabili tranne l'ultima. 
// In questo esempio, riduciamo "x int, y int" a "x, y int"
func add(x, y int) int {
	return x + y
}

// Una funzione può ritornare un qualsiasi numero di risultati. La funzione swap ritorna due stringhe. 
func swap(x, y string) (string, string) {
	return y, x
}

// I valori di ritorno in Go possono essere nominati ed agire come vere e proprie variabili. 
// Questi nomi dovrebbero essere usati per chiarire il significato dei valori di ritorno. 
// Un'istruzione return senza argomenti ritorna i valori correnti dei risultati. Ció è conosciuto come ritorno "naked" (nudo). 
// Ritorni "nudi" dovrebbero essere usati solo in piccole funzioni, come l'esempio qui mostrato. Esse possono nuocere alla leggibilitá in funzioni piú lunghe. 
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Diversamente da C, in Go l'assegnamento tra elementi di tipo differente richiede una conversione esplicita. 
func conversioneTipiSempreEsplicita() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func interferenzaDiTipo() {
	v := 42 // Provando a cambiare il valore iniziale di v nel codice cambia il tipo.
	fmt.Printf("v is of type %T\n", v)
}

// TIPI:
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte - alias per uint8
// rune - alias per int32 - rappresenta un "code point" Unicode
// float32 float64
// complex64 complex128


//Ciclo FOR 
func sommaNumeriDa0A10() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// Come in C o Java, puoi lasciare le istruzioni pre e post vuote diventandon di fatto un WHILE
// Se tralasci le condizioni del ciclo, si entrerà in un ciclo infinito, che quindi è espresso in modo molto compatto. 
//	for {
//	}
func forSenzaIstruzioniPreEPost() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Come for, la condizione if può cominciare con una piccola istruzione da eseguire prima della condizione. 
// Variabili dichiarate nell'istruzione rimangono nello scope solo fino alla fine dell'`if` 
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func whenIsFriday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func buongiornoBuonasera() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Buongiorno!")
	case t.Hour() < 17:
		fmt.Println("Buon pomeriggio.")
	default:
		fmt.Println("Buona sera.")
	}
}

func funzioniDifferite() {
	// Una istruzione di defer postpone l'esecuzione di una funzione fino a che la funzione associata ritorna. 
	// Gli argomenti della chiamata differita sono valutati immediatamente, ma la chiamata alla funzione non viene eseguita 
	// fino a che la funzione associata non ritorna.
	defer fmt.Println("world")

	fmt.Println("hello")
}

// Chiamate di funzioni differite sono inseriti in uno stack. Quando una funzione ritorna, le sue chiamate differite sono eseguite 
// nell'ordine last-in-first-out (LIFO) (ultimo entrato è il primo ad uscire) 
func deferSovrapposti() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
//Prova mario