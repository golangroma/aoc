# Advent Of Code - Golang Italia

Sfida "ufficiale" per l'[Advent of Code](https://adventofcode.com/) degli appassionati di Go in Italia! :it:

## Di che si tratta?

L'Advent of Code è un calendario dell'avvento con piccoli quiz di programmazione.  

In questo repository vogliamo raccogliere le soluzioni in *Go* della community italiana, e sfidare le varie community locali (Roma, Torino, Milano, Napoli e magari altre!).


## Come funziona?

Per partecipare aggiungetevi nel [`teams.yaml`](./teams.yaml), ed eventualmente **anche** in un "team".  

Sul sito ufficiale invece potrete trovare i quiz.

Una volta iscritti lanciate `./scripts/new-puzzle.sh <YEAR> <DAY>` che creerà una cartella `2022/dayXX/username` con un po' di scaffolding.

```
-> % ./scripts/new-puzzle.sh 2022 1                                                                                                  
You can now run 'cd ./2022/day01/enrichman' and work on your solution!
```

Una volta risolto il quiz aprite una PR con la vostra soluzione!

## Setup

Copiate il file `.env.sample` in `.env`, e modificatelo aggiungendo il vostro `USERNAME` ed il cookie `session` del sito dell'Advent of Code.  

In questo modo lo script riuscirà a scaricare il vostro input personalizzato da poter usare per testare la soluzione.

```
-> % cp .env.sample .env
```
```
USERNAME=enrichman
session=my-super-secret-session
```

## Cosa si vince?

La gloria eterna! E forse anche un piccolo premio, vedremo! :D

## Leaderboards

### User Leaderboard

| # | User  | | Team  | Submissions  | 
|---|-------|-|-------|--------------|
| 1 | ![https://github.com/enrichman.png?size=40](https://github.com/enrichman.png?size=40) | [enrichman](https://github.com/enrichman) | golangroma |   |

### Team Leaderboard

| # | Team  | Submissions  | 
|---|-------|--------------|
| 1 | golangroma | - |

### Private Leaderboard

_coming soon_

Per partecipare alla "private leaderboard" contattatemi per ricevere il codice di accesso.
