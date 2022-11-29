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


### Private Leaderboard

Per partecipare alla "private leaderboard" utilizzate il codice di accesso `<cinque><zero>4742-765328a3` su https://adventofcode.com/2022/leaderboard/private


### User Leaderboard

| # | User  | | Team  | Submissions  | 
|---|-------|-|-------|--------------|
| 1 | ![https://github.com/enrichman.png?size=60](https://github.com/enrichman.png?size=60) | [enrichman](https://github.com/enrichman) | golangroma | 0 (2022) <br /> 3 (2021) <br /> 0 (2020) |
| 2 | ![https://github.com/Al-Pragliola.png?size=60](https://github.com/Al-Pragliola.png?size=60) | [Al-Pragliola](https://github.com/Al-Pragliola) | golangnapoli | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 3 | ![https://github.com/alessio-perugini.png?size=60](https://github.com/alessio-perugini.png?size=60) | [alessio-perugini](https://github.com/alessio-perugini) | golangroma | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 4 | ![https://github.com/giulianopz.png?size=60](https://github.com/giulianopz.png?size=60) | [giulianopz](https://github.com/giulianopz) | golangroma | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 5 | ![https://github.com/lucianoq.png?size=60](https://github.com/lucianoq.png?size=60) | [lucianoq](https://github.com/lucianoq) | golangroma | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 6 | ![https://github.com/luigibarbato.png?size=60](https://github.com/luigibarbato.png?size=60) | [luigibarbato](https://github.com/luigibarbato) | golangnapoli | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 7 | ![https://github.com/mastrogiovanni.png?size=60](https://github.com/mastrogiovanni.png?size=60) | [mastrogiovanni](https://github.com/mastrogiovanni) | golangroma | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 8 | ![https://github.com/omissis.png?size=60](https://github.com/omissis.png?size=60) | [omissis](https://github.com/omissis) | golangtorino | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |

### Team Leaderboard

| # | Team  | Submissions  | 
|---|-------|--------------|
| 1 | golangroma | - |
| 2 | golangnapoli | - |
| 3 | golangtorino | - |
| 4 | golangmilano | - |
