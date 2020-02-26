# Discord Ban Bot

# How to build

```bash
docker build -t discordbanbot .
```


# Configuration
```bash
  --help         Show context-sensitive help (also try --help-long and --help-man).
  --discord-webhook=DISCORD-WEBHOOK
                 Webhookaddress
  --csvhour=7    Set the hour in which the script will run
  --csvminute=0  Set the minute in which the script will run
  --deadline="2020-08-18T20:23:00+01:00"
                 deadline till unban
```