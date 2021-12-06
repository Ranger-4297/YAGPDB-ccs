{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Join message in channel`
©️ Ranger 2021
MIT License
*/}}

{{/* Alternative join message */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{ $int := .Guild.MemberCount }}
{{ $ord := "th" }}
{{ $cent := toInt (mod $int 100) }}
{{ $dec := toInt (mod $int 10) }}
{{ if not (and (ge $cent 10) (le $cent 19)) }}
	{{ if eq $dec 1 }}
        {{ $ord = "st" }}
	{{ else if eq $dec 2 }}
        {{ $ord = "nd" }}
	{{ else if eq $dec 3 }}
        {{ $ord = "rd" }}
	{{ end }}
{{ end }}

{{if .UsernameHasInvite}}
    {{$silent := (execAdmin "ban" .User.ID "Advert in UserID blocked")}}
{{else}}
    {{sendMessage nil (complexMessage
            "content" (print .User.Mention )
            "embed" (cembed
            "title" (print "WELCOME")
            "image" (sdict "url" (print "https://api.no-api-key.com/api/v2/welcome?username=" (print "You%27re%20our%20" $int $ord "%20member") "&background=https://lovelytab.com/wp-content/uploads/2019/01/Tumblr-Aesthetic-Wallpapers-Free.jpg&user_image=" (.User.AvatarURL "256") "&text_heading=" (urlescape .User.String ) "&color=black"))
            "thumbnail" (sdict "url" (.User.AvatarURL "1020"))
        ))}}
{{end}}