{{/*
        Made by Ranger (779096217853886504)

    Trigger Type: `Command`
    Trigger: `Appeal`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$FormLink := "https://realsite.co.uk/appeal"}} {{/* Your servers application link */}}
{{$logchannel := 794365711614345267}} {{/* Log channel*/}}
{{$IconType := "gif"}} {{/* set to `png` if your server uses a static icon */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing except for the 21st line. Add your ban appeal link at the end. (: rawr */}}

{{sendMessage nil "Check your DM's for info!"}}
{{$embed := cembed
			"author" (sdict "url" (print "https://discord.com/channels/" .Guild.ID) "name" (print .Guild.Name " Ban appeals") "icon_url" (.User.AvatarURL "1024"))
			"description" (print  .User.Mention "\n\nHello " .User.Username ". You've requested " .Guild.Name "'s ban appeal form.\nThis means you are either looking to reverse your ban, or are simply taking a look at our appeal.\nPlease keep in mind that asking about the status of your appeal may warrant it's status as `denied`.\n[Ban appeal form](" $FormLink ")")
			"thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $IconType ))
			"footer" (sdict "text" " ")
			"timestamp" currentTime
			"color" 4645612
			}}
{{sendDM $embed}}


{{$logembed := cembed
			"author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
			"description" (print "**✉️ Ban appeal command notification**\n" .User.Mention " Has requested a ban appeal form.")
			"thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $IconType ))
			"footer" (sdict "text" " ")
			"timestamp" currentTime
			"color" 4645612
			}}
{{sendMessageNoEscape $logchannel $logembed}}