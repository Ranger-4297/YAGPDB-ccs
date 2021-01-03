# Disboard/Bump CCs
Fully functional, notifications, and bump messages that look like they're through YAGPDB! :0

## Features
- Basic ping and notification that you are able to bump 
- Completely customizable (if you know what you are doing)
- Compatibility with **DISBOARD** (see [Other Info](#Other-Info))
- Easy to understand and read

## Usage

`!d bump` - Will either bump the server (see image 1) or produce an error saying you cannot bump, and will output how long until the next bump (see image 2).

![Example Bump 1](https://cdn.discordapp.com/attachments/784132357002625047/795167404535185408/unknown.png)
![Example Bump 2](https://cdn.discordapp.com/attachments/784132357002625047/795171341384810516/unknown.png)

`Notify` - this is part of the `!d Bump` command, and isk what else to add here (I don't have the role added).
![Example Disboard](https://cdn.discordapp.com/attachments/784132357002625047/795179477574877194/unknown.png)

`!d Help` - Provides you with the basic DISBOARD help page through YAGPDB, nothing changed. (see image)
![Example Help](https://cdn.discordapp.com/attachments/784132357002625047/795175235329589278/unknown.png)

`!d Page` - Provides you with the basic DISBOARD server page through YAGPDB, nothing changed. (see image)
![Example Page](https://cdn.discordapp.com/attachments/784132357002625047/795176033036271636/unknown.png)

`Disboard` - sends an embed of the time until how long you can bump. Has a cooldown and multiple triggers (RegEx).
![Example Disboard](https://cdn.discordapp.com/attachments/784132357002625047/795176964042129438/unknown.png)

## Other Info
Regarding compatibility with DISBOARD; DISBOARD, the bot will still need send and read messages for your bump channele. YAGPDB does **not** bump the server. It just resends the embed(s).

These commands **are** standalone. You are able to use them without each other, except `BumpNotify.cc.go` to use this, you will need to add `BumpCommand.cc.go`

> *If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*


## credits

`Readme.me` - this document is an edited version of [NaruDevonte](https://github.com/NaruDevnote)'s [tag](https://github.com/NaruDevnote/yagpdb-ccs/tree/master/tags) custom command [README.md](https://github.com/NaruDevnote/yagpdb-ccs/blob/master/tags/README.md)

`BumpCommand.cc.go` - Credit for this command goes to [DZ-TM](https://github.com/DZ-TM/Yagpdb.xyz)'s [BumpCommand.go](https://github.com/DZ-TM/Yagpdb.xyz/blob/master/Commands/Bump/BumpCommand.go) for the base command+notify and [TheHDCrafter](https://github.com/TheHDCrafter/yagpdb-cc)'s DISBOARD command which I took inspiration off to make this, but isn't public
