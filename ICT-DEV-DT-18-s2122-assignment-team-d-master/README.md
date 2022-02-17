# Eindopdracht DT : Studentvolgsysteem

## Introductie
In deze applicatie gaan jullie een studentvolgsysteem bouwen en uitbreiden. Een gedeelte van de functionaliteit staat al (deels). In de loop van het blok gaan jullie de bugs in bestaande functionaliteit oplossen, kleine verbeteringen maken aan bestaande functionaliteit, en nieuwe functionaliteit toevoegen.

## Installatie
### MySQL
Maak in MySQL workbench een nieuwe database aan

### .env aanmaken
Maak een kopie van het bestand .env.example met de naam `.env`, en pas de `user`, `pass` en `dbname` aan naar de gebruikersnaam, het wachtwoord en de databasenaam van je lokale database.

### De database seeden
Nu kan je Go applicatie verbinding maken met de database. We hebben een programma meegeleverd om de database met voorbeelddata te vullen, een zogeheten *seeder*. Deze kan je uitvoeren met `go run ./seeder`

## Projectstructuur
### main.go
Hierin wordt de server gestart, en worden routes gekoppeld aan handlers

### handlers
Deze bestanden bevatten logica om HTML terug te geven

### helpers
Hierin staat een helper functie die als middleware wordt aangeroepen. Deze zorgt ervoor dat alleen ingelogde gebruikers de applicatie kunnen gebruiken.

### public
Hierin staan statische bestanden (voornamelijk css), die in de hele applicatie gebruikt kunnen worden, ongeacht of je ingelogd bent

### repositories
Hierin staat logica om met de database te werken

### templates
Hierin staan de HTML templates die worden gerendered in de handlers

### types
Hierin staan de structs die worden gebruikt in de applicatie
