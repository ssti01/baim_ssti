# Prezentacja

## Konspekt

1. Czym są szablony?
   - **Definicja**: Szablony to struktury pozwalające na tworzenie dokumentów lub innych formatów na podstawie wzorca i dostarczonych danych.
   - **Nie tylko HTML**: Mogą być wykorzystywane do generowania plików w różnych formatach, np. PDF, e-maili, JSON, XML itp.
   - **Rozszerzona składnia**: Szablony zawierają dodatkowe znaczniki lub konstrukcje pozwalające na dynamiczne generowanie treści.
   - **Malejąca popularność**: Wraz z rosnącą popularnością frameworków renderujących dane po stronie klienta, np. React, Vue, rola klasycznych szablonów serwerowych zmniejsza się.
   - **Elastyczność**: Dają programistom możliwość dynamicznego tworzenia treści na podstawie danych wejściowych, co jest niezbędne w wielu aplikacjach.
2. Czym są silniki szablonów?
   - **Definicja**: Narzędzia lub biblioteki przetwarzające szablony w celu wygenerowania ostatecznego dokumentu, często na podstawie dostarczonych danych.
   - **Własna składnia**: Każdy silnik wprowadza swój własny zestaw konstrukcji, jak pętle, warunki czy zmienne, np. Handlebars, EJS, Pug.
   - **Obsługa danych**: Silniki umożliwiają przekazywanie danych do szablonów i ich wykorzystanie w czasie renderowania.
   - **Bezpieczeństwo**: Same silniki szablonów nie są podatne na ataki – podatności pojawiają się w wyniku błędów w implementacji aplikacji.
   - **Popularność**: W przeszłości szeroko stosowane w aplikacjach serwerowych, dziś często współpracują z aplikacjami renderowanymi po stronie klienta.
3. Czym jest podatność SSTI?
   - **Definicja**: Server-Side Template Injection (SSTI) to podatność umożliwiająca atakującemu wprowadzenie własnej logiki do szablonu przetwarzanego po stronie serwera.
   - **Mechanizm**: Wynika z braku odpowiedniej walidacji lub filtrowania danych wejściowych, które zostają włączone do szablonu.
   - **Skutki**: Atakujący może uzyskać dostęp do danych serwera, wykonywać nieautoryzowane operacje, a nawet uruchamiać dowolny kod na serwerze.
   - **Przykłady**: Możliwość wstrzyknięcia wyrażeń w silniku Handlebars ({{this.constructor}}) lub wykonania poleceń w Jinja2 (Python).
4. Jak powstaje podatność SSTI?
   - **Kontrola treści szablonu**: Gdy użytkownik ma bezpośredni wpływ na treść szablonu lub na dane w nim używane.
   - **Brak separacji danych i logiki**: Silniki szablonów mogą umożliwiać definiowanie funkcji, wyrażeń czy wykonywanie kodu – co samo w sobie nie jest złe, ale wymaga odpowiedniej ochrony.
   - **Brak walidacji**: Dane wejściowe nie są sprawdzane, a następnie są bezpośrednio włączane do szablonu, umożliwiając wstrzyknięcie niebezpiecznej logiki.
   - **Przypadki szczególne**: Użycie eval() lub analogicznych funkcji w kodzie aplikacji, które generują szablon dynamicznie na podstawie wejścia.
5. Jak można wykorzystać podatność SSTI?
   - **Wstrzykiwanie kodu**: Atakujący może wprowadzić kod lub wyrażenia, które będą wykonywane po stronie serwera.
   - **Pozyskanie danych**: Możliwość wykradania poufnych informacji, takich jak dane użytkowników, konfiguracje serwera czy zmienne środowiskowe.
   - **Eksfiltracja informacji**: Wykorzystanie podatności do przekazania danych poza serwer, np. przez wykonywanie żądań HTTP.
   - **Utworzenie backdoora**: W niektórych przypadkach możliwe jest wprowadzenie trwałych zmian w aplikacji lub plikach serwera.
   - **Podwyższenie uprawnień**: Wykonanie kodu jako inny użytkownik lub uzyskanie dostępu do systemowych zasobów.
