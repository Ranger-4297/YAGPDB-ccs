{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(use(-?item)?)(\s+|\z)`

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

{{/* Use Item */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{$userdata := or (dbGet $userID "userEconData").Value (sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
{{$inventory := $userdata.inventory}}
{{if not $inventory }}
	{{$embed.Set "description" (print "Your inventory is empty :(\nYou should get some items from the shop!")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `Item` argument provided.\nSyntax is `" .Cmd "  <Item:Name>`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$name := index . 0}}
{{if not ($items.Get $name)}}
	{{$embed.Set "description" (print "This item doesn't exist\nUse `" $prefix "shop` to view the items!")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$item := $inventory.Get $name}}
{{$qty := $item.quantity}}
{{$role := $item.role}}
{{$nqty := sub (toInt $qty) 1}}
{{if eq (toInt $nqty) 0}}
	{{$inventory.Del $name}}
{{else}}
	{{$item.Set "quantity" $nqty}}
	{{$inventory.Set $name $item}}
	{{$userdata.Set "inventory" $inventory}}
{{end}}
{{if $role}}
	{{addRoleID $role}}
{{end}}
{{dbSet $userID "userEconData" $userdata}}
{{sendMessage nil $item.replyMsg}}