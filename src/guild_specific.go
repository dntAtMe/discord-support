package main

import (
	"github.com/bwmarrin/discordgo"
)

var helpCategories = []discordgo.SelectMenuOption{
	{
		Label: "Zgłoś naruszenie regulaminu",
		Value: "regulamin",
		Emoji: discordgo.ComponentEmoji{
			Name: "📝",
		},
		Default:     false,
		Description: "Tu możesz zgłosić naruszenie regulaminu przez innego gracza (...)",
	},
	{
		Label: "Apelacja od bana",
		Value: "ub",
		Emoji: discordgo.ComponentEmoji{
			Name: "⚖️",
		},
		Default:     false,
		Description: "Tu możesz złożyć apelację o zdjęcie nałożonej kary (...)",
	},
	{
		Label: "Lepszy start po CK Twojej głównej postaci",
		Value: "second-life",
		Emoji: discordgo.ComponentEmoji{
			Name: "💰",
		},
		Default:     false,
		Description: "Podanie o lepszy start nowej postaci po uśmierceniu poprzedniej.",
	},
	{
		Label: "Organizacja eventu",
		Value: "event",
		Emoji: discordgo.ComponentEmoji{
			Name: "🎫",
		},
		Default:     false,
		Description: "Miejsce na Twój pomysł z eventem dla graczy.",
	},
	{
		Label: "Podanie o biznes",
		Value: "biznes",
		Emoji: discordgo.ComponentEmoji{
			Name: "💼",
		},
		Default:     false,
		Description: "Jeśli masz pomysł na biznes który chciałbyś prowadzić (...)",
	},
	{
		Label: "Podanie o organizację",
		Value: "organizacja",
		Emoji: discordgo.ComponentEmoji{
			Name: "🔫",
		},
		Default:     false,
		Description: "Tu możesz złożyć podanie o zorganizowaną grupę przestępczą (...)",
	},
	{
		Label: "Podanie o FCK",
		Value: "fck",
		Emoji: discordgo.ComponentEmoji{
			Name: "💀",
		},
		Default:     false,
		Description: "Tu możesz ubiegać się o uśmiercenie postaci innego gracza.",
	},
	{
		Label: "Podanie o FCJ",
		Value: "fcj",
		Emoji: discordgo.ComponentEmoji{
			Name: "🚗",
		},
		Default:     false,
		Description: "Tu możesz ubiegać się o przepisanie na siebie pojazdu innego gracza",
	},
	{
		Label: "Donate",
		Value: "donate",
		Emoji: discordgo.ComponentEmoji{
			Name: "💰",
		},
		Default:     false,
		Description: "Pomoc w sprawach związanych z donacjami na serwer",
	},
	{
		Label: "Sprawa do CM'a",
		Value: "cm",
		Emoji: discordgo.ComponentEmoji{
			Name: "😉",
		},
		Default:     false,
		Description: `Na wrażliwe tematy, widoczne tylko dla CM.`,
	},
	{
		Label: "Inne",
		Value: "inne",
		Emoji: discordgo.ComponentEmoji{
			Name: "❓",
		},
		Default:     false,
		Description: `Jeśli żadna kategoria nie pasuje do Twojego tematu.`,
	},
}

