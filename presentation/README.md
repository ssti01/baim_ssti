# Prezentacja

## Konspekt

1. Czym są szablony?
   - Nie koniecznie muszą dotyczyć HTML.
   - Posiadają rozszerzoną składnię względem docelowego formatu.
   - Stają się coraz mniej popularne ze względu na renderowanie po stronie klienta.
   - ...
2. Czym są silniki szablonów?
   - Obsługują przetwarzanie szablonów z danymi.
   - Definiują własną składnię szablonu.
   - Nie są z natury podatne na żadne ataki.
   - ...
3. Czym jest podatność SSTI?
   - Występuje, gdy dane wejściowe użytkownika są niebezpieczne dołączone do szablonu.
   - Pozwala atakującym na wprowadzenie logiki która będzie wykonywać się po stronie serwera.
   - ...
4. Jak powstaje podatność SSTI?
   - Programista pozwala użytkownikowi kontrolować treść szablonu, który będzie przetwarzany.
   - Silniki szablonów pozwalają na zdefiniowanie w szablonie procedur prowadzących m. in. do wykonania kodu, co niekoniecznie jest złe same w sobie.
   - ...
5. Jak można wykorzystać podatność SSTI?
   - ...
6. Jak wykryć podatność SSTI?
   - ...
7. Jak prawidłowo używać silnika szablonów?
   - ...
8. Jak poprawnie zabezpieczać dane wejściowe?
   - ...
9. Jakie można wyciągnąć wnioski?
   - ...

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
4. Który z poniższych sposobów zabezpiecza aplikację przed podatnością SSTI?
   - Walidacja danych wejściowych z użyciem białej listy dopuszczalnych znaków
   - Wyłączenie obsługi zmiennych dynamicznych w szablonach
   - Szyfrowanie danych przesyłanych między serwerem a użytkownikiem
5. Jaki mechanizm najlepiej wykrywa potencjalne podatności SSTI w aplikacji?
   - Testy penetracyjne z użyciem automatycznych narzędzi
   - Logowanie i monitorowanie aktywności użytkownika
   - Stosowanie statycznej analizy kodu źródłowego
   
