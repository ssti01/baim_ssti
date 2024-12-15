# Ćwiczenia

## Docker

Przed przystąpieniem do wykonania każdego ćwiczenia należy w odpowiednim podfolderze zbudować obraz na podstawie `Dockerfile` oraz uruchomić kontener. Obraz najlepiej oznaczyć dowolną etykietą za pomocą flagi `-t`. Przy uruchamianiu trzeba ustawić przekierowanie odpowiedniego portu wewnątrz kontenera na `localhost` za pomocą flagi `-p`.

| Ćwiczenie | Port | TAG	  |
| --------- | ---- | ------------ |
| Python    | 3333 | python--ssti |
| Node.js   | 4444 | js-ssti	  |
| Go        | 5555 | go-stti	  |

Przydatna jest również flaga `-d`, która powoduje uruchomienie kontenera w tle i wypisanie na standardowe wyjście jego identyfikatora.

```bash
docker build -t <TAG> .
docker run -d -p <PORT>:<PORT> <TAG>
```

Po wykonaniu ćwiczenia kontener można zatrzymać i usunąć podając odpowiedni identyfikator.

```bash
docker stop <ID>
docker rm <ID>
```

## Python

1. Zapoznaj się z kodem aplikacji. Znajdź funkcję `render_template_string`. To w niej najczęściej występuje podatność SSTI.
2. Sprawdź, które pola formularza są podatne.
3. Za pomocą formularza wyświetl `SECRET_KEY` znajdujący się w zmiennej `config` przechowywującej konfigurację aplikacji.
4. Wykonaj komendę `echo $FLAG` na serwerze. Podpowiedź: https://www.youtube.com/watch?v=VBifwXFQJMQ (od minuty 2:40).

## Node.js

1. Wpisz dowolny tekst w pole formularza i zatwierdź klikając na przycisk `Submit`. Zobacz wynik tego działania.
2. Wejdź na stronę [HackTricks](https://book.hacktricks.xyz/pentesting-web/ssti-server-side-template-injection) i zapoznaj się z informacją dotyczącą wykorzystania podatności SSTI w przypadku silnika szablonów Handlebars.
3. Odpowiednio zmodyfikuj znaleziony złośliwy szablon tak, aby powodował on odczyt zmiennej środowiskowej `FLAG`.

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

1. Zobacz, co się stanie, jeśli w polu `Username` wpiszesz `Bogdan`, w `Template` szablon `Hi, {{.Username}}!`, a następnie zatwierdzisz formularz.
2. Otwórz plik `main.go` w podfolderze `go`, aby móc minimalnie zrozumieć działanie serwera oraz znaleźć podatność SSTI.
3. Znajdź w kodzie wywołanie funkcji `Execute` oraz definicję struktury której instancja została podana jako drugi argument.
4. Spróbuj zrozumieć działanie metody związanej z tą strukturą. Czy jest jakaś wartość, z której ona korzysta, a która jest pod twoją kontrolą?
5. Komenda `sh -c` wykonuje polecenie podane jako argument, tak więc przykładowo `sh -c "echo test"` jest w większości przypadków równoważne z `echo test`.
6. Język bash pozwala na wykonanie kilku poleceń w jednej linii za pomocą znaku `;`, na przykład `id; whoami; cat /etc/os-release`.
7. Ustawiając nazwę użytkownika na wartość wykorzystującą klasyczną podatność command injection oraz wywołując podatną metodę w kontrolowanym szablonie, odczytaj wartość zmiennej środowiskowej `FLAG`.
