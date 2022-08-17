{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(ecohelp)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Help */}}

{{/* Resoponse */}}
{{$embed := sdict}}
{{$embed.Set "timestamp" currentTime}}
{{$embed.Set "color" $successColor}}
{{with .CmdArgs}}
    {{$cmd := (index . 0) | lower}}
    {{$desc := ""}}
    {{$use := ""}}
    {{$alias := ""}}
    {{if eq $cmd "balance" "bal"}}
        {{$cmd = "Balance"}}
        {{$desc = "View your balance"}}
        {{$use = "balance"}}
        {{$alias = (print "`bal` " "`wallet` " "`money`")}}
    {{else if eq $cmd "leaderboard" "lb" "top"}}
        {{$cmd = "Leaderboard"}}
        {{$desc = "Displays the leaderboard for the server"}}
        {{$use = "leaderboard"}}
        {{$alias = (print "`lb` " "`top`")}}
    {{else if eq $cmd "coinflip"}}
        {{$cmd = "Coinflip"}}
        {{$desc = "Flips a coin, if you win you get 2x the bet amount"}}
        {{$use = "coinflip <Side:Heads/Tails> <Bet:Amount>"}}
        {{$alias = (print "`flipcoin` " "`cf` " "`fc`")}}
    {{else if eq $cmd "work"}}
        {{$cmd = "Work"}}
        {{$desc = "Work with no chance of being fined"}}
        {{$use = "work"}}
        {{$alias = (print "`work` " "`labour`")}}
    {{else if eq $cmd "crime" "commit-crime"}}
        {{$cmd = "Crime"}}
        {{$desc = "Commit a crime, higher risk for higher output"}}
        {{$use = "crime"}}
        {{$alias = "`commit-crime`"}}
    {{else if eq $cmd "rob" "steal"}}
        {{$cmd = "Rob"}}
        {{$desc = "Attempts to rob a user with a chance of failure"}}
        {{$use = "rob <User:Mention/ID>"}}
        {{$alias = (print "`mug` " "`steal`")}}
    {{else if eq $cmd "addmoney"}}
        {{$cmd = "Addmoney"}}
        {{$desc = "Addmoney to a members cash or bank balance"}}
        {{$use = "addmoney <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>"}}
        {{$alias = "`increase-money`"}}
    {{else if eq $cmd "deposit" "dep"}}
        {{$cmd = "Deposit"}}
        {{$desc = "Deposit money to your bank"}}
        {{$use = "deposit"}}
        {{$alias = "`dep`"}}
    {{else if eq $cmd "withdraw" "with"}}
        {{$cmd = "Withdraw"}}
        {{$desc = "Withdraw money to your cash"}}
        {{$use = "withdraw"}}
        {{$alias = "`with`"}}
    {{else if eq $cmd "givemoney"}}
        {{$cmd = "Givemoney"}}
        {{$desc = "Give another member some of your cash"}}
        {{$use = "givemoney <User:Mention/ID> <Amount:Amount>"}}
        {{$alias = "`loan-money`"}}
    {{else if eq $cmd "removemoney"}}
        {{$cmd = "Removemoney"}}
        {{$desc = "Removes money from a members cash or bank balance"}}
        {{$use = "removemoney <User:Mention/ID> <Destination:Cash/Bank> <Amount:Amount>"}}
        {{$alias = "`decrease-money`"}}
    {{else if eq $cmd "set" "configure"}}
        {{$cmd = "Set"}}
        {{$desc = "Configures the servers economy settings"}}
        {{$use = "set <Setting:String> <Value:String/Int/Duration>"}}
        {{$alias = "`Configure`"}}
    {{else if eq $cmd "viewsettings"}}
        {{$cmd = "Viewsettings"}}
        {{$desc = "Views the economy settings of the server"}}
        {{$use = "viewsettings"}}
    {{else if eq $cmd "create-item"}}
        {{$cmd = "Create-item"}}
        {{$desc = "Adds items to the shop"}}
        {{$use = "create-item <Name:String> <Price:Int> <Quantity:Int> <Description:String>"}}
        {{$alias = "`new-item`"}}
    {{else if eq $cmd "shop"}}
        {{$cmd = "Shop"}}
        {{$desc = "Views the servers shop"}}
        {{$use = "Shop [Page:Int]"}}
        {{$alias = "`store`"}}
    {{else}}
        {{$cmd = "Invalid command provided"}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
    {{$embed.Set "title" $cmd}}
    {{$embed.Set "description" $desc}}
    {{if $use}}
        {{$embed.Set "fields" (cslice 
            (sdict "name" "Usage" "value" (print "`" $use "`") "inline" true))}}
    {{end}}
    {{if $alias}}
        {{$embed.Set "fields" (($embed.Get "fields").Append (sdict "name" "Alias(es)" "value" (print $alias) "inline" true))}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "__**List of all commands**__\n\n**Informational**\n`Balance`\n`Leaderboard`\n\n**Income**\n`CoinFlip`\n`Work`\n`Crime`\n`Rob`\n\n**Management**\n`Addmoney`\n`Deposit`\n`Withdraw`\n`Givemoney`\n`Removemoney`\n\n**Settings**\n`Set`\n`Viewsettings`\n\n**Shop**\n`Create-item`\n`Shop`")}}
    {{$embed.Set "color" $successColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}