{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(buy-?item)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Buy item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$bal := $a.cash}}
		{{$inventory := sdict}}
		{{with (dbGet 0 "store")}}
			{{$info := sdict .Value}}
			{{$items := sdict}}
			{{if ($info.Get "Items")}}
				{{$items = sdict ($info.Get "Items")}}
				{{with $.CmdArgs}}
					{{$name := (index . 0)}}
					{{if $items.Get $name}}
						{{$item := $items.Get (index . 0)}}
						{{$price := $item.Get "price"}}
						{{$qty := $item.Get "quantity"}}
						{{$uqty := 0}} {{/* USER QUANTITY */}}
						{{$inventory := $a.Get "inventory"}}
						{{if ($inventory.Get $name)}} {{/* Checks if user has item with name*/}}
							{{$invitem := ($inventory.Get $name)}}
							{{$uqty = ($invitem.Get "qty")}}
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
								{{$info.Set "Items" $items}}
								{{dbSet 0 "store" $info}}
							{{else}}
								{{$item.Set "quantity" $nqty}}
								{{$items.Set $name $item}}
								{{$info.Set "Items" $items}}
								{{dbSet 0 "store" $info}}
							{{end}}
							{{$nbal := (sub $bal $price)}}
							{{$inventory := (sdict $name (sdict "desc" ($item.Get "desc") "quantity" (add $uqty $bqty)))}}
							{{$a.Set "cash" $nbal}}
							{{$a.Set "inventory" $inventory}}
							{{dbSet $userID "EconomyInfo" $a}}
							{{$embed.Set "description" (print "Purchase successful!\nYou're now a proud owner of a(n) " $name)}}
							{{$embed.Set "color" $successColor}}
						{{else}}
							{{$embed.Set "description" (print "You don't have enough money for this")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "Invalid item argument provided :(\nSyntax is `" $.Cmd " <Name> [Quantity:Int]`\n\nTo view all items, run the `" $prefix "shop` command.")}}
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
	{{else}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
		{{$embed.Set "description" (print "You were not in the economy database....adding you now\nPlease try again")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}