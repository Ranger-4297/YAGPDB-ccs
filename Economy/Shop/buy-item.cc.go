{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(buy(-?item)?)(\s+|\z)`

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

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
	{{$userdata := or (dbGet $userID "userEconData").Value (sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
	{{$inventory := $userdata.inventory}}
	{{with (dbGet 0 "store")}}
		{{$info := sdict .Value}}
		{{$items := sdict}}
		{{if ($info.Get "items")}}
			{{$items = sdict ($info.Get "items")}}
			{{with $.CmdArgs}}
				{{$name := (index . 0)}}
				{{if $items.Get $name}}
					{{$item := $items.Get (index . 0)}}
					{{$price := $item.Get "price"}}
					{{$qty := $item.Get "quantity"}}
					{{$uqty := 0}} {{/* USER QUANTITY */}}
					{{if ($inventory.Get $name)}} {{/* Checks if user has item with name*/}}
						{{$invitem := ($inventory.Get $name)}}
						{{$uqty = ($invitem.Get "quantity")}}
					{{end}}
					{{$bqty := 1}}
					{{if gt (len $.CmdArgs) 1}}
						{{$inp := (index . 1)}}
						{{if (toInt $inp)}}
							{{if lt (toInt $bqty) (toInt $qty)}}
								{{$bqty = (mult $bqty $inp)}}
								{{$price = (mult $price $inp)}}
							{{end}}
						{{else}}
							{{if (eq $inp "all" "max")}}
								{{$bqty = (mult (toInt $bqty) (toInt $qty))}}
								{{$price = (mult (toInt $price) (toInt $qty))}}
							{{end}}
						{{end}}
					{{end}}
					{{if le (toInt $price) (toInt $bal)}}
						{{$nqty := (sub $qty $bqty)}}
						{{if eq (toInt $nqty) 0}}
							{{$items.Del $name}}
						{{else}}
							{{$item.Set "quantity" $nqty}}
							{{$items.Set $name $item}}
						{{end}}
						{{$info.Set "items" $items}}
						{{dbSet 0 "store" $info}}
						{{$bal = (sub $bal $price)}}
						{{$inventory.Set $name (sdict "desc" ($item.Get "desc") "quantity" (add $uqty $bqty))}}
						{{$userdata.Set "inventory" $inventory}}
						{{$embed.Set "description" (print "Purchase successful!\nYou're now a proud owner of a(n) " $name)}}
						{{$embed.Set "color" $successColor}}
					{{else}}
						{{$embed.Set "description" (print "You don't have enough money for this")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "This item doesn't exist :( Create it with `" $prefix "create-item " $name "`\n\nTo view all items, run the `" $prefix "shop` command.")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No item argument provided :(\nSyntax is `" $.Cmd " <Name> [Quantity:Int]`\n\nTo view all items, run the `" $prefix "shop` command.")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "There are no items :(\nAdd some items with `" $prefix "create-item <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
	{{dbSet $userID "userEconData" $userdata}}
	{{dbSet $userID "cash" $bal}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}