6. Jak wykryć podatność SSTI?
   - **Testy penetracyjne**: Próba wstrzyknięcia typowych wyrażeń (np. {{7*7}}, {% print('test') %}) w polach wejściowych.
   - **Analiza kodu**: Przegląd miejsc, gdzie dane wejściowe są przetwarzane bezpośrednio w szablonach.
   - **Automatyczne skanery**: Narzędzia takie jak Burp Suite czy OWASP ZAP, które automatycznie identyfikują podatności.
   - **Monitorowanie zachowania aplikacji**: Nieoczekiwane wyniki renderowania szablonu mogą wskazywać na podatność.
7. Jak prawidłowo używać silnika szablonów?
   - **Minimalizacja funkcjonalności**: Wyłączenie funkcji, które nie są niezbędne, np. możliwości wykonywania kodu.
   - **Separacja danych i logiki**: W miarę możliwości przechowywanie logiki poza szablonami, a w szablonach jedynie prezentowanie danych.
   - **Stosowanie escape**: Automatyczne lub ręczne stosowanie odpowiednich metod escaping dla wyjścia HTML, JSON itp.
   - **Korzystanie z bibliotek wysokiego poziomu**: Frameworki oferujące wbudowane mechanizmy ochrony (np. Django, Flask).
8. Jak poprawnie zabezpieczać dane wejściowe?
   - **Walidacja danych**: Upewnienie się, że dane wejściowe spełniają oczekiwane kryteria (np. format, długość).
   - **Filtrowanie**: Usuwanie niebezpiecznych konstrukcji (np. znaków specjalnych).
   - **Ograniczanie wejść użytkownika**: Jeśli to możliwe, użycie predefiniowanych wartości zamiast pozwalania użytkownikowi na pełną kontrolę nad treścią.
   - **Wykorzystanie API**: Zamiast dynamicznego renderowania szablonów, przesyłanie danych do API i renderowanie po stronie klienta.
9. Jakie można wyciągnąć wnioski?
   - **Świadomość ryzyka**: Programiści muszą być świadomi potencjalnych zagrożeń związanych z obsługą danych użytkownika, szczególnie przy korzystaniu z silników szablonów.
   - **Bezpieczeństwo priorytetem**: Podatności takie jak SSTI wynikają z błędów w przetwarzaniu danych wejściowych. Odpowiednie walidacje i filtrowanie danych są kluczowe dla bezpieczeństwa aplikacji.
   - **Zasada ograniczonego zaufania**: Dane użytkownika nigdy nie powinny być bezpośrednio wykorzystywane do generowania szablonów, bez odpowiedniego zabezpieczenia.
   - **Ograniczanie uprawnień**: Warto ograniczać możliwości silnika szablonów, np. poprzez wyłączenie możliwości wykonywania kodu lub korzystania z funkcji, które nie są niezbędne w projekcie.
   - **Edukacja i testowanie**: Programiści powinni być edukowani w zakresie typowych podatności, a aplikacje powinny być regularnie testowane pod kątem bezpieczeństwa, np. za pomocą narzędzi do testów penetracyjnych.
   - **Wybór narzędzi**: Korzystanie z nowoczesnych rozwiązań, które minimalizują ryzyko podatności, np. frameworków z wbudowanymi mechanizmami ochrony, może znacznie zwiększyć bezpieczeństwo projektu.
   - **Przemyślane projektowanie aplikacji**: Renderowanie po stronie klienta (np. z wykorzystaniem frameworków takich jak React czy Angular) może w wielu przypadkach zastąpić tradycyjne szablony serwerowe, zmniejszając ryzyko związane z SSTI.
   - **Odpowiedzialne podejście do danych wejściowych**: Implementacja zasad takich jak walidacja, escaping czy stosowanie bezpiecznych bibliotek to fundament budowania odpornych na ataki aplikacji.
   - **Ulepszanie standardów**: Dbałość o dokumentację oraz dzielenie się wiedzą na temat podatności i ich zapobiegania w zespołach programistycznych może przyczynić się do tworzenia bezpieczniejszych aplikacji.

## Pytania

1. Który z poniższych sposobów działania aplikacji przedstawia podatność SSTI?
   - Niezabezpieczone dane wejściowe są łączone z zapytaniem do bazy danych
   - Niezabezpieczone dane wejściowe są przetwarzane przez silnik szablonów jako część szablonu
   - Pliki załadowane przez użytkownika są zapisywane na serwerze bez ograniczeń co do ich typu lub zawartości
