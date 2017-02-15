axxxxxxxxxxxxx

{{.user_list}}
{{ .test_data["name"] }}
{{ccTest "nice"}}
<table>
    <tr>
        <th>id</th>
        <th>名字</th>
        <th>年龄</th>
        <th>生日</th>
    </tr>

    {{ range $i, $user := .user_list }}
        {{if isRightType $user}}
        <tr>
            <td>{{$i}}</td>
            <td>{{$user}}</td>
            <td>{{.test_data}}</td>
        </tr>
        {{end}}
    {{end}}

</table>
