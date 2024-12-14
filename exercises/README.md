# Ćwiczenia

## Zasady

## Python

W Pythonie jednym z najpopularniejszych silników szablonów jest Jinja. Jest on domyślnie stosowany w bibliotece Flask służącej do tworzenia aplikacji internetowych. Chcąc wykorzystać podatność SSTI należy znaleźć sposób na wyjście ze środowiska Jinja, w którym większość funkcji Pythona jest niedostępna. Istnieją jednak obiekty które zawsze są dostępne i od których należy zacząć eksploatację. Są to między innymi `""`, `request` oraz `dict`.

### Docker

```bash
docker build -t python-ssti .
docker run -p 3333:3333 python-ssti
```

### Ćwiczenie

1. Zapoznaj się z kodem aplikacji. Znajdź funkcję `render_template_string`. To w niej najczęściej występuje podatność na SSTI.
2. Sprawdź, które pola formularza są podatne na ten atak.
3. Za pomocą formularza wyświetl `SECRET_KEY` znajdujący się w zmiennej `config` przechowywującej konfigurację aplikacji.
4. Wykonaj komendę `echo $FLAG` na serwerze. Podpowiedź: https://www.youtube.com/watch?v=VBifwXFQJMQ (od minuty 2:40).

## JavaScript

### Docker

```bash
docker build -t js-ssti .
docker run -p 4444:4444 js-ssti
```

### Ćwiczenie

1. Wejdź na stronę serwera JavaScript podaną na zajęciach i spróbuj wpisać dowolny tekst w pole formy i zatwierdź klikając na przycisk `Submit`, zobacz wynik tego działania.
2. Wejdź na stronę [HackTricks](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection) i zapoznaj się z informacją dotyczącą ataku SSTI w szablonie Handlebars.
3. Odpowiednią zmodyfikuj kod używany w tym ataku tak, aby dostać się do flagi. Podpowiedź: trzeba odczytać zmienną środowiskową `FLAG`.

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

### Docker

```bash
docker build -t go-ssti .
docker run -p 5555:5555 go-ssti
```

### Ćwiczenie

1. Zobacz, co się stanie, jeśli w polu `Username` wpiszesz `Bogdan`, w `Template` szablon `Hi, {{.Username}}!`, a następnie zatwierdzisz formularz.
2. Otwórz plik `main.go` w podfolderze `go`, aby móc minimalnie zrozumieć działanie serwera oraz znaleźć podatność SSTI.
3. Znajdź w kodzie wywołanie funkcji `Execute` oraz definicję struktury której instancja została podana jako drugi argument.
4. Spróbuj zrozumieć działanie metody związanej z tą strukturą. Czy jest jakaś wartość, z której ona korzysta, a która jest pod twoją kontrolą?
5. Komenda `sh -c` wykonuje polecenie podane jako argument, tak więc przykładowo `sh -c "echo test"` jest w większości przypadków równoważne z `echo test`.
6. Język bash pozwala na wykonanie kilku poleceń w jednej linii za pomocą znaku `;`, na przykład `id; whoami; cat /etc/os-release`.
7. Ustawiając nazwę użytkownika na wartość wykorzystującą klasyczną podatność command injection oraz wywołując podatną metodę w kontrolowanym szablonie, odczytaj wartość zmiennej środowiskowej `FLAG`.