var categoryDescriptions = map[string]string{
	"regulamin": `
>>> W trakcie rozgrywki napotkałeś zachowanie niezgodne z regulaminem? Zamiast pisać w prywatnej wiadomości, wyślij nam opis sytuacji (najlepiej wraz z nagraniem, jeśli dotyczy).
Dzięki temu zminimalizujemy ilość wewnętrznej komunikacji i szybciej zajmiemy się Twoim zgłoszeniem.

Czy chcesz zgłosić naruszenie regulaminu?
    `,

	"second-life": `
>>> W tym miejscu możesz napisać podanie o lepszy start Twojej nowej postaci jeżeli w ramach rozgrywki uśmierciłeś swoją poprzednią główną postać.

Liczymy na postacie, które były na serwerze kreowane przez długi czas, były rozpoznawalne i byliście do niej bardzo przywiązani a samo uśmiercenie miało wpływ na rozwój wydarzeń dla innych. Aczkolwiek nigdy nie mów nie - jeżeli Twoja postać nie spełnia wszystkich warunków a uważasz, że akcja jest warta wynagrodzenia nie wahaj się z nią podzielić.
    `,

	"event": `
>>> W tym miejscu możesz szczegółowo przedstawić swój pomysł na zorganizowanie eventu serwerowego dla graczy. Opisz swój pomysł dosyć szczegółowo, nie zapomnij o opisaniu istotnych informacji takich jak szacunkowa długość trwania eventu, target czy też propozycji nagród.
    `,

	"biznes": `
>>> Biznesy dzielą się na dwie kategorie:

- **Biznes czysto pod odgywanie**, potrzebny jako tło lub narzędzie IC. Nie potrzebuje on skryptów ani uwagi Administracji do działania. Można go założyć składając podanie o zarejestrowanie działalności w sądzie.
- **Biznes zarobkowy**. Wymaga własnego skryptu i podania, które musi zostać zaakceptowane. Decydując się na zalożenie takiego biznesu miej na uwadze że:
    - Właścicielem biznesu może zostać osoba która ma przegrane łącznie 168h (7dni) na wszystkich postaciach
    - Biznes który ma wykorzystywać zupełnie nową mechanikę potrzebuje do tego własnego skryptu, którego tworzenie może zająć trochę czasu
    - Trzecie i kolejne biznesy tego samego typu mają mniejszą szansę na zaakceptowanie
    - Biznesy które zajmują się pożyczkami/wynajmem/lotnictwem/bronią nie będą akceptowane

Po kliknięciu Tak zostanie utworzony nowy temat z szablonem.
`,

	"organizacja": `
>>> Aby starać się o własną organizację najlepiej jest zebrać małą grupę ludzi i bez pisania podania powoli zaczynać budować swoją reputację wśród innych graczy i administracji przez własne RP.
To jaką opinię ma dana grupa starająca się o własną organizację ma bardzo duży wpływ na to czy zostanie zaakceptowana.
Zanim napiszecie podanie dobrym pomysłem jest napisanie do jednego z Community Managerów, aby ustalić szczegóły i dowiedzieć się czy dana organizacja jest wolna lub czy jej założenie będzie miało sens. Po takiej rozmowie złóż podanie używając poniższego szablonu.

Aby aplikować o organizację, lider musi mieć przegrane na serwerze łącznie 168h (7 dni) na wszystkich postaciach.

**Wymagania jakie nalezy spełnić:**

- Organizacja musi bazować na jednej z lore GTA (wszystkie gry)
- Nie może nawiązywać bezpośrednio do jej lore, np. poprzez dane konkretnych postaci lub wydarzenia ze świata GTA
- Należy umieścić ją w rozsądnej odległości od innych organizacji, aby nie przeludnić jednego punktu na mapie
- Nie rozpatrujemy podań na "sety"
- Biznesy "pod przykrywkę" będą rozpatrywane osobno
    `,

	"ub": `
>>> Jeżeli uważasz, że zostałeś niesłusznie ukarany banem na naszym serwerze możesz w tym miejscu złożyć apelację, którą skrupulatnie rozpatrzymy.
Pamiętaj, że **nie jest to miejsce na prośby o skrócenie kary**, a apelacje powinieneś złożyć tylko i wyłącznie w przypadku jeżeli uważasz, że ban został nadany niesłusznie lub omyłkowo.

Czy chcesz złożyć podanie o UB?
    `,

	"fck": `
>>> Jeżeli Tobie lub Twojej organizacji przestępczej ktoś mocno zaszedł za skórę, możesz w tym miejscu złożyć podanie o permanentne uśmiercenie postaci danego gracza.
Podanie o FCK może złożyć każdy, aczkolwiek musi być bardzo dobrze uargumentowane. Miej na uwadze zabawę również innych graczy.

Czy chcesz złożyć podanie o FCK?
    `,

	"fcj": `
>>> Jeżeli chcesz z kimś wyrównać rachunki, ale powody nie są na tyle mocne, żeby ubiegać się o zabicie postaci; możesz spróbować FCJ.
Jest to dobry sposób na rozliczenie się z kimś, kto podpadł Tobie lub Twojej grupie.

Czy chcesz złożyć podanie o FCJ?
    `,

	"donate": `
>>> Tu możesz przekazać nam dodatkowe informacje związane z donacją lub zadać pytanie
    `,

	"cm": `
>>> Tematy widoczne tylko dla Community Managerów

Czy chcesz stworzyć nowy temat?
    `,

	"inne": `
>>> Skorzystaj z tej opcji jeśli nie widzisz kategorii dla tematu, jaki chcesz omówić z administracją.
Temat widoczny dla całej moderacji.

Czy chcesz stworzyć nowy temat?
    `,
}

