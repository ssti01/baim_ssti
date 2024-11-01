# Server Side Template Injection

## Zasady

1. Podczas wykorzystywania podatności dozwolone jest wykonywanie jedynie poniższych poleceń:
   - `id`
   - `whoami`
   - `cat flag.txt`

## Go

W języku programowania Go istnieją dwie natywne biblioteki do szablonów: `text/template` oraz `html/template`. Skupimy się na tej drugiej, ponieważ interesują nas aplikacje internetowe. Kluczowym elementem `html/template` jest metoda `Execute` struktury `Template`, która jako drugi argument przyjmuje dane, które będą przesłane do szablonu.

### Przykład

Załóżmy, że mamy tak zdefiniowaną strukturę `User`:

```go
type User struct {
	Id   int
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hi, %s!", u.Name)
}
```

Natomiast nasz szablon `html/template` wygląda tak:

```html
<h1>User {{.Id}}</h1>
<p>{{.Greet}}</p>
```

Jeśli zmienna `t` jest instancją `Template`, to wykonując poniższy kod:

```go
t.Execute(os.Stdout, User{91, "Bogdan"})
```

Na standardowe wyjście powinno zostać wypisane:

```html
<h1>User 91</h1>
<p>Hi, Bogdan!</p>
```

Jak widać, biblioteka `html/template` mając strukturę `User` potrafi skorzystać zarówno z pola `Id`, jak również z metody `Greet`, która po wykonaniu zwraca `string`.

### Ćwiczenie

1. Wejdź na stronę podaną podczas zajęć przeznaczoną dla języka Go i zapoznaj się z jej działaniem.
2. Otwórz plik `main.go` w folderze `go`, aby móc minimalnie zrozumieć działanie serwera oraz znaleźć podatność SSTI.
3. Znajdź w kodzie funkcję `Execute`, nazwę zmiennej podanej jako drugi argument, typ tej zmiennej oraz definicję struktury.
4. Przyjrzyj się, co oprócz nazwy użytkownika jest jeszcze w niej zdefiniowane. Spróbuj zrozumieć działanie jedynej metody. Czy jest jakaś wartość, z której ona korzysta, a która jest pod twoją kontrolą?
5. Jaki znak specjalny pozwala na wykonanie kilku poleceń w jednej linii w języku Bash?
6. Ustawiąjąc nazwę użytkownika na wartość wykorzystującą podstawową podatność command injection oraz wywołując podatną metodę w kontrolowanym szablonie, odczytaj wartość flagi.
