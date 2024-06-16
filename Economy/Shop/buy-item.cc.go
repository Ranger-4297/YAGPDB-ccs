{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(buy(-?item)?)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}
{{$ex := or (and (reFind "a_" .Guild.Icon) "gif") "png"}}
{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024"}}

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" (print .Guild.Name " Store")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
{{$userdata := or (dbGet $userID "userEconData").Value (sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
{{$inventory := $userdata.inventory}}
{{if .ExecData}}
	{{$userData := (dbGet .ExecData.user "userEconData").Value}}
	{{$item := .ExecData.itemName}}
	{{$inventory := $userData.inventory}}
	{{if $item = ($inventory.Get $item)}}
		{{if gt $item.quantity 1}}
			{{$item.Set "quantity" (sub $item.quantity 1)}}
		{{else}}
			{{$inventory.Del $item}}
		{{end}}
	{{end}}
	{{dbSet .ExecData.user "userEconData" $userData}}
	{{return}}
{{end}}
{{$store := (dbGet 0 "store").Value}}
{{$items := $store.items}}
{{if not $items}}
	{{$embed.Set "description" (print "The shop is empty :(\nAdd some items with `" $prefix "create-item`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `Item` argument provided.\nSyntax is `" .Cmd "  `\nSyntax is `" .Cmd " <Item:Name> [Quantity:Int/All]`\n\nTo view all items, run the `" $prefix "shop` command.")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$name := index .CmdArgs 0}}
{{if not ($items.Get $name)}}
	{{$embed.Set "description" (print "This item doesn't exist\nUse `" $prefix "shop` to view the items!")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$item := $items.Get $name}}
{{$price := $item.Get "price"}}
{{$shopQuantity := $item.Get "quantity"}}
{{$userQuantity := 0}}
{{if ($inventory.Get $name)}}
	{{$invitem := $inventory.Get $name}}
	{{$userQuantity = ($invitem.Get "quantity")}}
{{end}}
{{$buyQuantity := 1}}
{{if gt (len .CmdArgs) 1}}
	{{$buyQuantity = index .CmdArgs 1}}
	{{if $buyQuantity = toInt $buyQuantity}}
		{{if lt $buyQuantity 1}}
			{{$embed.Set "description" (print "Invalid `Quantity` argument provided :(\nSyntax is `" .Cmd " <Item:Name> [Quantity:Int/All]`")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
	{{else if $buyQuantity = lower $buyQuantity}}
		{{if not (eq (toString $buyQuantity) "max" "all")}}
			{{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" .Cmd " " $name " [Quantity:Int/All/Max]`")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$option := $buyQuantity}}
		{{$buyQuantity = div $bal $price}}
		{{$buyQuantity := (sdict "max" (div $bal $price) "all" $shopQuantity).Get $option}}
	{{end}}
{{end}}
{{if and (eq $name "chicken") (gt (toInt $buyQuantity) 1)}}
	{{$embed.Set "description" (print "You can only buy 1 " $name)}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if lt $buyQuantity $shopQuantity}}
	{{$embed.Set "description" (print "There's not enough of this in the shop to buy that much!")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$price = mult $buyQuantity $price}}
{{if lt $bal $price}}
	{{$embed.Set "description" (print "You don't have enough to buy this")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if (and $item.ID (eq (toInt $item.ID) (toInt .User.ID)))}}
	{{$embed.Set "description" (print "You cannot buy an item you have listed")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$userQuantity = add $userQuantity $buyQuantity}}
{{$bal = sub $bal $price}}
{{if $shopQuantity}}
	{{$shopQuantity = sub $shopQuantity $buyQuantity}}
	{{if eq $shopQuantity 0}}
		{{$items.Del $name}}
	{{else}}
		{{$item.Set "quantity" $shopQuantity}}
		{{$items.Set $name $item}}
	{{end}}
{{end}}
{{$exp := $item.expiry}}
{{$expires := "never"}}
{{if $exp}}
	{{$timeSeconds := toDuration (humanizeDurationSeconds (mult $exp .TimeSecond))}}
	{{$expires = (print "<t:" (currentTime.Add $timeSeconds).Unix ":f>")}}
{{end}}
{{$store.Set "items" $items}}
{{dbSet 0 "store" $store}}
{{if $inventory.Get $name}}
	{{$item = $inventory.Get $name}}
	{{$item.Set "quantity" $userQuantity}}
	{{$inventory.Set $name $item}}
{{else}}
	{{$inventory.Set $name (sdict "desc" $item.desc "quantity" $userQuantity "role" $item.role "replyMsg" $item.replyMsg "expiry" $exp)}}
{{end}}
{{$embed.Set "description" (print "You've bought  " (humanizeThousands $buyQuantity) " of " $name " for " $symbol (humanizeThousands $price) "!")}}
{{$embed.Set "color" $successColor}}
{{if $exp}}
	{{execCC .CCID nil $exp (sdict "user" .User.ID "itemName" $name "expiry" $exp)}}
{{end}}
{{dbSet $userID "userEconData" $userdata}}
{{dbSet $userID "cash" $bal}}
{{sendMessage nil (cembed $embed)}}