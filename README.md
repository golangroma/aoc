# Advent Of Code - Golang Italia

Sfida "ufficiale" per l'[Advent of Code](https://adventofcode.com/) degli appassionati di Go in Italia! :it:

## Di che si tratta?

L'Advent of Code è un calendario dell'avvento con piccoli quiz di programmazione.  

In questo repository vogliamo raccogliere le soluzioni in *Go* della community italiana, e sfidare le varie community locali (Roma, Torino, Milano, Napoli e magari altre!).


## Come funziona?

Per partecipare è sufficiente aggiungere il vostro username Github nel [`teams.yaml`](./teams.yaml) e, nel caso vogliate partecipare con un team, aggiungervi **anche** nella sezione "team".  

Ovviamente dovrete iscrivervi anche sul sito ufficiale, dove potrete trovare i quiz.

Una volta iscritti basterà aprire una PR con la vostra soluzione (sotto `2022/dayXX/username`, aggiungendo anche il vostro `input.txt` (che è personalizzato).

## Setup

Per comodità il repo ha un piccolo script `./scripts/new-puzzle.sh` per poter generare uno scaffolding uguale per tutti.  
Copiate il file `.env.sample` in `.env`, e modificatelo aggiungendo il vostro `USERNAME` ed il cookie `session` del sito dell'Advent of Code.
In questo modo lo script riuscirà in automatico anche a scaricare il vostro input.

```
-> % cp .env.sample .env
```
```
USERNAME=enrichman
session=my-super-secret-session
```

Lo script accetta l'anno ed il giorno di riferimento, creando una cartella con il vostro input ed alcuni file.
```
-> % ./scripts/new-puzzle.sh 2022 1                                                                                                  
You can now run 'cd ./2022/day01/enrichman' and work on your solution!
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

## Donazioni

Se l'idea ti è piaciuta puoi offrirmi un caffè! :coffee:
Ogni danaro verrà utilizzato nell'acquisto/stampa del premio (o premi?).
Tagga il messaggio con #AOC per identificare la donazione!

https://www.buymeacoffee.com/enrichman