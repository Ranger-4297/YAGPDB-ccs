{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(use(-?item)?)(\s+|\z)`

	©️ Ranger 2020-Present
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
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$userdata := or (dbGet $userID "userEconData").Value (sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
	{{if $inventory := $userdata.inventory}} 
		{{with $.CmdArgs}}
			{{$name := (index . 0)}}
			{{if $item := $inventory.Get $name}}
				{{$qty := $item.quantity}}
				{{$role := $item.role}}
				{{$nqty := (sub (toInt $qty) 1)}}
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
				{{$embed.Set "description" (print "You've just used " $name "!\nYou had " (humanizeThousands $qty) " and now have " (humanizeThousands $nqty) "\nIf there was a role associated with this item it has been assigned!")}}
				{{$embed.Set "color" $successColor}}
			{{else}}
				{{$embed.Set "description" (print "Invalid item argument provided :(\nSyntax is `" $.Cmd " <Name>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "No item argument provided :(\nSyntax is `" $.Cmd " <Name>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "You have no items in your inventory, purchase some from the shop!")}}
	{{end}}
	{{dbSet $userID "userEconData" $userdata}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}