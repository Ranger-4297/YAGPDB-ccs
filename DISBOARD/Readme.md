# Disboard/Bump CCs
Fully functional, notifications, and bump messages that look like they're through YAGPDB! :0
 
## Features
- Basic ping and notification that you are able to bump 
- Compatibility with **DISBOARD** (see [Other Info](#Other-Info))
- Easy to understand and read
- Logs for successful bumps

## Usage

`!d bump` - Notifies you when you can bump. Will either bump the server and log the bump or produce an error saying you cannot bump, and will output how long until the next bump.


`!d Help` - Provides you with the basic DISBOARD help page through YAGPDB with nothing changed.
**Note**: The support server shownis DISBOARD Support

`!d Page` - Provides you with the basic DISBOARD server page through YAGPDB with nothing changed.

`Disboard|Bump` - sends an embed of the time until how long you can bump. Features a cooldown and 2 triggers.

## Other Info
Regarding compatibility with DISBOARD; DISBOARD, the [YAGPDB] bot will still need send and read messages for your bump channel. YAGPDB does **not** bump the server, it just outputs the embed(s) as if it were.

These commands **are** standalone. You are able to use them without each other, however, to use [`Disboard.cc.go`](https://github.com/Ranger-4297/YAGPDB-ccs/blob/main/DISBOARD/DISBOARD.cc.go) you will need to add [`BumpCommand.cc.go`](https://github.com/Ranger-4297/YAGPDB-ccs/blob/main/DISBOARD/BumpCommand.cc.go)

<blockquote>If you find any bugs or issues, feel free to PR an issue or fix, alternatively contact me through the <a href="https://discord.gg/4uY54rw">YAGPDB support server</a> or VIA <a href="mailto:a.rhyker@gmail.com">email</a></blockquote>


## Credits

`BumpCommand.cc.go` - Credit for this command goes to [DZ-TM](https://github.com/DZ-TM) and [WickedWizard](https://github.com/WickedWizard3588) of which I took ideas and inspriation of off their own bump commands to make and [BlackWolf](https://github.com/BlackWolfWoof/yagpdb-cc)'s DISBOARD command which was the base idea of to make this, but isn't public.
