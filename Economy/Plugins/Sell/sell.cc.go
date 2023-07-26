{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(sell)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Sell  */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
	{{$userdata := or (dbGet $userID "userEconData").Value (sdict "settings" (sdict "balance" "yes" "inventory" "yes" "leaderboard" "yes" "trading" "yes") "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
    {{$inventory := $userdata.inventory}}
	{{$store := or (dbGet 0 "store").Value (sdict "items" sdict)}}
    {{$items := $store.items}}
    {{with $.CmdArgs}}
        {{$item := (index . 0)}}
        {{with $inventory.Get $item}}
            {{$invQuantity := .quantity}}
            {{$shopItem := $item}}
            {{if $items.Get $item}}
                {{$shopItem = (print $item "." $.User.Username)}}
            {{end}}
            {{$item := $inventory.Get $item}}
            {{if gt (len $.CmdArgs) 1}}
                {{$price := (index $.CmdArgs 1)}}
                {{if toInt $price}}
                    {{$sellQuantity := 1}}
                    {{$cont = 0}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$sellQuantity = (index . 1)}}
                        {{if toInt $sellQuantity}}
                            {{if ge (toInt $sellQuantity) 1}}
                                {{if le (toInt $sellQuantity) (toInt $invQuantity)}}
                                    {{$cont = true}}
                                    {{$invQuantity = sub $invQuantity $sellQuantity}}
                                {{else}}
                                    {{$embed.Set "description" (print "There's not enough of this in the shop to buy that much!")}}
                                    {{$embed.Set "color" $errorColor}}
                                {{end}}
                            {{else}}
                                {{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" $.Cmd "<Item:Item name> <Price:Int> [Quantity:Int/All]`")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{else}}
                            {{$sellQuantity = lower $sellQuantity}}
                            {{if eq (toString $sellQuantity) "all"}}
                                {{$sellQuantity = $invQuantity}}
                                {{$invQuantity = 0 }}
                                {{$cont = 1}}
                            {{else}}
                                {{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" $.Cmd "<Item:Item name> <Price:Int> [Quantity:Int/All]`")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                        {{end}}
                    {{end}}
                    {{if $cont}}
                        
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `price` argument provided.\nSyntax is `" $.Cmd " <Item:Item name> <Price:Int> [Quantity:Int/All]\n\nTo view your items, run " $prefix "inventory`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$embed.Set "description" (print "No `price` argument provided.\nSyntax is `" $.Cmd " <Item:Item name> <Price:Int> [Quantity:Int/All]\n\nTo view your items, run " $prefix "inventory`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "Invalid `item` argument provided.\nSyntax is `" $.Cmd " <Item:Item name> <Price:Int> [Quantity:Int/All]\n\nTo view your items, run " $prefix "inventory`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No `item` argument provided.\nSyntax is `" $.Cmd " <Item:Item name> <Price:Int> [Quantity:Int/All]")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
    {{dbSet 0 "store"}}
    {{dbSet $userID "userEconData" $userData}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}