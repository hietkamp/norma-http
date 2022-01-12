# NORMA Simple HTTP Server

Norma is de Engelse benaming van het sterrenbeeld winkelhaak. Norma is geimplementeerd als prototype om specificaties te verifiëren op volledigheid en correctheid. Het betreffen de specificaties voor gevalideerde vragen van het programma KIK-V. 

Norma Simple HTTP Server is geimplementeerd voor het netwerkpunt van een afnemer van de data. De gevalideerde vraag is namelijk een asynchroon proces waarbij de afnemer een gevalideerde vraag stelt en op een later moment het antwoord verkrijgt. 

| :point_up:    | De software is "as-is".                                                         |
|               | Er wordt geen ondersteuning op verleend en is niet voor bedoeld voor productie! |
|---------------|:--------------------------------------------------------------------------------|

## Installeren

```bash
git clone git@github.com:hietkamp/norma-http.git
```

## Opstarten

```bash
./manage go
```

## Docker

Een docker image kan gebouwd worden door 

```bash
./manage docker
```