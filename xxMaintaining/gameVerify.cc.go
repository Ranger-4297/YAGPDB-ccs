{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `.*`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
*/}}


{{/* Configuration values start */}}
{{$serverChannel := 1185342810312945804}}		{{/* channelID of the **tag** channel */}}
{{$serverRole := 1185703179254505582}}			{{/* roleID of the **tag** role */}}
{{$allianceChannel := 1185725114566836304}} 	{{/* channelID of the **alliance** channel */}}
{{$allianceRole := 1185727965397516379}}		{{/* roleID of the **alliance** role */}}
{{$nameChannel := 1185350063417983067}}			{{/* channelID of the **name** channel */}}
{{$nameRole := 1185703261613863013}}			{{/* roleID of the **name** role */}}
{{$rankChannel := 1185723603086491648}}			{{/* channelID of the **2in1** channel */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$user := .User.ID}}
{{$errorColor := 0xFF0000}}
{{$embedColour := 0x2b2d31}}
{{$waitResponseDB := toInt (dbGet $user "displayNameWaitResponse").Value}}
{{$stage := 0}}
{{$prefix := print `(` .ServerPrefix `|<@!?204255221017214977>\s*)(setup)`}} 

{{/* User setup */}}

{{/* Response */}}
{{if eq .Channel.ID $serverChannel $allianceChannel $nameChannel}}
	{{$embed := sdict "color" $embedColour "footer" (sdict "text" (print "Welcome to" .Server.Name) "icon_url" (.Guild.IconURL "1024") "timestamp" currentTime)}}
	{{if eq .Channel.ID $serverChannel}}
		{{if toInt .Message.Content}}
			{{$server := .Message.Content}}
			{{if (reFind `^[\d]{3,4}$` $server)}}
				{{if (reFind `^(\[[\d]{3,4}\])` .Member.Nick)}}
					{{editNickname (reReplace `(\[[\d]{3,4}\])` .Member.Nick (printf "[%s]" $server))}}
				{{else}}
					{{editNickname (printf "[%s] %s" $server (joinStr "" (split .User.Globalname " ")))}}
					{{addRoleID $serverRole}}
					{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "You've just verified! Now you need to join an alliance <#" $allianceChannel ">"))}}
					{{deleteMessage nil $m 45}}
				{{end}}
				{{try}}
					{{addReactions ":white_check_mark:"}}
				{{catch}}
					{{sendMessage nil "Cannot add reaction to user who has blocked bot. Nickname updated"}}
				{{end}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please input a 3-4 digit tag"}}
				{{deleteMessage nil $m 45}}
			{{end}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please input a numeric tag"}}
			{{deleteMessage nil $m 45}}
		{{end}}
		{{$embed.Set "description" (print "`Insert your server #### to proceed 🚩`\nThis is *STAMPED* to the bottom of the channel:smirk:")}}
	{{else if eq .Channel.ID $allianceChannel}}
		{{$alliance := .Message.Content}}
		{{if not (eq (lower $alliance) "skip")}}
			{{if (reFind `^[a-zA-Z\d]+$` $alliance)}}
				{{if (reFind `^[a-zA-Z\d]{3,4}$` $alliance)}}
					{{if (reFind `^(\[[\d]{3,4}\]) (\[[a-zA-Z\d]{3,4}\])` .Member.Nick)}}{{/* If user has an alliance */}}
						{{editNickname (reReplace ` (\[[a-zA-Z\d]{3,4}\])` .Member.Nick (printf " [%s]" $alliance))}}
					{{else if (reFind `^(\[[\d]{3,4}\]) (.{3,15})$` .Member.Nick)}}{{/* If user does not have an alliance */}}
						{{editNickname (reReplace `(\[[\d]{3,4}\])` .Member.Nick (printf "$1 [%s]$2" $alliance))}}
					{{end}}
					{{try}}
						{{addReactions ":white_check_mark:"}}
					{{catch}}
						{{sendMessage nil "Cannot add reaction to user who has blocked bot. Nickname updated"}}
					{{end}}
					{{addRoleID $allianceRole}}
					{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Your alliance name has been updated. Please make your way to update your name at <#" $nameChannel ">"))}}
					{{deleteMessage nil $m 45}}
				{{else}}
					{{deleteTrigger 0}}
					{{$m := sendMessageRetID nil "Please input an alliance between 3 and 4 characters (a-Z)"}}
					{{deleteMessage nil $m 45}}
				{{end}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please don't use special characters characters (a-Z)"}}
				{{deleteMessage nil $m 45}}
			{{end}}
		{{else}}
			{{addRoleID $allianceRole}}
			{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Please make your way to update your name at <#" $nameChannel ">"))}}
			{{deleteMessage nil $m 45}}
		{{end}}
		{{$embed.Set "description" (print "`Insert your alliance tag #### to proceed 🏰`\n`💥 No alliance? Type \"skip\" 💥`\nThis is *STAMPED* to the bottom of the channel:smirk:")}}
	{{else if eq .Channel.ID $nameChannel}}
		{{$name := .Message.Content}}
		{{if not (reFind `[^a-zA-Z\d\s:]` $name)}}
			{{if (reFind `^([a-zA-Z\d]{3,15})$` $name)}}
				{{editNickname (reReplace `^(\[[\d]{3,4}\]) (\[[a-zA-Z]{3,4}\]) (.{3,15})$` .Member.Nick (printf "$1 $2 %s" $name))}}
				{{try}}
					{{addReactions ":white_check_mark:"}}
				{{catch}}
					{{sendMessage nil "Cannot add reaction to user who has blocked bot. Nickname updated"}}
				{{end}}
				{{addRoleID $nameRole}}
				{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Your display name has been updated. Please make your way to the <#" $rankChannel ">"))}}
				{{deleteMessage nil $m 45}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please input a username between 3 and 15 characters"}}
				{{deleteMessage nil $m 45}}
			{{end}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please only use unicode characters"}}
			{{deleteMessage nil $m 45}}
		{{end}}
		{{$embed.Set "description" (print "`Insert your game name to proceed 🎮`\nThis is *STAMPED* to the bottom of the channel:smirk:")}}
	{{end}}
	{{if $db := dbGet .Channel.ID "stickymessage"}}
		{{deleteMessage .Channel.ID (toInt $db.Value) 0}}
	{{end}}
	{{$id := sendMessageRetID nil (cembed $embed)}}
	{{dbSet .Channel.ID "stickymessage" (str $id)}}
{{end}}