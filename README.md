# Advent Of Code - Golang Italia

Sfida "ufficiale" per l'[Advent of Code](https://adventofcode.com/) degli appassionati di Go in Italia! :it:

## Di che si tratta?

L'Advent of Code è un calendario dell'avvento con piccoli quiz di programmazione.  

In questo repository vogliamo raccogliere le soluzioni in *Go* della community italiana, e sfidare le varie community locali (Roma, Torino, Milano, Napoli e magari altre!).


## Come funziona?

Dopo esservi iscritti utilizzate il codice di accesso `<cinque><zero>4742-765328a3` su https://adventofcode.com/2022/leaderboard/private ed aggiungetevi nel [`teams.yaml`](./teams.yaml) con il vostro ID (lo potrete trovare nei vostri [settings](https://adventofcode.com/2022/settings)).  

Sul sito ufficiale troverete dal 1° al 25 Dicembre un nuovo quiz da risolvere. Dopo averlo risolto se volete potrete pubblicare qui la vostra soluzinoe, e condividerla con la community.


## Cosa si vince?

La gloria eterna! E forse anche un piccolo premio, vedremo! :D

## Leaderboards

### User Leaderboard

| # | Score | User  | | Stars | Team  | Submissions  | 
|---|-------|-------|-|-------|-------|--------------|
| 1 | 0 |  | Andrea Manzini | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ |  | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 2 | 0 | ![https://github.com/Al-Pragliola.png?size=60](https://github.com/Al-Pragliola.png?size=60) | [Al-Pragliola](https://github.com/Al-Pragliola) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 3 | 0 | ![https://github.com/alessio-perugini.png?size=60](https://github.com/alessio-perugini.png?size=60) | [alessio-perugini](https://github.com/alessio-perugini) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 4 | 0 | ![https://github.com/enrichman.png?size=60](https://github.com/enrichman.png?size=60) | [enrichman](https://github.com/enrichman) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 5 | 0 | ![https://github.com/giulianopz.png?size=60](https://github.com/giulianopz.png?size=60) | [giulianopz](https://github.com/giulianopz) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 6 | 0 | ![https://github.com/lucianoq.png?size=60](https://github.com/lucianoq.png?size=60) | [lucianoq](https://github.com/lucianoq) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 7 | 0 | ![https://github.com/luigibarbato.png?size=60](https://github.com/luigibarbato.png?size=60) | [luigibarbato](https://github.com/luigibarbato) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 8 | 0 | ![https://github.com/mastrogiovanni.png?size=60](https://github.com/mastrogiovanni.png?size=60) | [mastrogiovanni](https://github.com/mastrogiovanni) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |
| 9 | 0 | ![https://github.com/omissis.png?size=60](https://github.com/omissis.png?size=60) | [omissis](https://github.com/omissis) | ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖ | golangmilano | 0 (2022) <br /> 0 (2021) <br /> 0 (2020) |

### Team Leaderboard

| # | Team  | Submissions  | 
|---|-------|--------------|
| 1 | golangroma | - |
| 2 | golangnapoli | - |
| 3 | golangtorino | - |
| 4 | golangmilano | - |


## Submissions

Le soluzioni condivise NON concorrono alla classifica. Potete caricare la vostra sotto `<YYYY>/day<DD>/username`.  

Lo script `new-puzzle.sh` creerà un po' di scaffolding, scaricando anche il file di input da testare.

Per utilizzarlo copiate il file `.env.sample` in `.env`, e modificatelo aggiungendo il vostro `USERNAME` ed il cookie `session` del sito dell'Advent of Code.  

```
-> % cp .env.sample .env
```
```
# .env

USERNAME=enrichman
session=my-super-secret-session
```
```
-> % ./scripts/new-puzzle.sh 2022 1                                                                                                  
You can now run 'cd ./2022/day01/enrichman' and work on your solution!
```