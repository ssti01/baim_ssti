# Ćwiczenia

```bash
> git clone https://github.com/ssti01/baim_ssti.git
> cd baim_ssti/exercises
baim_ssti/exercises> ls
```

## Python

```bash
baim_ssti/exercises> cd python
baim_ssti/exercises/python> docker build -t python-ssti .
baim_ssti/exercises/python> docker run -p 3333:3333 python-ssti
server is running at http://localhost:3333
```

1. Zapoznaj się z kodem źródłowym aplikacji otwierając plik `main.py`. Znajdź wywołanie funkcji `render_template_string`. To przy jej użyciu najczęściej występuje podatność SSTI.
2. Sprawdź, które pola formularza jest podatne.
3. Za pomocą formularza wyświetl `SECRET_KEY` znajdujący się w zmiennej `config` przechowywującej konfigurację aplikacji.
4. Wprowadzając odpowiednie dane wejściowe wykonaj polecenie `echo $FLAG` na serwerze. Skorzystaj z [tego filmu](https://www.youtube.com/watch?v=VBifwXFQJMQ) (od 2:40).

## Node.js

```bash
baim_ssti/exercises> cd js
baim_ssti/exercises/js> docker build -t node-ssti .
baim_ssti/exercises/js> docker run -p 4444:4444 node-ssti
server is running at http://localhost:4444
```

1. Wpisz dowolny tekst w pole formularza i zatwierdź klikając na przycisk `Submit`. Zobacz wynik tego działania.
2. Wejdź na stronę [HackTricks](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection) dotyczącą podatności SSTI i zapoznaj się z informacją dotyczącą jej wykorzystania w przypadku silnika szablonów Handlebars.
3. Przedstawiony tam złośliwy szablon zawiera kluczową linię `return require('child_process').exec('whoami');`. W naszym przypadku nie spowoduje ona wykonania żadnego kodu ze względu m. in. na inne środowisko oraz brak `require`.
4. Dowiedz się, jak w Node.js odczytuje się zmienne środowiskowe i odpowiednio zmodyfikuj dane wejściowe tak, aby zwróciły zmienną środowiskową `FLAG`.

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

Na standardowe wyjście zostanie wypisane:

```html
<h1>User 91</h1>
<p>Hi, Bogdan!</p>
```

Jak widać, biblioteka `html/template` mając strukturę `User` potrafi skorzystać zarówno z pola `Id` typu `int`, jak również z metody `Greet`, która po wykonaniu zwraca `string`.

### Ćwiczenie

```bash
baim_ssti/exercises> cd go
baim_ssti/exercises/go> docker build -t go-ssti .
baim_ssti/exercises/go> docker run -p 5555:5555 go-ssti
server is running at http://localhost:5555
```

1. Zobacz, co się stanie, jeśli w polu `Username` wpiszesz `Bogdan`, w `Template` szablon `Hi, {{.Username}}!`, a następnie zatwierdzisz formularz.
2. Otwórz plik `main.go` aby móc minimalnie zrozumieć działanie serwera oraz znaleźć podatność SSTI.
3. Znajdź w kodzie wywołanie funkcji `Execute` oraz definicję struktury, której instancja została podana jako drugi argument.
4. Spróbuj zrozumieć działanie metody związanej z tą strukturą. Czy jest jakaś wartość, z której ona korzysta, a która jest pod twoją kontrolą?
5. Komenda `sh -c` wykonuje polecenie podane jako wartość flagi `-c`, tak więc przykładowo `sh -c "echo test"` jest równoważne z `echo test`.
6. Powłoka pozwala na wykonanie kilku poleceń w jednej linii za pomocą znaku `;`, na przykład `id; whoami; cat /etc/os-release`.
7. Ustawiając nazwę użytkownika na wartość wykorzystującą klasyczną podatność command injection oraz wywołując podatną metodę w kontrolowanym szablonie, odczytaj wartość zmiennej środowiskowej `FLAG`.
