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
	{{$embed := sdict "thumbnail" (sdict "url" (.Guild.IconURL "1024")) "color" $embedColour "footer" (sdict "text" (print "Welcome to " .Server.Name) "icon_url" (.Guild.IconURL "1024")) "timestamp" currentTime}}
	{{if eq .Channel.ID $serverChannel}}
		{{if toInt .Message.Content}}
			{{$server := .Message.Content}}
			{{if (reFind `^[\d]{3,4}$` $server)}}
				{{if (reFind `^(\[[\d]{3,4}\])` .Member.Nick)}}
					{{editNickname (reReplace `(\[[\d]{3,4}\])` .Member.Nick (printf "[%s]" $server))}}
				{{else}}
					{{editNickname (printf "[%s] %s" $server (joinStr "" (split .User.Globalname " ")))}}
					{{addRoleID $serverRole}}
				{{end}}
				{{try}}
					{{addReactions ":white_check_mark:"}}
				{{catch}}
					{{sendMessage nil "Cannot add reaction to user who has blocked bot. Nickname updated"}}
				{{end}}
				{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (printf "Awesome! %s You've just received your Server Tag! Now proceed to <#%d>. <1/5>" .User.Mention $allianceChannel))}}
				{{deleteMessage nil $m 45}}
				{{$m2 := sendMessageNoEscapeRetID $allianceChannel (printf "%s Enter your alliance tag here <2/5>" .User.Mention)}}
				{{deleteMessage $allianceChannel $m2 45}}
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
		{{$embed.Set "description" (print "`Step 1/5: Type your Server Number #### 📩`\n\n`Click `<#" $allianceChannel ">`to proceed in verification`\n\n*This is STAMPED to the bottom of the channel!*:smirk:")}}
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
					{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (printf "This is cool! %s you now have an Alliance Tag! Continue and proceed to <#%d>. <2/5>" .User.Mention $nameChannel))}}
					{{deleteMessage nil $m 45}}
					{{$m2 := sendMessageNoEscapeRetID $nameChannel (printf "%s type your Character Game Name over here! <3/5>" .User.Mention)}}
					{{deleteMessage $nameChannel $m2 45}}
				{{else}}
					{{deleteTrigger 0}}
					{{$m := sendMessageRetID nil "Please input a 3-4 character alliance tag"}}
					{{deleteMessage nil $m 45}}
				{{end}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please don't use special characters"}}
				{{deleteMessage nil $m 45}}
			{{end}}
		{{else}}
			{{addRoleID $allianceRole}}
			{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (print "Please make your way to update your name at <#" $nameChannel ">"))}}
			{{deleteMessage nil $m 45}}
		{{end}}
		{{$embed.Set "description" (print "`Step 2/5: Type your Alliance Tag #### 🏰`\n\n`Click `<#" $nameChannel ">`to proceed in verification`\n\n`💥 No alliance? Type \"skip\" 💥`\n\n*This is STAMPED to the bottom of the channel!*:smirk:")}}
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
				{{$m := sendMessageNoEscapeRetID nil (complexMessage "reply" .Message.ID "content" (printf "Perfect! %s you've updated your display name to your Character Game Name! Two more steps, proceed to <#%d>. <3/5>" .User.Mention $rankChannel))}}
				{{deleteMessage nil $m 45}}
				{{$m2 := sendMessageNoEscapeRetID $rankChannel (printf "%s select your in-game alliance rank. <4/5>" .User.Mention)}}
				{{deleteMessage $rankChannel $m2 45}}
			{{else}}
				{{deleteTrigger 0}}
				{{$m := sendMessageRetID nil "Please input a 3-15 character username"}}
				{{deleteMessage nil $m 45}}
			{{end}}
		{{else}}
			{{deleteTrigger 0}}
			{{$m := sendMessageRetID nil "Please don't use special characters"}}
			{{deleteMessage nil $m 45}}
		{{end}}
		{{$embed.Set "description" (print "`Step 3/5: Type your game name to proceed 🎮`\n\n`Click `<#" $rankChannel ">`to proceed in verification`\n\n*This is STAMPED to the bottom of the channel!*:smirk:")}}
	{{end}}
	{{if $db := dbGet .Channel.ID "stickymessage"}}
		{{deleteMessage .Channel.ID (toInt $db.Value) 0}}
	{{end}}
	{{$id := sendMessageRetID nil (cembed $embed)}}
	{{dbSet .Channel.ID "stickymessage" (str $id)}}
{{end}}