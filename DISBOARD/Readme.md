# Disboard/Bump CCs
Fully functional, notifications, and bump messages that look like they're through YAGPDB! :0
 
## Features
- Basic ping and notification that you are able to bump 
- Completely customizable (Given that you know what you are doing)
- Compatibility with **DISBOARD** (see [Other Info](#Other-Info))
- Easy to understand and read
- Logs for successful bumps
- More in the future!

## Usage

`!d bump` - Notifies you when you can bump (Image 1). Will either bump the server and log the bump (see image 2&3) or produce an error saying you cannot bump, and will output how long until the next bump (see image 4).

![Example Notify](https://cdn.discordapp.com/attachments/784132360399487066/809890885584027678/unknown.png)

![Example Bump 1](https://cdn.discordapp.com/attachments/784132360399487066/809888149794717716/unknown.png)

![Example Bump log](https://cdn.discordapp.com/attachments/784132360399487066/809889847509647423/unknown.png)

![Example Bump fail](https://cdn.discordapp.com/attachments/784132360399487066/809890534013009920/unknown.png)


`!d Help` - Provides you with the basic DISBOARD help page through YAGPDB with nothing changed. **Note**: The support server is DISBOARD Support (see image)

![Example Help](https://cdn.discordapp.com/attachments/784132360399487066/809891290639499274/unknown.png)

`!d Page` - Provides you with the basic DISBOARD server page through YAGPDB with nothing changed. (see image)

![Example Page](https://cdn.discordapp.com/attachments/784132360399487066/809891768114872400/unknown.png)

`Disboard|Bump` - sends an embed of the time until how long you can bump. Features a cooldown and 2 triggers.

![Example Disboard](https://cdn.discordapp.com/attachments/784132360399487066/809892129681440808/unknown.png)

## Other Info
Regarding compatibility with DISBOARD; DISBOARD, the bot will still need send and read messages for your bump channel. YAGPDB does **not** bump the server. It just resends the embed(s).

These commands **are** standalone. You are able to use them without each other, however, to use [`Disboard.cc.go`](https://github.com/Ranger-4297/YAGPDB-ccs/blob/main/DISBOARD/DISBOARD.cc.go) you will need to add [`BumpCommand.cc.go`](https://github.com/Ranger-4297/YAGPDB-ccs/blob/main/DISBOARD/BumpCommand.cc.go)

> *If you find any bugs or issues, feel free to PR an issue or fix, alternatively contact me through the [YAGPDB Support Server](https://discord.gg/SY7wn39SYD) or my server, [Dynamic](https://discord.gg/2WfF9JxuTU)*


## Credits

`Readme.md` - this document is an edited version of [NaruDevonte](https://github.com/NaruDevnote)'s [tag](https://github.com/NaruDevnote/yagpdb-ccs/tree/master/tags) custom command [README.md](https://github.com/NaruDevnote/yagpdb-ccs/blob/master/tags/README.md).

`BumpCommand.cc.go` - Credit for this command goes to [DZ-TM](https://github.com/DZ-TM) (Removed due to change in Disboard's Tos) and [WickedWizard](https://github.com/WickedWizard3588) (Removed due to change in Disboard's Tos) for the base command+notify and [TheHDCrafter](https://github.com/TheHDCrafter/yagpdb-cc)'s DISBOARD command which I took inspiration off to make this, but isn't public.
