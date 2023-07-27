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
	{{if not $.ExecData}}
		{{with (dbGet 0 "store")}}
			{{$shop := sdict .Value}}
			{{$items := sdict}}
			{{if ($shop.Get "items")}}
				{{$items = sdict ($shop.Get "items")}}
				{{with $.CmdArgs}}
					{{$name := (index . 0)}}
					{{if $items.Get $name}}
						{{$item := $items.Get $name}}
						{{$price := $item.Get "price"}}
						{{$shopQuantity := $item.Get "quantity"}}
						{{$userQuantity := 0}}
						{{if ($inventory.Get $name)}}
							{{$invitem := ($inventory.Get $name)}}
							{{$userQuantity = ($invitem.Get "quantity")}}
						{{end}}
						{{$buyQuantity := 1}}
						{{$cont := false}}
						{{if gt (len $.CmdArgs) 1}}
							{{$buyQuantity = (index . 1)}}
							{{if toInt $buyQuantity}}
								{{if ge (toInt $buyQuantity) 1}}
									{{if not (eq (toString $shopQuantity) "inf")}}
										{{if le (toInt $buyQuantity) (toInt $shopQuantity)}}
											{{$cont = true}}
											{{$shopQuantity = sub $shopQuantity $buyQuantity}}
										{{else}}
											{{$embed.Set "description" (print "There's not enough of this in the shop to buy that much!")}}
											{{$embed.Set "color" $errorColor}}
										{{end}}
									{{else}}
										{{$cont = true}}
									{{end}}
								{{else}}
									{{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" $.Cmd " " $name " [Quantity:Int/All]`")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$buyQuantity = lower $buyQuantity}}
								{{if eq (toString $buyQuantity) "all"}}
									{{if eq (toString $shopQuantity) "inf"}}
										{{$buyQuantity = div (toInt $bal) $price}}
										{{$cont = true}}
									{{else}}
										{{$buyQuantity = $shopQuantity}}
										{{$shopQuantity = sub $shopQuantity $buyQuantity}}
										{{$cont = true}}
									{{end}}
								{{else}}
									{{$embed.Set "description" (print "Invalid quantity argument provided :(\nSyntax is `" $.Cmd " " $name " [Quantity:Int/All[`")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{end}}
						{{else}}
							{{if not (eq (toString $shopQuantity) "inf")}}
								{{$shopQuantity = sub $shopQuantity 1}}
							{{end}}
							{{$buyQuantity = 1}}
							{{$cont = true}}
						{{end}}
						{{if $cont}}
							{{if not (and (eq $name "chicken") (gt (toInt $buyQuantity) 1))}}
								{{$price = mult $buyQuantity $price}}
								{{if ge $bal $price}}
									{{if not (and $item.ID (eq (toInt $item.ID) (toInt $.User.ID)))}}
										{{$userQuantity = add $userQuantity $buyQuantity}}
										{{$bal = sub $bal $price}}
										{{if not $shopQuantity}}
											{{$items.Del $name}}
										{{else}}
											{{$item.Set "quantity" $shopQuantity}}  
											{{$items.Set $name $item}}
										{{end}}
										{{$exp := $item.expiry}}
										{{$expires := "never"}}
										{{if (toDuration $exp)}}
											{{$timeSeconds := toDuration (humanizeDurationSeconds (mult $exp $.TimeSecond))}}
											{{$expires = (print "<t:" (currentTime.Add $timeSeconds).Unix ":f>")}}
										{{end}}
										{{$shop.Set "items" $items}}
										{{dbSet 0 "store" $shop}}
										{{if $inventory.Get $name}}
											{{$item = $inventory.Get $name}}
											{{$item.Set "quantity" $userQuantity}}
											{{$inventory.Set $name $item}}
										{{else}}
											{{$inventory.Set $name (sdict "desc" $item.desc "quantity" $userQuantity "role" $item.role "replyMsg" $item.replyMsg "expiry" $exp "expires" $expires)}}
										{{end}}
										{{$embed.Set "description" (print "You've bought  " $buyQuantity " of " $name " for " $symbol $price "!")}}
										{{$embed.Set "color" $successColor}}
										{{if $exp}}
											{{scheduleUniqueCC $.CCID nil $exp $name (sdict "user" $.User.ID "itemName" $name "expiry" $exp)}}
										{{end}}
									{{else}}
										{{$embed.Set "description" (print "You cannot buy your own item")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else}}
									{{$embed.Set "description" (print "You don't have enough to buy this :(")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "You can only buy 1 of the item " $name)}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "This item doesn't exist :( Create it with `" $prefix "create-item " $name "`\n\nTo view all items, run the `" $prefix "shop` command.")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "No item argument provided :(\nSyntax is `" $.Cmd " <Name> [Quantity:Int/All]`\n\nTo view all items, run the `" $prefix "shop` command.")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "There are no items :(\nAdd some items with `" $prefix "create-item <Name> <Price:Int> <Quantity:Int/Inf> <Description:String>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
		{{dbSet $userID "userEconData" $userdata}}
		{{dbSet $userID "cash" $bal}}
	{{else}}
		{{$userData := (dbGet $.ExecData.user "userEconData").Value}}
		{{$item := $.ExecData.itemName}}
		{{$inventory := $userData.inventory}}
		{{if $inventory.Get $item}}
			{{$inventory.Del $item}}
		{{end}}
		{{dbSet $.ExecData.user "userEconData" $userData}}
		{{cancelScheduledUniqueCC $.CCID $.ExecData.itemName}}
		{{return}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}