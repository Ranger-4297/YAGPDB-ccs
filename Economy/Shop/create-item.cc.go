{{/*
Made by ranger_4297 (765316548516380732)
Wait response logic by DZ (438789314101379072)

Trigger Type: `Regex`
Trigger: `.*`

©️ RhykerWells 2020-Present
GNU, GPLV3 License
Repository: https://github.com/Ranger-4297/YAGPDB-ccs

Note: Command is `create-item`/`new-item`. Use your severs default prefix
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$prefix := .ServerPrefix}}
{{$userID := .User.ID}}
{{$errorColor := 0xFF0000}}
{{$successColor := 0x00ff7b}}
{{$databaseValue := toInt (dbGet .User.ID "waitResponse").Value}}
{{$currentSection := 0}}
{{$trigger := print `(` .ServerPrefix `|<@!?204255221017214977>\s*)((create|new)-?item)`}} 

{{/* Create item */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" (print .Guild.Name " Store")) "title" "Item info" "footer" (sdict "text" "Type cancel to cancel the setup") "timestamp" currentTime "color" 0x0088CC}}
{{$store := or (dbGet 0 "store").Value (sdict "items" sdict)}}
{{$items := $store.items}}
{{$itemData := sdict "desc" "" "price" 0 "quantity" 0 "role" 0 "replyMsg" "" "expiry" 0}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if .ExecData}}
	{{sendMessage nil "Create-item was cancelled"}}
	{{dbDel .User.ID "waitResponse"}}
	{{dbDel 0 "createItem"}}
	{{return}}
{{end}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Del "footer"}}{{$embed.Del "title"}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$symbol := $economySettings.symbol}}
{{if not $databaseValue}}
	{{if reFind (print `\A(?i)` $trigger `(\s+|\z)`) .Message.Content}}
		{{if not (or (in $perms "Administrator") (in $perms "ManageServer"))}}
			{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$embed.Set "fields" (cslice (sdict "name" "Name" "value" "⠀⠀"))}}
		{{$currentSection = 1}}
		{{$embed := sendMessageRetID nil (complexMessage "content" "Please enter a name for the item (under 60 characters)" "embed" (cembed $embed))}}
		{{dbSet 0 "createItem" (sdict "embed" (str $embed) "item" (sdict "name" "" "data" $itemData) "user" (str $userID))}}
		{{scheduleUniqueCC .CCID nil 120 1 1}}
	{{end}}
{{end}}
{{$createItem := (dbGet 0 "createItem").Value}}
{{if and $createItem.user (eq (str $createItem.user) (str $userID))}}
	{{if eq (lower .Message.Content) "cancel"}}
		{{$currentSection = 0}}
		{{sendMessage nil "Create-item was cancelled"}}
		{{dbDel $userID "waitResponse"}}
		{{dbDel 0 "createItem"}}
		{{cancelScheduledUniqueCC .CCID 1}}
		{{return}}
	{{end}}
	{{if eq $databaseValue 1}}
		{{if gt (len (toRune .Message.Content)) 60}}
			{{$message := sendMessageRetID nil "Please try again and enter name under 60 characters"}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 2}}
			{{return}}
		{{end}}
		{{$name := .Message.Content}}
		{{$item := $createItem.item}}
		{{$item.Set "name" $name}}
		{{$embed.Set "fields" (cslice (sdict "name" "Name" "value" $name "inline" true))}}
		{{$currentSection = 1}}
		{{dbSet 0 "createItem" $createItem}}
		{{editMessage nil $createItem.embed (complexMessageEdit "content" "Please enter a price for the item" "embed" (cembed $embed))}}
		{{scheduleUniqueCC .CCID nil 120 1 2}}
	{{else if eq $databaseValue 2}}
		{{$price := .Message.Content}}
		{{if or (not (toInt $price)) (lt (toInt $price) 1)}}
			{{$message := sendMessageRetID nil "Please try again and enter a valid number"}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 3}}
			{{return}}
		{{end}}
		{{$price = toInt $price}}
		{{$item := $createItem.item}}
		{{$item.data.Set "price" $price}}
		{{$message := structToSdict (index (getMessage nil $createItem.embed).Embeds 0)}}
		{{$fields := sdict "name" "Price" "value" (print $symbol (humanizeThousands $price)) "inline" true}}
		{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $fields)}}
		{{$currentSection = 1}}
		{{dbSet 0 "createItem" $createItem}}
		{{editMessage nil $createItem.embed (complexMessageEdit "content" "Please enter a description for the item (under 200 characters)" "embed" (cembed $embed))}}
		{{scheduleUniqueCC .CCID nil 120 1 3}}
	{{else if or (eq $databaseValue 3) (eq $databaseValue 6)}}
		{{$dbOption := "desc"}}
		{{$fieldValue := "description"}}
		{{if eq $databaseValue 6}}
			{{$dbOption = "replyMsg"}}
			{{$fieldValue = "reply"}}
		{{end}}
		{{if gt (len (toRune .Message.Content)) 200}}
			{{$message := sendMessageRetID nil (print "Please try again and enter a " $fieldValue " under 200 characters")}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 3}}
			{{return}}
		{{end}}
		{{$item := $createItem.item}}
		{{$item.data.Set $dbOption .Message.Content}}
		{{$message := structToSdict (index (getMessage nil $createItem.embed).Embeds 0)}}
		{{$fields := sdict "name" $fieldValue "value" .Message.Content}}
		{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $fields)}}
		{{if eq $databaseValue 3}}
			{{editMessage nil $createItem.embed (complexMessageEdit "content" "How much of this item should the store stock?\nIf unlimited just reply `skip` or `inf`" "embed" (cembed $embed))}}
		{{else}}
			{{editMessage nil $createItem.embed (complexMessageEdit "content" "When should this item expire from the users inventory if not used\nIf never just reply `skip` or `inf`" "embed" (cembed $embed))}}
		{{end}}
		{{$currentSection = 1}}
		{{dbSet 0 "createItem" $createItem}}
		{{scheduleUniqueCC .CCID nil 120 1 4}}
	{{else if eq $databaseValue 4}}
		{{if not (or (gt (toInt .Message.Content) 0) (eq (lower .Message.Content) "inf" "skip"))}}
			{{$message := sendMessageRetID nil "Please try again and enter a valid number"}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 3}}
			{{return}}
		{{end}}
		{{$qty := .Message.Content}}
		{{if eq (lower $qty) "skip" "inf"}}
			{{$qty = 0}}
		{{end}}
		{{$item := $createItem.item}}
		{{$item.data.Set "quantity" $qty}}
		{{dbSet 0 "createItem" $createItem}}
		{{if not $qty}}
			{{$qty = "Infinite"}}
		{{end}}
		{{$message := structToSdict (index (getMessage nil $createItem.embed).Embeds 0)}}
		{{$fields := sdict "name" "Stock" "value" (str $qty) "inline" true}}
		{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $fields)}}
		{{$currentSection = 1}}
		{{editMessage nil $createItem.embed (complexMessageEdit "content" "What role do you want to be given when this item is used?\nType `skip` to skip this." "embed" (cembed $embed))}}
		{{scheduleUniqueCC .CCID nil 120 1 6}}
	{{else if eq $databaseValue 5}}
		{{$role := .Message.Content}}
		{{if not (or (getRole $role) (eq (lower $role) "skip"))}}
			{{$message := sendMessageRetID nil "Please try again and enter a valid role ID/mention, or type `skip` to skip this step. "}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 3}}
			{{return}}
		{{end}}
		{{if reFind `skip` $role}}
			{{$role = 0}}
		{{end}}
		{{$item := $createItem.item}}
		{{$item.data.Set "role" $role}}
		{{dbSet 0 "createItem" $createItem}}
		{{if $role}}
			{{$role = print "<@&" $role ">"}}
		{{else}}
			{{$role = "none"}}
		{{end}}
		{{$message := structToSdict (index (getMessage nil $createItem.embed).Embeds 0)}}
		{{$fields := sdict "name" "Role-given" "value" (str $role) "inline" true}}
		{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $fields)}}
		{{$currentSection = 1}}
		{{editMessage nil $createItem.embed (complexMessageEdit "content" "What message should I reply with when the item is used? (under 200 character)" "embed" (cembed $embed))}}
		{{scheduleUniqueCC .CCID nil 120 1 7}}
	{{else if eq $databaseValue 7}}
		{{$duration := .Message.Content}}
		{{if not (or (toDuration $duration) (eq (lower $duration) "skip" "inf"))}}
			{{$message := sendMessageRetID nil "Please try again with a valid duration"}}
			{{$currentSection = 0}}
			{{deleteTrigger 0}}
			{{deleteMessage nil $message 10}}
			{{scheduleUniqueCC .CCID nil 120 1 7}}
			{{return}}
		{{end}}
		{{$eV := "Never"}}
		{{$duration = toInt (toDuration $duration).Seconds}}
		{{if $duration}}
			{{$eV = humanizeDurationSeconds (mult $duration .TimeSecond)}}
		{{end}}
		{{$item := $createItem.item}}
		{{$item.data.Set "expiry" $duration}}
		{{$items.Set $createItem.item.name $createItem.item.data}}
		{{$message := structToSdict (index (getMessage nil $createItem.embed).Embeds 0)}}
		{{$fields := sdict "name" "Inventory expiry" "value" (str $eV) "inline" true}}
		{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $fields)}}
		{{$embed.Set "color" $successColor}}
		{{editMessage nil $createItem.embed (complexMessageEdit "content" "Item created! ✅" "embed" (cembed $embed))}}
		{{dbSet 0 "store" $store}}
		{{dbDel 0 "createItem"}}
		{{dbDel $userID "waitResponse"}}
		{{cancelScheduledUniqueCC .CCID 1}}
	{{end}}
{{end}}
{{if $currentSection}}
	{{dbSetExpire .User.ID "waitResponse" (str (add $databaseValue 1)) 120}}
{{end}}