{{/*
		Made by ranger_4297 (765316548516380732)
		Wait response logic by DZ (438789314101379072)

	Trigger Type: `Regex`
	Trigger: `.*`

	©️ Ranger 2020-Present
	GNU, GPLV3 License

	Note: Command is `setup`. Use your severs default prefix
*/}}

{{/* Configuration values start */}}
{{$serverChannel := 1185342810312945804}} {{/* channelID of the **tag** channel */}}
{{$serverRole := 1185703179254505582}} {{/* roleID of the **tag** role */}}
{{$allianceChannel := 1185725114566836304}} {{/* channelID of the **alliance** channel */}}
{{$allianceRole := 1185727965397516379}} {{/* roleID of the **alliance** role */}}
{{$nameChannel := 1185350063417983067}} {{/* channelID of the **name** channel */}}
{{$nameRole := 1185703261613863013}} {{/* roleID of the **name** role */}}
{{$multipleChannel := 1185342836518957187}} {{/* channelID of the **2in1** channel */}}
{{$rankChannel := 1185723603086491648}}
{{$enableMultiple := false}} {{/* Enable the 2in1 version */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User.ID}}
{{$errorColor := 0xFF0000}}
{{$waitResponseDB := toInt (dbGet $user "displayNameWaitResponse").Value}}
{{$stage := 0}}
{{$prefix := print `(` .ServerPrefix `|<@!?204255221017214977>\s*)(setup)`}} 

{{/* User setup */}}

