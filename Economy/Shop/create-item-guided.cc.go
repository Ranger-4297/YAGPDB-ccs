{{/*
		Made by Ranger (765316548516380732)
		Wait response logic by DZ (438789314101379072)

	Trigger Type: `Regex`
	Trigger: `.*`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* If your prefix isn't - change the value below to your prefix */}}
{{$trigger:=`(-|<@!?204255221017214977>\s*)((create|new)-?item)`}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$errorColor := 0xFF0000}}
{{$dbVal := toInt (dbGet .User.ID "waitResponse").Value}}
{{$cmdStage := 0}}

{{/* Create item */}}

{{/* Response */}}
{{$embed := sdict "title" "Item info" "footer" (sdict "text" "Type cancel to cancel the setup") "color" 0x00ff7b "timestamp" currentTime}}
{{$shop := or (dbGet 0 "store").Value (sdict "items" sdict)}}
{{$items := $shop.items}}
{{$itemData := sdict "desc" "" "price" 0 "quantity" 0 "role" 0 "reply-msg" ""}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if not .ExecData}}
	{{with $eco := (dbGet 0 "EconomySettings").Value}}
		{{$symbol := $eco.symbol}}
		{{if not $dbVal}}
			{{/* checks if message matches regex to begin tutorial */}}
			{{if reFind (print `\A(?i)` $trigger `(\s+|\z)`) $.Message.Content}}
				{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
					{{$embed.Set "fields" (cslice (sdict "name" "Name" "value" "⠀⠀"))}}
					{{$cmdStage = 1}}
					{{$embed := sendMessageRetID nil (complexMessage "content" "Please enter a name for the item (under 60 characters)" "embed" (cembed $embed))}}
					{{dbSet 0 "createItem" (sdict "embed" (toString $embed) "item" (sdict "name" "" "data" $itemData) "user" (toString $userID))}}
					{{scheduleUniqueCC $.CCID nil 120 1 1}}
				{{else}}
					{{$embed.Del "footer"}}{{$embed.Del "title"}}
					{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
					{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
					{{$embed.Set "color" $errorColor}}
					{{sendMessage nil (cembed $embed)}}
				{{end}}
			{{end}}
		{{else}}
			{{$createItem := (dbGet 0 "createItem").Value}}
			{{if and $createItem.user (eq (toString $createItem.user) (toString $userID))}}
				{{if eq $dbVal 1}}
					{{if and (le (len (toRune $.Message.Content)) 60) (not (eq $.Message.Content "cancel"))}}
						{{$name := $.Message.Content}}
						{{$item := $createItem.item}}
						{{$item.Set "name" $name}}
						{{dbSet 0 "createItem" $createItem}}
						{{$embed.Set "fields" (cslice (sdict "name" "Name" "value" $name "inline" true))}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Please enter a price for the item" "embed" (cembed $embed))}}
						{{$cmdStage = 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
						{{$cmdStage = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 2}}
				{{else if eq $dbVal 2}}
					{{with toInt $.Message.Content}}
						{{if gt . 0}}
							{{$price := .}}
							{{$item := $createItem.item}}
							{{$item.data.Set "price" $price}}
							{{dbSet 0 "createItem" $createItem}}
							{{$msg := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
							{{$field := sdict "name" "Price" "value" (print $symbol .) "inline" true}}
							{{$embed.Set "fields" ((cslice.AppendSlice $msg.Fields).Append $field)}}
							{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Please enter a description for the item (under 200 characters)" "embed" (cembed $embed))}}
							{{$cmdStage = 1}}
						{{else}}
							{{$msg := sendMessageRetID nil "Please try again and enter a valid number"}}
							{{deleteTrigger 0}}
							{{deleteMessage nil $msg 10}}
							{{$cmdStage = 0}}
						{{end}}
					{{else if not (eq (lower $.Message.Content) "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again and enter a valid number"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
						{{$cmdStage = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 3}}
				{{else if eq $dbVal 3}}
					{{if and (le (len (toRune $.Message.Content)) 200) (not (eq $.Message.Content "cancel"))}}
						{{$item := $createItem.item}}
						{{$item.data.Set "desc" $.Message.Content}}
						{{dbSet 0 "createItem" $createItem}}
						{{$msg := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$field := sdict "name" "Description" "value" $.Message.Content}}
						{{$embed.Set "fields" ((cslice.AppendSlice $msg.Fields).Append $field)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "How much of this item should the store stock?\nIf unlimited just reply `skip` or `inf`" "embed" (cembed $embed))}}
						{{$cmdStage = 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
						{{$cmdStage = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 4}}
				{{else if eq $dbVal 4}}
					{{if or (gt (toInt $.Message.Content) 0) (eq (lower $.Message.Content) "inf" "skip")}}
						{{$quantity := $.Message.Content}}
						{{if eq $quantity "skip"}}
							{{$quantity = "inf"}}
						{{end}}
						{{$item := $createItem.item}}
						{{$item.data.Set "quantity" $quantity}}
						{{$msg := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$field := sdict "name" "Stock" "value" (toString $quantity) "inline" true}}
						{{$embed.Set "fields" ((cslice.AppendSlice $msg.Fields).Append $field)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "What role do you want to be given when this item is used?\nType `skip` to skip this." "embed" (cembed $embed))}}
						{{$cmdStage = 1}}
						{{dbSet 0 "createItem" $createItem}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again and enter a valid number"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
						{{$cmdStage = 0}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 5}}
				{{else if eq $dbVal 5}}
					{{$role := $.Message.Content}}
					{{if or ($.Guild.GetRole (toInt64 $role)) (eq $role "skip")}}
						{{if eq $role "skip"}}
							{{$role = 0}}
						{{end}}
						{{$item := $createItem.item}}
						{{$item.data.Set "role" $role}}
						{{dbSet 0 "createItem" $createItem}}
						{{if not $role}}
							{{$role = "none"}}
						{{else}}
							{{$role = print "<@&" $role ">"}}
						{{end}}
						{{$msg := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$field := sdict "name" "Role-given" "value" (toString $role) "inline" true}}
						{{$embed.Set "fields" ((cslice.AppendSlice $msg.Fields).Append $field)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "What message should I reply with when the item is used? (under 100 character)" "embed" (cembed $embed))}}
						{{$cmdStage = 1}}
					{{else if not (eq (lower $role) "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again with a valid roleID"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
					{{end}}
					{{scheduleUniqueCC $.CCID nil 120 1 6}}
				{{else if eq $dbVal 6}}
					{{if and (le (len (toRune $.Message.Content)) 100) (not (eq $.Message.Content "cancel"))}}
						{{$item := $createItem.item}}
						{{$item.data.Set "replyMsg" $.Message.Content}}
						{{dbSet 0 "createItem" $createItem}}
						{{$msg := structToSdict (index (getMessage nil (dbGet 0 "createItem").Value.embed).Embeds 0)}}
						{{$field := sdict "name" "Reply" "value" $.Message.Content}}
						{{$embed.Set "fields" ((cslice.AppendSlice $msg.Fields).Append $field)}}
						{{editMessage nil (dbGet 0 "createItem").Value.embed (complexMessageEdit "content" "Item created! ✅" "embed" (cembed $embed))}}
						{{$cmdStage = 0}}
						{{$items.Set $createItem.item.name $createItem.item.data}}
						{{dbSet 0 "store" $shop}}
						{{dbDel 0 "createItem"}}
						{{dbDel $userID "waitResponse"}}
						{{cancelScheduledUniqueCC $.CCID 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$msg := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $msg 10}}
						{{$cmdStage = 0}}
						{{dbDel 0 "createItem"}}
						{{scheduleUniqueCC $.CCID nil 120 1 6}}
					{{end}}
				{{end}}
				{{if eq (lower $.Message.Content) "cancel"}}
					{{sendMessage nil "Create-item was cancelled"}}
					{{dbDel $userID "waitResponse"}}
					{{dbDel 0 "createItem"}}
					{{$cmdStage = 0}}
					{{cancelScheduledUniqueCC $.CCID 1}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Del "footer"}}{{$embed.Del "title"}}
		{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
		{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.ServerPrefix "set default`")}}
		{{$embed.Set "color" $errorColor}}
		{{sendMessage nil (cembed $embed)}}
	{{end}}
{{else}}
	{{.ExecData}}
	{{sendMessage nil "Create-item was cancelled"}}
	{{dbDel .User.ID "waitResponse"}}
	{{dbDel 0 "createItem"}}
{{end}}
{{if $cmdStage}}
	{{dbSetExpire .User.ID "waitResponse" (str (add $dbVal 1)) 120}}
{{end}}