var categoryCreationInfo = map[string]string{
	"regulamin": `
>>> Napisz czego dotyczy Twoje zgłoszenie. Koniecznie załącz zdjecia lub klipy. Bez dowodów rzadko jesteśmy w stanie zdecydować czy nastąpiło złamanie regulaminu.
    `,

	"second-life": `
>>> Zgłoszenie jest widoczne tylko dla Administracji.

W tym miejscu możesz napisać podanie o lepszy start Twojej nowej postaci jeżeli w ramach rozgrywki uśmierciłeś swoją poprzednią główną postać.

Liczymy na postacie, które były na serwerze kreowane przez długi czas, były rozpoznawalne i byliście do niej bardzo przywiązani a samo uśmiercenie miało wpływ na rozwój wydarzeń dla innych. Aczkolwiek nigdy nie mów nie - jeżeli Twoja postać nie spełnia wszystkich warunków a uważasz, że akcja jest warta wynagrodzenia nie wahaj się z nią podzielić.

Wzór:
Imię i nazwisko uśmierconej postaci:
Kiedy to się stało:
Jakie były okoliczności śmierci:
Jaki to miało wpływ na rozgrywkę innych:
Twoje oczekiwania:
    `,

	"event": `
>>> Zgłoszenie jest widoczne tylko dla Administracji oraz osób należących do naszego grona Event Team.

W tym miejscu możesz szczegółowo przedstawić swój pomysł na zorganizowanie eventu serwerowego dla graczy. Opisz swój pomysł dosyć szczegółowo, nie zapomnij o opisaniu istotnych informacji takich jak szacunkowa długość trwania eventu, target czy też propozycji nagród.

Jednocześnie przypominamy, że całość projektu powinna być zgodna z regulaminem serwera jak i jego przebieg. Jeżeli podanie zostanie rozpatrzone pomyślnie możesz liczyć na pomoc z naszej strony w postaci udostępnienia Tobie specjalnych narzędzi tylko dla graczy z Event Team, które pomogą Ci zorganizować event na własną rękę z minimalną lub też żadną ingerencją administracji jak i moderacji.
    `,

	"biznes": `
>>> Podanie jest widoczne tylko dla Administracji, dlatego niepotrzebna nam historia powstania lub geneza pomysłu, jedynie infomacje które przekonają nas do zaakceptowania właśnie tego biznesu.

- Czym ma zajmować się biznes
- Propozycja działania skryptu
- Pracownicy (Dane IC oraz Discord)
- Miejsce biznesu (oraz to czy potrzebny jest interior)
- Motywacja
    `,

	"organizacja": `
>>> Co ma zawierać podanie:

- Opis organizacji, jej nazwa, pochodzenie
- Członkowie (Dane IC oraz Discord)
- Proponowany teren i siedziba
- Proponowany ubiór (tylko na ghetto, inne organizacje nie powinny mieć narzuconego stylu)
- Motywacja
    `,

	"ub": `
>>> Podanie służy odwołaniu jedynie w sytuacji, kiedy ban został nałożony niesłusznie lub omyłkowo.
**Nie proś o skrócenie bana.**

**Numer bana**: wyświetlany jest kiedy próbujesz wejść na serwer np. AWKD-BDH2

    `,

	"fck": `
>>> Pisząc podanie staraj się pisać zwięźle i treściwie. Jako administratorzy nie chcemy czytać tego z przymusu tylko z przyjemnością, zainspirujcie nas to tego stopnia, że sami będziemy chcieli daną osobę zabić :).

Wzór podania:
**Kogo?:** Imię i nazwisko osoby zabijanej
**Kto?:** Imię i nazwisko postaci która wydaje wyrok (jeżeli jest to grupa wypisz uczestników i wyznacz osobę odpowiedzialną)
**Dlaczego?:** Uzasadnij dlaczego z perspektywy Twojej postaci zasługuje na śmierć
**Jak?:** Jak zamierzasz to zrobić? Nie musi to być wypunktowana lista zadań, opisz jak to zrobisz pomijając nieistotne szczegóły.
**Co potem?:** Ważne aby to nie był jednorazowy zryw, jaki wpływ na dalszą grę Twoją i innych będzie miała na akcja?
    `,

	"fcj": `
>>> Pisząc podanie pisz zwięźle i treściwie bez zbędnych informacji, których nie potrzebujemy. Przedstaw same fakty i powody, które będą mocno wskazywać na Twoją rację a tym samym rośnie szansa, że Twoje podanie zostanie rozpatrzone pozytywnie.

Wzór podania:
**Co?:** Rejestracja, model, marka, dane właściciela pojazdu
**Kto?:** Imię i Nazwisko postaci która ma zamiar ukraść to auto (jeżeli jest to grupa wypisz uczestników i wyznacz osobę odpowiedzialną)
**Dlaczego?:** Co ma na celu kradzież tego właśnie pojazdu? Co wniesie to do rozgrywki?
**Jak?:** Jak zamierzasz to zrobić? Nie musi to być wypunktowana lista zadań, opisz jak to zrobisz pomijając nieistotne szczegóły.
**Co potem?:** Ważne aby to nie był jednorazowy zryw, jaki wpływ na dalszą grę Twoją i innych będzie miała na akcja?
    `,

	"donate": `
>>> Tu możesz przekazać nam dodatkowe informacje związane z donacją lub zadać pytanie
    `,
	"cm": `
>>> Postaraj się opisać temat, jaki chcesz omówić.
    `,

	"inne": `
>>> Postaraj się opisać temat, jaki chcesz omówić.
    `,
}

var defaultCategoryRoles = []Role{roles["CommunityManager"], roles["ProjectManager"]}

var categoryRoles = map[string][]Role{
	"biznes":      {roles["Businessman"]},
	"inne":        {roles["ServerAdmin"]},
	"fck":         {roles["CrimeManager"]},
	"fcj":         {roles["CrimeManager"]},
	"organizacja": {roles["ServerAdmin"]},
	"ub":          {roles["ServerAdmin"]},
	"cm":          {},
	"second-life": {roles["ServerAdmin"]},
	"event":       {roles["CommunityManager"]},
	"regulamin":   {roles["ServerAdmin"]},
}
