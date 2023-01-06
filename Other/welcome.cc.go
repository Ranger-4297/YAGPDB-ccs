{{/* Variables */}}
{{$background := "https://i.pinimg.com/originals/d5/94/34/d5943460d6ec066d83a9838745df7742.jpg"}}
{{$accent := "red"}}

{{/* Calcuate placement of member */}}
{{$int := .Guild.MemberCount}}
{{$ord := "th"}}
{{$cent := toInt (mod $int 100)}}
{{$dec := toInt (mod $int 10)}}
{{if not (and (ge $cent 10) (le $cent 19))}}
    {{if eq $dec 1}}
        {{$ord = "st"}}
    {{else if eq $dec 2}}
        {{$ord = "nd"}}
    {{else if eq $dec 3}}
        {{$ord = "rd"}}
    {{end}}
{{end}}

{{/* Image */}}
{{$image := (print "https://rhykerw.com/API/welcome?background=" $background "&avatar=" (.User.AvatarURL "256") "&username=" (urlescape .User.String ) "&text=" (urlescape (print "You're our " $int $ord " member")) "&color=" (urlescape $accent))}}

{{/* Embed */}}
{{sendMessage nil (complexMessage
    "content" (print .User.Mention )
	"embed" (cembed
	"title" (print "WELCOME")
	"image" (sdict "url" $image)
	"color" 0x2f3136
))}}