# Server Side Template Injection

## Zasady

1. Podczas wykorzystywania podatności dozwolone jest wykonywanie jedynie poniższych poleceń:
   - `id`
   - `whoami`
   - `cat flag.txt`
2. Jeśli po wprowadzeniu jakichś danych strona zwraca błąd, który nie znika nawet po odświeżeniu, należy wyczyścić ciasteczka.

## Python

### Wstęp

W Pythonie jednym z najpopularniejszych silników szablonów jest Jinja2. Jest on domyślnie stosowany w bibliotece Flask służącej do tworzenia aplikacji internetowych. Chcąc wykorzystać podatność SSTI należy znaleźć sposób na wyjście ze środowiska Jinja2, w którym większość funkcji Pythona jest niedostępna. Istnieją jednak obiekty które zawsze są dostępne i od których należy zacząć eksploatację. Są to między innymi `[]`, `""` oraz `dict`.

## Go

W języku programowania Go istnieją dwie natywne biblioteki do szablonów: `text/template` oraz `html/template`. Skupimy się na tej drugiej, ponieważ interesują nas aplikacje internetowe. Kluczowym elementem `html/template` jest metoda `Execute` struktury `Template`, która jako drugi argument przyjmuje dane, które będą przesłane do szablonu.

### Wstęp

Załóżmy, że mamy tak zdefiniowaną strukturę `User`:

```go
type User struct {
	Id   int
	Name string
}

func (u User) Greet() string {
	return "Hi, " + u.Name + "!"
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

Jak widać, biblioteka `html/template` mając strukturę `User` potrafi skorzystać zarówno z pola `Id` typu `int`, jak również z metody `Greet`, która po wykonaniu zwraca `string`.

### Ćwiczenie

1. Wejdź na stronę podaną podczas zajęć przeznaczoną dla języka Go. Zobacz co się stanie, jeśli w polu `Username` wpiszesz `Bogdan`, w `Template` szablon `Hi, {{.Username}}!`, a następnie zatwierdzisz formularz.
2. Otwórz plik `main.go` w folderze `go`, aby móc minimalnie zrozumieć działanie serwera oraz znaleźć podatność SSTI.
3. Znajdź w kodzie funkcję `Execute`, nazwę zmiennej podanej jako drugi argument, typ tej zmiennej oraz definicję struktury.
4. Przyjrzyj się, co oprócz nazwy użytkownika jest jeszcze w niej zdefiniowane (nie zwracaj uwagi na pole `Template`). Spróbuj zrozumieć działanie jedynej metody. Czy jest jakaś wartość, z której ona korzysta, a która jest pod twoją kontrolą?
5. Komenda `bash -c` wykonuje skrypt podany jako argument, tak więc przykładowo `bash -c "echo test"` jest w większości przypadków równoważne z `echo test`.
6. Język bash pozwala na wykonanie kilku poleceń w jednej linii za pomocą znaku `;`, na przykład `id; whoami; cat /etc/os-release`.
7. Ustawiąjąc nazwę użytkownika na wartość wykorzystującą klasyczną podatność command injection oraz wywołując podatną metodę w kontrolowanym szablonie, odczytaj wartość flagi.