2. Która z poniższych podatności nie jest bezpośrednio ani pośrednio związana z SSTI?
   - RCE
   - LFI
   - XSS
3. Które z poniższych zachowań wskazuje, że aplikacja może być podatna na SSTI?
   - Pojawianie się wiadomości o błędach przy podaniu różnorodnych znaków specjalnych jako dane wejściowe
   - Pojawianie się alertu przy podaniu `<script>alert(1)</script>` jako dane wejściowe
   - Wysyłanie w odpowiedzi HTTP nagłówka `X-Powered-By: Express`
4. Dlaczego SSTI często występuje w aplikacjach korzystających z szablonów?
   - Silniki szablonów przetwarzają dynamiczne dane wejściowe, co może prowadzić do nieautoryzowanego wykonania kodu
   - Szablony są zaprojektowane wyłącznie do generowania statycznych treści
   - Dane wejściowe w szablonach nie są automatycznie zabezpieczone przez serwer
5. Który z przykładów nie daje zabiezpieczenia przed atakiem SSTI
   - Sanityzacja danych wejściowych
   - Ograniczenie na wprowadzenie danych przez użytkownika
   - Zbieranie logów danych wejściowych

## Fragmenty kodu

### Jinja

#### Szablon

```html
<h1>Hello, {{ name }}!</h1>
  {% for item in items %}
  <li>{{ item }}</li>
  {% endfor %}
</ul>
```

#### Serwer

```py
from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def home():
    data = {
        "name": "Bob",
        "items": ["Pepper", "Turmeric", "Ginger", "<script>alert(1);</script>"],
    }
    return render_template("index.html", **data)

if __name__ == "__main__":
    app.run(port=3000)
```

#### Wynik

```html
<h1>Hello, Bob!</h1>
<ul>
  <li>Pepper</li>
  <li>Turmeric</li>
  <li>Ginger</li>
  <li>&lt;script&gt;alert(1);&lt;/script&gt;</li>
</ul>
```

#### SSTI

##### Serwer

```py
from flask import Flask, request, render_template_string

app = Flask(__name__)

@app.route("/ssti", methods=["GET"])
def ssti():
    template = request.args.get("template", "")
    return render_template_string(template)

if __name__ == "__main__":
    app.run(port=3000)
```

##### Wykorzystanie

```py
{{ "".join("A" * 10**9) }}
{{ config["DATABASE_PASSWORD"] }}
{{ request.__class__.__mro__[1].__subclasses__()[40]("/etc/passwd").read() }}
```

##### Słaba sanityzacja

```py
template = template.replace("__", "")
template = template.replace("[", "").replace("]", "")
```

##### Omijanie filtrów

```py
{{ request["\x5f\x5fclass\x5f\x5f"] }}
{{ request|attr("__class__") }}
{{ "<script>alert(1);</script>"|safe }}
```

### EJS

#### Szablon

```html
<h1>Hello, <%= name %>!</h1>
<ul>
  <% items.forEach(item => { %>
  <li><%= item %></li>
  <% }); %>
</ul>
```

#### Serwer

```js
import express from "express";

const app = express();

app.set("view engine", "ejs");

app.get("/", (_, res) => {
  const data = {
    name: "Alice",
    items: ["Apples", "Bananas", "Cherries", "console.log(process.pid);"],
  };
  res.render("index", data);
});

app.listen(3000);
```

#### Wynik

```html
<h1>Hello, Alice!</h1>
<ul>
  <li>Apples</li>
  <li>Bananas</li>
  <li>Cherries</li>
  <li>console.log(process.pid);</li>
</ul>
```

#### SSTI

##### Serwer

```js
import { exec } from "node:child_process";
import express from "express";
import ejs from "ejs";

const app = express();

const execute = (command) => {
  exec(command);
};

const allocate = (size) => {
  Buffer.alloc(size, "A");
};

app.get("/", (req, res) => {
  res.send(ejs.render(req.query.template ?? "", { execute, allocate }));
});

app.listen(3000);
```

##### Wykorzystanie

```js
<%= execute("rm -rf /") %>
<%= allocate(10 ** 9) %>
```
