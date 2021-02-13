{{/*
        Made by Ranger (779096217853886504)

    Trigger Type: `Command`
    Trigger: `Apply`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$FormLink := "https://realsite.co.uk/application"}} {{/* Your servers application link */}}
{{$logchannel := 794365711614345267}}
{{$IconType := "gif"}} {{/* set to `png` if your server uses a static icon */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing except for the 21st line. Add your staff application link at the end. (: rawr */}}

{{sendMessage nil "Check your DM's for info!"}}
{{$embed := cembed
            "author" (sdict "url" (print "https://discord.com/channels/" .Guild.ID) "name" (print .Guild.Name " staff applications") "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Mention ",\nHello " .User.Username "\nThank you for taking interest in our servers staff team!\nPlease keep in mind that asking about the status of your application may warrant it's status as `denied`.\n[Application form](application)")
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $IconType ))
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 4645612
            }}
{{sendDM $embed}}


{{$logembed := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**✉️ Application command notification**\n" .User.Mention " Has requested a staff application.")
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $IconType))
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 4645612
            }}
{{sendMessageNoEscape $logchannel $logembed}}