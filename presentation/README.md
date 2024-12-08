# Prezentacja

## Konspekt

1. Czym jest szablon?
2. Czym jest silnik szablonów?
3. Czym jest podatność SSTI?
4. Jak powstaje podatność SSTI?
5. Jak można wykorzystać podatność SSTI?
6. Jak wykryć podatność SSTI?
7. Jak prawidłowo używać silnika szablonów?
8. Jak poprawnie zabezpieczać dane wejściowe?
9. Jakie można wyciągnąć wnioski?

## Pytania

1. Który z poniższych sposobów działania aplikacji przedstawia podatność SSTI?
   A. Niezabezpieczone dane wejściowe są łączone z zapytaniem do bazy danych
   B. Niezabezpieczone dane wejściowe są przetwarzane przez silnik szablonów jako część szablonu
   C. Pliki załadowane przez użytkownika są zapisywane na serwerze bez ograniczeń co do ich typu lub zawartości
2. Która z poniższych podatności nie jest bezpośrednio ani pośrednio związana z SSTI?
   A. RCE
   B. LFI
   C. XSS
3. Które z poniższych zachowań wskazuje, że aplikacja może być podatna na SSTI?
   A. Pojawianie się wiadomości o błędach przy podaniu różnorodnych znaków specjalnych jako dane wejściowe
   B. Pojawianie się alertu przy podaniu `<script>alert(1)</script>` jako dane wejściowe
   C. Wysyłanie w odpowiedzi HTTP nagłówka `X-Powered-By: Express`
