{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((create|new)-?item)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}

{{/* Create item */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{with (dbGet 0 "store")}}
			{{$store := sdict .Value}}
			{{with $.CmdArgs}}
				{{if gt (len $.CmdArgs) 0}}
					{{$name := (index . 0)}}
					{{if gt (len $.CmdArgs) 1}}
						{{if (index . 1) | toInt}}
							{{if ge (toInt (index . 1)) 1}}
								{{$price := (index . 1)}}
								{{if gt (len $.CmdArgs) 2}}
									{{$qty := ""}}
									{{if (index . 2) | toInt}}
										{{if ge (toInt (index . 2)) 1}}
											{{$qty = (index . 2)}}
										{{else}}
											{{$qty = "Infinite"}}
										{{end}}
									{{else}}
										{{if eq (lower (index . 2)) "infinite" "infinity" "inf"}}
											{{$qty = "Infinite"}}
										{{else}}
											{{$embed.Set "description" (print "Invalid `Quantity` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
											{{$embed.Set "color" $errorColor}}
										{{end}}
									{{end}}
									{{if gt (len $.CmdArgs) 3}}
										{{$description := (joinStr " " (slice $.CmdArgs 3))}}
										{{$items := sdict}}
										{{if ($store.Get "Items")}}
											{{$items = sdict ($store.Get "Items")}}
										{{else}}
											{{dbSet 0 "store" (sdict "Items" sdict)}}
											{{with (dbGet 0 "store")}}
												{{$store = sdict .Value}}
											{{end}}
											{{$items = sdict ($store.Get "Items")}}
										{{end}}
										{{$items.Set $name (sdict "desc" $description "price" $price "qty" $qty)}}
										{{$store.Set "Items" $items}}
										{{dbSet 0 "store" $store}}
										{{$embed.Set "description" (print "New item added to shop!")}}
										{{$embed.Set "fields" (cslice (sdict "Name" $name "value" (print "Description: " $description "\nPrice: " (humanizeThousands $price) "\nQuantity: " (humanizeThousands $qty)) "inline" false))}}
										{{$embed.Set "color" $successColor}}
									{{else}}
										{{$embed.Set "description" (print "No `description` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else}}
									{{$embed.Set "description" (print "No `Quantity` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "Invalid `Price` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `Price` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "No `Price` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "No `Name` argument provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No arguments provided.\nSyntax is `" $.Cmd " <Name> <Price:Int> <Quantity:Int> <Description:String>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}