{{/* Response */}}
{{if and (eq .Channel.ID $multipleChannel) $enableMultiple}} {{/* Command based 2 in 1 */}}
	{{$embed := sdict "title" "User data" "footer" (sdict "text" "Type cancel to cancel your assignment") "color" 0x00ff7b "timestamp" currentTime}}
	{{if not .ExecData}}
		{{if not $waitResponseDB}}
			{{if reFind (print `\A(?i)` $prefix `(\s+|\z)`) $.Message.Content}}
				{{$embed.Set "fields" (cslice (sdict "name" "Server tag" "value" "⠀⠀"))}}
				{{$stage = 1}}
				{{$embedID := sendMessageRetID nil (complexMessage "content" "Please enter your servers tag (000-9999)" "embed" (cembed $embed))}}
				{{dbSet $user "displayName" (sdict "embed" (toString $embedID) "userData" (sdict  "tag" 0 "name" "") "user" (toString $user))}}
				{{scheduleUniqueCC .CCID nil 120 1 1}}
			{{end}}
		{{else}}
			{{$userData := (dbGet $user "displayName").Value}}
			{{if and $userData.user (eq (toString $userData.user) (toString $user))}}
				{{if eq $waitResponseDB 1}}
					{{if and (reFind `\A(?i)([\d]{3,4})(\s+|\z)` $.Message.Content) (not (eq $.Message.Content "cancel"))}}
						{{$tag := (reFind `\A(?i)([\d]{3,4})(\s+|\z)` $.Message.Content) | toInt}}
						{{$data := $userData.userData}}
						{{$data.Set "tag" $tag}}
						{{dbSet $user "displayName" $userData}}
						{{$embed.Set "fields" (cslice (sdict "name" "Server tag" "value" (toString $tag) "inline" true))}}
						{{editMessage nil $userData.embed (complexMessageEdit "content" "Please enter your display name (between 3 & 15 characters)" "embed" (cembed $embed))}}
						{{$stage = 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$m := sendMessageRetID nil "Please try again and input a 3-4 digit tag"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$stage = 0}}
					{{end}}
					{{scheduleUniqueCC .CCID nil 120 1 2}}
				{{else if eq $waitResponseDB 2}}
					{{scheduleUniqueCC .CCID nil 120 1 2}}
					{{if and (ge (len (toRune $.Message.Content)) 3) (le (len (toRune $.Message.Content)) 15) (not (eq $.Message.Content "cancel"))}}
						{{$name := $.Message.Content}}
						{{$data := $userData.userData}}
						{{$message := structToSdict (index (getMessage nil $userData.embed).Embeds 0)}}
						{{$field := sdict "name" "Game username" "value" (toString $name) "inline" true}}
						{{$embed.Set "fields" ((cslice.AppendSlice $message.Fields).Append $field)}}
						{{$embed.Del "footer"}}
						{{editMessage nil $userData.embed (complexMessageEdit "content" "User submitted ✅" "embed" (cembed $embed))}}
						{{dbDel 0 "displayName"}}
						{{dbDel $user "displayNameWaitResponse"}}
						{{editNickname (printf "[%d] | %s" $data.tag $name)}}
						{{addRoleID $serverRole}}
						{{addRoleID $nameRole}}
						{{cancelScheduledUniqueCC .CCID 1}}
					{{else if not (eq $.Message.Content "cancel")}}
						{{$m := sendMessageRetID nil "Please try again with a username of 3-15 characters"}}
						{{deleteTrigger 0}}
						{{deleteMessage nil $m 10}}
						{{$stage = 0}}
					{{end}}
				{{end}}
				{{if eq (lower $.Message.Content) "cancel"}}
					{{sendMessage nil (print "Server display cancelled. Please begin with the `" .ServerPrefix "setup` command")}}
					{{dbDel $user "displayName"}}
					{{dbDel $user "displayNameWaitResponse"}}
					{{$stage = 0}}
					{{cancelScheduledUniqueCC .CCID 1}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{sendMessage nil (print "Server display cancelled. Please send the `" .ServerPrefix "setup` command to set up your tag and username")}}
		{{dbDel $user "displayNameWaitResponse"}}
		{{dbDel $user "displayName"}}
	{{end}}
	{{if $stage}}
		{{dbSetExpire $user "displayNameWaitResponse" (str (add $waitResponseDB 1)) 120}}
	{{end}}
{{else if eq .Channel.ID $serverChannel}}
	{{if toInt .Message.Content}}
		{{$server := .Message.Content}}
		{{if (reFind `^[\d]{3,4}$` $server)}}
			{{if (reFind `^(\[[\d]{3,4}\]) (\[[a-zA-Z\d]{3,4}\])` .Member.Nick)}}
				{{editNickname (reReplace `(\[[\d]{3,4}\])` .Member.Nick (print "[" $server "]"))}}
			{{else}}
				{{editNickname (printf "[%s] %s" $server (joinStr "" (split .User.Globalname " ")))}}
				{{addRoleID $serverRole}}
				{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "You've just verified! Now you need to join an alliance <#" $allianceChannel ">"))}}
				{{deleteMessage nil $m 10}}
			{{end}}
			{{addReactions ":white_check_mark:"}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please input a 3-4 digit tag"}}
			{{deleteMessage nil $m 10}}
		{{end}}
	{{else}}
		{{deleteTrigger 0}}
		{{$m := sendMessageRetID nil "Please input a numeric tag"}}
		{{deleteMessage nil $m 10}}
	{{end}}
{{else if eq .Channel.ID $allianceChannel}}
	{{$alliance := .Message.Content}}
	{{if not (eq (lower $alliance) "skip")}}
		{{if (reFind `^[a-zA-Z\d]+$` $alliance)}}
			{{if (reFind `^[a-zA-Z\d]{3,4}$` $alliance)}}
				{{if (reFind `^(\[[\d]{3,4}\]) (\[[a-zA-Z\d]{3,4}\])` .Member.Nick)}}
					{{editNickname (reReplace ` (\[[a-zA-Z\d]{3,4}\])` .Member.Nick (printf " [%s]" $alliance))}}
				{{else if (reFind `^(\[[\d]{3,4}\]) ([a-zA-Z\d]{3,15})` .Member.Nick)}}
					{{editNickname (reReplace `(\[[\d]{3,4}\]) ([a-zA-Z\d]{3,15})` .Member.Nick (printf "$1 [%s] $2" $alliance))}}
				{{end}}
				{{addReactions ":white_check_mark:"}}
				{{addRoleID $allianceRole}}
				{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Your alliance name has been updated. Please make your way to update your name at <#" $nameChannel ">"))}}
				{{deleteMessage nil $m 10}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please input an alliance between 3 and 4 characters (a-Z)"}}
				{{deleteMessage nil $m 10}}
			{{end}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please only use alphabetical characters (a-Z)"}}
			{{deleteMessage nil $m 10}}
		{{end}}
	{{else}}
		{{addRoleID $allianceRole}}
		{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Please make your way to update your name at <#" $nameChannel ">"))}}
		{{deleteMessage nil $m 10}}
	{{end}}
	{{$embed := sdict "title" "`[!]` **Important** `[!]`" "description" "IF YOU ARE NOT PART OF AN ALLIANCE. TYPE `skip`"}}
	{{if $db := dbGet $allianceChannel "stickymessage"}}
		{{deleteMessage $allianceChannel (toInt $db.Value) 0}}
	{{end}}
	{{$id := sendMessageRetID nil (cembed $embed)}}
	{{dbSet $allianceChannel "stickymessage" (str $id)}}
{{else if eq .Channel.ID $nameChannel}}
	{{$name := .Message.Content}}
	{{if not (reFind `[^a-zA-Z\d\s:]` $name)}}
		{{if (reFind `^([a-zA-Z\d]{3,15})$` $name)}}
			{{editNickname (reReplace `([a-zA-Z\d]{3,15})$` .Member.Nick $name)}}
			{{addReactions ":white_check_mark:"}}
			{{addRoleID $nameRole}}
			{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Your display name has been updated. Please make your way to the <#" $rankChannel ">"))}}
			{{deleteMessage nil $m 10}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please input a username between 3 and 15 characters"}}
			{{deleteMessage nil $m 10}}
		{{end}}
	{{else}}
		{{deleteTrigger 0}}
		{{$m := sendMessageRetID nil "Please only use alphanumeric characters (a-Z/0-9)"}}
		{{deleteMessage nil $m 10}}
	{{end}}
{{end}}